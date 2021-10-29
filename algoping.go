package main

import (
	"fmt"
	"github.com/bastinrobin/algoping/compile"
)

func main() {
	
	// Get list of arguments from the command prompt 
	// eg: algoping language, filepath, stdin, stdout
	// expected return output, time, memory
	file := "/Users/developer/Desktop/trial/test.py"
	inFile := "/Users/developer/Desktop/trial/input.txt"
	outFile := "/Users/developer/Desktop/trial/output.txt"
	errFile := "/Users/developer/Desktop/trial/error.txt"
	lang := "/usr/bin/python"
	result, elapsed, memory := compile.Sandbox(lang, file, inFile, outFile, errFile)
	fmt.Println("Result: ", result)
	fmt.Println("Time: ", elapsed)
	fmt.Println("Mem: ", memory)
}
