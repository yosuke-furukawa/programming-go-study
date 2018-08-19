package ftp

type PasswordCommand struct {
	Operation string
	Arguments []string
}

func CreatePasswordCommand(operation string, arguments []string) (Command, error) {
	return PasswordCommand{
		operation,
		arguments,
	}, nil
}

func (user PasswordCommand) GetOp() string {
	return user.Operation
}

func (user PasswordCommand) GetArgs() []string {
	return user.Arguments
}

func (user PasswordCommand) Eval() (string, error) {
	return USER_LOGGED_IN, nil
}
