package task3

import "container/heap"

type intHeap struct {
	values []int
	isMax  bool
}

func (h intHeap) Len() int { return len(h.values) }
func (h intHeap) Less(i, j int) bool {
	if h.isMax {
		return h.values[i] > h.values[j]
	}
	return h.values[i] < h.values[j]
}
func (h intHeap) Swap(i, j int)   { h.values[i], h.values[j] = h.values[j], h.values[i] }
func (h *intHeap) Push(value any) { h.values = append(h.values, value.(int)) }
func (h *intHeap) Pop() any {
	last := len(h.values) - 1
	value := h.values[last]
	h.values = h.values[:last]
	return value
}

func kthSmallest(data []int, k int) (int, bool) {
	return kthElement(data, k, false)
}

func kthLargest(data []int, k int) (int, bool) {
	return kthElement(data, k, true)
}

func kthElement(data []int, k int, largest bool) (int, bool) {
	if k < 1 || k > len(data) {
		return 0, false
	}
	h := &intHeap{values: append([]int(nil), data[:k]...), isMax: !largest}
	heap.Init(h)
	for _, value := range data[k:] {
		if (!largest && value < h.values[0]) || (largest && value > h.values[0]) {
			h.values[0] = value
			heap.Fix(h, 0)
		}
	}
	return h.values[0], true
}
