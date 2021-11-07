package study

import "fmt"

func PXSBX() {

	var row = 10
	var column = 20
	var letter = "#"

	for j := 1; j <= row; j++ {
		getRow(row-j, " ")
		getRow(column, letter)
		fmt.Println()
	}
}

func getRow(a int, b string) {
	for i := 1; i <= a; i++ {
		fmt.Printf("%s", b)
	}
}
