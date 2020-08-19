package main

func distributeCandies(candies int, numPeople int) []int {
	candyDistribution := make([]int, numPeople)
	i := 0
	candy := 0
	for candies != 0 && candies > candy {
		candy++
		candies -= candy
		candyDistribution[i] += candy
		i++
		if i == numPeople {
			i = 0
		}
	}
	if candies != 0 {
		candyDistribution[i] += candies
	}
	return candyDistribution
}

// func main() {
// 	fmt.Println(distributeCandies(11, 3))
// }
