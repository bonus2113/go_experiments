package main

import (
	"fmt"
	"net"
	"encoding/gob"
)

type P struct {
	M, N int64
}

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	p := 0
	dec.Decode(&p)

	encoder := gob.NewEncoder(conn)
    encoder.Encode(p * 2)

	fmt.Printf("Received : %+v\n", p);
}

func main() {
	fmt.Println("start");
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
	    // handle error
	}
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
	}
}