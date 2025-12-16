package main

import (
	"fmt"
	fileutils "go/files/fileutils"
	"os"
)

func main(){
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadTextFile(filePath)
	// A nil error represents no error.
	if err == nil{
		fmt.Println(c)
		// newContent := "Original: " + c + " \n Double the original: " + c + c
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", c,c,c)
		fileutils.WriteToFile(filePath + ".output.txt",newContent)
	} else {
		fmt.Printf("Error %v",err)
	}
}