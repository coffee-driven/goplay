package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Message struct {
	Author string
	Msg    string
	Ts     time.Time
}

func NewMessage(author string, message string) *Message {
	return &Message{
		Author: author,
		Msg:    message,
		Ts:     time.Now(),
	}
}

type Sender struct {
	PeerAddress net.IPAddr
	PeerPort    int
}

func NewSender(ip net.IP, port int) *Sender {
	return &Sender{
		PeerAddress: ip,
		PeerPort:    port,
	}
}

func (s *Sender) RunService(c chan []byte) {
	conn, err := net.DialIP("tcp", s.PeerAddress, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		msg := <-c
		if _, err := conn.Write(msg); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

}

func main() {
	fmt.Println("running")
}
