package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceExample(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = sliceExample(slice)
	assert.Equal(t, []int{2, 4, 6, 8, 10}, slice)
}

func TestAddElements(t *testing.T) {
	var slice []int

	slice = addElements(slice, 1)
	assert.Equal(t, []int{1}, slice)

	slice = addElements(slice, 2)
	assert.Equal(t, []int{1, 2}, slice)
}

func TestCopy(t *testing.T) {
	slice := []int{1, 2}

	copied := copySlice(slice)
	assert.Equal(t, slice, copied)

	copied = append(copied, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, copied)
	assert.Equal(t, []int{1, 2}, slice)
}

func TestCopy_WhenEmptySlice(t *testing.T) {
	slice := []int{}

	copied := copySlice(slice)
	assert.Equal(t, slice, copied)

	copied = append(copied, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.Equal(t, []int{3, 4, 5, 6, 7, 8, 9, 10}, copied)
	assert.Equal(t, []int{}, slice)
}

func TestRemove(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice = removeElement(slice, 0) //first
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}, slice)

	slice = removeElement(slice, 1) //not first or last
	assert.Equal(t, []int{2, 4, 5, 6, 7, 8, 9, 10}, slice)

	slice = removeElement(slice, 7) //last
	assert.Equal(t, []int{2, 4, 5, 6, 7, 8, 9}, slice)

	slice = removeElement(slice, 7) //out of borders
	assert.Equal(t, []int{2, 4, 5, 6, 7, 8, 9}, slice)

	slice = removeElement(slice, -1) //out of borders
	assert.Equal(t, []int{2, 4, 5, 6, 7, 8, 9}, slice)
}
