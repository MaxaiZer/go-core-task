package main

import "fmt"

type StringIntMap struct {
	values map[string]int
}

func NewStringIntMap(cap ...int) *StringIntMap {
	capacity := 0
	if len(cap) > 0 {
		capacity = cap[0]
	}
	return &StringIntMap{values: make(map[string]int, capacity)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.values[key] = value
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.values[key]
	return value, ok
}

func (m *StringIntMap) Remove(key string) {
	delete(m.values, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copied := make(map[string]int)
	for k, v := range m.values {
		copied[k] = v
	}
	return copied
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.values[key]
	return ok
}

func main() {
	myMap := NewStringIntMap()
	key := "myKey"

	myMap.Add(key, 1)

	v, ok := myMap.Get(key)
	fmt.Printf("key %s: v: %v ok: %v\n", key, v, ok)

	ok = myMap.Exists(key)
	fmt.Printf("key %s: exists: %v\n", key, ok)

	copied := myMap.Copy()
	fmt.Printf("copied %v\n", copied)

	myMap.Remove(key)
	ok = myMap.Exists(key)
	fmt.Printf("key %s: exists: %v\n", key, ok)
}
