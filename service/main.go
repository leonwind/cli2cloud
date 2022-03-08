package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"service/api"
)

const configFile = "config.yaml"

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

/*
const (
	port  = ":50051"
	dbUrl = "postgres://cli2cloud:123@postgres:5432/cli2cloud"
)
*/

func readConfig() Config {
	var config = Config{}

	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Can't read %s %v\n", configFile, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("Error while unmarshalling", err)
	}
	log.Println(config)

	return config
}

func main() {
	config := readConfig()
	log.Println("DB URL: ", config.Database.Url)
	service, err := api.NewServer(config.Database.Url)
	if err != nil {
		log.Fatal("Cant create server", err)
	}

	if err := service.Start(config.Service.Port); err != nil {
		log.Fatal("Can't start server", err)
	}
}
