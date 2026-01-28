package main

import "fmt"

func Map(data []int, action func(int) int) []int {
	if data == nil {
		return nil
	}

	result := make([]int, 0, len(data))
	for _, value := range data {
		result = append(result, action(value))
	}
	return result
}

func Filter(data []int, action func(int) bool) []int {
	if data == nil {
		return nil
	}

	result := make([]int, 0)
	for _, value := range data {
		if action(value) {
			result = append(result, value)
		}
	}
	return result
}

func Reduce(data []int, initial int, action func(int, int) int) int {
	for _, v := range data {
		initial = action(initial, v)
	}
	return initial
}

func main() {
	a := Reduce(
		[]int{50, 50},
		100,
		func(lhs, rhs int) int {
			return lhs + rhs
		},
	)

	fmt.Println(a)
}
