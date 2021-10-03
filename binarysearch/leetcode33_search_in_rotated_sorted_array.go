// first do binary search for the pivot point of array
// from where array is rotated

// after that adjust left and right pointer accordingly 
// and do the search in that half

func search(nums []int, target int) int {
    res := -1
    numsLen := len(nums)
    // [4,5,6,7,0,1,2]
    // here idx of 7 will be pivot
    l, r := 0, numsLen-1
    for l < r {
        mid := (l+r)/2
        if nums[mid] > nums[r] {
            l = mid+1
        } else {
            r = mid
        }
    }
    pivotIdx := l-1
    l, r = 0, numsLen-1
    if pivotIdx != -1 {
        if target == nums[pivotIdx] {
            return pivotIdx
        } else if target >= nums[0] && target < nums[pivotIdx] {
            l, r = 0, pivotIdx-1
        } else {
            l, r = pivotIdx+1, numsLen-1
        }
    }
    
    for l <= r {
        mid := (l+r)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            r = mid-1
        } else {
            l = mid+1
        }
    }
    return res
}