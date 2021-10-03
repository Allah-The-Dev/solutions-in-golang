// idea is either left half of right half from mid will always be sorted
// 
func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (r+l)/2
        if nums[mid] == target {
            return mid
        }
        // check if left half is sorted and if yes does no exist in that part
        if nums[l] <= nums[mid] {
            if target >= nums[l] && target < nums[mid] {
                r = mid-1
            } else {
                l = mid+1
            }
        } else {
            if target > nums[mid] && target <= nums[r] {
                l = mid+1
            } else {
                r = mid-1
            }
        }
    }
    return -1
}