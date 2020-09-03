package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func launchSSM(destInstance string) {
	destination := destInstance
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	time.Sleep(1 * time.Second)
	awscliPath := getAWSCLIpath()

	fmt.Printf(">> Starting a new ssm session to %s\n", destination)
	proc, err := os.StartProcess(awscliPath, []string{"aws", "ssm", "start-session", "--target", destination}, &pa)

	if err != nil {
		panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Printf("<< Exited ssm session: %s\n", state.String())
}

func getAWSCLIpath() string {
    cmd := exec.Command("which", "aws")
	path, err := cmd.CombinedOutput()
	fmt.Printf(">>AWS CLI executable PATH %s\n", path)

    if err != nil {
		fmt.Printf("aws cli executable not found by sshaws.\n Failed with %s\n", err)
    }
    return strings.TrimSuffix(string(path), "\n")
}
