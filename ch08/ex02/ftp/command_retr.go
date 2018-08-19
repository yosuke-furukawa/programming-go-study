package ftp

import "fmt"

type RETRCommand struct {
	Operation string
	Argument  string
	*Conn
}

func CreateRETRCommand(operation, argument string, conn *Conn) (Command, error) {
	return RETRCommand{
		operation,
		argument,
		conn,
	}, nil
}

func (command RETRCommand) GetOp() string {
	return command.Operation
}

func (command RETRCommand) GetArgs() []string {
	return []string{}
}

func (command RETRCommand) Eval() (string, error) {
	results, err := command.Conn.cwd.Get(command.Argument)
	if err != nil {
		return "", err
	}
	command.Conn.Write(fmt.Sprintf(RETR_PRERESPONSE, command.Argument, len(results)))
	dataConn := command.Conn.dataConn
	defer dataConn.Close()
	dataConn.Write(results)
	return FILE_TRANSFER_COMPLETE, nil
}
