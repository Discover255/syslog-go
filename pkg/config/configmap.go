package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type MQConfig struct {
	Name  string
	Kind  string
	Hosts []string
	Topic string
}

type RuleConfig struct {
	Name     string
	Hosts    []string
	Encoding string
	Keywords []string
	MQs      []string
}

type DefaultSettings struct {
	Encoding string
	MQ       string
}

type ConfigMap struct {
	MQs     []MQConfig
	Rules   []RuleConfig
	Default DefaultSettings
}

var GlobaConfigMap *ConfigMap

func ReadConfig(filename *string) {
	content, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("Read Config File Error %v", err)
	}
	yaml.Unmarshal(content, &GlobaConfigMap)
}
