package main

//WordDictionary ...
type WordDictionary struct {
	root *Node
}

//Node ...
type Node struct {
	parent   *Node
	children map[rune]*Node
	value    rune
}

//Constructor ... Initialize your data structure here. */
func TrieConstructor() *WordDictionary {
	return &WordDictionary{
		root: &Node{
			parent:   nil,
			children: make(map[rune]*Node),
			value:    0,
		},
	}
}

//AddWord ...  Adds a word into the data structure. */
func (wd *WordDictionary) AddWord(word string) {
	currentNode := wd.root
	for _, ch := range word {
		if node, ok := currentNode.children[ch]; ok {
			currentNode = node
		} else {
			currentNode.children[ch] = &Node{
				parent:   currentNode,
				children: make(map[rune]*Node),
				value:    ch,
			}
			currentNode = currentNode.children[ch]
		}
	}
}

// Search ... Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (wd *WordDictionary) Search(word string) (*Node, bool) {
	currentNode := wd.root
	for _, ch := range word {
		if _, ok := currentNode.children[ch]; ok {
			currentNode = currentNode.children[ch]
		} else {
			return nil, false
		}
	}
	return currentNode, true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// func main() {
// 	obj := Constructor()
// 	obj.AddWord("oyooyooyo")
// 	fmt.Printf("%#v\n", *obj)
// 	fmt.Println(obj.Search("oyoo"))
// 	fmt.Println(obj.Search("oy"))
// }
