func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	maxWater := 0

	for left < right {
		h := min(heights[left], heights[right])
		width := right - left

		maxWater = max(maxWater, h*width)

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}
