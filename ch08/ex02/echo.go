package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echoHandler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	writer.WriteString("220 FTP server (nodeftpd) ready\n")
	writer.Flush()
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(line))
	io.Copy(conn, conn)
}

func main() {
	sock, e := net.Listen("tcp", "localhost:5000")
	if e != nil {
		log.Fatal(e)
	}
	for {
		conn, e := sock.Accept()
		if e != nil {
			log.Fatal(e)
		}
		go echoHandler(conn)
	}
}
