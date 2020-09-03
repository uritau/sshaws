package params

import (
	"flag"
	"fmt"
	"os"

	"github.com/uritau/sshaws/helpers"
	"github.com/uritau/sshaws/version"
)

var (
	app            string
	env            string
	name           string
	region         string
	displayVersion bool
	username       string
	silent         bool
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
		defaultUser   = ""
		usageUser     = "SSH login name"
		defaultSilent = false
		usageSilent   = "Show only IP"
		defaultSSH = false
		usageSSH   = "Use SSH instead of SSM"

	)

	flag.StringVar(&app, "app", defaultApp, usageApp)
	flag.StringVar(&env, "env", defaultEnv, usageEnv)
	flag.StringVar(&name, "name", defaultName, usageName)
	flag.BoolVar(&silent, "silent", defaultSilent, usageSilent)
	flag.BoolVar(&ssh, "ssh", defaultSSH, usageSSH+" [short mode]")
	flag.StringVar(&app, "a", defaultApp, usageApp+" [short mode]")
	flag.StringVar(&env, "e", defaultEnv, usageEnv+" [short mode]")
	flag.StringVar(&name, "n", defaultName, usageName+" [short mode]")
	flag.StringVar(&region, "region", defaultRegion, usageRegion)
	flag.BoolVar(&displayVersion, "version", false, "Display app version")
	flag.StringVar(&username, "l", defaultUser, usageUser)
}

//Read function
func Read() helpers.Configuration {
	flag.Parse()
	if displayVersion {
		fmt.Printf("sshaws version %s \n", version.Get())
		os.Exit(0)
	}

	if flag.NArg() != 0 {
		name = flag.Args()[0]
	}
	return *helpers.NewConfiguration(region, env, app, name, username, silent, ssh)
}
