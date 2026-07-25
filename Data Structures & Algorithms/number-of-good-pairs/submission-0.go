func numIdenticalPairs(nums []int) int {
    count := make(map[int]int)
	result := 0

	for _, n := range nums {
		result += count[n]
		count[n]++
	}

	return result
}