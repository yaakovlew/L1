package first_task

import "fmt"

//Пример структуры Action
type Action struct {
	Age    int
	Gender bool
	//...
}

//Метод для структуры Action
func (act Action) Run() {
	fmt.Println("He/she run")
}

//Пример структуры Human, со встроенной структурой Action и т.о мы можем использовать Human.Run
type Human struct {
	Action
	//...
}


