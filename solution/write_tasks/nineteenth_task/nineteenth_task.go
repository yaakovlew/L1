package nineteenth_task

func Nineteenth(s string) string {
	//Представление в виде массива rune
	runes := []rune(s)
	//В цикле меняем элементы: первый = последний, последний = первый
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
