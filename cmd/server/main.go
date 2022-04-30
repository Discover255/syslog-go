package server

import (
	"fmt"
	"log"
	"net"
)

func StartListen() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 30514})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Syslog server is listening...")
	buffer := make([]byte, 1024*4)
	var raw []byte
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println("failed to read UDP", err)
		}
		raw = make([]byte, n)
		copy(raw, buffer[:n])
		go func() {
			fmt.Println(remoteAddr)
			fmt.Println(string(raw[:]))
		}()
	}
}
