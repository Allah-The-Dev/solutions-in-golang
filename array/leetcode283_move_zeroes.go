// method - 1 : 2 pointer approach
func moveZeroes(nums []int)  {
    l, r := 0, 0
    for l < len(nums) && r < len(nums) {
        if nums[l] == 0 && nums[r] != 0 && l < r {
            nums[l], nums[r] = nums[r], 0
            l++
            r = l+1
        } else if nums[l] != 0 {
            l++
        } else {
            r++
        }
    }
}

// method - 2 : move all numbers to left in first pass
// in 2nd pass make remaining indices zero using slow index
func moveZeroes(nums []int)  {
    slowIdx := 0
    for _, num := range nums {
        if num != 0 {
            nums[slowIdx] = num
            slowIdx++
        }
    }
    for ; slowIdx < len(nums); slowIdx++ {
        nums[slowIdx] = 0
    }
}