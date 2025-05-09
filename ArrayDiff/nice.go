package main

func arrayDiff(a, b []int) (r []int) {
	m := map[int]bool{}
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			r = append(r, v)
		}
	}
	return // a-b
}
