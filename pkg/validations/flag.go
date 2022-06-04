package validations

func FlagIsValidClient(args []string) bool {
	return len(args) == 2
}
