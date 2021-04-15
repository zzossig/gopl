/*
	Implement a concurrent `File Transfer Protocol (FTP)` server.
	The server should interpret commands from each client such as `cd` to chagne directore,
	`ls` to list a directory,
	`get` to send the contents of a file,
	and `close` to close the connection.
	You can use the standard ftp command as the client, or write your own.
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var fp = "."

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	ch := make(chan string)

	go clientWriter(conn, ch)
	ch <- ">>"

loop:
	for scanner.Scan() {
		if scanner.Err() != nil {
			ch <- scanner.Err().Error()
			break
		}

		var sb strings.Builder
		var files []string

		fields := strings.Split(scanner.Text(), " ")
		if len(fields) == 0 {
			ch <- "no cmd found"
			continue loop
		}

		cmd := fields[0]
		sb.Reset()

		switch cmd {
		case "get":
			if len(fields) < 2 {
				ch <- "no arg found"
				continue loop
			}
			arg := fields[1]
			fp = filepath.Join(fp, arg)
			err := transferFile(conn, fp, ch)

			if err != nil {
				ch <- err.Error()
				continue loop
			}
		case "ls":
			err := filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
				files = append(files, path)
				return err
			})
			if err != nil {
				ch <- err.Error()
				continue loop
			}

			printFiles(files, &sb, ch)
		case "cd":
			if len(fields) < 2 {
				ch <- "no arg found"
				continue loop
			}
			arg := fields[1]
			fp = filepath.Join(fp, arg)

			err := filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
				files = append(files, path)
				return err
			})
			if err != nil {
				ch <- err.Error()
				continue loop
			}

			printFiles(files, &sb, ch)
		case "close":
			ch <- "closing connection..."
			return
		default:
			ch <- fmt.Sprintf("unsupported command: %s\n", cmd)
		}

		ch <- ">>"
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprint(conn, msg)
	}
}

func printFiles(files []string, sb *strings.Builder, ch chan<- string) {
	for _, file := range files {
		sb.WriteString(file)
		sb.WriteString("\n")
	}
	ch <- sb.String()
}

func transferFile(conn net.Conn, fp string, ch chan<- string) error {
	ch <- "starting transfer file"
	file, err := os.Open(fp)
	if err != nil {
		return fmt.Errorf("cannot read file: %s", fp)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("cannot read file: %s", fp)
	}
	fileSize := fmt.Sprintf("%d", fileInfo.Size())
	fileName := fileInfo.Name()

	ch <- "================\n"
	ch <- fmt.Sprintf("file size: %s\n", fileSize)
	ch <- fmt.Sprintf("file name: %s\n", fileName)
	ch <- "================\n"

	sendBuffer := make([]byte, 1024)
	fmt.Println("sending...")
	ch <- "signal"

	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			fmt.Println("end of file.")
			break
		}

		conn.Write(sendBuffer)
	}

	return nil
}
