package study

import (
    "math/rand"
    "study/util"
    "time"
)

func BuXiPai(n int, p1 []int, p2 []int, p3 []int, dipai []int) {
    var a []int   //最终数组
    var b [][]int //大数组
    for i := 1; i <= 13; i++ {
        var c []int //打包数组
        t := util.RandInt(0, 100)
        if int(t) <= n {
            for j := 1; j <= 4; j++ {
                c = append(c, i)
            }
            b = append(b, c)
        } else {
            c = append(c, i)
            for q := 1; q <= 4; q++ {
                b = append(b, c)
            }
        }
    }
    t := util.RandInt(0, 100)
    var wangzha = []int{14, 15}
    var dawang = []int{15}
    var xiaowang = []int{14}

    if int(t) <= n {
        b = append(b, wangzha)
    } else {
        b = append(b, dawang)
        b = append(b, xiaowang)
    }

    ShuffleUser(b)
    for _, v := range b {
        for _, pai := range v {
            a = append(a, pai)
        }
    }
    util.PrintVar(a)

    p1 = a[:17]
    p2 = a[17:34]
    p3 = a[34:51]
    dipai = a[51:]
    util.PrintVar(p1, p2, p3, dipai)
}

func ShuffleUser(pai [][]int) [][]int {
    if len(pai) <= 0 {
        return pai
    }
    rand.Seed(time.Now().Unix())
    for i := len(pai) - 1; i >= 0; i-- {
        num := rand.Intn(len(pai))
        pai[i], pai[num] = pai[num], pai[i]
    }
    return pai
}