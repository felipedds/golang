package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func shellCommand(command string) (error, string, string) {
	cmd := exec.Command("bash", "-c", command)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return err, stdout.String(), stderr.String()
}

func shellScript() {
	cmd := exec.Command("bash", "script.sh")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Output:\n", stdout.String())
}

func main() {
	// Shell Command
	err, out, errout := shellCommand("ls -ltr")
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	fmt.Printf("stdout-> %s\n", out)
	fmt.Println("stderr->")
	fmt.Println(errout)

	// Shell Script
	shellScript()
}
