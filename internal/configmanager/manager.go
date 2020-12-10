package configmanager

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// The config manager handles interactions between the CLI and the
// actual backend storage for the config files

type Authority struct {
	Hostname    string `yaml:"hostname"`
	ID          string `yaml:"id"`
	Secret      string `yaml:"secret"`
	Certificate string `yaml:"certificate"`
	ConfigPath  string `yaml:"-"`
}

type Manager struct {
	root string
}

func New(root string) *Manager {
	return &Manager{
		root: root,
	}
}

func (cm *Manager) Authority(name string) (*Authority, error) {
	path := cm.pathForName(name)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil
	}

	authority := &Authority{ConfigPath: path}
	err = yaml.Unmarshal(file, authority)

	return authority, err
}

func (cm *Manager) SaveAuthority(name string, authority *Authority) error {
	path := cm.pathForName(name)
	os.MkdirAll(filepath.Dir(path), 0700)

	yaml, err := yaml.Marshal(authority)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, yaml, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (cm *Manager) pathForName(name string) string {
	return fmt.Sprintf("/%s/authorities/%s.yaml", cm.root, name)
}
