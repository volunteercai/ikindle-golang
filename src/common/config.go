package common

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var c *Config = nil

type Config struct {
	Web struct {
		Port int `yaml:"port"`
	} `yaml:"web"`
	Db struct {
		Url      string `yaml:"url"`
		Db       string `yaml:"db"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"db"`
}

func GetConfig() *Config {
	if c == nil {
		panic("配置未初始化")
	}
	return c
}

func LoadConfig(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}

	c = config
}
