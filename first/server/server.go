package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ln, err := net.Listen("tcp", "localhost:8069")

	if err != nil {
		log.Fatal(err)
	}

	for {

		conn, err := ln.Accept()

		if err != nil {
			continue
		}
		wg.Add(1)
		go handler(conn)

	}
}

func handler(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, "connected to server\n")
		if err != nil {
			break
		}
		time.Sleep(time.Second * 5)

	}

}
