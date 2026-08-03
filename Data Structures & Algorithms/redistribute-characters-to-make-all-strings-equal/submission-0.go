func makeEqual(words []string) bool {
    var count [26]int
	for _, w := range words {
		for _, c := range w {
			count[c - 'a']++
		}
	}
	
	n := len(words)
	for _, c := range count {
		if c%n != 0 {
			return false
		}
	}
	return true
}