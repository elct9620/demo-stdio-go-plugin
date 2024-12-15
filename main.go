package main

import (
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/internal/plugin"
	"github.com/elct9620/demo-stdio-go-plugin/pkg/sdk"
)

func main() {
	p, err := plugin.NewPlugin("./plugin-bin/xml")
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

	products := []sdk.Item{
		{
			Name:  "Apple",
			Price: 40,
		},
		{
			Name:  "Banana",
			Price: 30,
		},
	}

	res, err := client.Encode(&sdk.EncodeRequest{
		Items: products,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res.Result))
}
