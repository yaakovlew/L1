package first_task

import "fmt"

type Action struct {
	Age    int
	Gender bool
	//...
}

func (act Action) Run() {
	fmt.Println("He/she run")
}

type Human struct {
	Action
	//...
}

/*h := new(Human)
h.Run*/
