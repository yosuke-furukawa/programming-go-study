package ftp

import "fmt"

type SizeCommand struct {
	Operation string
	Argument  string
	*Conn
}

func CreateSizeCommand(operation, argument string, conn *Conn) (Command, error) {
	return SizeCommand{
		operation,
		argument,
		conn,
	}, nil
}

func (command SizeCommand) GetOp() string {
	return command.Operation
}

func (command SizeCommand) GetArgs() []string {
	return []string{}
}

func (command SizeCommand) Eval() (string, error) {
	finfo, err := command.Conn.cwd.Stat(command.Argument)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(SIZE_RESPONSE, finfo.Size()), nil
}
