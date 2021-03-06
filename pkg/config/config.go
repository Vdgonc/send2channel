package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	// CONFIGFILE - the path of configfile
	CONFIGFILE = "/etc/send2channel/config.yml"
)

// Message - message template from config file
type Message struct {
	Init    string `yaml:"init"`
	Success string `yaml:"success"`
	Fail    string `yaml:"fail"`
}

// Fallback - message fallback from config file
type Fallback struct {
	Init    string `yaml:"init"`
	Success string `yaml:"success"`
	Fail    string `yaml:"fail"`
}

// Color - color templete for messages
type Color struct {
	Init    string `yaml:"init"`
	Success string `yaml:"success"`
	Fail    string `yaml:"fail"`
}

// Config - config type from config.yml
type Config struct {
	Channels []string `yaml:"channels"`
	Title    string   `yaml:"title"`
	Fallback Fallback `yaml:"fallback"`
	Message  Message  `yaml:"message"`
	Color    Color    `yaml:"color"`
	Footer   string   `yaml:"footer"`
}

// New - return config type
func New() *Config {
	var config Config

	yamlFile, _ := ioutil.ReadFile(CONFIGFILE)

	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatal(err)
	}

	return &config

}
