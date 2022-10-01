package main

import (
	. "fmt"
	"reflect"
	"strconv"
)

var (
	name, course = "Ben", "course name"
	module       = "4"
	clip         = 2
)

func main() {

	Println("Name and course are", name, "and", course, ".")
	Println("Module and clip are", module, "and", clip, ".")
	Println("Name is type ", reflect.TypeOf(name))
	Println("Module is type", reflect.TypeOf(module))
	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		Println("module plus clip is", total)
	}
}
