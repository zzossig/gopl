/*
	Modify `clock2` to accept a port number, and write a program, `clockwall`, that acts as a client of several clock servers at once,
	reading the times from each one and displaying the results in a table, akin to the wall of clocks seen in some business offices.
	If you have access to `geographically` distributed computers, run instances remotely;
	otherwise run local instances on different ports with fake time zones.

	``` text
			$ TZ=US/Eastern    ./clock2 -port 8010 &
			$ TZ=Asia/Tokyo    ./clock2 -port 8020 &
			$ TZ=Europe/London ./clock2 -port 8030 &
			$ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
	```
*/

// usage: go run main.go NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func handleConn(c net.Conn, name string) {
	defer c.Close()

	var location *time.Location
	var err error

	switch name {
	case "NewYork":
		location, err = time.LoadLocation("US/Eastern")
	case "London":
		location, err = time.LoadLocation("Europe/London")
	case "Tokyo":
		location, err = time.LoadLocation("Asia/Tokyo")
	}

	if err != nil {
		panic(err)
	}

	for {
		_, err := io.WriteString(c, fmt.Sprintf("[%s]%s", name, time.Now().In(location).Format("15:04:05\n")))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan struct{})
	for _, arg := range os.Args[1:] {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			log.Fatal("usage: go run main.go NewYork=localhost:8010 ...")
		}

		go listen(kv[0], kv[1])
	}
	<-ch
}

func listen(name, port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, name) // handle connections concurrently
	}
}
