func evalRPN(tokens []string) int {

	operators := map[string]func(int, int) int {
		"+" : func(a, b int) int { return a + b },
		"-" : func(a, b int) int { return a - b },
		"*" : func(a, b int) int { return a * b },
		"/" : func(a, b int) int { return a / b },
	}

	stack := []int{}

	for _, token := range tokens {
		if op, ok := operators[token]; ok {
			b := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			a := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]

			stack = append(stack, op(a, b)) 
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
