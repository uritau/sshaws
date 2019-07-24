package main

import (
	"fmt"
	"os"

	"github.com/uritau/sshaws/helpers"
)

func showInstanceList(instanceList []helpers.Instance, user string) {
	if len(instanceList) == 0 {
		fmt.Printf("There are no instances matching your request.\n")
		os.Exit(0)
	}
	for idx, inst := range instanceList {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
	if len(instanceList) == 1 {
		fmt.Printf("\n\n")
		launchSSH(instanceList[0].IP, user)
		os.Exit(0)
	}
}
