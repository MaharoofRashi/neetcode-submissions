func calPoints(operations []string) int {
	stack := []int{}

	for _, operation := range operations {
		if operation == "+" {
			n := len(stack)
			stack = append(stack, stack[n-1]+stack[n-2])
		} else if operation == "D" {
			n := len(stack)
			stack = append(stack, stack[n-1]*2)
		} else if operation == "C" {
			stack = stack[:len(stack)-1]
		} else {
			x, _ := strconv.Atoi(operation)
			stack = append(stack, x)
		}
	}

	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum 
}
