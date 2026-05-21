func majorityElement(nums []int) int {
    m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	result, maxCount := 0, 0

	for elem, count := range m {
		if count > maxCount {
			maxCount = count
			result = elem
		}
	}

	return result
}
