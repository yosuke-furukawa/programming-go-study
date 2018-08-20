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

var wg sync.WaitGroup

func echo(c net.Conn, shout string, delay time.Duration, wait chan struct{}) {
	wg.Add(1)
	wait <- struct{}{}
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		wait := make(chan struct{})
		go echo(c, input.Text(), 1*time.Second, wait)
		go func(c net.Conn, wait chan struct{}) {
			<-wait
			wg.Wait()
			tcpConn := c.(*net.TCPConn)
			tcpConn.CloseWrite()
		}(c, wait)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:8000"))
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
