package study

import "fmt"

func TiaoChu(from, to int) {
	for i := from; i <=to; i++ {
		fmt.Printf("%d\n", i)
		if i > 25 {
			break
		}

	}

}
