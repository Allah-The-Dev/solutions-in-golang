package main

func threeConsecutiveOdds(arr []int) bool {
	for i, _ := range arr {
		if i < len(arr) && i+1 < len(arr) && i+2 < len(arr) {
			if arr[i]%2 != 0 && arr[i+1]%2 != 0 && arr[i+2]%2 != 0 {
				return true
			}
		}
	}
	return false
}
