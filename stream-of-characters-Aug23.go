package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//StreamChecker ...
type StreamChecker struct {
	next           map[rune]*StreamChecker
	isWord         bool
	searchStrSoFar strings.Builder
}

//Constructor ...
func Constructor(words []string) StreamChecker {
	rootNode := StreamChecker{make(map[rune]*StreamChecker), false, strings.Builder{}}
	for _, word := range words {
		rootNode.Append(word)
	}
	return rootNode
}

//Append ...
func (strChecker *StreamChecker) Append(word string) {
	node := strChecker
	for len(word) > 0 {
		//run = rune
		run, runSize := utf8.DecodeLastRuneInString(word)
		if next, ok := node.next[run]; ok {
			node = next
		} else {
			node.next[run] = &StreamChecker{next: make(map[rune]*StreamChecker), isWord: false}
			node = node.next[run]
		}
		word = word[:len(word)-runSize]
	}
	node.isWord = true
}

//Query ...
func (strChecker *StreamChecker) Query(letter byte) bool {
	node := strChecker
	//keep appending the coming byte in existing search string
	node.searchStrSoFar.WriteByte(letter)
	wordToSearch := node.searchStrSoFar.String()
	//fmt.Println(wordToSearch)
	for len(wordToSearch) > 0 {
		run, runSize := utf8.DecodeLastRuneInString(wordToSearch)
		node = node.next[run]
		if node != nil && node.isWord {
			return true
		}
		if node == nil {
			return false
		}
		wordToSearch = wordToSearch[:len(wordToSearch)-runSize]
	}
	return false
}

func main() {
	obj := Constructor([]string{"baa", "aa", "aaaa", "abbbb", "aba"})
	fmt.Println(obj.Query(byte(97)))
	fmt.Println(obj.Query(byte(97)))
	fmt.Println(obj.Query(byte(97)))
	fmt.Println(obj.Query(byte(100)))
	fmt.Println(obj.Query(byte(101)))
	fmt.Println(obj.Query(byte(102)))
}
