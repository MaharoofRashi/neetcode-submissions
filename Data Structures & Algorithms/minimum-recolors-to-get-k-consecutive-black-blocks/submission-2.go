func minimumRecolors(blocks string, k int) int {
	whites := 0

	for i := 0; i < k; i++ {
		if blocks[i] == 'W'{
			whites++
		}
	}

	minOps := whites

	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			whites++
		}
		if blocks[i-k] == 'W' {
			whites--
		}
		if whites < minOps {
			minOps = whites
		}
	}

	return minOps
}
