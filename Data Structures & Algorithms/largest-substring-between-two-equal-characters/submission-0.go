func maxLengthBetweenEqualCharacters(s string) int {
    first := make(map[byte]int)
	result := -1
	
	for i := 0; i < len(s); i++ {
		c := s[i]
		if idx, ok := first[c]; ok {
			if i-idx-1 > result {
				result = i - idx - 1
			}
		} else {
			first[c] = i
		}
	}

	return result
}