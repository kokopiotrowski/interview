package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	RestPort    string `yaml:"restport"`
	GRPCPort    string `yaml:"grpcport"`
	Version     string `yaml:"version"`
	Environment string `yaml:"environment"`
}

var config *Conf

func GetConf() *Conf {
	if config != nil {
		return config
	}
	c := Conf{}

	yamlFile, err := ioutil.ReadFile("configuration.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	config = &c
	return config
}
