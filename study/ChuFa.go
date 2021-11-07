package study

import "errors"

func ChuFa(x int, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("你的除数为零哦")
		return
	}

	z = x / y
	return
}

