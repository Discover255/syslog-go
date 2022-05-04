package server

import (
	"testing"
	"time"

	"github.com/discover255/syslog-go/internal/client"
)

func TestUDP(t *testing.T) {
	go StartListen()
	time.Sleep(time.Second * 2)
	client.Send("你好，世界！")
	time.Sleep(time.Second * 2)
}

func TestStartListen(t *testing.T) {
	StartListen()
}
