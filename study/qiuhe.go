package study

import "fmt"

func QiuHe() {
	var from, to int

	from = 1
	to = 10

	var sum = 0
	for i := from; i <= to; i++ {
		sum = sum + i
		//fmt.Printf("这一行累加了数值：%d，得到的值为：%d\n", i, sum)
	}

	fmt.Printf("从%d累加到%d的总和为：%d", from, to, sum)
}
