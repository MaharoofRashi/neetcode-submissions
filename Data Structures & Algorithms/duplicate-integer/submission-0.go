func hasDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _,v := range nums {
        m[v] = true
    }
    if len(nums) != len(m) {
        return true
    } else {
        return false
    }
}
