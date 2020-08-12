package main

import (
	"sort"
)

func hIndexLinear(citations []int) int {

	sort.Ints(citations)
	result := 0
	i := 0
	for ; i < len(citations) && len(citations[i:]) >= citations[i]; i++ {
		result = citations[i]
	}

	if res2 := len(citations) - i; res2 > result {
		return res2
	}
	return result
}

// func main() {
// 	fmt.Println(hIndex([]int{0, 1}))
// }
