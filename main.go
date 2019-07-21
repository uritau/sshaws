package main

func main() {
	config := readParams()
	env, app, name, region := ReturnConfiguration(config)
	rawInstanceList := filterInstances(region, env, app, name)
	instances := getInstancesInfo(rawInstanceList)
	showInstanceList(instances)
	selectedInstance := selectInstanceIndex(instances)
	launchSsh(instances[selectedInstance].IP)
}
