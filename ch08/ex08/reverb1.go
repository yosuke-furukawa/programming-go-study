package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	ticker := time.NewTicker(10 * time.Second)
	scan := make(chan bool)
	for {
		go func() {
			if input.Scan() {
				scan <- true
			}
		}()
		select {
		case <-ticker.C:
			c.Close()

		case <-scan:
			ticker = time.NewTicker(10 * time.Second)
			echo(c, input.Text(), 1*time.Second)
		}
	}
	c.Close()
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
