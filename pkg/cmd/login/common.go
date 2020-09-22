package login

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

// Capture SIGINT signal to prevent that Control+C causes session termination
func ignoreInterruptionSignal() {
	c := make(chan os.Signal, syscall.SIGINT)
	signal.Notify(c, os.Interrupt)
}

func getCommandPath(command string) string {
	cmd := exec.Command("which", command)
	path, err := cmd.CombinedOutput()
	fmt.Printf(">>%s executable PATH %s\n", command, path)

	if err != nil {
		fmt.Printf("%s executable not found by sshaws.\n Failed with %s\n", command, err)
	}
	return strings.TrimSuffix(string(path), "\n")
}

func getRegionFromAZ(az string) string {
	return az[:len(az)-1]
}