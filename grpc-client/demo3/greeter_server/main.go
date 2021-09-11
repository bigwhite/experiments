/*
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var port int

func init() {
	flag.IntVar(&port, "port", 50051, "listen port")
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func register(host string, port uint64) (func() error, error) {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //When namespace is public, fill in the blank string here.
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log/"+strconv.Itoa(int(port))),
		constant.WithCacheDir("/tmp/nacos/cache/"+strconv.Itoa(int(port))),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("debug"),
	)

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "localhost",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		return nil, err
	}

	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        port,
		ServiceName: "demo3-service",
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shenyang"},
		//ClusterName: "cluster-a", // default value is DEFAULT
		GroupName: "group-a", // default value is DEFAULT_GROUP
	})
	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errors.New("register instance failed")
	}

	deregisterFunc := func() error {
		success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          host,
			Port:        port,
			ServiceName: "demo3-service",
			Ephemeral:   true,
			Cluster:     "cluster-a", // default value is DEFAULT
			GroupName:   "group-a",   // default value is DEFAULT_GROUP
		})
		if err != nil {
			fmt.Println("unregister instance error:", err)
			return err
		}

		if !success {
			fmt.Println("unregister instance failed")
			return errors.New("unregister instance failed")
		}
		fmt.Println("unregister instance ok")

		return nil
	}

	return deregisterFunc, nil
}

func main() {
	flag.Parse()

	// register this serive to nacos
	df, err := register("localhost", uint64(port))
	if err != nil {
		panic(err)
		return
	}
	defer df()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	var c = make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	s.Stop()
	fmt.Println("exit")
}
