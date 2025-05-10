package main

import (
	"fmt"
)

type Cartage struct {
	start int
	end   int
}

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&list[i])
	}

	cartages := make([]Cartage, 0)
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			for k := j + 1; k < cnt; k++ {
				if list[j]-list[i] == list[k]-list[j] {
					cartages = getAllCartages(cartages, i, k, cnt)
				}
			}
		}
	}

	fmt.Println(len(cartages))
}

func getAllCartages(cartages []Cartage, start, end, cnt int) []Cartage {
	for i := end; i < cnt; i++ {
		if !inCartage(cartages, start, i) {
			cartages = append(cartages, Cartage{start: start, end: i})
		}
	}

	return cartages
}

func inCartage(cartages []Cartage, start, end int) bool {
	for _, v := range cartages {
		if v.end == end && v.start == start {
			return true
		}
	}

	return false
}
