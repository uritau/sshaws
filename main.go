package main

import (
	"sshaws/helpers"
	"sshaws/params"
)

func main() {
	config := params.Read()
	env, app, name, region := helpers.ReturnConfiguration(config)
	rawInstanceList := filterInstances(region, env, app, name)
	instances := getInstancesInfo(rawInstanceList)
	showInstanceList(instances)
	selectedInstance := selectInstanceIndex(instances)
	launchSsh(instances[selectedInstance].IP)
}
