package Heap

import "fmt"

type maxHeap struct {
	list []int
}

func (h *maxHeap) insert(val int) {
	h.list = append(h.list, val)
	h.heapifyUp()
}

func (h *maxHeap) heapifyUp() {
	idx := len(h.list) - 1
	for h.list[idx] > h.list[(idx-1)/2] {
		h.list[idx], h.list[(idx-1)/2] = h.list[(idx-1)/2], h.list[idx]
		idx = (idx - 1) / 2
	}
}

func (h *maxHeap) poll() int {
	max := h.list[0]
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	h.heapifyDown()
	return max
}

func (h *maxHeap) heapifyDown() {
	idx := 0
	for idx < len(h.list) {
		leftIdx := 2*idx + 1
		rightIdx := 2*idx + 2
		if leftIdx >= len(h.list) {
			break
		}
		maxIdx := leftIdx
		if rightIdx < len(h.list) && h.list[rightIdx] > h.list[leftIdx] {
			maxIdx = rightIdx
		}
		if h.list[idx] > h.list[maxIdx] {
			break
		}
		h.list[idx], h.list[maxIdx] = h.list[maxIdx], h.list[idx]
		idx = maxIdx
	}
}

func (h *maxHeap) peek() int {
	return h.list[0]
}

func (h *maxHeap) size() int {
	return len(h.list)
}

func (h *maxHeap) isEmpty() bool {
	return len(h.list) == 0
}

func (h *maxHeap) print() {
	fmt.Println(h.list)
}

func NewMaxHeap() *maxHeap {
	return &maxHeap{}
}
