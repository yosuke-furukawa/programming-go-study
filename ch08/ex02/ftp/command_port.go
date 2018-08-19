package ftp

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type PORTCommand struct {
	Operation string
	Address   string
	Conn      *Conn
}

func CreatePORTCommand(operation string, argument string, conn *Conn) (Command, error) {
	s := strings.Split(argument, ",")
	if len(s) != 6 {
		return nil, fmt.Errorf("parse error")
	}

	p1, err := strconv.Atoi(s[4])
	if err != nil {
		return nil, err
	}

	p2, err := strconv.Atoi(s[5])
	if err != nil {
		return nil, err
	}

	port := p1*256 + p2
	address := fmt.Sprintf("%s.%s.%s.%s:%d", s[0], s[1], s[2], s[3], port)
	dataConn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	conn.DataConn(dataConn)
	return PORTCommand{
		operation,
		address,
		conn,
	}, nil
}

func (command PORTCommand) GetOp() string {
	return command.Operation
}

func (command PORTCommand) GetArgs() []string {
	return []string{command.Address}
}

func (command PORTCommand) Eval() (string, error) {
	return COMMAND_OK, nil
}
