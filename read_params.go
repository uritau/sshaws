package main

import "flag"

func read_params() configuration {
	app := flag.String("app", "*", "Tag Application of the instance")
	env := flag.String("env", "*", "Tag Environment of the application")
	name := flag.String("name", "*", "Name of the instance")
	region := flag.String("region", "eu-west-1", "Default AWS region")
	flag.Parse()
	return *NewConfiguration(*region, *env, *app, *name)
}
