package login

import (
	"fmt"
	"os"
)

func showIPsList(instanceList []Instance) {
	if len(instanceList) == 0 {
		os.Exit(0)
	}
	for _, inst := range instanceList {
		fmt.Printf("%s\n", inst.IP)
	}
	os.Exit(0)
}
