func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]bool)
	L, maxLen := 0, 0

	for R := 0; R < len(s); R++ {
		for window[s[R]] {
			delete(window, s[L])
			L++
		}

		window[s[R]] = true
		maxLen = max(maxLen, R-L+1)

	}

	return maxLen
}
