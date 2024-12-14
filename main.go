package main

import (
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/internal/plugin"
)

type Request struct {
	Msg string
}

func main() {
	p, err := plugin.NewPlugin("/bin/cat")
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

	var reply Request
	err = client.Call("Echo.Ping", &Request{Msg: "Hello, World!"}, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reply.Msg)

	err = client.Call("Echo.Ping", &Request{Msg: "Another Message"}, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(reply.Msg)
}
