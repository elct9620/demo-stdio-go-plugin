package plugin

import (
	"fmt"
	"net/rpc"
	"os/exec"

	"github.com/elct9620/demo-stdio-go-plugin/pkg/sdk"
)

type Client struct {
	cmd *exec.Cmd
	rpc *rpc.Client
}

func NewClient(cmd *exec.Cmd) (*Client, error) {
	conn, err := newConn(cmd)
	if err != nil {
		return nil, fmt.Errorf("Error creating connection for plugin %s: %v", cmd.Path, err)
	}

	return &Client{
		cmd: cmd,
		rpc: rpc.NewClient(conn),
	}, nil
}

func (c *Client) Close() (err error) {
	defer func() {
		err = c.cmd.Process.Kill()
	}()

	return c.rpc.Close()
}

func (c *Client) Ping(msg string) (reply string, err error) {
	var req sdk.EchoRequest
	err = c.rpc.Call("Echo.Ping", sdk.EchoRequest{Msg: msg}, &req)
	if err != nil {
		return "", fmt.Errorf("Error calling Ping: %v", err)
	}

	return req.Msg, nil
}
