package ftp

import (
	"fmt"
)

type MDTMCommand struct {
	Operation string
	Argument  string
	*Conn
}

func CreateMDTMCommand(operation, argument string, conn *Conn) (Command, error) {
	return MDTMCommand{
		operation,
		argument,
		conn,
	}, nil
}

func (command MDTMCommand) GetOp() string {
	return command.Operation
}

func (command MDTMCommand) GetArgs() []string {
	return []string{}
}

func (command MDTMCommand) Eval() (string, error) {
	finfo, err := command.Conn.cwd.Stat(command.Argument)
	if err != nil {
		return "", err
	}
	mdtm := finfo.ModTime().Format("20060102150405")
	return fmt.Sprintf(MDTM_RESPONSE, mdtm), nil
}
