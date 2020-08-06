package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	duplicates := []int{}
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] >= 0 {
			nums[num-1] = -nums[num-1]
		} else {
			duplicates = append(duplicates, num)
		}
		fmt.Printf("%#v\n", nums)
	}
	return duplicates
}

// func main() {
// 	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
// 	fmt.Println(findDuplicates(nums))
// }
