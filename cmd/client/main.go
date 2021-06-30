package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/khunafin/magazine/internal/producer"
)

func main() {
	fmt.Println("Start client...")
	serverAddr := flag.String("server", ":7000", "server addr")
	count := flag.Int("count", 5, "message count")
	flag.Parse()
	conn, err := net.DialTimeout("tcp", *serverAddr, time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	p := producer.New(conn, *count)
	p.Produce()

	fmt.Println("Client stopped")
}
