package septemberleetcodechallange

import "math"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i, num := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if int(math.Abs(float64(nums[j]-num))) <= t {
				return true
			}
		}
	}
	return false
}
