package twentieth_task

import (
	"strings"
)

func Twentieth(s string) string {
	newString := strings.Split(s, " ")
	res := ""
	mas := []string{}
	for i := len(newString) - 1; i >= 0; i-- {
		mas = append(mas, newString[i])
	}
	res = strings.Join(mas, " ")
	return res
}
