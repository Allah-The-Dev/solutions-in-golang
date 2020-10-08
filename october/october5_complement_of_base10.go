package october

func bitwiseComplement(N int) int {
	sum := 1
	for N > sum {
		sum = sum*2 + 1
	}
	return sum - N
}
