package main

type node struct {
	id       int
	priority int
	object   interface{}
}

type Heap struct {
	list []node
}

func NewHeap() *Heap {
	return &Heap{
		list: make([]node, 0),
	}
}

func (h *Heap) Add(id int, priority int, object interface{}) {
	h.list = append(h.list, node{
		id:       id,
		priority: priority,
		object:   object,
	})

	i := len(h.list) - 1
	parent := (i - 1) / 2

	for i > 0 && h.list[parent].priority < h.list[i].priority {
		h.list[i], h.list[parent] = h.list[parent], h.list[i]

		i = parent
		parent = (i - 1) / 2
	}
}

func (h *Heap) SetPriority(nodeId int, newPriority int) {
	for i := 0; i < len(h.list); i++ {
		if h.list[i].id == nodeId {
			h.list[i].priority = newPriority
		}
	}
	h.heapify(0)
}

func (h *Heap) Get() interface{} {
	if len(h.list) <= 0 {
		return nil
	}

	object := h.list[0].object
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	h.heapify(0)

	return object
}

func (h *Heap) heapify(i int) {
	var leftChild int
	var rightChild int
	var largestChild int

	for {
		leftChild = 2*i + 1
		rightChild = 2*i + 2
		largestChild = i

		if leftChild < len(h.list) && h.list[leftChild].priority > h.list[largestChild].priority {
			largestChild = leftChild
		}

		if rightChild < len(h.list) && h.list[rightChild].priority > h.list[largestChild].priority {
			largestChild = rightChild
		}

		if largestChild == i {
			break
		}

		h.list[i], h.list[largestChild] = h.list[largestChild], h.list[i]
	}
}
