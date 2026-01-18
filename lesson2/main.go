package main

type CircularQueue struct {
	values   []int
	length   int
	capacity int
	readIdx  int
	writeIdx int
}

func NewCircularQueue(size int) CircularQueue {
	q := CircularQueue{
		values:   make([]int, size, size),
		length:   0,
		capacity: size,
		writeIdx: 0,
		readIdx:  0,
	}
	return q
}

func (q *CircularQueue) Push(value int) bool {
	if q.Full() {
		return false
	}

	q.values[q.writeIdx] = value

	q.writeIdx++
	if q.writeIdx > q.capacity-1 {
		q.writeIdx = 0
	}

	q.length++

	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.values[q.readIdx] = 0

	q.readIdx++
	if q.readIdx > q.capacity-1 {
		q.readIdx = 0
	}

	q.length--

	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.readIdx]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	idx := (q.writeIdx - 1 + q.capacity) % q.capacity // it's too tricky, I broke my brain
	return q.values[idx]
}

func (q *CircularQueue) Empty() bool {
	return q.length <= 0
}

func (q *CircularQueue) Full() bool {
	return q.length == q.capacity
}

func main() {
	q := NewCircularQueue(3)
	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Push(4) // nope
	q.Front() // 1
	q.Back() // 3

	q.Pop()
	q.Pop()
	q.Pop()

	q.Empty() // true 
}
