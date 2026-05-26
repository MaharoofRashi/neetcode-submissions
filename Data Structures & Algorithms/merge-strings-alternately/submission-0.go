func mergeAlternately(word1 string, word2 string) string {
	var ans strings.Builder
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		ans.WriteByte(word1[i])
		ans.WriteByte(word2[j])
		i++
		j++
	}

	ans.WriteString(word1[i:])
	ans.WriteString(word2[j:])
	
	return ans.String()
}
