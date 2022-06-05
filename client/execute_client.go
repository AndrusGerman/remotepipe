package client

import (
	"log"
	"os"

	"github.com/AndrusGerman/remotepipe/pkg/validations"
)

func Execute(args []string) {
	var flagValid = validations.FlagIsValidClient(args)
	if !flagValid {
		log.Println("client: flag not valid")
		os.Exit(1)
	}
	send_comand(args[0], args[1])
}
