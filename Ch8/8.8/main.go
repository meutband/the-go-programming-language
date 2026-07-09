// This file contains an echo server that timeouts after
// 10 seconds
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const timeout = 10 * time.Second

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {

	inputs := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			inputs <- input.Text()
		}
		if input.Err() != nil {
			log.Println("scan error:", input.Err())
		}
	}()

	timer := time.NewTimer(timeout)
	for {
		select {
		case in := <-inputs:
			timer.Reset(timeout)
			echo(c, in, 1*time.Second)
		case <-timer.C:
			return
		}
	}
}

func main() {

	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}
