package study

import "fmt"

func JiuJiu() {
	for i := 1; i <= 9; i++ {
		for j := i; j <= 9; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
