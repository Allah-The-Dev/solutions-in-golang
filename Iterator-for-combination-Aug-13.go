package main

import (
	"fmt"
)

//CombinationIterator ...
type CombinationIterator struct {
	Characters []rune
	Queue      []string
}

//Constructor ...
func Constructor(characters string, combinationLength int) CombinationIterator {
	combinationIterator := CombinationIterator{[]rune(characters), []string{}}
	combinationIterator.GetAllCombination(0, combinationLength, "")
	return combinationIterator
}

//GetAllCombination ...
func (ci *CombinationIterator) GetAllCombination(start, length int, txt string) {
	if length == 0 {
		ci.Queue = append(ci.Queue, txt)
		return
	}
	fmt.Printf("start = %d, length = %d, txt = %s\n", start, len(ci.Characters)-length, txt)
	for i := start; i <= len(ci.Characters)-length; i++ {
		txt = txt + string(ci.Characters[i])
		fmt.Printf("method call is GetAllCombination(%d, %d, %s)\n", i+1, length-1, txt)
		ci.GetAllCombination(i+1, length-1, txt)
		txt = txt[:len(txt)-1]
	}
}

//Next ...
func (ci *CombinationIterator) Next() string {
	next := ci.Queue[0]
	ci.Queue = ci.Queue[1:]
	return next
}

//HasNext ...
func (ci *CombinationIterator) HasNext() bool {
	return len(ci.Queue) != 0
}

// func main() {
// 	obj := Constructor("abcde", 3)
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.HasNext())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.Next())
// 	fmt.Println(obj.HasNext())
// }
