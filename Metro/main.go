package main

import "fmt"

type line struct {
	start   int
	timeout int
}

type coming struct {
	lineNum int
	time    int
}

func main() {
	var cntLines, cntComings int
	fmt.Scanln(&cntLines)

	lines := make([]line, cntLines)
	for i := 0; i < cntLines; i++ {
		l := line{}
		fmt.Scan(&l.start, &l.timeout)
		lines[i] = l
	}

	fmt.Scanln(&cntComings)

	comings := make([]coming, cntComings)
	for i := 0; i < cntComings; i++ {
		c := coming{}
		fmt.Scan(&c.lineNum, &c.time)
		comings[i] = c
	}

	for _, v := range comings {
		fmt.Println(checkTime(lines[v.lineNum-1], v.time))
	}
}

func checkTime(l line, time int) int {
	if (time-l.start)%l.timeout == 0 {
		return time
	}
	if time < l.start {
		return l.start
	}

	n := (time-l.start)/l.timeout + 1
	return l.start + n*l.timeout
}
