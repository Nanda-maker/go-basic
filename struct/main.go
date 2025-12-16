package main

import (
	"fmt"
	"go/struct/data"
)

func main(){

	// Create instructor using struct literal
	max := data.Instructor {Id: 3, FirstName: "john"} 
	max.LastName = "adc"

// Create instructor using factory
	kyle := data.NewInstructor("kyln","Simpson")

// Create course with embedded instructor
	goCourse := data.Course{Id: 2, Name: "Go Basic", Instructor: max }
	fmt.Printf("%v",goCourse)

	// swiftWS := data.Workshop {Course: data.Course{Name: "Swift for ios", Instructor: max},Date: time.Now()}
	// Create workshop (demonstrates embedding)
	swiftWS := data.NewWorkshop("Swift with ios",max) // when we are embedding types we are also embedding methods of those types
	  // all the properties from course

	fmt.Printf("%v",swiftWS)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS
	for _, course := range courses {
		fmt.Println(course)
	}

// Print using methods
	print(max.Print())
	print(kyle.Print())
}
