package main

import (
	"fmt"
	"os"
	"time"
)

func launchSSH(destIP string, username string) {
	destination := destIP
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	time.Sleep(1 * time.Second)
	if username != "" {
		destination = username + "@" + destination
	}

	fmt.Printf(">> Starting a new ssh session to %s\n", destination)
	proc, err := os.StartProcess("/usr/bin/ssh", []string{"ssh", destination}, &pa)

	if err != nil {
		panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Printf("<< Exited ssh session: %s\n", state.String())
}
