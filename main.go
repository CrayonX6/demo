package main

import (
    "math/rand"
    "strings"
    "study/study"
    "time"
)

func main() {

    rand.Seed(time.Now().Unix())

    /*
    	var a []int
    	a = []int{8, 4, 10, 6, 2, 1, 3}
    	max, min := study.MinMax(a)
    	fmt.Printf("最小值是：%d  最大值是：%d", min, max)

    	fmt.Println()

    	study.Nest(1)
    */

    /*
    	from := 20
    	study.TiaoChu(from, 30)
    */

    /*
    	s:="hello"
    	fmt.Print(s[1:3])
    */

    /*
    	// 使用工具包打印随机数/字符串
    	var ShuZi = util.RandInt(0, 100)
    	var ZiFu = util.RandStr(6)
    	util.PrintVar(ShuZi, ZiFu)
    */

    /*
    	// 使用工具包打印今天是礼拜几
    	util.PrintVar(util.TodayIsWeekTh(carbon.Now()))
    */

    /*
    	// 下载远程文件
    	var fileUrl = "https://avatars.githubusercontent.com/u/13517412?s=64&v=4"
    	abc := util.DownloadRemoteFile(fileUrl, "/Users/rinin/Downloads/a.jpg", 1)
    	util.PrintVar(abc)
    */

    /*
    	// 返回错误类型的示例
    	shang, err := study.ChuFa(10, 0)
    	if err == nil {
    		util.PrintVar(fmt.Sprintf("两数相除为：%d", shang))
    	} else {
    		util.PrintVar(err.Error())
    	}
    */

    /*
    	var x = "111"
    	var y = 10086

    	// 调用结构体
    	study.Tab(util.ToInt(x), util.ToStr(y))
    */

    //study.FangFa()
    //study.ABC(9)

    /*
    	list := study.RecursionDirectory("/Users/rinin/Downloads/poker", func(name string) bool {
    		return !strings.Contains(name, ".png")
    	})
    	util.PrintVar(list)

    */

    /*
    	filename := "/Users/rinin/Downloads/poker"
    	result := study.RecursionDirectory(filename, IsNotPng)

    	for _, filename := range result {
    		err := study.ResizeImage(filename, filename, 500, 700)
    		if err != nil {
    			util.PrintVar(err)
    		}
    	}

    */

    /*study.MapToString(map[string]interface{}{
      	"name":  "leon",
      	"age":   20,
      	"skill": "golang",
      })

    */
    //study.StringToMap("https://www.baidu.com/sayHello?name=abc&age&skill=php")

    //file := "../../../../Downloads/poker/joker_big.png"
    //util.PrintVar(filepath.Ext(file))

    /*seq := []string{"a", "b", "c", "d", "e", "f", "g"}
      // 指定删除位置
      index := 3
      // 输出删除位置之前和之后的元素
      fmt.Println(seq[:index], seq[index+1:])*/

    /*
    	go func() {
    		study.Mainaicha()
    	}()
    	study.Shuadouyin(10)
    */
    /*
       var p1 []int
       var p2 []int
       var p3 []int
       var dipai []int
       study.BuXiPai(90, p1, p2, p3, dipai)
    */

    /* study.Defer_call()
     */

    /*
       var a = 1
       var b = 2
       fmt.Println(a, b)

       b = a
       fmt.Println(a, b)

       c := &a
       fmt.Println(a, *c)

       a = 100
       fmt.Println(a, *c)
    */

    /*
       list := Abc(10, func(x int) bool{
           return x % 5 == 0
       })
       util.PrintVar(list)
    */

    /*
    filename := "/Users/rinin/Downloads/poker"
    result := study.RecursionDirectory(filename, IsNotPng)
    fmt.Print(result)
     */

    study.Triangle()

    //study.JiuJiu()
    //fmt.Print(strings.HasPrefix("select d.deptno, tmp1.avg_sal avg_sal1, tmp2.avg_sal avg_sal2\n  from dept d\n  left join tmp1\n    on d.deptno = tmp1.deptno\n  left join tmp2\n    on d.deptno = tmp2.deptno;", "select") && !strings.Contains("select d.deptno, tmp1.avg_sal avg_sal1, tmp2.avg_sal avg_sal2\n  from dept d\n  left join tmp1\n    on d.deptno = tmp1.deptno\n  left join tmp2\n    on d.deptno = tmp2.deptno;", "for update"))
}

func Abc(max int, fn func(x int) bool) []int {
    var result []int
    for i := 1; i <= max; i++ {
        if fn(i) {
            result = append(result, i)
        }
    }
    return result
}

func IsNotPng(name string) bool {
    return !strings.Contains(name, ".png")
}
