func replaceElements(arr []int) []int {
    for i := 0; i < len(arr) - 1; i++ {
        maxRight := arr[i+1]
        for j := i + 1; j < len(arr); j++ {
            if arr[j] > maxRight {
                maxRight = arr[j]
            }
        }
        arr[i] = maxRight
    }
    
    arr[len(arr)-1] = -1 
    return arr
}