package client

import (
	"fmt"
	"net"
	"testing"
)

func TestSend(t *testing.T) {
	conn, err := net.Dial("udp", "127.0.0.1:30514")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("sending msg...")
	fmt.Fprint(conn, "hello")
	// buffer := make([]byte, 4*1024)
	// _, err = bufio.NewReader(conn).Read(buffer)
}
