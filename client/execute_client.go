package client

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/AndrusGerman/remotepipe/pkg/validations"
)

func Execute(args []string) {
	var id = fmt.Sprint(time.Now().Unix(), rand.Int31())
	var flagValid = validations.FlagIsValidClient(args)
	if !flagValid {
		log.Println("client: flag not valid")
		os.Exit(1)
	}

	send_comand(id, args[0], args[1])
}
