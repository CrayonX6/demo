package study

import (
	"fmt"
	"net/url"
	"strings"
	"study/util"
)

func MapToString(m map[string]interface{}) string {
	var str string
	for k, v := range m {
		s := fmt.Sprintf("%s=%s", k, util.ToStr(v))

		str = fmt.Sprintf("%s&%s", str, s)

	}
	str = strings.Trim(str, "&")

	//util.PrintVar(str)

	return str
}

func StringToMap(str string) map[string]interface{} {

	/*
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["aa"] = 3

	var a map[string]interface{}
	a= map[string]interface{}{
		"":nil,
	}
	*/
	x, _ := url.Parse(str)
	var Mymap = make(map[string]interface{})
	a := strings.Split(x.RawQuery, "&")
	for _, v := range a {
		b := strings.Split(v, "=")
		if len(b) < 2 {
			continue
		}

		util.PrintVar(b)

		Mymap[b[0]] = b[1]
	}
	util.PrintVar(Mymap)
	return Mymap
}
