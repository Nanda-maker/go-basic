package main

import (
	"fmt"
	"go/io/data"
)

// golobal variable
//var url = "https://go.dev"
var name = "Go master"
// function that will initialize the func when we start the app
func init(){
	fmt.Println("A")
}
func init(){
	fmt.Println("B")
}
// first indentifier and second type
func calculateTax(price float32) (float32,float32,string) {
	return price * 0.09, price * 0.02,""
}
func calculateTaxWithName(price float32) (stateTax float32,cityTaxx float32){
stateTax = price * 0.09
cityTaxx = price * 0.02
	return 
}
func birthday(pointerAge *int){
	if (*pointerAge > 140){
		panic("Too old to be true")
	}
	fmt.Printf("the pointer is %v and the value is %v\n",pointerAge, *pointerAge)
	*pointerAge++
}

func main(){

	fmt.Println(data.MaxSpeed)
	_,cityTax,_ := calculateTax(100)  // ":=" works only within the function - shorcut for var and type
	fmt.Println(cityTax)

	_,cityTaxx := calculateTaxWithName(100)
	fmt.Println(cityTaxx)
	// function-scoped variables
	// const maxSpeed = 60.   0
	//  message := "Hello from a module"
	//  price := 34. 4
	//  var isready bool
	// print(message,price,url,isready)

	defer fmt.Println("Bye!") //execution at the end until the function call that line.
	age := 22
	birthday(&age)
	fmt.Println(age) 
	// fmt.Println(age)  // print the memory address
	PrintData()
} 