func maxAreaOfIsland(grid [][]int) int {
    res := 0
    for i, row := range grid {
        for j, num := range row {
            if num == 1 {
                res = max(res, dfs(i, j, grid))
            }
        }
    }
    return res
}

func dfs(i, j int, grid [][]int) int {
    if isNotValidIdx(i, j, grid) || grid[i][j] == 0 {
        return 0
    }
    grid[i][j] = 0
    
    lArea := dfs(i, j-1, grid)
    rArea := dfs(i, j+1, grid)
    tArea := dfs(i-1, j, grid)
    bArea := dfs(i+1, j, grid)
    
    return 1+lArea+rArea+tArea+bArea
}

func isNotValidIdx(i, j int, grid [][]int) bool {
    return i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0])
}

func max(a, b int) int {
    if a > b { return a }
    return b
}