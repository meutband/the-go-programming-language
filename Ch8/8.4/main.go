// This package counts the goroutines from stimulating an echo
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(conn net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(conn)

	for input.Scan() {
		wg.Add(1)
		go func() {
			echo(conn, input.Text(), 1*time.Second)
			wg.Done()
		}()
	}

	wg.Wait()
	if c, ok := conn.(*net.TCPConn); ok {
		c.CloseWrite()
	} else {
		conn.Close()
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
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}
