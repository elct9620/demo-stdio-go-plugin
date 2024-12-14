package main

import (
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/pkg/sdk"
)

type Request struct {
	Msg string
}

type Echo struct{}

func (e *Echo) Ping(req Request, reply *Request) error {
	reply.Msg = fmt.Sprintf("Echo: %s", req.Msg)
	return nil
}

func main() {
	plugin := sdk.NewPlugin()
	plugin.Register(&Echo{})
	plugin.Start()
}
