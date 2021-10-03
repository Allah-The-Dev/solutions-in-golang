/*
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
*/

func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    currRes := []int{}
    backtrack(candidates, currRes, 0, 0, target, &res)
    return res
}

func backtrack(candidates, currRes []int, currIdx, currSum, target int, res *[][]int) {
    if currSum == target {
        temp := make([]int, len(currRes))
        copy(temp, currRes)
        *res = append(*res, temp)
        return
    }
    for i := currIdx; i < len(candidates); i++ {
        num := candidates[i]
        if currSum + num <= target {
            currSum += num
            currRes = append(currRes, num)
            backtrack(candidates, currRes, i, currSum, target, res)
            currSum -= num
            currRes = currRes[:len(currRes)-1]
        }
    }
}