func minWindow(s string, t string) string {

	if len(t) > len(s) {
		return ""
	}

	need := map[byte]int{}
	window := map[byte]int{}

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have, required := 0, len(need)
	L := 0
	resL, resR := 0, 0
	resLen := len(s) + 1

	for R := 0; R < len(s); R++ {
		ch := s[R]
		window[ch]++

		if need[ch] > 0 && need[ch] == window[ch] {
			have++
		}

		for have == required {
			if (R - L + 1) < resLen {
				resLen = R - L + 1
				resL, resR = L, R
			}

			leftCh := s[L]
			window[leftCh]--
			L++

			if need[leftCh] > 0 && window[leftCh] < need[leftCh] {
				have--
			}
		}
	}

	if resLen == len(s)+1 {
		return ""
	}

    return s[resL : resR+1]
}
