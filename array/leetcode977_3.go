// take 2 pointers from left and right of array
// whichever have max absolute value put it in result from last
func sortedSquares(nums []int) []int {
    numsLen := len(nums)
    res := make([]int, numsLen)
    resIdx := numsLen-1
    l, r := 0, numsLen-1
    for l <= r {
        if abs(nums[l]) > abs(nums[r]) {
            res[resIdx] = nums[l]*nums[l]
            l++
        } else {
            res[resIdx] = nums[r]*nums[r]
            r--
        }
        resIdx--
    }
    return res
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}