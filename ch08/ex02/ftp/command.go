package ftp

import (
	"fmt"
	"log"
	"strings"
)

type Command interface {
	GetOp() string
	GetArgs() []string
	Eval() (string, error)
}

func ParseCommand(message string, conn *Conn) (Command, error) {
	log.Println(message)
	m := strings.Split(message, " ")
	if len(m) == 0 {
		return nil, fmt.Errorf("no command line")
	}

	switch m[0] {
	case "USER":
		return CreateUserCommand(m[0], m[1:])
	case "PASS":
		return CreatePasswordCommand(m[0], m[1:])
	case "SYST":
		return CreateSystCommand(m[0])
	case "PWD":
		return CreatePWDCommand(m[0], conn)
	case "EPRT":
		return CreateEPRTCommand(m[0], m[1], conn)
	case "PORT":
		return CreatePORTCommand(m[0], m[1], conn)
	case "LIST":
		return CreateLISTCommand(m[0], conn)
	case "CWD":
		return CreateCWDCommand(m[0], m[1], conn)
	case "TYPE":
		return CreateTypeCommand(m[0], m[1], conn)
	case "SIZE":
		return CreateSizeCommand(m[0], m[1], conn)
	case "RETR":
		return CreateRETRCommand(m[0], m[1], conn)
	case "MDTM":
		return CreateMDTMCommand(m[0], m[1], conn)
	case "QUIT":
		return CreateQuitCommand(m[0], conn)
	case "FEAT", "EPSV", "PASV":
		return CreateUnimplementedCommand(m[0])
	default:
		return CreateUnimplementedCommand(m[0])

	}
}
