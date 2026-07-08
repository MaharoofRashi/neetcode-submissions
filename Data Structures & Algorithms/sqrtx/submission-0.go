func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	i := 1
	for i*i <= x {
		i++
	}
	return i - 1
}
