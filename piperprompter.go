package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Incorrect usage: Please append the command to be run.")
		return
	}

	command := exec.Command(args[0], args[1:]...)

	fmt.Printf("Enter filename: ")
	var filename string
	fmt.Scanln(&filename)

	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, out.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Saved to %s!\n", filename)

}
