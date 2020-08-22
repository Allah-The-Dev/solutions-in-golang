package main

func minOperations(n int) int {
	midIndex := n / 2
	if n%2 != 0 {
		return midIndex * (midIndex + 1)
	}
	return midIndex * midIndex
}
