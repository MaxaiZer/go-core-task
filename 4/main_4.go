package main

import "fmt"

func elemsOnlyInFirst(slice1 []string, slice2 []string) []string {

	res := make([]string, 0)
	firstItems := make(map[string]bool)

	for _, item := range slice1 {
		firstItems[item] = true
	}
	for _, item := range slice2 {
		if _, ok := firstItems[item]; ok {
			firstItems[item] = false
		}
	}
	for key, value := range firstItems {
		if value {
			res = append(res, key)
		}
	}
	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(elemsOnlyInFirst(slice1, slice2))
}
