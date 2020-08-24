package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	Rects        [][]int
	NumPts       int
	RectCumCount []int //this will hold no of points upto that rect //it'll be a cumulative sum
	UsedRand     []bool
}

//ConstructorR ...
func ConstructorR(rects [][]int) Solution {
	cumSumArr := make([]int, len(rects))
	totalNumPts := 0
	for i, rect := range rects {
		totalNumPts += (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
		cumSumArr[i] = totalNumPts
	}
	usedRand := make([]bool, totalNumPts)
	sol := Solution{rects, totalNumPts, cumSumArr, usedRand}
	return sol
}

//Pick ...
func (sol *Solution) Pick() []int {
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(sol.NumPts - 1)
	fmt.Println("randIdx : ", randIdx)

	//find rectangle no where this point belongs
	l, r := 0, len(sol.Rects)-1
	for l < r {
		mid := (l + r) / 2
		if sol.RectCumCount[mid] <= randIdx {
			l = mid + 1
		} else {
			r = mid
		}
	}
	//fmt.Println("l => ", l)
	rect := sol.Rects[l]
	xPts := rect[2] - rect[0] + 1
	yPts := rect[3] - rect[1] + 1
	ptsInRect := xPts * yPts
	startPt := sol.RectCumCount[l] - ptsInRect
	offsetOfReqPt := randIdx - startPt
	//fmt.Printf("xPts=>%d, y=>%d, ptsInRect=>%d, startPt=>%d, offsetOfReqPt=>%d\n", xPts, yPts, ptsInRect, startPt, offsetOfReqPt)
	//fmt.Printf("x=>%d, y=>%d\n", (offsetOfReqPt % xPts), (offsetOfReqPt / xPts))
	x := rect[0] + offsetOfReqPt%xPts
	y := rect[1] + offsetOfReqPt/xPts
	return []int{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */

// func main() {
// 	obj := Constructor([][]int{{-2, -2, -1, -1}, {1, 0, 3, 0}})
// 	fmt.Printf("total points => %d\n", obj.NumPts)
// 	param1 := obj.Pick()
// 	fmt.Printf("%#v\n", param1)
// 	param1 = obj.Pick()
// 	fmt.Printf("%#v\n", param1)
// 	param1 = obj.Pick()
// 	fmt.Printf("%#v\n", param1)
// }
