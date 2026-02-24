package main

import (
	"fmt"
	"unsafe"
)

func trace(ptr uintptr, result *[]uintptr, visited map[uintptr]bool) {
	next := *(*uintptr)(unsafe.Pointer(ptr))
	if next != 0 && !visited[next] {
		visited[next] = true
		*result = append(*result, next)
		trace(next, result, visited)
	}
}

func Trace(stacks [][]uintptr) []uintptr {
	visited := make(map[uintptr]bool)
	result := make([]uintptr, 0)

	for _, stack := range stacks {
		for _, ptr := range stack {
			if ptr != 0x00 && !visited[ptr] {
				visited[ptr] = true
				result = append(result, ptr)

				pointersChain := make([]uintptr, 0)
				trace(ptr, &pointersChain, visited)

				result = append(result, pointersChain...)
			}
		}
	}

	return result
}

func main() {
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
	pointers := Trace(stacks)

	fmt.Println(expectedPointers)
	fmt.Println(pointers)
}
