func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	left, right := 0, n-1

	for i := n-1; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			ans[i] = nums[left] * nums[left]
			left++
		} else {
			ans[i] = nums[right] * nums[right]
			right--
		}
	}

	return ans
}



func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}