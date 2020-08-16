package main

import (
	"fmt"
	"sort"
)

//sort 2d array based on valeu of the first column
func sort2DArray() {
	matrix := [4][2]int{
		{3, 4},
		{2, 3},
		{1, 2},
		{1, 3},
	}
	sort.SliceStable(matrix[:], func(i, j int) bool {
		if matrix[i][1] == matrix[j][1] {
			return false
		}
		return matrix[i][1] < matrix[j][1]
	})

	fmt.Println(matrix)
}

func main() {
	sort2DArray()
}
