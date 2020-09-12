package septemberleetcodechallange

func PartitionLabels(S string) []int {
	chEndCache := map[rune]int{}
	//store the max index of each character in cache
	for i, ch := range S {
		if value, ok := chEndCache[ch]; ok {
			if i > value {
				chEndCache[ch] = i
			}
		} else {
			chEndCache[ch] = i
		}
	}
	//now use 2 pointer approach to find the least length string
	//where each character falls only once
	l, r := 0, 0
	res := []int{}
	for i, ch := range S {
		if chEndCache[ch] > r {
			r = chEndCache[ch]
		}
		if r == i {
			res = append(res, r-l+1)
			l = i + 1
		}
	}
	return res
}
