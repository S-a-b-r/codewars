package main

import (
	"fmt"
	"math"
)

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	list := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&list[i])
	}

	points := 0
	for len(list) > 1 {
		var pp int
		list, pp = removeMaxPotencial(list)
		points += pp
	}

	fmt.Println(points)
}

func removeMaxPotencial(list []int) ([]int, int) {
	potential := make(map[int]float64)
	for i := 0; i+1 < len(list); i++ {
		if i > 1 && i+2 < len(list) {
			potential[i] = math.Abs(float64(list[i]-list[i+1])) + math.Abs(float64(list[i-1]-list[i+2]))
			continue
		}
		potential[i] = math.Abs(float64(list[i] - list[i+1]))
	}

	idx := findMaxPotencial(potential)
	points := int(math.Abs(float64(list[idx] - list[idx+1])))
	return append(list[:idx], list[idx+2:]...), points
}

func findMaxPotencial(p map[int]float64) int {
	var maxV float64
	var maxI int
	for i, v := range p {
		if v > maxV {
			maxI = i
			maxV = v
		}
	}

	return maxI
}
