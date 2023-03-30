package ai

import (
	"fmt"
	"net"
	"os"
)

var (
	addr string
)

func Init() {
	fmt.Println("[AI] Init")
	addr = os.Getenv("AI_PORT")
}

func SendMessage(msg string) (string, error) {
	// connect to this socket
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// send to socket
	conn.Write([]byte(msg))
	fmt.Println([]byte(msg))

	// listen for reply
	bs := make([]byte, 1024)
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	}
}
