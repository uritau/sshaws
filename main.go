package main

import (
	"sshaws/helpers"
	"sshaws/params"
)

func main() {
	config := params.Read()
	env, app, name, region, user := helpers.ReturnConfiguration(config)
	rawInstanceList := filterInstances(region, env, app, name)
	instances := getInstancesInfo(rawInstanceList)
	showInstanceList(instances, user)
	selectedInstance := selectInstanceIndex(instances)
	launchSSH(instances[selectedInstance].IP, user)
}
