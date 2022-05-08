// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	loc, err1 := time.LoadLocation(os.Getenv("TZ"))
	fmt.Println("user is connected to: " + os.Getenv("TZ"))

	if err1 != nil {
		fmt.Println("using local")
		loc, err1 = time.LoadLocation("Local")
	}
	
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-port" {
		return
	}

	listener, err := net.Listen("tcp", "localhost:"+os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}