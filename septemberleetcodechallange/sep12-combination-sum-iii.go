package septemberleetcodechallange

func getSum3Combnations(start, remainingN, reqArrSize int, currCombination []int, result *[][]int) {
	if remainingN == 0 && len(currCombination) == reqArrSize {
		tempRes := make([]int, reqArrSize)
		copy(tempRes, currCombination)
		*result = append(*result, tempRes)
		return
	}
	for i := start; i <= 9; i++ {
		if remainingN-i < 0 {
			continue
		}
		currCombination = append(currCombination, i)
		getSum3Combnations(i+1, remainingN-i, reqArrSize, currCombination, result)
		currCombination = currCombination[:len(currCombination)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	getSum3Combnations(1, n, k, []int{}, &result)
	return result
}
