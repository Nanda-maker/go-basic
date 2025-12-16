package data

import "fmt"

type Duration float32

type Course struct{
	Id int
	Name string
	Slug string
	Legacy bool
	Duration   Duration      // custom type (float32 alias)
    Instructor Instructor    // composition
}

func (c Course) SignUp() bool {
	return true
}


// change the string representation of your structure (Implementing the stringer Interface)
// If your struct has a String() method, printing it automatically uses your custom format.
// String() method - Customizes how Course prints (used by fmt.Printf with %v)
// Returns formatted string: "---CourseName --- (InstructorFirstName)"
func (c Course) String() string {
	return fmt.Sprintf("---%v --- (%v)\n",c.Name,c.Instructor.FirstName)
}