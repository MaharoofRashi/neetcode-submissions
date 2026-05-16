func checkInclusion(s1 string, s2 string) bool {
	// create a map of s1 with byte:int
	m := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
	}

	// intialize L 
	L := 0

	// write the R loop to check are return the result
	for R := 0; R < len(s2); R++ {
		m[s2[R]]--

		if R - L + 1 > len(s1) {
			m[s2[L]]++
			L++
		}

		if R - L + 1 == len(s1) {
			allZero := true
			for _, v := range m {
				if v != 0 {
					allZero = false
					break
				}
			}
			if allZero {
				return true
			}
		}
	}

	return false
}
