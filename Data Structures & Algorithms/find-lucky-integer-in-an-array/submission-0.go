func findLucky(arr []int) int {
	count := make(map[int]int)
	for _, v := range arr {
		count[v]++
	}

	ans := -1
	for num, freq := range count {
		if num == freq && num > ans {
			ans = num
		}
	}

	return ans
}
