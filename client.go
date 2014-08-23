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
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := gob.NewEncoder(conn)
    p := 3
    encoder.Encode(p)

    dec := gob.NewDecoder(conn)
    dec.Decode(&p)
    fmt.Printf("Received : %+v", p);
    
    conn.Close()
    fmt.Println("done");
}