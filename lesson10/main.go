package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func findFirstFreeByte(memory []byte, pointers []unsafe.Pointer) (memoryIndex int, err error) {
	for i := 0; i < len(memory); i++ {
		isFree := true

		for j := 0; j < len(pointers); j++ {
			if &memory[i] == (*byte)(pointers[j]) {
				isFree = false
			}
		}

		if isFree {
			return i, nil
		}
	}

	return 0, errors.New("no free memory left")
}

func findFirstUsedByte(memory []byte, pointers []unsafe.Pointer, memoryStartSearchIndex int) (memoryIndex, pointerIndex int, err error) {
	for i := memoryStartSearchIndex; i < len(memory); i++ {
		for j := 0; j < len(pointers); j++ {
			if &memory[i] == (*byte)(pointers[j]) {
				return i, j, nil
			}
		}
	}

	// if we are here, no used bytes after `memoryStartSearchIndex` are found
	return 0, 0, errors.New("no used bytes found")
}

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	for {
		freeIndex, err := findFirstFreeByte(memory, pointers)
		if err != nil {
			return // no need in defragmentation because of memory if full
		}

		usedIndex, pointerIndex, err := findFirstUsedByte(memory, pointers, freeIndex)
		if err != nil {
			return // no need in defragmentation because of no bytes are in use (after `freeIndex`)
		}

		memory[freeIndex] = memory[usedIndex]
		pointers[pointerIndex] = unsafe.Pointer(&memory[freeIndex])
		memory[usedIndex] = 0x0
	}
}

func main() {
	var fragmentedMemory = []byte{
		0xFF, 0x00, 0x00, 0x00,
		0x00, 0xFF, 0x00, 0x00,
		0x00, 0x00, 0xFF, 0x00,
		0x00, 0x00, 0x00, 0xFF,
	}

	var fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&fragmentedMemory[0]),
		unsafe.Pointer(&fragmentedMemory[5]),
		unsafe.Pointer(&fragmentedMemory[10]),
		unsafe.Pointer(&fragmentedMemory[15]),
	}

	fmt.Println(fragmentedMemory, fragmentedPointers)

	Defragment(fragmentedMemory, fragmentedPointers)

	fmt.Println(fragmentedMemory, fragmentedPointers)
}

/*

Первая версия: свободные байты ищем, начиная с начала памяти. Занятые - с конца. Переставляем занятые байты
в свободные ячейки. Как только индекс свободного байта больше, чем индекс занятого, значит всё дефрагментировано.

Не подходит потому, что меняется порядок байт после дефрагментации.

func findFirstFreeByte(memory []byte, pointers []unsafe.Pointer) (memoryIndex int, err error) {
	for i := range memory {
		isFree := true

		for j := range pointers{
			if &memory[i] == (*byte)(pointers[j]) {
				isFree = false
			}
		}

		if isFree {
			return i, nil
		}
	}
	return 0, errors.New("no free memory left")
}

func findLastUsedByte(memory []byte, pointers []unsafe.Pointer) (memoryIndex, pointerIndex int, err error) {
	for i := len(memory) - 1; i > 0; i-- {
		for j := range pointers{
			if &memory[i] == (*byte)(pointers[j]) {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("no used bytes found")
}

func Defragment(memory []byte, pointers []unsafe.Pointer) {
	for {
		freeMemoryIndex, err := findFirstFreeByte(memory, pointers)
		if err != nil {
			log.Println("No free memory left")
			return
		}

		usedMemoryIndex, usedPointerIndex, err := findLastUsedByte(memory, pointers)
		if err != nil {
			log.Println("No memory used")
			return
		}

		if usedMemoryIndex < freeMemoryIndex { // memory is defragmented
			return
		}

		memory[freeMemoryIndex] = memory[usedMemoryIndex]
		pointers[usedPointerIndex] = unsafe.Pointer(&memory[freeMemoryIndex])
		memory[usedMemoryIndex] = 0x0
	}
}

func main() {
	var fragmentedMemory = []byte{
		0xFF, 0x00, 0x00, 0x00,
		0x00, 0xFF, 0x00, 0x00,
		0x00, 0x00, 0xFF, 0x00,
		0x00, 0x00, 0x00, 0xFF,
	}

	var fragmentedPointers = []unsafe.Pointer{
		unsafe.Pointer(&fragmentedMemory[0]),
		unsafe.Pointer(&fragmentedMemory[5]),
		unsafe.Pointer(&fragmentedMemory[10]),
		unsafe.Pointer(&fragmentedMemory[15]),
	}

	fmt.Println(fragmentedMemory, fragmentedPointers)

	Defragment(fragmentedMemory, fragmentedPointers)

	fmt.Println(fragmentedMemory, fragmentedPointers)
}
*/
