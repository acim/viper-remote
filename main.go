package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	viper.AddRemoteProvider("consul", "consul:8500", "config/app")
	viper.SetConfigType("yaml")

	for {
		err := viper.ReadRemoteConfig()
		if err != nil {
			fmt.Println("Please set key config/app in Consul with YML ")
		}
		fmt.Println(viper.Get("port"))
		time.Sleep(5 * time.Second)
	}
}
