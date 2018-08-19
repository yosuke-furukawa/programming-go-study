package ftp

import (
	"bufio"
	"log"
	"net"
)

type Conn struct {
	conn         net.Conn
	dataConn     net.Conn
	reader       *bufio.Reader
	writer       *bufio.Writer
	cwd          *Cwd
	transferType string
}

func NewConn(conn net.Conn) *Conn {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	cwd := NewCwd()
	return &Conn{
		conn,
		nil,
		reader,
		writer,
		cwd,
		"",
	}
}

func (conn *Conn) DataConn(dataConn net.Conn) {
	conn.dataConn = dataConn
}

func (conn *Conn) TransferType(transferType string) {
	conn.transferType = transferType
}

func (conn *Conn) Read(msg chan string, errorMsg chan error) {
	line, _, err := conn.reader.ReadLine()
	if err != nil {
		errorMsg <- err
		return
	}
	msg <- string(line)
}

func (conn *Conn) Write(message string) {
	log.Println(message)
	conn.writer.WriteString(message)
	conn.writer.Flush()
}
