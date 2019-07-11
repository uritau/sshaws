package main

func main() {
	config := read_params()
	env, app, name, region := ReturnConfiguration(config)
	raw_instance_list := filter_instances(region, env, app, name)
	instances := get_instances_info(raw_instance_list)
	show_instance_list(instances)
	selected_instance := select_instance_index(instances)
	launch_ssh(instances[selected_instance].IP)
}
