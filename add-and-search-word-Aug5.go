package main

import (
	"fmt"
	"solutions-in-golang/model"
	"solutions-in-golang/septemberleetcodechallange"
)

// sep - 8
func main() {
	tree := &model.TreeNode{Val: 1}
	tree.Left = &model.TreeNode{Val: 0, Left: &model.TreeNode{Val: 0}, Right: &model.TreeNode{Val: 1}}
	tree.Right = &model.TreeNode{Val: 1, Left: &model.TreeNode{Val: 0}, Right: &model.TreeNode{Val: 1}}
	fmt.Println(septemberleetcodechallange.SumRootToLeaf(tree))
}

// sep - 7
// func main() {
// 	fmt.Println(septemberleetcodechallange.WordPattern("abba", "dog dog dog dog"))
// }

// sep - 6
// func main() {
// 	A := [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}
// 	B := [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}
// 	fmt.Println(septemberleetcodechallange.LargestOverlap(A, B))
// }

// //sep - 3
// func main() {
// 	fmt.Println(septemberleetcodechallange.PartitionLabels("ababcbacadefegdehijhklij"))
// }

// import (
// 	"fmt"
// 	"solutions-in-golang/septemberleetcodechallange"
// )

// func main() {
// 	fmt.Println(septemberleetcodechallange.RepeatedSubstringPattern("abcabcab"))
// }

// //sep - 1 :largest time
// func main() {
// 	fmt.Println(septemberleetcodechallange.LargestTimeFromDigits([]int{1, 2, 3, 4}))
// }

// import (
// 	"fmt"
// 	"solutions-in-golang/pancakesorting"
// )

// import (
// 	"fmt"
// 	"solutions-in-golang/addandsearchword"
// )

// func main() {
// 	obj := addandsearchword.Constructor()
// 	obj.AddWord("bca")
// 	obj.AddWord("mno")
// 	obj.AddWord("xyz")
// 	fmt.Println(obj.Search("..a"))
// 	fmt.Println(obj.Search("..."))
// 	fmt.Println(obj.Search("x..z"))
// 	fmt.Println(obj.Search("..p"))
// 	fmt.Println(obj.Search("mno"))
// 	obj.AddWord("nntiantzzodd")
// 	fmt.Println(obj.Search(".nti.nt..o.."))
// }

// func main() {
// 	fmt.Println(pancakesorting.PancakeSort([]int{3, 2, 4, 1}))
// }
