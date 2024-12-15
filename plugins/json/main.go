package main

import (
	"encoding/json"
	"fmt"

	"github.com/elct9620/demo-stdio-go-plugin/pkg/sdk"
)

var _ sdk.EchoService = &Echo{}

type Echo struct{}

func (e *Echo) Ping(req sdk.EchoRequest, reply *sdk.EchoResponse) error {
	reply.Msg = fmt.Sprintf("Echo: %s", req.Msg)
	return nil
}

type Encoder struct{}

func (e *Encoder) Encode(req sdk.EncodeRequest, reply *sdk.EncodeResponse) error {
	res, err := json.Marshal(req.Items)
	if err != nil {
		return err
	}

	reply.Result = res
	return nil
}

func main() {
	plugin := sdk.NewPlugin()
	plugin.Register(&Echo{})
	plugin.Register(&Encoder{})
	plugin.Start()
}
