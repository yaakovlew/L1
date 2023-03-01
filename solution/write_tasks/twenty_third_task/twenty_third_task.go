package twenty_third_task

func TwentyThird(s []string, i int) []string {
	//добавление срезов с 0 по i и c i+1 по последний
	newString := append(s[:i], s[i+1:]...)
	return newString
}
