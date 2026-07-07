func arrangeCoins(n int) int {
	row := 0;

	for n > 0 {
		row++
		n -= row
	}

	if n < 0 {
		row--
	}

	return row
}
