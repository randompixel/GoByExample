package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Provide arguments as slice
	args := []string{"ls", "-a", "-l", "-h"}

	// Get Environment variables
	env := os.Environ()

	// If successful, execution of go stops here and replaced by exec
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
