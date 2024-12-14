package main

import (
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/pkg/sdk"
)

var _ sdk.EchoService = &Echo{}

type Echo struct{}

func (e *Echo) Ping(req sdk.EchoRequest, reply *sdk.EchoResponse) error {
	reply.Msg = fmt.Sprintf("Echo: %s", req.Msg)
	return nil
}

func main() {
	plugin := sdk.NewPlugin()
	plugin.Register(&Echo{})
	plugin.Start()
}
