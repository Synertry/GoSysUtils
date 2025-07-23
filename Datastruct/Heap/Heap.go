// I cooked this up 3 years ago on the 24th of October, 2022
// I got 0 use for it currently, so testing will be minimal

package Heap

type Heap interface {
	insert(int)
	heapifyUp()
	poll() int
	heapifyDown()
	peek() int
	size() int
	isEmpty() bool
	print()
}

func insert(h Heap, val int) {
	h.insert(val)
}

func heapifyUp(h Heap) {
	h.heapifyUp()
}

func heapifyDown(h Heap) {
	h.heapifyDown()
}

func peek(h Heap) int {
	return h.peek()
}

func size(h Heap) int {
	return h.size()
}

func isEmpty(h Heap) bool {
	return h.isEmpty()
}

func print(h Heap) {
	h.print()
}
