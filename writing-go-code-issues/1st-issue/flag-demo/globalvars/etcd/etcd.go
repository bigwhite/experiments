package etcd

import (
	"fmt"

	"../config"
)

func EtcdProxy() {
	fmt.Println(config.Endpoints, config.User, config.Password)
}
