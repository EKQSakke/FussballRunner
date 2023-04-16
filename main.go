package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Started")
	commandPath := "G:/Projects/Godot/builds/Fussball.exe"
	out, err := exec.Command(commandPath).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
