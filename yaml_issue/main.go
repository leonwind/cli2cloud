package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

const configFile = "config.yaml"

var ymlString = "config:\n  service:\n    port: \":50051\"\n\n  database:\n    url: \"postgres://cli2cloud:123@postgres:5432/cli2cloud\""

type Config struct {
	Service  ServiceConfig  `yaml:"service"`
	Database DatabaseConfig `yaml:"database"`
}

type ServiceConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Url string `yaml:"url"`
}

func readConfig() Config {
	var config = Config{}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Can't read %s %v\n", configFile, err)
	}
	log.Println(yamlFile)
	yamlFile = []byte(ymlString)
	log.Println(yamlFile)

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("Error while unmarshalling", err)
	}
	log.Println(config)

	return config
}

func main() {
	_ = readConfig()
}
