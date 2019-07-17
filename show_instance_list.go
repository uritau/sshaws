package main

import "fmt"
import "os"

func show_instance_list(instance_list []instance) {
	if len(instance_list) == 0 {
		fmt.Printf("There are no instances matching your request.\n")
		os.Exit(0)
	}
	for idx, inst := range instance_list {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
	if len(instance_list) == 1 {
		fmt.Printf("\n\n")
		launch_ssh(instance_list[0].IP)
		os.Exit(0)
	}
}
