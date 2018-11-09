package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	c := api.DefaultConfig()
	c.Address = "consul:8500"
	consul, err := api.NewClient(c)
	if err != nil {
		panic(err)
	}

	services, _, err := consul.Catalog().Service("app", "", &api.QueryOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(services)

	select {}
}
