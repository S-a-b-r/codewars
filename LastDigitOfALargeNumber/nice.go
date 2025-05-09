package main

import "strconv"

func lastDigit(n1, n2 string) int {
	var l = map[int][]int{
		0: {0},
		1: {1},
		2: {6, 2, 4, 8},
		3: {1, 3, 9, 7},
		4: {6, 4},
		5: {5},
		6: {6},
		7: {1, 7, 9, 3},
		8: {6, 8, 4, 2},
		9: {1, 9},
	}

	//last digit of n1
	x, _ := strconv.Atoi(n1[len(n1)-1:])

	if len(n2) > 2 {
		n2 = n2[len(n2)-2:]
	}

	y, _ := strconv.Atoi(n2)

	return l[x][int(y)%len(l[x])]
}
