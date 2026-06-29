func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			leftOk := i == 0 || flowerbed[i-1] == 0
			rightOk := i == len(flowerbed) - 1 || flowerbed[i+1] == 0

			if leftOk && rightOk {
				flowerbed[i] = 1
				count++
			}
		}
	}

	return count >= n
}