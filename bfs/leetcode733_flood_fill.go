// do flood fill with new color 
// approach : using BFS
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    rowL, colL := len(image), len(image[0])
    queue := [][]int{{sr, sc}}
    for len(queue) > 0 {
        currQueueLen := len(queue)
        for i := 1; i <= currQueueLen; i++ {
            first := queue[0]
            row, col := first[0], first[1]
            l, r, t, b := []int{row, col-1}, []int{row, col+1}, []int{row-1, col}, []int{row+1, col}
            if l[1] >= 0 && image[l[0]][l[1]] == image[row][col] && image[l[0]][l[1]] != newColor {
                queue = append(queue, l)
            }
            if r[1] < colL && image[r[0]][r[1]] == image[row][col] && image[r[0]][r[1]] != newColor {
                queue = append(queue, r)
            }
            if t[0] >= 0 && image[t[0]][t[1]] == image[row][col] && image[t[0]][t[1]] != newColor {
                queue = append(queue, t)
            }
            if b[0] < rowL && image[b[0]][b[1]] == image[row][col] && image[b[0]][b[1]] != newColor {
                queue = append(queue, b)
            }
            image[row][col] = newColor
            //fmt.Println(queue)
            queue = queue[1:]
        }
    }
    return image
}