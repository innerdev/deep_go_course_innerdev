package main

import "fmt"

type CircularQueue struct {
	values    []int
	readIndex int
	length    int
}

func NewCircularQueue(size int) CircularQueue {
	if size <= 0 {
		panic("Size should be > 0")
	}

	q := CircularQueue{
		values: make([]int, size),
	}

	return q
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	q.values[(q.readIndex+q.length)%cap(q.values)] = value
	q.length++

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.readIndex = (q.readIndex + 1) % cap(q.values)
	q.length--

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}

	return q.values[q.readIndex]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}

	idx := (q.length + q.readIndex - 1) % cap(q.values)

	return q.values[idx]
}

func (q *CircularQueue) Empty() bool {
	return q.length == 0
}

func (q *CircularQueue) Full() bool {
	return q.length == cap(q.values)
}

func (q *CircularQueue) Debug() {
	fmt.Println(q.values, cap(q.values), q.readIndex, q.length)
}

func main() {
	q := NewCircularQueue(3)
	q.Debug()

	q.Push(1)
	q.Debug()

	q.Push(2)
	q.Debug()

	q.Push(3)
	q.Debug()

	q.Push(4) // nope
	q.Debug()

	fmt.Println(q.Front()) // 1
	q.Debug()

	fmt.Println(q.Back()) // 3
	q.Debug()

	q.Pop()
	q.Debug()

	q.Pop()
	q.Debug()

	q.Pop()
	q.Debug()

	q.Empty() // true
	q.Debug()
}
