func dailyTemperatures(temperatures []int) []int {

	// brute for O(n^2)
	result := []int{}
	

	for i := 0; i < len(temperatures); i++ {
		found := false
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result = append(result, j - i)
				found = true
				break
			}
		}
		if !found {
			result = append(result, 0)
		}
	}

	return result
}
