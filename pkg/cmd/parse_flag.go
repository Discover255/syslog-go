package cmd

import (
	"flag"
	"fmt"
)

func GetConfigLocation() *string {
	yconfig := flag.String("f", "/etc/syslog-go/sample.yaml", "yaml config for syslog-go")
	flag.Parse()
	fmt.Println("yaml config location:", *yconfig)
	return yconfig
}
