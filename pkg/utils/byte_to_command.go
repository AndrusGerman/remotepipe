package utils

import "strings"

type Command struct {
	Command string
	Flags   []string
}

// I need to fix the commands inside quotes
func ByteToCommand(commandByte []byte) *Command {
	var comandSplit = strings.Split(string(commandByte), " ")
	var resp = new(Command)
	resp.Command = comandSplit[0]
	if len(comandSplit) == 1 {
		return resp
	}
	resp.Flags = comandSplit[1:]
	return resp
}
