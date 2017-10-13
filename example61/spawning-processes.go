package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	// Create helper to process
	dateCmd := exec.Command("date")

	// Run process and collect it's output
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()                 // connect to cmd stdIn
	grepOut, _ := grepCmd.StdoutPipe()               // connect to cmd stdOut
	grepCmd.Start()                                  // start process
	grepIn.Write([]byte("hello grep\ngoodbye grep")) // write to input
	grepIn.Close()                                   // close input stream
	grepBytes, _ := ioutil.ReadAll(grepOut)          // read output
	grepCmd.Wait()                                   // wait for process to exit
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Need to define commands as list, not a single string
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
