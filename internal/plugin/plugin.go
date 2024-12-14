package plugin

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
)

type Plugin struct {
	name string
	path string
}

func NewPlugin(pluginPath string) (*Plugin, error) {
	name := path.Base(pluginPath)
	absPath, err := filepath.Abs(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("Error getting absolute path for plugin %s: %v", pluginPath, err)
	}

	return &Plugin{
		name: name,
		path: absPath,
	}, nil
}

func (p *Plugin) Name() string {
	return p.name
}

func (p *Plugin) Call(service string, args any) (any, error) {
	cmd := exec.Command(p.path)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Error executing plugin %s: %v", p.name, err)
	}

	return out, nil
}
