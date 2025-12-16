package fileio

import "os"

func ReadTextFile(filename string) (string,error) {
	content, err := os.ReadFile(filename)
	if err !=nil {
// we couldn't read the file
return "", err
	} else {
		//Read operation susccessful
		return string(content),nil
	}
}
