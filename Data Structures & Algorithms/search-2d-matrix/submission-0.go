func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		left, right := 0, len(matrix[i]) - 1
		for left <= right {
			mid := left + (right - left) / 2 

			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false

}
