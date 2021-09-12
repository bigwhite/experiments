package main

import (
	"fmt"
	"sync"

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc/resolver"
)

type NacosResolver struct {
	namingClient naming_client.INamingClient
	cc           resolver.ClientConn
	namespace    string
	group        string
	serviceName  string
	sync.Mutex
}

func (r *NacosResolver) doResolve(opts resolver.ResolveNowOptions) {
	instances, err := r.namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: r.serviceName,
		GroupName:   r.group,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(instances) == 0 {
		fmt.Printf("service %s has zero instance\n", r.serviceName)
		return
	}

	// update cc.States
	var addrs []resolver.Address
	for i, inst := range instances {
		if (!inst.Enable) || (inst.Weight == 0) {
			continue
		}

		addr := resolver.Address{
			Addr:       fmt.Sprintf("%s:%d", inst.Ip, inst.Port),
			ServerName: fmt.Sprintf("instance-%d", i+1),
		}
		addr.Attributes = addr.Attributes.WithValues("weight", int(inst.Weight))
		addrs = append(addrs, addr)
	}

	if len(addrs) == 0 {
		fmt.Printf("service %s has zero valid instance\n", r.serviceName)
	}

	newState := resolver.State{
		Addresses: addrs,
	}

	r.Lock()
	r.cc.UpdateState(newState)
	r.Unlock()
}

func (r *NacosResolver) ResolveNow(opts resolver.ResolveNowOptions) {
	r.doResolve(opts)
}

func (r *NacosResolver) Close() {
	r.namingClient.Unsubscribe(&vo.SubscribeParam{
		ServiceName: r.serviceName,
		GroupName:   r.group,
	})
}

func (r *NacosResolver) watch() {
	r.namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: r.serviceName,
		GroupName:   r.group,
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			fmt.Printf("subcallback: %#v\n", services)
			r.doResolve(resolver.ResolveNowOptions{})
		},
	})
}
