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
func GetCommandByStr(commandStr string) *Command {
	var cm = new(Command)

	// is unix complex
	if IsUnix() && TextContainOne(commandStr, "&&", "'", "|", ";") {
		return cm.get_command_unix_complex(commandStr)
	}
	// basic command
	return cm.get_command_basic(commandStr)
}

func (ctx *Command) get_command_unix_complex(commandStr string) *Command {
	ctx.Command = "sh"
	ctx.Flags = append(ctx.Flags, "-c")
	ctx.Flags = append(ctx.Flags, commandStr)
	return ctx
}

func (ctx *Command) get_command_basic(commandStr string) *Command {
	var comandSplit = strings.Split(commandStr, " ")
	ctx.Command = comandSplit[0]
	if len(comandSplit) == 1 {
		return ctx
	}
	ctx.Flags = comandSplit[1:]
	return ctx
}

func (ctx *Command) Send(conn io.Writer) error {
	return json.NewEncoder(conn).Encode(ctx)
}

func (ctx *Command) Get(conn io.Reader) error {
	return json.NewDecoder(conn).Decode(ctx)
}
