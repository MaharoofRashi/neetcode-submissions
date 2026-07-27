func maxNumberOfBalloons(text string) int {
    freq := make(map[rune]int)
	for _, ch := range text {
		freq[ch]++
	}

	b := freq['b']
	a := freq['a']
	l := freq['l'] / 2
	o := freq['o'] / 2
	n := freq['n']

	result := b
	if a < result {
		result = a
	}
	if l < result {
		result = l
	}
	if o < result {
		result = o
	}
	if n < result {
		result = n
	}

	return result
}