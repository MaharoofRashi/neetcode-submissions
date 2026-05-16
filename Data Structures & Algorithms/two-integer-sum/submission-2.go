func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0,1}
	}

	result := []int{0,0}
    for i, _ := range nums {
		for k := 0; k < len(nums); k++ {
			if nums[i] + nums[k] == target && i != k{
				result[0] = k
				result[1] = i
			}
		}
	}
	return result
}
