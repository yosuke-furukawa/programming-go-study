package ftp

type TypeCommand struct {
	Operation string
	Argument  string
	*Conn
}

func CreateTypeCommand(operation, argument string, conn *Conn) (Command, error) {
	conn.TransferType(argument)
	return TypeCommand{
		operation,
		argument,
		conn,
	}, nil
}

func (command TypeCommand) GetOp() string {
	return command.Operation
}

func (command TypeCommand) GetArgs() []string {
	return []string{}
}

func (command TypeCommand) Eval() (string, error) {
	return COMMAND_OK, nil
}
