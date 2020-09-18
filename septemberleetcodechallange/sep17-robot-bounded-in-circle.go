package septemberleetcodechallange

func isRobotBounded(instructions string) bool {
	direction := [2]int{0, 1}
	start := [2]int{0, 0}

	for _, instruction := range instructions {
		if instruction == 'G' {
			start[0] += direction[0]
			start[1] += direction[1]
		} else if instruction == 'R' {
			direction[0], direction[1] = direction[1], -direction[0]
		} else if instruction == 'L' {
			direction[0], direction[1] = -direction[1], direction[0]
		}
	}

	return start == [2]int{0, 0} || direction != [2]int{0, 1}
}
