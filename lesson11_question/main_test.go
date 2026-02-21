package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)


/*
Домашнее задание №11

В домашнем задании нужно реализовать функцию трассировки объектов.

Трассировка объектов будет осуществляться со стеков, рекурсивно проходя все объекты в куче. Важно написать
реализацию алгоритма так, чтобы не было повторных обходов одних и тех же объектов. Циклов не будет, проверку
на это реализовывать не нужно. Для обхода объектов можно использовать алгоритмы DFS или BFS.

Различать указатели и свободные участки памяти нужно будет при помощи значений в памяти, если значение непустое -
значит это указатель, который указывает на какой-то другой участок память.
*/

func TestTrace(t *testing.T) {
	var heapObjects = []int{
		0x00, 0x00, 0x00, 0x00, 0x00,
	}

	var heapPointer1 *int = &heapObjects[1]
	var heapPointer2 *int = &heapObjects[2]
	var heapPointer3 *int = nil
	var heapPointer4 **int = &heapPointer3

	var stacks = [][]uintptr{
		{
			uintptr(unsafe.Pointer(&heapPointer1)), 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[0])),
			0x00, 0x00, 0x00, 0x00,
		},
		{
			uintptr(unsafe.Pointer(&heapPointer2)), 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[1])),
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[2])),
			uintptr(unsafe.Pointer(&heapPointer4)), 0x00, 0x00, 0x00,
		},
		{
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, uintptr(unsafe.Pointer(&heapObjects[3])),
		},
	}

	pointers := Trace(stacks) // Try to move this line under `expectedPoiners` definition
	expectedPointers := []uintptr{
		uintptr(unsafe.Pointer(&heapPointer1)),
		uintptr(unsafe.Pointer(&heapObjects[0])),
		uintptr(unsafe.Pointer(&heapPointer2)),
		uintptr(unsafe.Pointer(&heapObjects[1])),
		uintptr(unsafe.Pointer(&heapObjects[2])),
		uintptr(unsafe.Pointer(&heapPointer4)),
		uintptr(unsafe.Pointer(&heapPointer3)),
		uintptr(unsafe.Pointer(&heapObjects[3])),
	}
	
	fmt.Println("Expected:", expectedPointers)
	fmt.Println("Stack 0 0:", stacks[0][0])

	assert.True(t, reflect.DeepEqual(expectedPointers, pointers))
}

func Trace(stacks [][]uintptr) []uintptr {
	fmt.Println("Stack 0 0 from Trace", stacks[0][0])
	return []uintptr{}
}

func main() {
}
