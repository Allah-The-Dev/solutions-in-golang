package main

import (
	"sort"
)

func sortArrayBasedOnLastIndex(matrix [][]int) {
	sort.SliceStable(matrix[:], func(i, j int) bool {
		if matrix[i][1] == matrix[j][1] {
			return false
		}
		return matrix[i][1] < matrix[j][1]
	})
}

func eraseOverlapIntervals(intervals [][]int) int {
	sortArrayBasedOnLastIndex(intervals)
	close := 0
	count := 0
	for i := range intervals {
		if intervals[i][0] > close {
			close = intervals[i][1]
		} else {
			count++
		}
	}
	return count
}

// func main() {
// 	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
// }
