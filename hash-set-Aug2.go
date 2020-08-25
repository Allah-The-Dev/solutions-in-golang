package main

//MyHashSet ...
type MyHashSet struct {
	Buckets [][]int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Buckets: make([][]int, 15000),
	}
}

func getHashKey(key int) int {
	return key % 15000
}

func (myHashSet *MyHashSet) Add(key int) {
	//applying chaining mechanism
	hashKey := getHashKey(key)
	if innerChain := myHashSet.Buckets[hashKey]; innerChain == nil {
		innerChain = []int{key}
		myHashSet.Buckets[hashKey] = innerChain
	} else {
		innerChainContainsKey := false
		for _, keyInChain := range innerChain {
			if keyInChain == key {
				innerChainContainsKey = true
				break
			}
		}
		if !innerChainContainsKey {
			innerChain = append(innerChain, key)
			myHashSet.Buckets[hashKey] = innerChain
		}
	}
}

func (myHashSet *MyHashSet) Remove(key int) {
	hashKey := getHashKey(key)
	if innerChain := myHashSet.Buckets[hashKey]; innerChain != nil {
		for i, keyInChain := range innerChain {
			if keyInChain == key {
				innerChain2 := innerChain[:i]
				innerChain = append(innerChain2, innerChain[i+1:]...)
				myHashSet.Buckets[hashKey] = innerChain
				break
			}
		}
	}
}

/** Returns true if this set contains the specified element */
func (myHashSet *MyHashSet) Contains(key int) bool {
	hashKey := getHashKey(key)
	if innerChain := myHashSet.Buckets[hashKey]; innerChain != nil {
		for _, keyInChain := range innerChain {
			if keyInChain == key {
				return true
			}
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
