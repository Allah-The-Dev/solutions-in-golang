package pancakesorting

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func PancakeSort(A []int) []int {
	resultArr := []int{}
	lastIndexToConsider := len(A)
	for i := 0; i < len(A); i++ {
		//find the max upto lastindex
		max := 0
		indexOfMax := 0
		for j := 0; j < lastIndexToConsider; j++ {
			if A[j] > max {
				max = A[j]
				indexOfMax = j
			}
		}
		//do 2 times reverse
		//1. from 0 -> indexOfMax
		reversed := reverse(A[:indexOfMax+1])
		A = append(reversed, A[indexOfMax+1:]...)
		resultArr = append(resultArr, indexOfMax+1)
		//2. from 0 -> last
		reversed = reverse(A[:lastIndexToConsider])
		A = append(reversed, A[lastIndexToConsider:]...)
		resultArr = append(resultArr, lastIndexToConsider)
		lastIndexToConsider--
	}
	return resultArr
}
