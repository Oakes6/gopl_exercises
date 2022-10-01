// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		_, err = io.WriteString(c, fmt.Sprintf("FROM %s", timezone))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := os.Args[1]
	fmt.Printf("PORT: %s", port)
	timezone := os.Args[2]
	fmt.Printf(" Timezone: %s", timezone)
	// check if port is int
	// if port.(int) != false {
	// 	fmt.Println("ERROR")
	// }
	address := fmt.Sprintf("localhost:%v", port)
	fmt.Println("LISTENING ON " + address)
	listener, err := net.Listen("ftp", address)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, timezone) // handle connections concurrently
	}
	//!-
}
