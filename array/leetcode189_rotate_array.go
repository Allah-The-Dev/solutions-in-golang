// [1,2,3,4,5,6,7], 3
func rotate(nums []int, k int)  {
    numsLen := len(nums)
    totalRotation := k % numsLen
    if totalRotation == 0 { return }
    reverse(nums, 0, numsLen-1) // [7,6,5,4,3,2,1]
    reverse(nums, 0, totalRotation-1) // [5,6,7,4,3,2,1] 
    reverse(nums, totalRotation, numsLen-1) // [5,6,7,1,2,3,4]
}

func reverse(nums []int, l, r int) {
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}