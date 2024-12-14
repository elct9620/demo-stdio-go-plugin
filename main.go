package main

import (
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/internal/plugin"
)

func main() {
	p, err := plugin.NewPlugin("./plugin-bin/json")
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

	reply, err := client.Ping("Hello World")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

	reply, err = client.Ping("Another Message")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reply)
}
