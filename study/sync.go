package study

import (
	"fmt"
	"time"
)

func Mainaicha() {
	fmt.Printf("正在等待奶茶制作完成\n")
	time.Sleep(time.Second * 5)
	fmt.Printf("奶茶制作完成\n")
}

func Shuadouyin(times int) {
	if times <= 0 {
		return
	}
	fmt.Printf("刷抖音中\n")
	time.Sleep(time.Second)
	Shuadouyin(times - 1)
}
