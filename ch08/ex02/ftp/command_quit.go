package ftp

type QuitCommand struct {
	Operation string
	*Conn
}

func CreateQuitCommand(operation string, conn *Conn) (Command, error) {
	return QuitCommand{
		operation,
		conn,
	}, nil
}

func (command QuitCommand) GetOp() string {
	return command.Operation
}

func (command QuitCommand) GetArgs() []string {
	return []string{}
}

func (command QuitCommand) Eval() (string, error) {
	return CLOSE_REPONSE, nil
}
