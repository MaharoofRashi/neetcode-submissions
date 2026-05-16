func maxProfit(prices []int) int {
	L, R := 0, 1
	maxP := 0

	for R < len(prices) {
		if prices[R] > prices[L] {
			profit := prices[R] - prices[L]
			maxP = max(maxP, profit)
		} else {
			L = R
		}
		R++
	}
	return maxP
}
