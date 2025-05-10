package main

import "fmt"

type Stack struct {
	Items []byte
}

func NewStack() *Stack {
	return &Stack{Items: make([]byte, 0)}
}

func (s *Stack) Push(item byte) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() byte {
	if len(s.Items) == 0 {
		return 0
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item
}

func (s *Stack) Peek() byte {
	if len(s.Items) == 0 {
		return 0
	}
	return s.Items[len(s.Items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	var str string
	fmt.Scanln(&str)

	if len(str) != 2*n {
		return
	}

	st := NewStack()
	for i := 0; i < len(str); i++ {
		if str[i] == ')' && st.Peek() == '(' {
			st.Pop()
			continue
		}
		st.Push(str[i])
	}

	balance := 0
	openNeeded := 0
	closeNeeded := 0

	// Подсчет баланса и определение необходимых изменений
	for _, char := range st.Items {
		if char == '(' {
			balance++
		} else {
			if balance > 0 {
				balance--
			} else {
				closeNeeded++
			}
		}
	}
	openNeeded = balance

	totalChanges := (openNeeded + closeNeeded) / 2

	cost := 0
	if a < b {
		cost = totalChanges * a
	} else {
		cost = totalChanges * b
	}

	fmt.Println(cost)
}
