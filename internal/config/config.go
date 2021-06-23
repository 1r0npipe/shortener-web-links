package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var (
	ErrDecodeYAML = errors.New("can't decode yaml file")
)

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"server"`
	DB struct {
		Address  string `yaml:"dbAddress"`
		Port     string `yaml:"dbPort"`
		Username string `yaml:"dbUsername"`
		Password string `yaml:"dbpasswor"`
	} `yaml:"db"`
}

func ReadNewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(file, config); err != nil {
		return nil, ErrDecodeYAML
	}
	return config, nil
}
