package main

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 2, 2, 3}
	arr2 := []int{3}

	fmt.Println(ArrayDiff(arr1, arr2))
}

func ArrayDiff(a, b []int) []int {
	c := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if !slicesContains(b, a[i]) {
			c = append(c, a[i])
		}
	}

	return c
}

func slicesContains(slice []int, item int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == item {
			return true
		}
	}

	return false
}
