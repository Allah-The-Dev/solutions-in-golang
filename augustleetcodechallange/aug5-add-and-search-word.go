package augustleetcodechallange

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

func searchRecursively(word []rune, wd *WordDictionary) bool {
	if wd == nil {
		return false
	}
	if len(word) == 1 {
		if word[0] == '.' {
			for _, child := range wd.children {
				if child.isWord {
					return true
				}
			}
			return false
		} else if child, ok := wd.children[word[0]]; ok && child.isWord {
			return true
		} else {
			return false
		}
	} else {
		if word[0] == '.' {
			word = word[1:]
			for _, child := range wd.children {
				if searchRecursively(word, child) {
					return true
				}
			}
		} else {
			firstRune := word[0]
			if child, ok := wd.children[firstRune]; ok {
				word = word[1:]
				if searchRecursively(word, child) {
					return true
				}
			}
		}
	}
	return false
}

//Search ... Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (wd *WordDictionary) Search(word string) bool {
	return searchRecursively([]rune(word), wd)
}
