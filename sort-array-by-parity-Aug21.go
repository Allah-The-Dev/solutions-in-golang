package main

func sortArrayByParity(A []int) []int {
	if len(A) == 1 {
		return A
	}
	evenNoPtr := 0
	oddNoPtr := 0
	for oddNoPtr < len(A) {
		if A[oddNoPtr]%2 == 0 {
			A[evenNoPtr], A[oddNoPtr] = A[oddNoPtr], A[evenNoPtr]
			evenNoPtr++
			oddNoPtr++
		} else {
			oddNoPtr++
		}
	}
	return A
}
