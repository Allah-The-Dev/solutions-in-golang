package main

import (
	"fmt"
	"math"
)

func titleToNumber(s string) int {
	char := []rune(s)
	if len(s) == 1 {
		return int(char[0] - 64)
	}
	columnNo := 0
	for i := 1; i < len(s); i++ {
		columnNo += int(math.Pow(26, float64(i)))
	}
	for i, ch := range char {
		if len(s) == (i + 1) {
			columnNo += int(ch - 64 - 1)
			continue
		}
		columnNo += int(ch-64-1) * int(math.Pow(26, float64((len(s)-(i+1)))))
	}
	return columnNo + 1
}

func main() {
	fmt.Println(titleToNumber("AAA"))
}
