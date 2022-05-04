package client

import (
	// "bufio"
	"fmt"
	"net"
)

func Send(msg string) {
	conn, err := net.Dial("udp", "127.0.0.1:30514")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println("sending msg...")
	fmt.Fprint(conn, msg)
}
