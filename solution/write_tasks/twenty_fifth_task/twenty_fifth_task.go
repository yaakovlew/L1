package twenty_fifth_task

import "time"

func sleep(t time.Duration) {
	<-time.After(t)
}
