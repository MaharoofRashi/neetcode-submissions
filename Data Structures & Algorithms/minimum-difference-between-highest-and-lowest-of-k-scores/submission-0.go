func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)

    minDiff := math.MaxInt32

    for L := 0; L+k-1 < len(nums); L++ {
        R := L + k - 1
        diff := nums[R] - nums[L]
        if diff < minDiff {
            minDiff = diff
        }
    }

    return minDiff
}