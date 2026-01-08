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
    p := m //n is a copy of m
 
    fmt.Println("n is equal to m: ", p == m) // => true
    m[0] = -1                                //only m is modified
    fmt.Println("n is equal to m: ", p == m) // => false

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

		// declaring a string slice, by default is initialized with nil or uninitialized
	var citiess []string
 
	fmt.Println("cities is equal to nil: ", citiess == nil) // -> cities is equal to nil:  true
	fmt.Printf("cities: %#v\n", cities)                    // -> cities: []string(nil)
 
	// we can not assign elements to nil slice:
	// cities[0] = "Paris" // -> runtime error
 
	// declaring a slice using a slice literal
	numbers := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers)         // => [2 3 4 5]
 
	// creating a slice using the make() built-in function
	// creating a slice with 2 int elements initialized with zero.
	nums := make([]int, 2) // the same as []int{0, 0}.
	fmt.Println(nums)      // => [0 0]
 
	// declaring a slice using a slice literal
	type namess []string
	friends := namess{"Dan", "Maria"}
	fmt.Println(friends)
 
	// getting an element from the slice
	x := numbers[0]
	fmt.Println("x is", x) // => x is 2
 
	// modifying an element of the slice
	numbers[1] = 200
	fmt.Printf("%#v\n", numbers) // => []int{2, 200, 4, 5}
 
	// iterating over a slice
	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}
 
	//iterating over a slice (C/C++, Java Style)
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index: %v, value: %v\n", i, numbers[i])
 
	}
 
	// slices with the same element type can be assigned to each other
	var n []int
	n = numbers
	_ = n
 
	//** COMPARING SLICES **//
	// a Go slice can only be compared to nil
 
	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true
 
	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false
 
	// we can not compare slices using the equal (=) operator
	// fmt.Println(nn == mm) //error -> slice can only be compared to nil
 
	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
    // it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
    a, b := []int{1, 2, 3}, []int{1, 2, 3}
    var eq bool = true
	if len(a) != len(b) {
		eq = false
	}
 
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}
	
	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
	// append() returns a new slice after appending a value to its end
    numbers = append(numbers, 10)
    fmt.Println(numbers) //-> [2 3 10]

	// The append() function creates a new backing array with a larger capacity to avoid creating a new backing array when the next append() is called.
	n1 := []int{10, 20, 30, 40}
    n1 = append(n1, 100)
    fmt.Println(len(n1), cap(n1))

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
    // if the slices don't have the same no of elements, it copies the minimum of length of the slices
    src := []int{10, 20, 30}
    dst := make([]int, len(src))
    nnm:= copy(dst, src)
    fmt.Println(src, dst, nnm) // => [10 20 30] [10 20 30] 3
 
}


 