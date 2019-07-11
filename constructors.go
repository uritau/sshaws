package main

type configuration struct {
	App    string
	Env    string
	Region string
	Name   string
}

func NewConfiguration(region, env, app, name string) *configuration {
	config := new(configuration)
	config.App = app
	config.Env = env
	config.Name = name
	config.Region = region
	return config
}

type instance struct {
	Name string
	IP   string
	ID   string
	Size string
}

func NewInstance(name, ip, id, size string) *instance {
	new_instance := new(instance)
	new_instance.Name = name
	new_instance.IP = ip
	new_instance.ID = id
	new_instance.Size = size
	return new_instance
}
