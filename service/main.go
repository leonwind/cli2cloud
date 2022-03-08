package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"service/api"
)

const configFile = "config.yaml"

type Config struct {
	Service  *ServiceConfig  `yaml:"service"`
	Database *DatabaseConfig `yaml:"database"`
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

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("Error while unmarshalling", err)
	}

	return config
}

func main() {
	config := readConfig()
	dbUrl := (*config.Database).Url
	port := (*config.Service).Port
	log.Println(dbUrl, port)

	service, err := api.NewServer(dbUrl)
	if err != nil {
		log.Fatal("Cant create server", err)
	}

	if err := service.Start(port); err != nil {
		log.Fatal("Can't start server", err)
	}
}
