package ftp

type CWDCommand struct {
	Operation string
	Argument  string
	Conn      *Conn
}

func CreateCWDCommand(operation, argument string, conn *Conn) (Command, error) {
	return CWDCommand{
		operation,
		argument,
		conn,
	}, nil
}

func (command CWDCommand) GetOp() string {
	return command.Operation
}

func (command CWDCommand) GetArgs() []string {
	return []string{command.Argument}
}

func (command CWDCommand) Eval() (string, error) {
	command.Conn.cwd.Chdir(command.Argument)
	return CWD_RESPONSE, nil
}
