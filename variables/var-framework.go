package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	course := "course name"
	module := "4"
	clip := 2
	//courseComplete := false

	fmt.Println("Name and course are", name, "and", course, ".")
	fmt.Println("Module and clip are", module, "and", clip, ".")
}
