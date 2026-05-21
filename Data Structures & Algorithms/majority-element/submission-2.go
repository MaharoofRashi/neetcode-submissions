func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}


	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 0
		}
	}

	result := 0

	for i, v := range m {
		if v > result {
			result = i
		}
	}

	return result
    
}
