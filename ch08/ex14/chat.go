package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const timeout = 5 * time.Minute

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name string
	ch   chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			list := []string{}
			for cli := range clients {
				list = append(list, cli.name)
			}
			cli.ch <- "current clients: " + strings.Join(list, ",")

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	br := bufio.NewReader(conn)
	line, _, err := br.ReadLine()
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}

	who := string(line)
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{who, ch}

	timer := time.AfterFunc(timeout, func() {
		conn.Close()
	})

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timer.Reset(timeout)
		messages <- who + ": " + input.Text()
	}
	timer.Stop()

	messages <- who + " has left"
	leaving <- client{who, ch}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
