package study

import (
	"fmt"
	"study/util"
)

func FangFa() {
	name := "leon"
	var result = SaySomething(name, "你好") // （"Leon","你好"）-> 实参（实际参数，有具体值的参数）
	util.PrintVar(result)
}





func SaySomething(who string, hua string) string { // （变量 who、hua） -> 形参（形式参数）
	return fmt.Sprintf("%s说了：%s", who, hua)
}
