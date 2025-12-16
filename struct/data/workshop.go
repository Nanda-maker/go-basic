package data

import "time"

// Embedding = putting one struct inside another, so that its fields & methods become part of the bigger struct.
//Inheritance-like behavior
type Workshop struct{
	Course        // embedding one type into another type - embedding course within the workshop
	Instructor
	Date time.Time
}

func (c Workshop) SignUp() bool {
	return true
}


func NewWorkshop(name string, instructor Instructor) Workshop{
	w := Workshop{}
	w.Name = name
	// w.Instructor.FirstName 
	w.Instructor = instructor
	return w

}
