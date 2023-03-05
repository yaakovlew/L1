package fifteenth_task

import "sync"

func createHugeString(a int) (s string) {
	s = ""
	//....
	return
}

var justString string
var m sync.Mutex

func someFunc() {
	v := createHugeString(1 << 10)
	m.Lock()
	justString = v[:len(v)]
	m.Unlock()
}


