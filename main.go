package main

import (
	"fmt"
	"strconv"
)

func select_instance_index(instance_list []instance) int {
	var input string
	fmt.Println("Which one do you want to shh in?")
	fmt.Scanln(&input)
	index, err := strconv.Atoi(input)
	if err != nil || index > len(instance_list)-1 || index < 0 {
		fmt.Println("Please enter a valid number.")
		index = select_instance_index(instance_list)
	}
	return index
}

func show_instance_list(instance_list []instance) {
	for idx, inst := range instance_list {
		fmt.Printf("[%d] %s - %s (%s, %s) \n", idx, inst.Name, inst.IP, inst.ID, inst.Size)
	}
}

func main() {
	config := read_params()
	region := config.Region
	env := config.Env
	app := config.App
	name := config.Name
	resp := filter_instances(region, env, app, name)
	instance_list := get_instances_info(resp)
	show_instance_list(instance_list)
	instance_number := select_instance_index(instance_list)
	launch_ssh(instance_list[instance_number].IP)
}
