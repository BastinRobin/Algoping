package compile

import (
	"fmt"
	"os"
	"os/exec"
)

// Sandbox executes the code and returns output 
func Sandbox(lang string, path string, inFile string, outFile string, errFile string) (string, string, string){

	filePath := fmt.Sprintf("%v <%v", path, inFile)
	cmd := exec.Command(lang, filePath)
	cmd.Stderr := os.Stderr
	return "Finished", "time", "mem"
}


