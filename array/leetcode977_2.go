// approach two 
// find the mid i.e. point where nums[i] is -ve and nums[i+1] is +ve
// then consider left and right of array as 2 different array
// and then try merging them using 2 pointer approach
func sortedSquares(nums []int) []int {
    numsLen := len(nums)
    
    idxOfMid := -1
    for i := 0; i < numsLen; i++ {
        if i == 0 && nums[0] >= 0 { 
            idxOfMid = 0 
            break 
        } // all nos are positive
        if i+1 < numsLen && nums[i] < 0 && nums[i+1] >=0 {
            idxOfMid = i+1
            break
        }
        if i == numsLen-1 {
            idxOfMid = numsLen
        }
    }
    //fmt.Println(idxOfMid)
    l, r := idxOfMid-1, idxOfMid
    
    res := make([]int, numsLen)
    resIdx := 0
    for l >= 0 && r < numsLen {
        numL, numR := abs(nums[l]), abs(nums[r])
        if numL <= numR {
            res[resIdx] = numL*numL
            l--
        } else {
            res[resIdx] = numR*numR
            r++
        }
        resIdx++
    }
    
    for l >= 0 {
        res[resIdx] = nums[l]*nums[l]
        resIdx++
        l--
    }
    for r < numsLen {
        res[resIdx] = nums[r]*nums[r]
        resIdx++
        r++
    }
    
    return res
}

func abs(a int) int {
    if a < 0 { return -a }
    return a
}