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

func NewOrderedMap() OrderedMap {
	return OrderedMap{}
}

func (m *OrderedMap) Insert (key, value int) {
	m.root = m.insert(m.root, key, value)
}

func (m *OrderedMap) insert(n *node, key, value int) *node {
	if n == nil {
		m.size++
		return &node{
			key: key,
			value: value,
		}
	}

	if key == n.key {
		n.value = value
	}

	if key > n.key {
		n.right = m.insert(n.right, key, value)
	}

	if key < n.key {
		n.left = m.insert(n.left, key, value)
	}

	return n
}

func (m *OrderedMap) Contains (key int) bool {
	isFound := m.contains(m.root, key)
	return isFound
}

func (m *OrderedMap) contains(n *node, key int) bool {
	if n == nil {
		return false
	}

	if key == n.key {
		return true
	}

	if key > n.key {
		return m.contains(n.right, key)
	}

	return m.contains(n.left, key)
}

func (m *OrderedMap) Erase(key int) {
	m.root = m.erase(m.root, key)
}

func (m *OrderedMap) erase(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = m.erase(n.left, key)
	}

	if key > n.key {
		n.right = m.erase(n.right, key)
	}

	// no childs
	if key == n.key && n.left == nil && n.right == nil {
		m.size--
		return nil
	}

	// one child (left or right)
	if key == n.key && n.left == nil && n.right != nil {
		m.size--
		return n.right
	}

	if key == n.key && n.left != nil && n.right == nil {
		m.size--
		return n.left
	}

	// two childs
	if key == n.key && n.left != nil && n.right != nil {
		m.size--
		minInRightSubtree := minNode(n.right)
		n.key = minInRightSubtree.key // "replace" erasing node with min node
		n.value = minInRightSubtree.value
		n.right = m.erase(n.right, minInRightSubtree.key) // remove min node
	}

	return n
}

func minNode(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	m.traversal(m.root, action)
}

func (m *OrderedMap) traversal(n *node, action func(int, int)) {
	if n != nil {
		m.traversal(n.left, action)
		action(n.key, n.value)
		m.traversal(n.right, action)
	}
}

func (m *OrderedMap) Size() int {
	return m.size
}

func (m *OrderedMap) Debug() [][2]int {
	var nodeData [][2]int
	m.ForEach(func(key, value int) {
		nodeData = append(nodeData, [2]int{key, value})
	})
	fmt.Println(nodeData)
	return nodeData
}


func main() {
	m := NewOrderedMap()
	m.Insert(10, 10)
	m.Insert(4, 4)
	m.Insert(5, 5)
	m.Insert(2, 2)
	m.Insert(20, 20)
	m.Insert(18, 18)
	m.Insert(22, 22)
	m.Insert(16, 16)
	m.Insert(14, 14)

	m.Debug()
	// m.Insert(18, 18)

	// fmt.Println(m.Contains(14))
	// fmt.Println(m.Contains(20))
	// m.Erase(12)
	m.Erase(18)
	m.Debug()

	// m.Insert(2, 2)
	// m.Insert(3, 3)
	// fmt.Println(m.Contains(2)) // true
	// fmt.Println(m.Contains(4)) // false
	// m.Erase(2)
	// fmt.Println(m.Contains(2)) // false
	// m.Debug()
}


/* Ugly double pointer implementation:


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

func (t *OrderedMap) Size() int {
	return t.size
}

func (m *OrderedMap) Debug() {
	var keys []int
	m.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	fmt.Println(keys, len(keys))
}
*/
