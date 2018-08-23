package ftp

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	host string
	port int
}

func FtpHandler(c net.Conn) {
	log.Println("start!!")
	defer c.Close()
	conn := NewConn(c)

	conn.Write(SERVER_HELLO)
	msg := make(chan string)
	errorMsg := make(chan error)
	go conn.Read(msg, errorMsg)
	for {
		select {
		case message := <-msg:
			command, err := ParseCommand(message, conn)
			if err != nil {
				log.Println(err)
				conn.Write(fmt.Sprintf("%s", err))
			}
			result, err := command.Eval()
			if err != nil {
				conn.Write(fmt.Sprintf("%s", err))
			}
			conn.Write(result)
			if result == CLOSE_REPONSE {
				conn.conn.Close()
				return
			}
			go conn.Read(msg, errorMsg)
		case err := <-errorMsg:
			conn.Write(fmt.Sprintf(NOT_SUPPORTED_HASDETAIL, err.Error()))
			conn.conn.Close()
		}
	}
}
