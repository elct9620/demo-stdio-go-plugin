package plugin

import (
	"io"
	"os/exec"
)

type conn struct {
	io.Reader
	io.WriteCloser
}

func newConn(cmd *exec.Cmd) (*conn, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	return &conn{
		Reader:      stdout,
		WriteCloser: stdin,
	}, nil
}
