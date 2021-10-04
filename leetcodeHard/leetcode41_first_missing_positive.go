// algo 
// 1. make all negative nos zero
// 2. for left out nos in a separte iteration, make nos in that index -ve
// 3. use nums as cache

func firstMissingPositive(nums []int) int {
    // make all -ve nos zero
    for i, num := range nums {
        if num < 0 {
            nums[i] = 0
        }
    }
    numsLen := len(nums)
    //now use nums array as cache
    for _, num := range nums {
        absOfNum := abs(num)
        idxForNum := absOfNum-1
        if idxForNum < numsLen && idxForNum >= 0 && nums[absOfNum-1] >= 0 {
            if nums[idxForNum] == 0 {
                nums[idxForNum] = -numsLen
            } else {
                nums[idxForNum] = -nums[idxForNum]
            }
        }
    }
    // fmt.Println(nums)
    // now first leftover positive nos is the first missing positive
    for i, num := range nums {
        if num >= 0 {
            return i+1
        }
    }
    return numsLen+1
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}