package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type clock struct {
	location string
	address  string
}

func main() {
	args := os.Args[1:]
	clocks := []clock{}
	for _, arg := range args {
		a := strings.Split(arg, "=")
		location := a[0]
		address := a[1]
		clocks = append(clocks, clock{
			location,
			address,
		})
	}

	time := make(chan string)
	for _, clock := range clocks {
		conn, err := net.Dial("tcp", clock.address)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(conn)
		go displayTime(clock.location, scanner, time)
	}

	for {
		select {
		case t := <-time:
			fmt.Println(t)
		}

	}
}

func displayTime(location string, scanner *bufio.Scanner, time chan string) {
	for scanner.Scan() {
		time <- fmt.Sprintf("%s: %s", location, scanner.Text())
	}
}
