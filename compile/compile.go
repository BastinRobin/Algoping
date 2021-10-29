package compile

import (
	"fmt"
	"os/exec"
)

// Sandbox executes the code and returns output 
func Sandbox(lang string, path string, inFile string, outFile string, errFile string) (string, string, string){
	
	fullPath := fmt.Sprintf("%v <%v 1>%v 2>%v", path, inFile, outFile, errFile)
	fmt.Println(fullPath)	
	exec.Command(lang, fullPath)
	return "Finished", "time", "mem"
}


