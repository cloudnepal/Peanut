// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package definition

import (
	"gopkg.in/yaml.v2"
)

// DockerCompose type
type DockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Networks map[string]string  `yaml:"networks,omitempty"`
	Volumes  map[string]string  `yaml:"volumes,omitempty"`
}

// Service type
type Service struct {
	Image     string   `yaml:"image"`
	Volumes   []string `yaml:"volumes,omitempty"`
	Networks  []string `yaml:"networks,omitempty"`
	DependsOn []string `yaml:"depends_on,omitempty"`
	Ports     []string `yaml:"ports,omitempty"`
	Restart   string   `yaml:"restart,omitempty"`
}

// ToString converts object to a string
func (d *DockerCompose) ToString() (string, error) {
	o, err := yaml.Marshal(&d)

	if err != nil {
		return "", err
	}

	return string(o), nil
}
