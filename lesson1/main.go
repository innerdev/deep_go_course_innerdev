package main

import (
	"fmt"
	"unsafe"
)

func ToLittleEndian(number uint32) uint32 {
	return ((number & 0xff000000) >> 24) |
		((number & 0x00ff0000) >> 8) |
		((number & 0x0000ff00) << 8) |
		((number & 0x000000ff) << 24)
}

func ToLittleEndianFirstVersionWithPointers(number uint32) uint32 {
	numberSize := unsafe.Sizeof(number) // but I have no clue about generics yet
	numberPtr := unsafe.Pointer(&number)

	for i := range numberSize / 2 {
		a := (*uint8)(unsafe.Add(numberPtr, i))
		b := (*uint8)(unsafe.Add(numberPtr, numberSize-i-1))
		*a, *b = *b, *a
	}

	return number
}

func MyLittleEndianCheck() bool {
	var n uint = 0x00FF
	return *(*uint8)(unsafe.Pointer(&n)) == 0xFF
}

func main() {
	var source uint32 = 0xDDCCBBAA

	fmt.Printf("%08x\n", source&0xf)
	fmt.Printf("%d\n", unsafe.Sizeof((uint8)(source&0xf))) // Still 4 bytes, but why? :thinking:

	fmt.Println("Is machine Little Endian?", MyLittleEndianCheck())
	fmt.Printf("Changing 0x%08X to 0x%08X\n", source, ToLittleEndian(source))
}
