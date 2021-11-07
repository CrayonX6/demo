package study

func MinMax(list []int) (int, int) {
	var max = 0
	var min = 100
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
		if list[i] < min {
			min = list[i]
		}
	}

	return max, min
}
