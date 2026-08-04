func longestPalindrome(s string) int {
    counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}

	length := 0
	hasOdd := false
	for _, cnt := range counts {
		length += (cnt/2) * 2
		if cnt%2 == 1 {
			hasOdd = true
		}
	}

	if hasOdd {
		length++
	}

	return length
}