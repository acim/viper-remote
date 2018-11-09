package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type config struct {
	Port int
	mux  sync.Mutex
}

func main() {
	rv := viper.New()
	rv.AddRemoteProvider("consul", "consul:8500", "config/app")
	rv.SetConfigType("yaml")
	err := rv.ReadRemoteConfig()
	if err != nil {
		fmt.Println("Please set key config/app in Consul")
	}

	c := &config{}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			err := rv.WatchRemoteConfig()
			if err != nil {
				fmt.Println("Please set key config/app in Consul")
				continue
			}

			c.mux.Lock()
			rv.Unmarshal(c)
			c.mux.Unlock()
		}
	}()

	for {
		time.Sleep(time.Second * 5)
		fmt.Printf("%#v\n", c)
	}
}
