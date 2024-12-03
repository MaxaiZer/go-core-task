package main

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestIntersect1(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	res, slice := areIntersect(a, b)
	slices.Sort(slice)

	assert.True(t, res)
	assert.Equal(t, []int{3, 64}, slice)
}

func TestIntersect2(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{66, 2, 4, 43}

	res, slice := areIntersect(a, b)

	assert.False(t, res)
	assert.Equal(t, []int{}, slice)
}
