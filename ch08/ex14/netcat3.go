package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	tcpConn := conn.(*net.TCPConn)

	fmt.Fprint(os.Stdout, "name: ")
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(tcpConn, string(line))
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	tcpConn.CloseWrite()

	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
