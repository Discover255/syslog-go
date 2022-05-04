package config

import (
	"fmt"
	"testing"
)

type User struct {
	Name       string
	Occupation string
}

type User2 struct {
	Users []string
}

func TestReadConfig(t *testing.T) {
	// yfile, _ := ioutil.ReadFile("./sample.yaml")
	// data := ConfigMap{}
	// // data := make(map[string]interface{})
	// yaml.Unmarshal(yfile, &data)
	// fmt.Println(data)
	// fmt.Println(*(data.Rules[0].MQs[0]))
	filename := "./sample.yaml"
	ReadConfig(&filename)
	fmt.Println(*GlobaConfigMap)
}
