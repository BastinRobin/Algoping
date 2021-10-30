package compile

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "os/exec"
)

// Execute the command from sandbox
func Sandbox(lang string, path string, inFile string, outFile string, errFile string) {
       
           inputReader, err := os.Open(inFile)
           if err != nil {
                log.Fatal(err) 
           }
           
           
           cmd := exec.Command(lang, path)
           cmd.Stdin = bufio.NewReader(inputReader) 

           // open the out file for redirecting standard output
           outfile, err := os.Create(outFile)
           if err != nil {
             panic(err)
           }
           defer outfile.Close()
           cmd.Stdout = outfile
           
           // open the error file for redirecting standard error
           errorfile, err := os.Create(errFile)
           if err != nil {
             panic(err)
           }
           defer errorfile.Close()
           cmd.Stderr = errorfile

           err = cmd.Run()
           if err != nil {
            fmt.Printf("Error: %v", err)
            os.Exit(1)
           }
}
