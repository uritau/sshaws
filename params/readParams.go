package params

import (
	"flag"
	"fmt"
	"os"
	"sshaws/helpers"
	"sshaws/version"
)

var (
	app            string
	env            string
	name           string
	region         string
	displayVersion bool
)

func init() {
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

	flag.StringVar(&app, "app", defaultApp, usageApp)
	flag.StringVar(&env, "env", defaultEnv, usageEnv)
	flag.StringVar(&name, "name", defaultName, usageName)
	flag.StringVar(&app, "a", defaultApp, usageApp+" [short mode]")
	flag.StringVar(&env, "e", defaultEnv, usageEnv+" [short mode]")
	flag.StringVar(&name, "n", defaultName, usageName+" [short mode]")
	flag.StringVar(&region, "region", defaultRegion, usageRegion)
	flag.BoolVar(&displayVersion, "version", false, "Display app version")
	flag.BoolVar(&displayVersion, "v", false, "Display app version")
}

func Read() helpers.Configuration {
	flag.Parse()
	if displayVersion {
		fmt.Printf("sshaws version %s \n", version.Get())
		os.Exit(0)
	}

	if flag.NArg() != 0 {
		name = flag.Args()[0]
	}
	return *helpers.NewConfiguration(region, env, app, name)
}
