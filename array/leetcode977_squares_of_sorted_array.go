// approach one 
// sorting array once more with absolute value of nos
func sortedSquares(nums []int) []int {
    quickSort(nums, 0, len(nums)-1)
    for i, num := range nums {
        nums[i] = num*num
    }
    return nums
}

func quickSort(nums []int, start, end int) {
    if start < end {
        pIndex := partition(nums, start, end)
        quickSort(nums, start, pIndex-1)
        quickSort(nums, pIndex+1, end)
    }
}

func partition(nums []int, start, end int) int {
    pivot, pIndex := abs(nums[end]), start
    for i := start; i < end; i++ {
        if abs(nums[i]) <= pivot {
            nums[i], nums[pIndex] = nums[pIndex], nums[i]
            pIndex++
        }
    }
    nums[end], nums[pIndex] = nums[pIndex], nums[end]
    return pIndex
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}