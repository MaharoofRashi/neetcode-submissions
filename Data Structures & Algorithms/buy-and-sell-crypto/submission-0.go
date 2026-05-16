func maxProfit(prices []int) int {
	maxP := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]

			if profit > maxP {
				maxP = profit
			}
		}
	}

	return maxP
}
