package septemberleetcodechallange

func maxProduct(nums []int) int {
	reverseOfNums := make([]int, len(nums))
	for i, j := 0, len(nums)-1; j >= 0; i, j = i+1, j-1 {
		reverseOfNums[i] = nums[j]
	}

	for i, _ := range nums {
		if i == 0 {
			continue
		}
		if nums[i-1] != 0 {
			nums[i] *= nums[i-1]
		}
		if reverseOfNums[i-1] != 0 {
			reverseOfNums[i] *= reverseOfNums[i-1]
		}
	}
	//fmt.Println(nums)
	//fmt.Println(reverseOfNums)
	nums = append(nums, reverseOfNums...)
	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}
