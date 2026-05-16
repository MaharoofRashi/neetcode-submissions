func longestConsecutive(nums []int) int {

	// edge case of nums = 0

	if len(nums) == 0 {
		return 0
	}

	// sort to accending
	sort.Slice(nums, func (a, b int) bool {
		return nums[a] < nums[b]
	})

	// loop through the array and count ++

	count := 1
	longest := 1

	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i+1] {
			continue
		} else if nums[i + 1] - nums[i] == 1 {
			count++
			longest = max(longest, count)
		} else {
			count = 1
		}
	}

	return longest
}
