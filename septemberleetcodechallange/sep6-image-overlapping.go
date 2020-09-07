package septemberleetcodechallange

func LargestOverlap(A [][]int, B [][]int) int {

	type trslationIdx [2]int

	onesIdxsInA := [][]int{}
	onesIdxsInB := [][]int{}

	for i, row := range A {
		for j, num := range row {
			if num == 1 {
				temp := []int{i, j}
				onesIdxsInA = append(onesIdxsInA, temp)
			}
			if B[i][j] == 1 {
				temp := []int{i, j}
				onesIdxsInB = append(onesIdxsInB, temp)
			}
		}
	}

	translationMap := make(map[trslationIdx]int)

	curTranslateMax := 0

	for _, idx := range onesIdxsInA {
		for _, jdx := range onesIdxsInB {
			idxObj := trslationIdx{jdx[0] - idx[0], jdx[1] - idx[1]}
			if count, ok := translationMap[idxObj]; ok {
				translationMap[idxObj] = count + 1
			} else {
				translationMap[idxObj] = 1
			}
			if translationMap[idxObj] > curTranslateMax {
				curTranslateMax = translationMap[idxObj]
			}
		}
	}
	return curTranslateMax
}
