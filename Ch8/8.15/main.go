// This file is a server that lets clients chat with each other.
// This file does not hang on broadcast client writes.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	msg  chan<- string // an outgoing message channel
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	timeout  = 5 * time.Minute
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
				case cli.msg <- msg:
				default:
				}
			}

		case cli := <-entering:
			clients[cli] = true
			cli.msg <- "Clients Presents:"
			for c := range clients {
				cli.msg <- c.name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

func handleConn(conn net.Conn) {
	toutChan := make(chan struct{})
	inputChan := make(chan string)
	out := make(chan string)
	go clientWriter(conn, out)
	in := make(chan string)
	go clientReader(conn, in)

	var who string
	timer := time.NewTimer(timeout)

	out <- "Enter your name:"
	select {
	case name := <-in:
		who = name
		timer.Reset(timeout)
	case <-timer.C:
		out <- "Timeout... Exiting"
		return
	}

	cli := client{out, who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	go func() {
		<-timer.C
		close(toutChan)
	}()

	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			inputChan <- input.Text()
		}
		if input.Err() != nil {
			log.Println("scan error:", input.Err())
		}
		close(toutChan)
	}()
loop:
	for {
		select {
		case text := <-inputChan:
			messages <- who + ": " + text
			timer.Reset(timeout)
		case <-toutChan:
			break loop
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
	if input.Err() != nil {
		log.Println("scan error:", input.Err())
	}
}

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
