package main

import "flag"

func read_params() configuration {
	const (
		defaultApp    = "*"
		usageApp      = "Tag Application of the instance"
		defaultEnv    = "*"
		usageEnv      = "Tag Environment of the instance"
		defaultName   = "*"
		usageName     = "Instance Name"
		defaultRegion = "eu-west-1"
		usageRegion   = "AWS Region"
	)
	var app string
	var env string
	var name string
	var region string

	flag.StringVar(&app, "app", defaultApp, usageApp)
	flag.StringVar(&env, "env", defaultEnv, usageEnv)
	flag.StringVar(&name, "name", defaultName, usageName)
	flag.StringVar(&app, "a", defaultApp, usageApp+" [short mode]")
	flag.StringVar(&env, "e", defaultEnv, usageEnv+" [short mode]")
	flag.StringVar(&name, "n", defaultName, usageName+" [short mode]")
	flag.StringVar(&region, "region", defaultRegion, usageRegion)
	flag.Parse()
	return *NewConfiguration(region, env, app, name)
}
