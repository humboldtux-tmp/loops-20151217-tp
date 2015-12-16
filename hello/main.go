package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Printf("hello loOPS-20151217\n")
	fmt.Printf("running pkg-config --cflags python2 now...\n")
	cmd := exec.Command("pkg-config", "--cflags", "python2")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}
