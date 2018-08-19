package ftp

type SystCommand struct {
	Operation string
}

func CreateSystCommand(operation string) (Command, error) {
	return SystCommand{
		operation,
	}, nil
}

func (user SystCommand) GetOp() string {
	return user.Operation
}

func (user SystCommand) GetArgs() []string {
	return []string{}
}

func (user SystCommand) Eval() (string, error) {
	return SYST_RESPONSE, nil
}
