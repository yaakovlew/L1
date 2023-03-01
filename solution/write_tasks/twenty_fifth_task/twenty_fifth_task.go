package twenty_fifth_task

import "time"

func TwentyFifth(t time.Duration) {
	// остановка на время t
	<-time.After(t * time.Second)
}
