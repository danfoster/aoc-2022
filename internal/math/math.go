package math

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
