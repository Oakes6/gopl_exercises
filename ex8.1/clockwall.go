package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	for _, v := range os.Args[1:] {
		go connect(v)
	}
	for {
		time.Sleep(time.Minute)
	}
}

func connect(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

// from Chapter 8 : netcat1
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
