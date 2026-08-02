func minOperations(s string) int {
    count0 := 0
	for i, c := range s {
		expected := byte('0')
		if i%2 == 1 {
			expected = byte('1')
		}

		if byte(c) != expected {
			count0++
		}
	}
	count1 := len(s) - count0
	if count0 < count1 {
		return count0
	}
	return count1
}