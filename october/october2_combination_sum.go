package october

func CombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	backtrack(&res, candidates, []int{}, target, 0)
	return res
}

func backtrack(res *[][]int, candidates, curRes []int, target, idx int) {
	if target == 0 {
		tempRes := make([]int, len(curRes))
		copy(tempRes, curRes)
		*res = append(*res, tempRes)
		return
	}
	if target < 0 {
		return
	}
	for i := idx; i < len(candidates); i++ {
		value := candidates[i]
		curRes = append(curRes, value)
		backtrack(res, candidates, curRes, target-value, i)
		curRes = curRes[:len(curRes)-1]
	}
}
