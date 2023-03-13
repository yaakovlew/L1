package fifteenth_task

func createHugeString(a int) (s string) {
	s = ""
	//....
	return
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//копируем даннеы, чтобы не ссылались на одну и ту же строку
	justString = string(v[:100])
}

func Fifthen() {
	someFunc()
}
