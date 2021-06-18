package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	System struct {
		HTTPRequestTimeout int    `yaml:"http_request_timeout" json:"http_request_timeout"`
		CommandTimeout     int    `yaml:"command_timeout" json:"command_timeout"`
		Version            string `yaml:"version" json:"version"`
		Log                struct {
			Path         string `yaml:"path" json:"path"`
			Level        string `yaml:"level" json:"level"`
			MaxAge       int    `yaml:"max_age" json:"max_age"`
			RotationTime int    `yaml:"rotation_time" json:"rotation_time"`
		} `yaml:"log" json:"log"`
	} `yaml:"system" json:"system"`
	Database struct {
		Host     string `yaml:"host" json:"host"`
		Port     int    `yaml:"post" json:"post"`
		Username string `yaml:"username" json:"username"`
		Password string `yaml:"password" json:"password"`
	} `yaml:"database" json:"database"`
}

var configuration *Configuration

func LoadConfiguration() error {
	fmt.Println("读取文件")
	current_path, err := os.Getwd()
	if err != nil {
		log.Fatal("load config failed")
	}
	config_path := path.Join(current_path, "config", "config.yaml")
	data, err := ioutil.ReadFile(config_path)
	if err != nil {
		return err
	}

	// var config Configuration
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return err
	}
	return nil
}

func GetConfiguration() *Configuration {
	return configuration
}
