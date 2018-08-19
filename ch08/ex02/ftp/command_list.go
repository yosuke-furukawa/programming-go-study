package ftp

import (
	"bufio"
	"io/ioutil"
	"strings"
)

type LISTCommand struct {
	Operation string
	Conn      *Conn
}

func CreateLISTCommand(operation string, conn *Conn) (Command, error) {
	return LISTCommand{
		operation,
		conn,
	}, nil
}

func (command LISTCommand) GetOp() string {
	return command.Operation
}

func (command LISTCommand) GetArgs() []string {
	return []string{}
}

func (command LISTCommand) Eval() (string, error) {
	command.Conn.Write(LIST_PRERESPONSE)
	pwd := command.Conn.cwd.Pwd()
	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		return "", nil
	}
	results := []string{}
	for _, file := range files {
		results = append(results, file.Name())
	}
	f := strings.Join(results, "\r\n")
	f += "\r\n"
	dataConn := command.Conn.dataConn
	defer dataConn.Close()
	writer := bufio.NewWriter(dataConn)
	writer.WriteString(f)
	writer.Flush()
	return FILE_TRANSFER_COMPLETE, nil
}
