package main

import (
	"log"
	"net"

	"github.com/yosuke-furukawa/programming-go-study/ch08/ex02/ftp"
)

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
		go ftp.FtpHandler(conn)
	}
}
