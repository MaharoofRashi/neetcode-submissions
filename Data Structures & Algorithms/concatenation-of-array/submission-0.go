func getConcatenation(nums []int) []int {
	ans := []int(nums)

	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[i])
	}

	return ans
    
}
