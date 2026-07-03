func containsNearbyDuplicate(nums []int, k int) bool {
	lastSeen := make(map[int]int)

	for i, num := range nums {
		prevIndex, exists := lastSeen[num]

		if exists {
			if i-prevIndex <= k {
				return true
			}
		}

		lastSeen[num] = i
	}

	return false
}
