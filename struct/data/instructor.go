package data

//A struct is like a container that groups related data together.
type Instructor struct {
	Id int
	FirstName string
	LastName string
	Score int
}
// design pattern to create a factory - a function
// Factory Pattern:NewInstructor() - Constructor function that creates and returns an Instructor instance
func NewInstructor(name string, lastname string) Instructor {
	return Instructor{FirstName: name, LastName: lastname}
}



// func (i Instructor) Print() string{
// 	return fmt.Sprintf("%v, %v (%d)",i.FirstName, i.LastName, i.Score)
// }