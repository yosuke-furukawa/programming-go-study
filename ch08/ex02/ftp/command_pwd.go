package ftp

import (
	"fmt"
	"log"
)

type PWDCommand struct {
	Operation string
	Conn      *Conn
}

func CreatePWDCommand(operation string, conn *Conn) (Command, error) {
	return PWDCommand{
		operation,
		conn,
	}, nil
}

func (pwd PWDCommand) GetOp() string {
	return pwd.Operation
}

func (pwd PWDCommand) GetArgs() []string {
	return []string{}
}

func (pwd PWDCommand) Eval() (string, error) {
	cwd := pwd.Conn.cwd.Pwd()
	log.Println(cwd)
	return fmt.Sprintf(PWD_RESPONSE, cwd), nil
}
