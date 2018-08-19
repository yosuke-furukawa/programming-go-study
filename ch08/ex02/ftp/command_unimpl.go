package ftp

type UnimplementedCommand struct {
	Operation string
}

func CreateUnimplementedCommand(operation string) (Command, error) {
	return UnimplementedCommand{
		operation,
	}, nil
}

func (command UnimplementedCommand) GetOp() string {
	return command.Operation
}

func (command UnimplementedCommand) GetArgs() []string {
	return []string{}
}

func (command UnimplementedCommand) Eval() (string, error) {
	return NOT_SUPPORTED_PERMANENTLY, nil
}
