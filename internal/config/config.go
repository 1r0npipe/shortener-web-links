package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	ErrFileRead   = errors.New("can't read config file")
	ErrDecodeYAML = errors.New("can't decode yaml file")
)

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"server"`
}

func ReadNewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, ErrFileRead
	}
	if err := yaml.Unmarshal(file, config); err != nil {
		return nil, ErrDecodeYAML
	}
	return config, nil
}
