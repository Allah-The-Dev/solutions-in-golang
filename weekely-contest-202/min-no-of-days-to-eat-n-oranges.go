package main

import "container/list"

func minDays(n int) int {
	days := 0
	//solve this problem with tree approach + BFS
	queueOfRemains := list.New()
	queueOfRemains.PushBack(n)
	visited := make(map[int]struct{})
	for queueOfRemains.Len() != 0 {
		//every day counts for a single day
		days++
		queueLen := queueOfRemains.Len()
		for i := 1; i <= queueLen; i++ {
			node := queueOfRemains.Front()
			nodeVal := node.Value.(int)
			if nodeVal == 1 {
				return days
			}
			if _, ok := visited[nodeVal]; !ok {
				if nodeVal%3 == 0 {
					queueOfRemains.PushBack(nodeVal / 3)
				}
				if nodeVal%2 == 0 {
					queueOfRemains.PushBack(nodeVal / 2)
				}
				queueOfRemains.PushBack(nodeVal - 1)
			}
			visited[nodeVal] = struct{}{}
			queueOfRemains.Remove(node)
		}
	}
	return days
}
