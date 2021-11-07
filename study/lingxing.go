package study

import (
	"fmt"
	"strings"
)

func LingXing() {

	getPositive(5, "#")
	getNegative(4, "#")

	fmt.Print("\n\n\n\n\n")
	x := printRow(10, "@", true)
	fmt.Print(x)
	//printRow(10, "@", false)
}

func getPositive(m int, n string) {

	for j := 1; j <= m; j++ {
		getRow(m-j, " ")
		getRow(2*j-1, n)
		getRow(m-j, " ")
		fmt.Println()
	}
}
func getNegative(s int, t string) {

	for j := 1; j <= s; j++ {
		getRow(j-1, " ")
		getRow(2*(s-j)+1, t)
		getRow(j-1, " ")
		fmt.Println()
	}
}

func printRow(rowTotal int, letter string, asc bool) string {
	var x int
	var lx string
	for i := 1; i <= rowTotal; i++ {
		if asc == true {
			x = i
		} else {
			x = rowTotal - i + 1
		}
		spaceString := strings.Repeat(" ", rowTotal-x)
		letterString := strings.Repeat(letter, x*2-1)
		lx += fmt.Sprintf("%s%s\n", spaceString, letterString)
	}
	return lx
}

