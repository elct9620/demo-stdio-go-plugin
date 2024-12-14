package plugin

import (
	"fmt"
	"net/rpc"
	"os/exec"
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

func (c *Client) Call(service string, args any, reply any) error {
	err := c.rpc.Call(service, args, reply)
	if err != nil {
		return err
	}

	return nil
}
