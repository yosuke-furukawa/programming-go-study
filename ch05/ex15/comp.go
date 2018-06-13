package ex15

func Max1(vals ...int) int {
	if len(vals) <= 0 {
		panic("required argument")
	}
	max := vals[0]
	for _, val := range vals[1:] {
		if max < val {
			max = val
		}
	}
	return max
}
func Max2(v int, vals ...int) int {
	max := v
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func Min1(vals ...int) int {
	if len(vals) <= 0 {
		panic("required argument")
	}
	min := vals[0]
	for _, val := range vals[1:] {
		if min > val {
			min = val
		}
	}
	return min
}

func Min2(v int, vals ...int) int {
	min := v
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}
