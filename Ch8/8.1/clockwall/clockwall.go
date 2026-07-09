// This file listens to multiple clock servers concurrently
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

var clockTimes map[string]string

func getTime(place, ip string) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := bufio.NewReader(conn)
	for {

		blob, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading", err)
			continue
		}

		clockTimes[place] = string(blob)
		// printTime()
	}
}

func printTime() {
	clearScreen()
	const format string = "%v\t%v\n"

	names := make([]string, 0, len(clockTimes))
	for name := range clockTimes { // get names as sort key
		names = append(names, name)
	}
	sort.Strings(names)

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Place", "Time")
	fmt.Fprintf(tw, format, "-----", "----")
	for _, n := range names {
		fmt.Fprintf(tw, format, n, clockTimes[n])
	}
	tw.Flush()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	if len(os.Args) <= 1 {
		log.Fatal("usage: clockwall Place=IP:Port ...")
	}

	clockTimes = make(map[string]string)

	counts := 0
	for _, clockIP := range os.Args[1:] {

		clock := strings.Split(clockIP, "=")
		if len(clock) != 2 {
			fmt.Printf("cannot handle arg: %s\n", clock)
			continue
		}

		clockTimes[clock[0]] = "TBD"

		go getTime(clock[0], clock[1])
		counts++
	}

	for {
		time.Sleep(time.Second)
		printTime()
	}

}
