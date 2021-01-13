package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	for _, v := range os.Args {
		go connect(v)
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
