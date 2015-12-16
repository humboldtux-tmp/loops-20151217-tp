package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", rootHandle)
	port := ":8080"
	log.Printf("listening on: http://localhost%s\n", port)
	err := http.ListenAndServe("0.0.0.0"+port, nil)
	if err != nil {
		log.Fatalf("error closing web server: %v\n", err)
	}
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello LoOPS 20151217!\n")
	fmt.Fprintf(w, "\n\n--- running external command...\n\n>>> pkg-config --cflags python2\n")
	cmd := exec.Command("pkg-config", "--cflags", "python2")
	cmd.Stdout = w
	cmd.Stderr = w
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
	}
}
