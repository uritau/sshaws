package login

import (
	"fmt"
	"os"
)

func showInstanceList(instanceList []Instance, user string) {
	if len(instanceList) == 0 {
		fmt.Printf("There are no instances matching your request.\n")
		os.Exit(0)
	}
	for idx, inst := range instanceList {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
}
