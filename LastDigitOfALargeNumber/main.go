package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(LastDigit("6971586188927822805771146473742", "11573003345756015180549423511417543625794"))
}

func LastDigit(n1, n2 string) int {
	if len(n1) == 1 && len(n2) == 1 {
		if n1[0] == '0' && n2[0] == '0' {
			return 1
		}
	}

	x, _ := strconv.Atoi(string(n1[len(n1)-1]))

	var y int
	if len(n2) < 2 {
		y, _ = strconv.Atoi(string(n2[len(n2)-1]))
	} else {
		y, _ = strconv.Atoi(n2[len(n2)-2:])
	}

	switch x {
	case 1:
		return x
	case 2:
		switch y % 4 {
		case 0:
			return 6
		case 1:
			return 2
		case 2:
			return 4
		case 3:
			return 8
		}
	case 3:
		switch y % 4 {
		case 0:
			return 1
		case 1:
			return 3
		case 2:
			return 9
		case 3:
			return 7
		}
	case 4:
		switch y % 2 {
		case 0:
			return 6
		case 1:
			return 4
		}
	case 5:
		return 5
	case 6:
		return 6
	case 7:
		switch y % 4 {
		case 0:
			return 1
		case 1:
			return 7
		case 2:
			return 9
		case 3:
			return 3
		}
	case 8:
		switch y % 4 {
		case 0:
			return 6
		case 1:
			return 8
		case 2:
			return 4
		case 3:
			return 2
		}
	case 9:
		switch y % 2 {
		case 0:
			return 1
		case 1:
			return 9
		}
	case 0:
		return 0
	}

	return 0
}
