func arrangeCoins(n int) int {
	return int((-1 + math.Sqrt(1+8*float64(n)))/2)
}
