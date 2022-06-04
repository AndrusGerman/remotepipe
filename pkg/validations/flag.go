package validations

func FlagValidClient(args []string) bool {
	return len(args) != 2
}
