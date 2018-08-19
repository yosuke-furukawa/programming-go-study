package ftp

type FeatCommand struct {
	Operation string
}

func CreateFeatCommand(operation string) (Command, error) {
	return FeatCommand{
		operation,
	}, nil
}

func (user FeatCommand) GetOp() string {
	return user.Operation
}

func (user FeatCommand) GetArgs() []string {
	return []string{}
}

func (user FeatCommand) Eval() (string, error) {
	return FEAT_RESPONSE, nil
}
