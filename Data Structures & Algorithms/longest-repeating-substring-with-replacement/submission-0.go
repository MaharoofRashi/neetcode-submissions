func characterReplacement(s string, k int) int {

	count := make(map[byte]int)
	maxCount, maxLen := 0, 0
	L := 0


	for R := 0; R < len(s); R++ {
		count[s[R]]++
		if count[s[R]] > maxCount {
			maxCount = count[s[R]]
		}

		for (R - L + 1) - maxCount > k {
			count[s[L]]--
			L++
		}

		if R - L + 1 > maxLen {
			maxLen = R - L + 1
		}
	}

	return maxLen
}
