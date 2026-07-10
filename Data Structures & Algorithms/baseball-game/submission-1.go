func calPoints(operations []string) int {
	stack := []int{}
	
	for _, operation := range operations {
		switch operation {
			case "+":
				stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
			case "D":
				stack = append(stack, stack[len(stack)-1]*2)
			case "C":
				stack = stack[:len(stack)-1]
			default:
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
