package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8069")

	if err != nil {
		fmt.Println("your server is down dude ^^")
	}

	defer conn.Close()

	handler(conn)
}

func handler(conn net.Conn) {
	// print("---")
	// bs, err := io.ReadAll(conn)
	// fmt.Println(bs)

	_, err := io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(bs))
}
