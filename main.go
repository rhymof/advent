package main

import (
	"fmt"
	"time"
)

func main() {
	if cowntdown(time.Now()) == 0 {
		fmt.Println("Merry, Christmas!")
	}
}

func cowntdown(t time.Time) int {
	if t.Month() != time.December {
		return -1
	}
	if t.Day() < 25 {
		return 25 - t.Day()
	} else if t.Day() == 25 {
		return 0
	}
	return -1
}
