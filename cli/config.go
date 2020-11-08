package cli

import (
	_ "gopkg.in/yaml.v2"
)

var ConfigFile string = ""

type Config struct {
	APIVersion string `yaml:"apiVersion"`
	ID         string `yaml:"id"`

	Bind struct {
		IP   string `yaml:"ip"`
		Port uint16 `yaml:"port"`
	} `yaml:"bind"`
}
