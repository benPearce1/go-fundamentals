package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	name := "Ben"
	course := "course name"
	module := "4"
	clip := 2
	//courseComplete := false

	fmt.Println("Name and course are", name, "and", course, ".")
	fmt.Println("Module and clip are", module, "and", clip, ".")
	fmt.Println("Name is type ", reflect.TypeOf(name))
	fmt.Println("Module is type", reflect.TypeOf(module))
	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("module plus clip is", total)
	}

	fmt.Println("Memory address of 'course' variable is", &course)
	var ptr *string = &course
	fmt.Println("Pointing course variable at address", ptr, "which holds value,", *ptr)
}
