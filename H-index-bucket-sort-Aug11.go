package main

import "fmt"

func hIndexBucketSearch(citations []int) int {
	bucket := make([]int, len(citations)+1)
	for _, citation := range citations {
		if citation >= len(citations) {
			bucket[len(citations)]++
		} else {
			bucket[citation]++
		}
	}

	totalCitationsSoFar := 0
	for i := len(citations); i >= 0; i-- {
		totalCitationsSoFar += bucket[i]
		if totalCitationsSoFar >= i {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(hIndexBucketSearch([]int{0, 3, 1, 7, 8}))
}
