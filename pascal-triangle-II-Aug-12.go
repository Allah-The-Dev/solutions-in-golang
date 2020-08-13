package main

import "fmt"

func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}
	targetRow := make([]int, rowIndex+1)
	//no of columns to iterate
	for i := 0; i <= rowIndex; i++ {
		//row items iterations
		for j := i; j >= 0; j-- {
			if j == 0 {
				targetRow[j] = 1
				continue
			}
			targetRow[j] = targetRow[j] + targetRow[j-1]
		}
	}
	return targetRow
}

func main() {
	fmt.Println(getRow(6))
}
