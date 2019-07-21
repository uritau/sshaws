package main

import "fmt"
import "os"

func showInstanceList(instanceList []instance) {
	if len(instanceList) == 0 {
		fmt.Printf("There are no instances matching your request.\n")
		os.Exit(0)
	}
	for idx, inst := range instanceList {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
	if len(instanceList) == 1 {
		fmt.Printf("\n\n")
		launchSsh(instanceList[0].IP)
		os.Exit(0)
	}
}
