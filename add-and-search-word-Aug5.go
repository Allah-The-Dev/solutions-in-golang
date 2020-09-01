package main

// import (
// 	"fmt"
// 	"solutions-in-golang/pancakesorting"
// )

import (
	"fmt"
	"solutions-in-golang/addandsearchword"
)

func main() {
	obj := addandsearchword.Constructor()
	// obj.AddWord("bca")
	// obj.AddWord("mno")
	// obj.AddWord("xyz")
	// fmt.Println(obj.Search("..a"))
	// fmt.Println(obj.Search("..."))
	// fmt.Println(obj.Search("x..z"))
	// fmt.Println(obj.Search("..p"))
	// fmt.Println(obj.Search("mno"))
	obj.AddWord("nntiantzzodd")
	fmt.Println(obj.Search(".nti.nt..o.."))
}

// func main() {
// 	fmt.Println(pancakesorting.PancakeSort([]int{3, 2, 4, 1}))
// }
