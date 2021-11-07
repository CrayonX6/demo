package study

import (
    "study/util"
)

type Person struct {
    Name  string   `json:"name"`  // 姓名
    Age   int      `json:"age"`   // 年龄
    Skill []string `json:"skill"` // 拥有的能力
}

func (a *Person) hasSkill(s string) bool {
    exist, _ := util.InArray(s, a.Skill)
    return exist
}

func Tab(x int, y string) {
    var group = []Person{
        {
            Name:  "leon",
            Age:   20,
            Skill: []string{"普通话", "英语", "粤语"},
        },
        {
            Name:  "lisa",
            Age:   24,
            Skill: []string{"普通话", "上海话"},
        },
        {
            Name:  "张三",
            Age:   50,
            Skill: []string{"潮汕话", "英语"},
        },
    }
    for _, value := range group {
        if value.hasSkill("英语") {
            util.PrintVar(value.Name)
        }
    }
    util.PrintVar(util.InterfaceToJson(group))
}

/*
var MyDictionay = map[interface{}]interface{}{
	"zhangsan": "19",
	"lisi": "24",
}
*/
