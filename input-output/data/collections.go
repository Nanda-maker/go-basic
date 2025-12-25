package data

var Countries [10]string // array of strings with fixed size 10
var Slice []int // slice of integers - dynamic size --they are actually chunks of arrays
var Codes map[int]string  // map of int keys and string values

func init(){
	Countries[0] = "Argentina"
	Countries[1] = "Brazil"
	Countries[2] = "Brazil"
	// qty := len(Countries)
}

 