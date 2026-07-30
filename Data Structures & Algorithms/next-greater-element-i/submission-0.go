func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))

	for i, x := range nums1 {
		result[i] = -1
		found := false
		
		for _, y := range nums2 {
			if found && y > x {
				result[i] = y
				break
			}
			if y == x {
				found = true
			}
		}
	}


	return result
}
