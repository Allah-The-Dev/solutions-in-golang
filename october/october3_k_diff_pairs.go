package october

import "sort"

func findPairs(nums []int, k int) int {

	sort.Ints(nums)

	exists := make(map[int]struct{})

	for i, num := range nums {
		if _, ok := exists[num]; ok {
			continue
		}
		if found := searchForPair(k+num, nums[i+1:]); found {
			exists[num] = struct{}{}
		}
	}

	return len(exists)
}

func searchForPair(n int, nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == n {
			return true
		} else if nums[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
