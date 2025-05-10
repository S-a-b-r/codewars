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

	repeats := make(map[int]int)

	for i := range list {
		repeats[list[i]]++

		for repeats[list[i]] > 1 && list[i] != 0 {
			divTwo(list, repeats, i)
		}
	}

	fmt.Println(len(repeats))
}

func divTwo(list []int, repeats map[int]int, num int) {
	repeats[list[num]]--
	list[num] = int(math.Floor(float64(list[num]) / 2))
	repeats[list[num]]++
}
