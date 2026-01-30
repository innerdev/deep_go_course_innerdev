package main

import "fmt"

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	head *node
}

func (l *LinkedList) Add(value int) {
	if l.head == nil {
		l.head = &node{value: value}
		return
	}

	n := l.head
	for n.next != nil {
		n = n.next
	}

	n.next = &node{value: value}
}

func (l *LinkedList) Remove(value int) {
	l.head = l.remove(l.head, value)
}

func (l *LinkedList) remove(n *node, value int) *node {
	if n == nil {
		return nil
	}

	if n.value == value {
		if n.next == nil { // last elem
			return nil	
		}

		if n.next != nil { // not last, "reorder"
			return n.next
		}
	}

	n.next = l.remove(n.next, value)

	return n
}

func (l *LinkedList) Print() {
	n := l.head
	for n != nil {
		fmt.Printf("%d ", n.value)
		n = n.next
	}
}

func main() {
	l := LinkedList{}
	l.Add(1)
	l.Add(2)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Remove(5)
	l.Remove(4)
	l.Remove(3)
	l.Remove(2)
	l.Print()
}
