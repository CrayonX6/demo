package study

import (
	"fmt"
	"time"
)

func Nest(t int) {

	if t > 100 {
		return
	}

	time.Sleep(time.Millisecond*100)
	fmt.Println(t)
	t += 1
	Nest(t)
}
