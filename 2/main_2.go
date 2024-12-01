package main

import (
	"fmt"
	"math/rand"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func sliceExample[T Integer](slice []T) []T {
	res := make([]T, 0)
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			res = append(res, slice[i])
		}
	}
	return res
}

func addElements[T any](slice []T, elem T) []T {
	return append(slice, elem)
}

func copySlice[T any](slice []T) []T {
	res := make([]T, len(slice))
	copy(res, slice)
	return res
}

func removeElement[T any](slice []T, idx int) []T {
	if idx < 0 || idx >= len(slice) {
		return slice
	}
	return append(slice[:idx], slice[idx+1:]...)
}

func main() {

	size := 10
	elems := make([]int, 0, size)
	for i := 0; i < size; i++ {
		elems = append(elems, rand.Int()%100)
	}

	fmt.Printf("created array: %v\n", elems)

	elems = sliceExample(elems)
	fmt.Printf("array with only even numbers: %v\n", elems)

	elems = addElements(elems, 42)
	fmt.Printf("array after adding 42: %v\n", elems)

	copied := copySlice(elems)
	fmt.Printf("copied array: %v\n", copied)

	copied = removeElement(copied, 0)
	fmt.Printf("copied array after removing at 0 idx: %v\n", copied)
	fmt.Printf("original array: %v\n", elems)
}
