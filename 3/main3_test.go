package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	myMap := NewStringIntMap()
	key := "key"
	value := 1

	myMap.Add(key, value)

	assert.True(t, myMap.Exists(key))
	v, ok := myMap.Get(key)
	assert.True(t, ok)
	assert.Equal(t, value, v)
	assert.Equal(t, value, myMap.values[key])
}

func TestAdd_Recurring(t *testing.T) {
	myMap := NewStringIntMap()
	key := "key"
	value := 1

	myMap.Add(key, value)
	myMap.Add(key, value)

	assert.True(t, myMap.Exists(key))
	v, ok := myMap.Get(key)
	assert.True(t, ok)
	assert.Equal(t, value, v)
	assert.Equal(t, value, myMap.values[key])
}

func TestRemove(t *testing.T) {
	myMap := NewStringIntMap()
	key := "key"
	value := 1

	myMap.Add(key, value)
	myMap.Remove(key)

	assert.False(t, myMap.Exists(key))
	v, ok := myMap.Get(key)
	assert.False(t, ok)
	assert.Equal(t, 0, v)
	assert.Equal(t, 0, myMap.values[key])
}

func TestCopy(t *testing.T) {
	myMap := NewStringIntMap()
	key := "key"
	value := 1

	myMap.Add(key, value)
	copied := myMap.Copy()

	v, ok := copied[key]
	assert.True(t, ok)
	assert.Equal(t, value, v)

	myMap.Remove(key)
	v, ok = copied[key]
	assert.True(t, ok)
	assert.Equal(t, value, v)
}
