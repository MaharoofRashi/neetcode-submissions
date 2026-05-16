func isPalindrome(s string) bool {

	// make the string to lowercase
	str := strings.ToLower(s)

	// filter only the alphabetic characters
	var filtered []byte
	for i := 0; i < len(str); i++ {
		c := str[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			filtered = append(filtered, c)
		}
	}

	// do 2 pointer using single loop
	i, j := 0, len(filtered) - 1
	for i < j {
		if filtered[i] != filtered[j] {
			return false
		}
		i++
		j--
	}


	// at this point if not returned bool as false, then return true because it passed the checks
	return true

}
