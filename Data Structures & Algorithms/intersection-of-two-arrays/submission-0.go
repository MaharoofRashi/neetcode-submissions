func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	for _, n := range nums1 {
		set1[n] = true
	}

	result := []int{}
	added := make(map[int]bool)
	for _, n := range nums2 {
		if set1[n] && !added[n] {
			result = append(result, n)
			added[n] = true
		}
	}

	return result
}
