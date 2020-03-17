package main

import (
	"container/heap"
	"fmt"
)

// An StringHeap is a min-heap of Strings.
type StringHeap []string

func (h StringHeap) Len() int               { return len(h) }
func (h StringHeap) Less(i int, j int) bool { return h[i] < h[j] }
func (h StringHeap) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }

func (h *StringHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(string))
}

func (h *StringHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several Strings Stringo an StringHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &StringHeap{"A", "B", "C"}
	heap.Init(h)
	heap.Push(h, "D")
	fmt.Printf("Minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
