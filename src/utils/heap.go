package utils

import (
	"container/heap"
)

// Item represents an item in the priority queue.
type Item struct {
	Value map[string]interface{}
	Index int // This is purely for user's custom use and is ignored by the heap
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	items    []*Item
	lessFunc func(a, b map[string]interface{}) int
}

// NewPriorityQueue creates a new PriorityQueue with the given comparison function.
func NewPriorityQueue(sorter func(a, b map[string]interface{}) int) *PriorityQueue {
	pq := &PriorityQueue{
		items:    make([]*Item, 0),
		lessFunc: sorter,
	}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue) Len() int { return len(pq.items) }

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.lessFunc(pq.items[i].Value, pq.items[j].Value) < 0
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	pq.items = old[0 : n-1]
	return item
}

// PushItem wraps heap.Push to make it easier to use
func (pq *PriorityQueue) PushItem(item *Item) {
	heap.Push(pq, item)
}

// PopItem wraps heap.Pop to make it easier to use
func (pq *PriorityQueue) PopItem() *Item {
	return heap.Pop(pq).(*Item)
}
