package twenty_third_task

func TwentyThird(s []string, i int) []string {
	newString := append(s[:i], s[i+1:]...)
	return newString
}
