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
	newInstance := new(instance)
	newInstance.Name = name
	newInstance.IP = ip
	newInstance.ID = id
	newInstance.Size = size
	return newInstance
}

type instanceFilter struct {
	Tag             string
	Flag            string
	FlagShort       string
	FlagDescription string
	DefaultValue    string
}

func NewInstanceFilter(tag, flag, flagShort, flagDescription, defaultValue string) *instanceFilter {
	newInstanceFilter := new(instanceFilter)
	newInstanceFilter.Tag = tag
	newInstanceFilter.Flag = flag
	newInstanceFilter.FlagShort = flagShort
	newInstanceFilter.FlagDescription = flagDescription
	newInstanceFilter.DefaultValue = defaultValue
	return newInstanceFilter
}
