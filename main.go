package main

import (
	"flag"
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/internal/entity"
	"github.com/elct9620/demo-stdio-go-plugin/internal/plugin"
)

func main() {
	// Section: entrypoint
	var pluginName = "json"
	flag.StringVar(&pluginName, "plugin", pluginName, "plugin name")
	flag.Parse()

	manager := plugin.NewManager()
	err := manager.Discover("./plugin-bin")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Section: controller
	p, err := manager.Get(pluginName)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := p.Client()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// Section: usecase
	products := []*entity.Product{
		entity.NewProduct("Apple", 10),
		entity.NewProduct("Banana", 20),
	}

	res, err := client.Encode(products)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res))
}
