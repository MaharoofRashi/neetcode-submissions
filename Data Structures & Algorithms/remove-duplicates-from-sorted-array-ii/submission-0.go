func removeDuplicates(nums []int) int {
    n := len(nums)
	
	if n <= 2 {
		return n
	}

	k := 2
	for i := 2; i < n; i++ {
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}