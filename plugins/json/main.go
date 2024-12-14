package main

import (
	"fmt"
	"io"
	"net/rpc"
	"os"
)

type Request struct {
	Msg string
}

type Echo struct{}

func (e *Echo) Ping(req Request, reply *Request) error {
	reply.Msg = fmt.Sprintf("Echo: %s", req.Msg)
	return nil
}

type conn struct {
	io.Reader
	io.WriteCloser
}

func main() {
	conn := &conn{
		Reader:      os.Stdin,
		WriteCloser: os.Stdout,
	}

	server := rpc.NewServer()
	defer conn.Close()

	server.Register(&Echo{})
	server.ServeConn(conn)
}
