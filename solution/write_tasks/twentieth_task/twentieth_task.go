package twentieth_task

import (
	"strings"
)

func Twentieth(s string) string {
	//разделение строки на слова
	newString := strings.Split(s, " ")
	res := ""
	mas := []string{}
	//добавление слов с конца в новую строку
	for i := len(newString) - 1; i >= 0; i-- {
		mas = append(mas, newString[i])
	}
	//массив строк -> в строку
	res = strings.Join(mas, " ")
	return res
}
