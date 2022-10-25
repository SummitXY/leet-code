package sort

import (
	"container/heap"
)

func HeapSortTopK(a []int, k int) int {

	heapLen := len(a) - k + 1
	hp := &Tree{}
	heap.Init(hp)
	for _, v := range a {
		heap.Push(hp, v)
		if hp.Len() > heapLen {
			heap.Pop(hp)
		}
	}

	res := heap.Pop(hp)
	return res.(int)
}

type Tree []int

func (t Tree) Len() int {
	return len(t)
}

func (t Tree) Less(i, j int) bool {
	return t[i] > t[j]
}

func (t *Tree) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *Tree) Push(v interface{}) {
	*t = append(*t, v.(int))
}

func (t *Tree) Pop() interface{} {
	n := t.Len()
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}

