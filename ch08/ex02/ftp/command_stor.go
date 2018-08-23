package ftp

import (
	"io/ioutil"
)

type STORCommand struct {
	Operation string
	Argument  string
	*Conn
}

func CreateSTORCommand(operation, argument string, conn *Conn) (Command, error) {
	return STORCommand{
		operation,
		argument,
		conn,
	}, nil
}

func (command STORCommand) GetOp() string {
	return command.Operation
}

func (command STORCommand) GetArgs() []string {
	return []string{}
}

func (command STORCommand) Eval() (string, error) {
	command.Conn.Write(STOR_PRERESPONSE)
	dataConn := command.Conn.dataConn
	defer dataConn.Close()
	content, err := ioutil.ReadAll(dataConn)
	if err != nil {
		return "", err
	}
	command.Conn.cwd.Put(command.Argument, content)

	return FILE_TRANSFER_COMPLETE, nil
}
