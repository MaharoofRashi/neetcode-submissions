func replaceElements(arr []int) []int {
    maxRight := -1
    
    for i := len(arr) - 1; i >= 0; i-- {
        newMax := max(maxRight, arr[i])
        arr[i] = maxRight
        maxRight = newMax
    }
    
    return arr
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}