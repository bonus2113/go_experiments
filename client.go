package main

import (
    "fmt"
    "log"
    "net"
    "encoding/gob"
)

type P struct {
    M, N int64
}

func main() {
    fmt.Println("start client");
    conn, err := net.Dial("tcp", "54.200.222.128:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := gob.NewEncoder(conn)
    p := 3
    encoder.Encode(p)

    dec := gob.NewDecoder(conn)
    dec.Decode(&p)
    fmt.Printf("Received : %+v\n", p);

    conn.Close()
    fmt.Println("done");
}