package main

import (
	"github.com/uritau/sshaws/helpers"
	"github.com/uritau/sshaws/params"
)

func main() {
	config := params.Read()
	env, app, name, region, user, silent := helpers.ReturnConfiguration(config)
	rawInstanceList := filterInstances(region, env, app, name, silent)
	instances := getInstancesInfo(rawInstanceList)
	if silent {
		showIPsList(instances)
	} else {
		showInstanceList(instances, user)
	}
	selectedInstance := selectInstanceIndex(instances)
	launchSSH(instances[selectedInstance].IP, user)
}
