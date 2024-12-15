package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Manager struct {
	plugins map[string]*Plugin
}

func NewManager() *Manager {
	return &Manager{
		plugins: make(map[string]*Plugin),
	}
}

func (m *Manager) Discover(path string) error {
	return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		_, err = exec.LookPath(p)
		if err != nil {
			return nil
		}

		plugin, err := NewPlugin(p)
		if err != nil {
			return err
		}

		m.plugins[plugin.Name()] = plugin
		return nil
	})
}

func (m *Manager) Get(name string) (*Plugin, error) {
	plugin, ok := m.plugins[name]
	if !ok {
		return nil, fmt.Errorf("Plugin %s not found", name)
	}

	return plugin, nil
}
