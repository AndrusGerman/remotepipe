package client

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AndrusGerman/remotepipe/pkg/validations"
)

func Execute(args []string) {
	var id = fmt.Sprint(time.Now().Unix())
	var flagValid = validations.FlagIsValidClient(args)
	if !flagValid {
		log.Println("client: flag not valid")
		os.Exit(1)
	}

	send_comand(id, args[0], args[1])
}
