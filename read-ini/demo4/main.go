package main

import (
	"github.com/bigwhite/readini/pkg/pkg1"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("demo")
	viper.SetConfigType("ini")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	pkg1.Foo()
}
