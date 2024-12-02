package main

import "fmt"

func areIntersect(slice1 []int, slice2 []int) (bool, []int) {

	values := make(map[int]bool)
	res := make([]int, 0)

	for _, value := range slice1 {
		values[value] = true
	}
	for _, value := range slice2 {
		if values[value] {
			res = append(res, value)
		}
	}

	return len(res) > 0, res
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(areIntersect(a, b))
}
