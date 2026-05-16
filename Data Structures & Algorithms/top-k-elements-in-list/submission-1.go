func topKFrequent(nums []int, k int) []int {

	// do the counting using map
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	// convert to slice of the values and make it to the desceding order.

	values := []int{}

	for i, _ := range counts {
		values = append(values, i)
	}

	sort.Slice(values, func (a, b int) bool {
		return counts[values[a]] > counts[values[b]]
	})

	// return the countings based on k 

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, values[i])
	}
	return result
}
