package main

import (
	"fmt"
	"os"
)

func launch_ssh(dest_ip string) {
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}

	fmt.Printf(">> Starting a new ssh session to %s\n", dest_ip)
	proc, err := os.StartProcess("/usr/bin/ssh", []string{"ssh", dest_ip}, &pa)
	if err != nil {
		panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Printf("<< Exited ssh session: %s\n", state.String())
}
