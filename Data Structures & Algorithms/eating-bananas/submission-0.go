func canFinish(piles []int, k int, h int) bool {
	totalHours := 0

	for _, pile := range piles {
		totalHours += (pile + k - 1) / k
	}

	return totalHours <= h
}

func minEatingSpeed(piles []int, h int) int {
	low, high := 1, 0

	for _, pile := range piles {
		if pile > high {
			high = pile
		}
	}

	result := high 

	for low <= high {
		mid := (low + high) / 2

		if canFinish(piles, mid, h) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}
