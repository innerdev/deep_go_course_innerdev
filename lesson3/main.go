package main

import (
	"fmt"
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	str := COWBuffer{
		data: data,
		refs: new(int), // have no idea how to inline creation and set of pointer
	}
	*str.refs = 1

	return str
}

func (b *COWBuffer) Clone() COWBuffer {
	if b.refs == nil && b.data == nil {
		panic("Closed COWBuffer can't be cloned.")
	}

	newCow := COWBuffer{
		data: b.data,
		refs: b.refs,
	}

	*newCow.refs++

	return newCow
}

func (b *COWBuffer) Close() {
	if b.refs == nil && b.data == nil {
		panic("Closed COWBuffer can't be closed.")
	}

	*b.refs--
	b.data = nil
	b.refs = nil
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if b.refs == nil && b.data == nil {
		return false
	}

	if index < 0 || index > len(b.data) - 1 {
		return false
	}

	if *b.refs > 1 {
		*b.refs-- // decrement

		old := b.data // copy underlying array
		b.data = make([]byte, len(old))
		copy(b.data, old)

		b.refs = new(int) // new element is only one
		*b.refs = 1
	}

	b.data[index] = value

	return true
}

func (b *COWBuffer) String() string {
	if b.refs == nil && b.data == nil {
		panic("Closed COWBuffer can't be converted to string.")
	}

	return unsafe.String(unsafe.SliceData(b.data), len(b.data))  // or *(*string)(unsafe.Pointer(&b.data))
}

func (b *COWBuffer) Debug() {
	if b.refs == nil && b.data == nil {
		panic("Can't debug closed buffer.")
	}

	fmt.Println("data:", string(b.data), "dataptr:", unsafe.SliceData(b.data), "refs:", b.refs, "refs p", *b.refs)
}

func main() {
	c := NewCOWBuffer([]byte("hello"))
	c1 := c.Clone()
	c2 := c.Clone()
	c3 := c.Clone()
	// c2.Update(2, 'B')
	c1.Debug()
	c2.Debug()
	c3.Debug()
	c3.Update(0, 'A')
	c3.Debug()
	c2.Debug()

	// c.Debug()
	//
	// c.Update(0, 'A')
	// c.Debug()
	//
	// c2 := c.Clone()
	// c.Debug()
	//
	// c2.Update(2, 'B')
	// c.Debug()
	//
	// c.Close()
	// c.Debug() // panic (it's ok)
}
