package main

import (
	"fmt"
)

type OrderedMap struct {
	root  *node
	size int
}

type node struct {
	left  *node
	right *node
	key int
	value int
}

// CONSTRUCT ------------------------------------------------

func NewOrderedMap() OrderedMap {
	return OrderedMap{
		root: nil,
		size: 0,
	}
}

// INSERT ---------------------------------------------------

func (t *OrderedMap) Insert (key, value int) {
	t.insertRecursively(&t.root, key, value)	
}

func (m *OrderedMap) insertRecursively(n **node, key, value int) {
	if (*n == nil) { // if *node is nil - SUBtree is empty
		*n = &node {
			key: key,
			value: value,
		}
		m.size++
		return
	}

	if key == (*(*n)).key {
 		(*(*n)).value = value
		return
	}

	if key < (*(*n)).key {
		m.insertRecursively(&(*n).left, key, value)
	}

	if key > (*(*n)).key {
		m.insertRecursively(&(*n).right, key, value)
	}
}

// SEARCH ----------------------------------------------------

func (m *OrderedMap) Contains (key int) bool {
	_, _, isFound := m.searchNodeRecursively(&m.root, nil, key)
	return isFound  
}

// **node, **parentNode, and isFound are needed for `Erase()` func
func (m *OrderedMap) searchNodeRecursively(n **node, parent **node, key int) (**node, **node, bool) {
	if (*n == nil) {
		return nil, nil, false
	}

	if key == (*n).key {
		return n, parent, true
	}

	if key < (*n).key {
		return m.searchNodeRecursively(&(*n).left, n, key)
	}

	if key > (*n).key {
		return m.searchNodeRecursively(&(*n).right, n, key)
	}

	return nil, nil, false
}

// TREE TRAVERSAL -------------------------------------------

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.traversal(m.root, action)	
}

func (m *OrderedMap) traversal(n *node, action func(int, int)) {
	if n != nil {
		m.traversal(n.left, action)

		action(n.key, n.value) // we can move it to change traversal order

		m.traversal(n.right, action)
	}
}

// DELETE --------------------------------------------------

func (m *OrderedMap) Erase(key int) {
	node, parent, isFound := m.searchNodeRecursively(&m.root, nil, key)
	if isFound == false {
		return
	}

	m.erase(node, parent) // **node, **node
}

func (m *OrderedMap) erase(n **node, p **node) {
	// if try delete root tree
	if *n == m.root {
		m.root = nil
		m.size = 0
		return
	}

	// if no childs, just remove
	if *n != nil && (*n).left == nil && (*n).right == nil {
		*n = nil
		m.size--
		return
	}

	// if only one child 
	if *n != nil && (*n).left == nil && (*n).right != nil {
		if (*n).right.key > (*p).key { // from what side we came from?
			(*p).right = (*n).right
		} else {
			(*p).left = (*n).right
		}
		m.size--
		return
	}

	if *n != nil && (*n).left != nil && (*n).right == nil {
		if (*n).left.key > (*p).key {
			(*p).right = (*n).left
		} else {
			(*p).left = (*n).left
		}
		m.size--
		return
	}

	// if both childs are in place we should find min element in right subtree
	// set key / value to removing element and remove min element
	if *n != nil && (*n).left != nil && (*n).right != nil {
		minNode := m.findMinNode(&(*n).right)
		(*(*n)).key = (*minNode).key 
		(*(*n)).value = (*minNode).value 
		*minNode = nil
		m.size--
		return
	}
}

func (m *OrderedMap) findMinNode(n **node) **node {
	for true {
		if (*n).right == nil {
			return n
		}
		m.findMinNode(&(*n).right)
	}
	return nil
}

// SIZE -----------------------------------------------------------------------------------------------------------

func (t *OrderedMap) Size() int {
	return t.size
}

// DEBUG -----------------------------------------------------------------------------------------------------------

func (m *OrderedMap) Debug() {
	var keys []int
	m.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	fmt.Println(keys, len(keys))
}

// MAIN -----------------------------------------------------------------------------------------------------------

func main() {
	m := NewOrderedMap()
	m.Insert(1, 1)
	m.Insert(2, 2)
	m.Insert(3, 3)
	fmt.Println(m.Contains(2)) // true
	fmt.Println(m.Contains(4)) // false
	m.Erase(2)
	fmt.Println(m.Contains(2)) // false
	m.Debug()
}
