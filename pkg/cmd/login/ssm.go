package login

import (
	"fmt"
	"os"
	"time"
)

func launchSSM(destInstance string) {
	destination := destInstance
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	time.Sleep(1 * time.Second)
	awscliPath := getCommandPath("aws")

	fmt.Printf(">> Starting a new ssm session to %s\n", destination)
	ignoreInterruptionSignal()
	proc, err := os.StartProcess(awscliPath, []string{"aws", "ssm", "start-session", "--target", destination, "--document-name", "AWS-StartInteractiveCommand", "--parameters", "command=\"bash -l\""}, &pa)

	if err != nil {
		panic(err)
	}
	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Printf("<< Exited ssm session: %s\n", state.String())
}


