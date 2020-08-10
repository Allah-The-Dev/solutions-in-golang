package main

type coordinate struct {
	x, y int
}

func orangesRotting(grid [][]int) int {

	queue := []coordinate{}
	timeRequired := 0
	totalOranges := 0
	rotten := 0

	for i, row := range grid {
		for j, val := range row {
			if val == 2 {
				queue = append(queue, coordinate{i, j})
				rotten++
			}
			if val == 1 || val == 2 {
				totalOranges++
			}
		}
	}

	if totalOranges == rotten {
		return 0
	}

	for len(queue) != 0 {
		queueSizeToIterate := len(queue)
		//for every iteration add 1
		if rotten == totalOranges {
			return timeRequired
		}

		timeRequired++
		//iterate over all the items in queue
		for i := 0; i < queueSizeToIterate; i++ {
			coordinates := queue[0]
			queue = queue[1:]
			//check left
			if (coordinates.y-1) >= 0 && grid[coordinates.x][coordinates.y-1] == 1 {
				grid[coordinates.x][coordinates.y-1] = 2
				queue = append(queue, coordinate{coordinates.x, coordinates.y - 1})
				rotten++
			}
			//check top
			if (coordinates.x-1) >= 0 && grid[coordinates.x-1][coordinates.y] == 1 {
				grid[coordinates.x-1][coordinates.y] = 2
				queue = append(queue, coordinate{coordinates.x - 1, coordinates.y})
				rotten++
			}
			//check right
			if (coordinates.y+1 <= len(grid[0])-1) && grid[coordinates.x][coordinates.y+1] == 1 {
				grid[coordinates.x][coordinates.y+1] = 2
				queue = append(queue, coordinate{coordinates.x, coordinates.y + 1})
				rotten++
			}
			//check bottom
			if (coordinates.x+1 <= len(grid)-1) && grid[coordinates.x+1][coordinates.y] == 1 {
				grid[coordinates.x+1][coordinates.y] = 2
				queue = append(queue, coordinate{coordinates.x + 1, coordinates.y})
				rotten++
			}
		}
	}

	return -1
}

// func main() {
// 	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
// 	fmt.Println(orangesRotting(grid))
// }
