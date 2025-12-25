package main

import "fmt" 

func PrintData(){
	fmt.Print("Hello ") //print() - is only used for development purposes since this func. is not safe that it is not gaurantee it will work in every platform & OS. just quick way to debug.	
	fmt.Println("World")
	fmt.Println(name)
}