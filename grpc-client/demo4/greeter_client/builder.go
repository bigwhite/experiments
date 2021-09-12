package main

import (
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&NacosBuilder{})
}

type NacosBuilder struct{}

func initNamingClient(host, port, namespace, group string) (naming_client.INamingClient, error) {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		/*
			constant.WithLogDir("/tmp/nacos/log/"+strconv.Itoa(int(port))),
			constant.WithCacheDir("/tmp/nacos/cache/"+strconv.Itoa(int(port))),
			constant.WithRotateTime("1h"),
			constant.WithMaxAge(3),
			constant.WithLogLevel("debug"),
		*/
	)

	portNum, _ := strconv.Atoi(port)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      host,
			ContextPath: "/nacos",
			Port:        uint64(portNum),
			Scheme:      "http",
		},
	}

	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	return namingClient, err
}

func (nb *NacosBuilder) Build(target resolver.Target,
	cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {

	// use info in target to access naming service
	// parse the target.endpoint
	// target.Endpoint - localhost:8848/public/DEFAULT_GROUP/serviceName, the addr of naming service :nacos endpoint
	sl := strings.Split(target.Endpoint, "/")
	nacosAddr := sl[0]
	namespace := sl[1]
	group := sl[2]
	serviceName := sl[3]
	sl1 := strings.Split(nacosAddr, ":")
	host := sl1[0]
	port := sl1[1]
	namingClient, err := initNamingClient(host, port, namespace, group)
	if err != nil {
		return nil, err
	}

	r := &NacosResolver{
		namingClient: namingClient,
		cc:           cc,
		namespace:    namespace,
		group:        group,
		serviceName:  serviceName,
	}

	// initialize the cc's states
	r.ResolveNow(resolver.ResolveNowOptions{})

	// subscribe and watch
	r.watch()
	return r, nil
}

func (nb *NacosBuilder) Scheme() string {
	return "nacos"
}
