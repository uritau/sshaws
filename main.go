package main

import (
	"github.com/uritau/sshaws/helpers"
	"github.com/uritau/sshaws/params"
)

func main() {
	config := params.Read()
	env, app, name, region, user, silent, ssh := helpers.ReturnConfiguration(config)
	rawInstanceList := filterInstances(region, env, app, name, silent, ssh)
	instances := getInstancesInfo(rawInstanceList)
	if silent {
		showIPsList(instances)
	} else {
		showInstanceList(instances, user)
	}
	selectedInstance := selectInstanceIndex(instances)

	if ssh {
		launchSSH(instances[selectedInstance].IP, user)
	} else {
		launchSSM(instances[selectedInstance].ID)
	}
}
