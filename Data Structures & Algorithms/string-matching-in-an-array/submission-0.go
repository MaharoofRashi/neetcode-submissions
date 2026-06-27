func stringMatching(words []string) []string {
    result := []string{}

	for i, word := range words {
		for j, other := range words {
			if i != j && strings.Contains(other, word) {
				result = append(result, word)
				break
			}
		}
	}

	return result
}