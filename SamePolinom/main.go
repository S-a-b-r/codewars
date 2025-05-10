package main

import "fmt"

func main() {
	if checkNearlyPoly("abca") {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}

func checkNearlyPoly(inpStr string) bool {
	if inpStr[0] != inpStr[3] {
		return isPoly(inpStr[0:3]) || isPoly(inpStr[1:])
	}
	if inpStr[0] != inpStr[2] {
		return isPoly(inpStr[0:2]+string(inpStr[3])) || isPoly(inpStr[1:])
	}
	return false
}

func isPoly(str string) bool {
	if str[0] != str[2] {
		return false
	}
	return true
}
