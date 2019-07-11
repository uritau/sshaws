package main

import "fmt"

func show_instance_list(instance_list []instance) {
	for idx, inst := range instance_list {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
}
