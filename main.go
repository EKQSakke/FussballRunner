package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Started")
	commandPath := "G:/Projects/Godot/builds/Fussball.exe"
	cmd := exec.Command(commandPath)
	if err := cmd.Start(); err != nil {
		log.Println("Failed to start")
	}
	fmt.Println("Aloo")
	// how to kill ->
	// if err := cmd.Process.Kill(); err != nil {
	//	log.Fatal("failed to kill process: ", err)
	//	}
	if err := cmd.Wait(); err != nil {
		log.Printf("Cmd returned error: %v", err)
	}
}
