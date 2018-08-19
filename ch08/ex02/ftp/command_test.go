package ftp

import (
	"testing"
)

func TestParseUserCommand(t *testing.T) {
	tests := []struct {
		message   string
		operation string
		arguments []string
	}{
		{
			"USER yosuke",
			"USER",
			[]string{"yosuke"},
		},
	}

	for _, test := range tests {
		command, err := ParseCommand(test.message)
		if err != nil {
			t.Errorf("Error! %s", err)
		}
		if command.GetOp() != test.operation {
			t.Errorf("Error! operation %s is different expected %s", command.GetOp(), test.operation)
		}

		for i, arg := range command.GetArgs() {
			if test.arguments[i] != arg {
				t.Errorf("Error! argument %s is different expected %s", arg, test.arguments[i])
			}
		}
	}
}
