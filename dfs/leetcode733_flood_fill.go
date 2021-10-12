func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image[sr][sc] == newColor { return image }
    fillColor(image, sr, sc, image[sr][sc], newColor)
    return image
}

func fillColor(image [][]int, sr, sc, color, newColor int) {
    if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != color {
        return
    }
    image[sr][sc] = newColor
    fillColor(image, sr-1, sc, color, newColor)
    fillColor(image, sr+1, sc, color, newColor)
    fillColor(image, sr, sc-1, color, newColor)
    fillColor(image, sr, sc+1, color, newColor)
}