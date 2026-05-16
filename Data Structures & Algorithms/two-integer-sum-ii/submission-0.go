func twoSum(numbers []int, target int) []int {

	// edge case:- if the len(numbers) == 2 then return 0,1
	// if len(numbers) == 2 {
	// 	return []int{0,1}
	// }

	// double loop the array and find the result and return it by +1 of the indies

	i, j := 0, len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
