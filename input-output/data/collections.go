package data

import (
	"fmt"
	"strings"
)

var Countries [10]string // array of strings with fixed size 10
var Slice []int // slice of integers - dynamic size --they are actually chunks of arrays
var Codes map[int]string  // map of int keys and string values
// declare an array on multiple lines for better readability
var a6 = [...]int{1,
	2,
	3,
} //the ending comma is mandatory when initializing the array on multiple lines and the closing curly brace is on its own line
 var numbers [4]int

func init(){
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[2] = "Brazil"
	// qty := len(Countries)
	_ = a6
	   //  = operator creates a copy of an array.
    // the arrays are not connected and are saved in different memory locations
    m := [3]int{1, 2, 3}
    n := m //n is a copy of m
 
    fmt.Println("n is equal to m: ", n == m) // => true
    m[0] = -1                                //only m is modified
    fmt.Println("n is equal to m: ", n == m) // => false

	fmt.Println(strings.Repeat("#",20))
	  // changing an array
    // we can't add or remove elements from the array (it's fixed length)
    numbers[0] = 7              //changing first element (index 0)
    fmt.Printf("%v\n", numbers) // -> [7 0 0 0]

	  // declaring a multi-dimensional arrays (array of arrays or matrix)
    balances := [2][3]int{
       {5, 6, 7}, 
        {8, 9, 10},
    }
 
    for _, arr := range balances {
        for _, value := range arr {
            fmt.Printf("%d ", value)
        }
        fmt.Println("")
    }

	 // each key corresponds to an index of the array
    grades := [3]int{ //the keyed elements can be in any order
        1: 10,
        0: 5,
        2: 7,
    }
    fmt.Println(grades) // -> [5 10 7]
 
    accounts := [3]int{
        2: 50,
    }
    fmt.Println(accounts) //[0 0 50]
	
	names := [...]string{
        4: "Dan",
    }
    fmt.Println(len(names)) // -> 5
	fmt.Println(names)
 
    // un unkeyed element gets its index from the last keyed element
    cities := [...]string{
        5:        "Paris",
        "London", // this is at index 6
        1:        "NYC",
    }
    fmt.Printf("%#v\n", cities) // -> [7]string{"", "NYC", "", "", "", "Paris", "London"}


}

 