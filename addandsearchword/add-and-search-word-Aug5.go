package addandsearchword

import (
	"container/list"
)

//WordDictionary ...
type WordDictionary struct {
	value    rune
	children map[rune]*WordDictionary
	isWord   bool
}

//Constructor ... Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{' ', map[rune]*WordDictionary{}, false}
}

//AddWord ... Adds a word into the data structure. */
func (wd *WordDictionary) AddWord(word string) {
	currentNode := wd
	for _, ch := range word {
		if currentNode.children[ch] != nil {
			currentNode = currentNode.children[ch]
		} else {
			currentNode.children[ch] = &WordDictionary{ch, map[rune]*WordDictionary{}, false}
			currentNode = currentNode.children[ch]
		}
	}
	currentNode.isWord = true
}

//Search ... Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (wd *WordDictionary) Search(word string) bool {
	currentNode := wd
	queue := list.New()
	ch := rune(word[0])
	if ch == '.' {
		for _, child := range currentNode.children {
			if child.isWord && len(word) == 1 {
				return true
			}
			queue.PushBack(child)
		}
	} else {
		if child, ok := currentNode.children[ch]; ok {
			if child.isWord && len(word) == 1 && ch == child.value {
				return true
			}
			queue.PushBack(child)
		} else {
			return false
		}
	}
	for i, ch := range word {
		if i == 0 {
			continue
		}
		currentQLen := queue.Len()
		if ch == '.' {
			for j := 1; j <= currentQLen; j++ {
				queueNode := queue.Front().Value.(*WordDictionary)
				for _, nodeFromQ := range queueNode.children {
					if nodeFromQ.isWord && i == len(word)-1 {
						return true
					}
					queue.PushBack(nodeFromQ)
				}
				queue.Remove(queue.Front())
			}
		} else {
			nodeFound := false
			for j := 1; j <= currentQLen && !nodeFound; j++ {
				queueNode := queue.Front().Value.(*WordDictionary)
				queue.Remove(queue.Front())
				for _, nodeFromQueue := range queueNode.children {
					if ch == nodeFromQueue.value {
						if nodeFromQueue.isWord && i == len(word)-1 {
							return true
						}
						queue.Init()
						queue.PushBack(nodeFromQueue)
						nodeFound = true
						break
					}
				}
			}
			if !nodeFound {
				return false
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
