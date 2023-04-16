package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Started")

	out, err := exec.Command("fussballerinner.exe").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
