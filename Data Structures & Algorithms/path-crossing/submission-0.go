func isPathCrossing(path string) bool {
    visited := make(map[[2]int]bool)
	x, y := 0, 0
	visited[[2]int{0,0}] = true
	
	for _, c := range path {
		switch c {
			case 'N':
				y++
			case 'S':
				y--
			case 'E':
				x++
			case 'W':
				x--
		}

		p := [2]int{x, y}
		if visited[p] {
			return true
		}
		visited[p] = true
	}

	return false
}