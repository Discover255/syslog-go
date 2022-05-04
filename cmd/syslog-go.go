package main

import (
	"fmt"

	"github.com/discover255/syslog-go/pkg/cmd"
	"github.com/discover255/syslog-go/pkg/config"
	"github.com/discover255/syslog-go/pkg/server"
)

// import "github.com/discover255/syslog-go/pkg/server"

func main() {
	config_filename := cmd.GetConfigLocation()
	config.ReadConfig(config_filename)
	fmt.Println(*(config.GlobaConfigMap))
	server.StartListen()
}
