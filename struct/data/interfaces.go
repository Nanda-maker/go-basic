package data

// interfaces just a list of methods that then we can use as a type
// implicit implementation - polymorphism -- looping through lots of objects
type Signable interface{
	SignUp() bool
} 