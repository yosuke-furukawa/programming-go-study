package ftp

import (
	"fmt"
	"net"
	"strings"
)

type EPRTCommand struct {
	Operation string
	Address   string
	Conn      *Conn
}

func CreateEPRTCommand(operation string, argument string, conn *Conn) (Command, error) {
	s := strings.Split(argument[1:len(argument)-1], "|")
	if len(s) < 3 {
		return nil, fmt.Errorf("parse error")
	}

	address := s[1] + ":" + s[2]
	dataConn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	conn.DataConn(dataConn)
	return EPRTCommand{
		operation,
		address,
		conn,
	}, nil
}

func (command EPRTCommand) GetOp() string {
	return command.Operation
}

func (command EPRTCommand) GetArgs() []string {
	return []string{command.Address}
}

func (command EPRTCommand) Eval() (string, error) {
	return COMMAND_OK, nil
}
