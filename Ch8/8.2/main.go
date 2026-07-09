// This file creates a ftp server
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var dir = flag.String("dir", ".", "start directory to read")
var port = flag.String("port", "8080", "port to listen to")
var path string

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func listFiles(folder string) []string {
	fls, err := os.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	var files []string
	for _, f := range fls {
		files = append(files, f.Name())
	}
	return files
}

func listenFTP(conn net.Conn) {

	// reset starting directory every new connection
	err := os.Chdir(path)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")
		switch cmds[0] {
		case "ls":

			loc := *dir
			if len(cmds) == 2 {
				loc = cmds[1]
			}

			files := listFiles(loc)
			mustCopy(conn, strings.NewReader(fmt.Sprintf("%s\n", strings.Join(files, ","))))
		case "cd":

			loc := cmds[1]
			err := os.Chdir(loc)
			if err != nil {
				mustCopy(conn, strings.NewReader(fmt.Sprintf("%s\n", err.Error())))
			}
		case "get":

			f := cmds[1]
			file, err := os.Open(f)
			if err != nil {
				mustCopy(conn, strings.NewReader(fmt.Sprintf("%s\n", err.Error())))
				continue
			}
			mustCopy(conn, file)
		case "close":
			return
		default:
			help := "invalid command -\nls <file> : list contents (default to -dir flag)\ncd <folder> : change directory\nget <file>: get file content\n" +
				"close: close connection\n"
			mustCopy(conn, strings.NewReader(help))
		}
	}
}

func main() {

	flag.Parse()

	var err error
	path, err = filepath.Abs(*dir)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}
		go listenFTP(conn)
	}
}
