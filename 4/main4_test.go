package main

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestElemsOnlyInFist(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	res := elemsOnlyInFirst(slice1, slice2)
	slices.Sort(res)
	assert.Equal(t, []string{"43", "apple", "cherry", "gno1", "lead"}, res)
}

func TestElemsOnlyInFist_WithRecurringElems1(t *testing.T) {
	slice1 := []string{"apple", "apple", "cherry"}
	slice2 := []string{"cherry"}
	res := elemsOnlyInFirst(slice1, slice2)
	assert.Equal(t, []string{"apple"}, res)
}

func TestElemsOnlyInFist_WithRecurringElems2(t *testing.T) {
	slice1 := []string{"apple", "apple", "cherry", "cherry"}
	slice2 := []string{"cherry"}
	res := elemsOnlyInFirst(slice1, slice2)
	assert.Equal(t, []string{"apple"}, res)
}

func TestElemsOnlyInFist_NoCommonElems(t *testing.T) {
	slice1 := []string{"42", "43", "44"}
	slice2 := []string{"45"}
	res := elemsOnlyInFirst(slice1, slice2)
	assert.Equal(t, slice1, res)
}
