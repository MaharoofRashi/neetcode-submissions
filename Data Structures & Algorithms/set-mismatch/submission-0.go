func findErrorNums(nums []int) []int {
    n := len(nums)
	count := make([]int, n+1)
	for _, v := range nums {
		count[v]++
	}

	var dup, missing int
	for i := 0; i <= n; i++ {
		if count[i] == 2 {
			dup = i
		} else if count[i] == 0 {
			missing = i
		}
	}

	return []int{dup, missing}
}