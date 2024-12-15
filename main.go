package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/elct9620/demo-stdio-go-plugin/internal/controller"
	"github.com/elct9620/demo-stdio-go-plugin/internal/plugin"
	"github.com/elct9620/demo-stdio-go-plugin/internal/usecase"
)

type StdoutPresenter struct{}

func (p *StdoutPresenter) Render(data []byte) error {
	fmt.Println("")
	fmt.Println("Encoded data:")
	fmt.Println(string(data))
	return nil
}

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

	presenter := &StdoutPresenter{}
	usecase := usecase.NewEncodeProduct(client, presenter)

	controller := controller.NewStdin(usecase)
	err = controller.Handle(context.Background(), os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
}
