package main

import "fmt"

func main() {

	if dddLength, cawLength := 275, 30; dddLength > cawLength {
		fmt.Println("ddd is longer")
	} else if dddLength == cawLength {
		fmt.Println("same length")
	} else {
		fmt.Println("caw is longer")
	}
}
