package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) != 3{
		fmt.Println("Usage: go run. input.txt output.txt")
		os.Exit(1)
	}

	input := os.Args[1]
	ouput := os.Args[2]

	data , err := os.ReadFile(input)
	if err != nil{
		fmt.Println("Error reading file",err)
		os.Exit(1)
	}

	content := string(data)
	result := applyFunctions(content)

	err = os.WriteFile(ouput, []byte(result+"\n"), 0644)
	if err != nil{
		fmt.Println("Error writing to file", err)
		os.Exit(1)
	}
}