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
		refs: new(int),
	}

	*newCow.refs = 1
	*b.refs++

	return newCow
}

func (b *COWBuffer) Close() {
	if b.refs == nil && b.data == nil {
		panic("Closed COWBuffer can't be closed.")
	}

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
		*b.refs--
		old := b.data
		b.data = make([]byte, len(old))
		copy(b.data, old)
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

	fmt.Println("data:", string(b.data), "refs:", b.refs, "refs p", *b.refs)
}

func main() {
	c := NewCOWBuffer([]byte("hello"))
	c.Debug()

	c.Update(0, 'A')
	c.Debug()

	c2 := c.Clone()
	c.Debug()

	c2.Update(2, 'B')
	c.Debug()

	c.Close()
	c.Debug() // panic (it's ok)
}
