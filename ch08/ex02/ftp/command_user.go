package ftp

type UserCommand struct {
	Operation string
	Arguments []string
}

func CreateUserCommand(operation string, arguments []string) (Command, error) {
	return UserCommand{
		operation,
		arguments,
	}, nil
}

func (user UserCommand) GetOp() string {
	return user.Operation
}

func (user UserCommand) GetArgs() []string {
	return user.Arguments
}

func (user UserCommand) Eval() (string, error) {
	return USER_NAME_OK, nil
}
