func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	for i := 0; i < k; i++ {
		last := nums[n-1]

		for j := n-1; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = last
	}
}
