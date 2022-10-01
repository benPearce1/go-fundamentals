package main

import (
	"fmt"
)

func main() {
	name := "Ben"
	course := "course name"
	fmt.Println("\nHi", name, "your current course is", course)
	updateCourse(&course)
	fmt.Println("Current course is", course)
}

func updateCourse(course *string) string {
	*course = "Another course"
	fmt.Println("Updated course to", *course)
	return *course
}
