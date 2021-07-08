package pkg1

import (
	"fmt"

	"github.com/spf13/viper"
)

func Foo() {
	id := viper.GetString("server.id")
	fmt.Printf("id = [%s]\n", id)
	tlsPort := viper.GetInt("server.tls_port")
	fmt.Printf("tls_port = [%d]\n", tlsPort)

	logPath := viper.GetString("log.path")
	fmt.Printf("path = [%s]\n", logPath)
	if viper.IsSet("log.path1") {
		logPath1 := viper.GetString("log.path1")
		fmt.Printf("path1 = [%s]\n", logPath1)
	} else {
		fmt.Printf("log.path1 is not found\n")
	}
	logPath1 := viper.GetString("log.path1")
	fmt.Printf("path1 = [%s]\n", logPath1) // 获得空字符串
}
