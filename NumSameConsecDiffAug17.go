package main

var result = []int{}

func getNosForI(initialDigit, N, K, currentNo, currentNosLength int) {
	if currentNosLength == N {
		result = append(result, currentNo)
		return
	}
	next := (currentNo % 10) + K
	if next <= 9 {
		currentNosLength++
		getNosForI(initialDigit, N, K, (currentNo*10)+next, currentNosLength)
		currentNosLength--
	}
	next = (currentNo % 10) - K
	if K != 0 && next >= 0 {
		currentNosLength++
		getNosForI(initialDigit, N, K, (currentNo*10)+next, currentNosLength)
		currentNosLength--
	}
}

func numsSameConsecDiff(N int, K int) []int {
	result = []int{}
	for i := 1; i <= 9; i++ {
		getNosForI(i, N, K, i, 1)
	}
	if N == 1 {
		result = append(result, 0)
	}
	return result
}

// func main() {
// 	fmt.Println(numsSameConsecDiff(2, 1))
// }
