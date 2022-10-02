package main

import (
	"fmt"
)

func main() {
	bestFinish := finishes(12, 5, 6, 4, 3, 3)
	fmt.Println(bestFinish)
}

func finishes(list ...int) int {
	best := list[0]
	for _, i := range list {
		if i < best {
			best = i
		}
	}

	return best
}
