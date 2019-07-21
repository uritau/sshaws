package main

import "flag"
import "fmt"
import "os"

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
	var displayVersion bool

	flag.StringVar(&app, "app", defaultApp, usageApp)
	flag.StringVar(&env, "env", defaultEnv, usageEnv)
	flag.StringVar(&name, "name", defaultName, usageName)
	flag.StringVar(&app, "a", defaultApp, usageApp+" [short mode]")
	flag.StringVar(&env, "e", defaultEnv, usageEnv+" [short mode]")
	flag.StringVar(&name, "n", defaultName, usageName+" [short mode]")
	flag.StringVar(&region, "region", defaultRegion, usageRegion)
	flag.BoolVar(&displayVersion, "version", false, "Display app version")
	flag.BoolVar(&displayVersion, "v", false, "Display app version")
	flag.Parse()
	if displayVersion {
		fmt.Printf("sshaws version %s \n", get_version())
		fmt.Printf("Original repository: https://github.com/uritau/sshaws/\n")
		os.Exit(0)
	}

	if flag.NArg() != 0 {
		name = flag.Args()[0]
	}
	return *NewConfiguration(region, env, app, name)
}
