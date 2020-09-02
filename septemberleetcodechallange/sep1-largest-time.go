package septemberleetcodechallange

import (
	"fmt"
	"strconv"
)

func backtrack(permArr *[][]int, nums []int, curArr []int, visited []bool) {
	if len(curArr) == 4 {
		copyCurArray := make([]int, 4)
		copy(copyCurArray, curArr)
		*permArr = append(*permArr, copyCurArray)
		fmt.Printf("returned from stack : backtrack(%v, %v)\n", curArr, visited)
	}

	for i, num := range nums {
		if visited[i] {
			continue
		}
		curArr = append(curArr, num)
		visited[i] = true
		fmt.Printf("call stack : backtrack(%v, %v)\n", curArr, visited)
		backtrack(permArr, nums, curArr, visited)
		fmt.Printf("before currArr %v\n", curArr)
		curArr = curArr[:len(curArr)-1]
		fmt.Printf("after currArr %v\n", curArr)
		visited[i] = false
	}
}

func permutations(nums []int) [][]int {
	permArr := make([][]int, 0, 24)
	backtrack(&permArr, nums, make([]int, 0, 4), make([]bool, 4))
	return permArr
}

//LargestTimeFromDigits ...
func LargestTimeFromDigits(A []int) string {
	//0 = index of permArr to use to store resultant permutations
	permArr := permutations(A)
	bestMins := -1
	resultIdxInPermArr := -1
	for i, nums := range permArr {
		totalHours := nums[0]*10 + nums[1]
		totalMins := nums[2]*10 + nums[3]
		if totalHours < 24 && totalMins < 60 {
			if totalHours*60+totalMins > bestMins {
				bestMins = totalHours*60 + totalMins
				resultIdxInPermArr = i
			}
		}
	}
	if resultIdxInPermArr != -1 {
		nums := permArr[resultIdxInPermArr]
		return strconv.Itoa(nums[0]) + strconv.Itoa(nums[1]) + ":" + strconv.Itoa(nums[2]) + strconv.Itoa(nums[3])
	}
	return ""
}
