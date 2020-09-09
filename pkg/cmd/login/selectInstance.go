package login

import (
	"fmt"
	"strconv"
)

func selectInstanceIndex(instanceList []Instance) int {
	var input string
	var index int
	if len(instanceList) == 1 {
		fmt.Printf("\n\n")
		index = 0
	} else {
		fmt.Println("\nWhich one do you want to ssh in?")
		fmt.Scanln(&input)
		index, err := strconv.Atoi(input)
		if err != nil || index > len(instanceList)-1 || index < 0 {
			fmt.Println("Please enter a valid number.")
			index = selectInstanceIndex(instanceList)
		}
	}
	return index
}
