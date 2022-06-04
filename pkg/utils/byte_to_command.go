package utils

import (
	"encoding/json"
	"io"
	"strings"
)

type Command struct {
	Command string
	Flags   []string
}

// I need to fix the commands inside quotes
func StringToCommand(commandStr string) *Command {
	var comandSplit = strings.Split(commandStr, " ")
	var resp = new(Command)
	resp.Command = comandSplit[0]
	if len(comandSplit) == 1 {
		return resp
	}
	resp.Flags = comandSplit[1:]
	return resp
}

func (ctx *Command) Send(conn io.Writer) error {
	return json.NewEncoder(conn).Encode(ctx)

}

func (ctx *Command) Get(conn io.Reader) error {
	return json.NewDecoder(conn).Decode(ctx)
}
