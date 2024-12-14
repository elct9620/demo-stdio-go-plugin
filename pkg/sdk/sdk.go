package sdk

import (
	"io"
	"net/rpc"
	"os"
)

type conn struct {
	io.Reader
	io.WriteCloser
}

type Operation string

type Plugin struct {
	server *rpc.Server
}

func NewPlugin() *Plugin {
	return &Plugin{
		server: rpc.NewServer(),
	}
}

func (p *Plugin) Register(service any) error {
	return p.server.Register(service)
}

func (p *Plugin) Start() {
	conn := &conn{
		Reader:      os.Stdin,
		WriteCloser: os.Stdout,
	}

	p.server.ServeConn(conn)
	conn.Close()
}
