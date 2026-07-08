func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	lo, hi := 0, x/2
	ans := 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		
		if mid * mid == x {
			return mid
		} else if mid * mid < x {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return ans
}
