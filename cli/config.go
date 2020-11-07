package cli

import (
	_ "gopkg.in/yaml.v2"
)

var ConfigFile string = ""

type Config struct {
	APIVersion string `yaml:"apiVersion"`
	ID         string `yaml:"id"`
}
