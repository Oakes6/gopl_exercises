// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//!+broadcaster
type client struct {
	channel chan<- string // an outgoing message channel
	name    string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.channel <- msg:
				default:
					//skip if client channel is blocked
				}
			}

		case cli := <-entering:
			clients[cli] = true
			var connectedClients strings.Builder
			connectedClients.WriteString("Clients connected: ")
			for cli := range clients {
				connectedClients.WriteString(fmt.Sprintf("\n- %s", cli.name))
			}
			cli.channel <- connectedClients.String()

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.channel)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	// who := conn.RemoteAddr().String()
	fmt.Fprint(conn, "What is the name of this client: ")
	whoInput := bufio.NewScanner(conn)
	whoInput.Scan()
	who := whoInput.Text()
	newClient := client{
		channel: ch,
		name:    who,
	}

	newClient.channel <- "You are " + who
	messages <- who + " has arrived"
	entering <- newClient

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- newClient.name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- newClient
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
