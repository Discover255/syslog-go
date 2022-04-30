package client

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:30514")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sending msg...")
	fmt.Fprint(conn, "hello")
	buffer := make([]byte, 4*1024)
	_, err = bufio.NewReader(conn).Read(buffer)
	conn.Close()
}
