package main

import "fmt"

func main(){
	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0")
	fmt.Println("Which operation do you want to perform? (add, substract, multiply, divide)")
	fmt.Scanf("%s", &operation) //if a function to change not pass by value we need to add "&" - send the address of the variable
	fmt.Println("Enter first number")
	fmt.Scanf("%d",&number1)
	fmt.Println("Enter Second number")
	fmt.Scanf("%d",&number2)
	switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "substract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
		
	}

}