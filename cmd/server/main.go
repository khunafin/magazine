package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/khunafin/magazine/internal/consumer"
)

const (
	walFileName = "wal.dat"
)

func main() {
	listenAddr := flag.String("listen", ":7000", "listen addr")
	fmt.Println("Start server...")
	listener, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		panic(err)
	}

	walFile, err := os.OpenFile(walFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	defer walFile.Close()

	if err != nil {
		panic(err)
	}
	hndlr := consumer.New()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go func() {
			defer walFile.Sync()
			fmt.Println("Client connected")
			hndlr.Handle(conn, walFile)
			fmt.Println("Client disconnected")
		}()
	}
}
