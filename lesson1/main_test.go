package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
		"my test case #6": {
			number: 0xF0000000,
			result: 0x000000F0,
		},
		"my test case #7": {
			number: 0xFF000000,
			result: 0x000000FF,
		},
		"my test case #8": {
			number: 0xAABBCCDD,
			result: 0xDDCCBBAA,
		},
		"my test case #9": {
			number: 0xA0B0C0D0,
			result: 0xD0C0B0A0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
