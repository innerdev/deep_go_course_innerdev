package main

import (
	"fmt"
	"unsafe"
)

func ToLittleEndian(number uint32) uint32 {
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

	fmt.Println("Is machine Little Endian?", MyLittleEndianCheck())
	fmt.Printf("Changing 0x%08X to 0x%08X\n", source, ToLittleEndian(source))
}
