package main

import (
	"github.com/bastinrobin/algoping/compile"
)

func main() {
	
	// Get list of arguments from the command prompt 
	// eg: algoping language, filepath, stdin, stdout
	// expected return output, time, memory
	file := "/Users/developer/Desktop/trial/com.py"
	inFile := "/Users/developer/Desktop/trial/input.txt"
	outFile := "/Users/developer/Desktop/trial/output.txt"
	errFile := "/Users/developer/Desktop/trial/error.txt"
	lang := "/usr/local/bin/python3"
	compile.Sandbox(lang, file, inFile, outFile, errFile)
}
