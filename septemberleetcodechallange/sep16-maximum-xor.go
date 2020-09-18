package septemberleetcodechallange

import (
	"fmt"
	"strconv"
	"strings"
)

type TrieNode struct {
	children map[rune]*TrieNode
}

func NewTrie() *TrieNode {
	return &TrieNode{
		children: map[rune]*TrieNode{},
	}
}

func (root *TrieNode) InsertBitsInTrie(num int) {
	binOfNum := fmt.Sprintf("%032b", num)
	currentNode := root
	for _, bit := range binOfNum {
		if _, ok := currentNode.children[bit]; !ok {
			currentNode.children[bit] = &TrieNode{
				children: map[rune]*TrieNode{},
			}
		}
		currentNode = currentNode.children[bit]
	}
}

func (root *TrieNode) MaxXORWithCurrentNo(num int) int {
	binOfNum := fmt.Sprintf("%032b", num)
	currentNode := root
	var oppoBit rune
	var maxXOR strings.Builder
	for _, bit := range binOfNum {
		if bit == '0' {
			oppoBit = '1'
		} else if bit == '1' {
			oppoBit = '0'
		}
		if value, ok := currentNode.children[oppoBit]; ok {
			maxXOR.WriteRune(oppoBit)
			currentNode = value
		} else {
			maxXOR.WriteRune(bit)
			currentNode = currentNode.children[bit]
		}
	}
	decimalOfMaxXORFoundInTrie, _ := strconv.ParseInt(maxXOR.String(), 2, 32)
	return int(decimalOfMaxXORFoundInTrie) ^ num
}

func FindMaximumXOR(nums []int) int {
	trieRoot := NewTrie()
	for _, num := range nums {
		trieRoot.InsertBitsInTrie(num)
	}
	maxXOR := 0
	for _, num := range nums {
		if currentMaxXOR := trieRoot.MaxXORWithCurrentNo(num); currentMaxXOR > maxXOR {
			maxXOR = currentMaxXOR
		}
	}
	return maxXOR
}
