package server

import (
	"fmt"
	"log"
	"net"

	"github.com/discover255/syslog-go/pkg/matcher"
	"github.com/discover255/syslog-go/pkg/mqclient"
)

func StartListen() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 30514})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Syslog server is listening...")
	buffer := make([]byte, 1024*16)
	var raw []byte
	mqclient.ProducerInit()
	defer mqclient.Close()
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("failed to read UDP", err)
		}
		raw = make([]byte, n)
		copy(raw, buffer[:n])
		go matcher.MatchAndSend(raw, remoteAddr)
	}
}
