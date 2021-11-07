package study

import "study/util"

func Defer_call()  {
    defer func() {util.PrintVar("1")}()
    defer func() {util.PrintVar("2")}()
    defer func() {util.PrintVar("3")}()
    //panic("触发异常")

}
