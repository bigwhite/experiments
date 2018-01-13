package etcd

import (
	"flag"
	"fmt"
)

func EtcdProxy() {
	endpoints := flag.Lookup("endpoints").Value.(flag.Getter).Get().(string)
	user := flag.Lookup("user").Value.(flag.Getter).Get().(string)
	password := flag.Lookup("password").Value.(flag.Getter).Get().(string)

	fmt.Println(endpoints, user, password)
}
