func commonChars(words []string) []string {
    common := [26]int{}
	for _, c := range words[0] {
		common[c-'a']++
	}

	for _, w := range words[1:] {
		count := [26]int{}
		for _, c := range w {
			count[c-'a']++
		}
		for i := 0; i < 26; i++ {
			if count[i] < common[i] {
				common[i] = count[i]
			}
		}
	}

	result := []string{}
	for i := 0; i < 26; i++ {
		for j := 0; j < common[i]; j++ {
			result = append(result, string(rune('a'+i)))
		}
	}
	return result
}