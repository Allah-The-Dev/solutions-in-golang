package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	leftMaxProfit := make([]int, len(prices))
	leftMin := prices[0]
	for i, num := range prices {
		if i == 0 {
			continue
		}
		if num-leftMin > leftMaxProfit[i-1] {
			leftMaxProfit[i] = num - leftMin
		} else {
			leftMaxProfit[i] = leftMaxProfit[i-1]
		}
		if leftMin > num {
			leftMin = num
		}
	}
	rightMaxProfit := make([]int, len(prices))
	rightMax := prices[len(prices)-1]
	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 {
			continue
		}
		if rightMax-prices[i] > rightMaxProfit[i+1] {
			rightMaxProfit[i] = rightMax - prices[i]
		} else {
			rightMaxProfit[i] = rightMaxProfit[i+1]
		}
		if rightMax < prices[i] {
			rightMax = prices[i]
		}
	}
	fmt.Println(leftMaxProfit, rightMaxProfit)
	totalProfit := leftMaxProfit[0] + rightMaxProfit[0]
	for i := range leftMaxProfit {
		if totalProfit < leftMaxProfit[i]+rightMaxProfit[i] {
			totalProfit = leftMaxProfit[i] + rightMaxProfit[i]
		}
	}
	return totalProfit
}

// func main() {
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
// }
