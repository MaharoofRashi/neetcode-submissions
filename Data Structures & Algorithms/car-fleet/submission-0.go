func carFleet(target int, position []int, speed []int) int {

	// Step 1
	n := len(position)

	type Car struct {
		pos int
		spd int
	}

	cars := make([]Car, n)

	for i := 0; i < n; i++ {
		cars[i] = Car{position[i], speed[i]}
	}

	sort.Slice(cars, func (a, b int) bool {
		return cars[a].pos > cars[b].pos
	})

	// Step 2
	stack := []float64{}
	
	for _, car := range cars {
		time := float64(target - car.pos) / float64(car.spd)
		
		if len(stack) == 0 || time > stack[len(stack) - 1] {
			stack = append(stack, time)
		}
	}

	return len(stack)

}
