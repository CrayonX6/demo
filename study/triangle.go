package study

import "fmt"

func Triangle() {

	var row = 5
	var letter = "#"

	for j := 1; j <= row; j++ {
		getRow(row-j, " ")
		getRow(2*j-1, letter)
		getRow(row-j, " ")
		fmt.Println()
	}

	/*
		for j := 1; j <= row; j++ {
			getRow(j-1, " ")
			getRow(2*(row-j)+1, letter)
			getRow(j-1, " ")
			fmt.Println()
		}
	*/
}
