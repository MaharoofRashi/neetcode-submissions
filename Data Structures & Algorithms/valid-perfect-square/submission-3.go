func isPerfectSquare(num int) bool {
	left, right := 1, num

	for left <= right {
		mid := left + (right - left)/2
		sum := mid * mid

		if sum == num {
			return true
		} else if sum < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
