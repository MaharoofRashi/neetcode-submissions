func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter++
			if counter > result {
				result = counter
			}
		} else {
			counter = 0
		}
	}

	return result
}
