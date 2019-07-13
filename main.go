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

func define_filters() []instanceFilter {
	var filterList []instanceFilter
	filterList = append(filterList, *NewInstanceFilter("Name", "name", "n", "Instance Name", "*"))
	filterList = append(filterList, *NewInstanceFilter("Application", "app", "a", "Tag Application of the instance", "*"))
	filterList = append(filterList, *NewInstanceFilter("Environment", "env", "e", "Tag Environment of the instance", "*"))
	return filterList
}

func main() {

	filters := define_filters()

	config := read_params()
	region := config.Region
	resp := filter_instances(filters, region)
	instance_list := get_instances_info(resp)
	show_instance_list(instance_list)
	instance_number := select_instance_index(instance_list)
	launch_ssh(instance_list[instance_number].IP)
}
