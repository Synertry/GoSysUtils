package Heap

import "fmt"

type minHeap struct {
	list []int
}

func (h *minHeap) insert(val int) {
	h.list = append(h.list, val)
	h.heapifyUp()
}

func (h *minHeap) heapifyUp() {
	idx := len(h.list) - 1
	for h.list[idx] < h.list[(idx-1)/2] {
		h.list[idx], h.list[(idx-1)/2] = h.list[(idx-1)/2], h.list[idx]
		idx = (idx - 1) / 2
	}
}

func (h *minHeap) poll() int {
	min := h.list[0]
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	h.heapifyDown()
	return min
}

func (h *minHeap) heapifyDown() {
	idx := 0
	for idx < len(h.list) {
		leftIdx := 2*idx + 1
		rightIdx := 2*idx + 2
		if leftIdx >= len(h.list) {
			break
		}
		minIdx := leftIdx
		if rightIdx < len(h.list) && h.list[rightIdx] < h.list[leftIdx] {
			minIdx = rightIdx
		}
		if h.list[idx] < h.list[minIdx] {
			break
		}
		h.list[idx], h.list[minIdx] = h.list[minIdx], h.list[idx]
		idx = minIdx
	}
}

func (h *minHeap) peek() int {
	return h.list[0]
}

func (h *minHeap) size() int {
	return len(h.list)
}

func (h *minHeap) isEmpty() bool {
	return len(h.list) == 0
}

func (h *minHeap) print() {
	fmt.Println(h.list)
}

func NewMinHeap() *minHeap {
	return &minHeap{}
}
