package main

import (
	"flag"
	"fmt"
	"os"

	cmd "sshaws/pkg/cmd"
	auth "sshaws/pkg/cmd/login"
)

func main() {
	var (
		name           string
		region         string
		displayVersion bool
		username       string
		silent         bool
		ssh            bool
		pushKey		   bool
	)

	const (
		defaultName   	= "*"
		usageName     	= "Instance Name"
		defaultRegion 	= "eu-west-1"
		usageRegion   	= "AWS Region"
		defaultUser   	= ""
		usageUser     	= "SSH login name"
		defaultSilent 	= false
		usageSilent   	= "Show only IP"
		defaultSSH    	= false
		usageSSH      	= "Use SSH instead of SSM"
		defaultPushKey 	= false
		usagePushKey	= "Push temporal public key to instance"
	)

	flag.StringVar(&name, "name", defaultName, usageName)
	flag.BoolVar(&silent, "silent", defaultSilent, usageSilent)
	flag.BoolVar(&ssh, "ssh", defaultSSH, usageSSH+" [short mode]")
	flag.BoolVar(&pushKey, "k", defaultPushKey, usagePushKey+" [short mode]")
	flag.StringVar(&name, "n", defaultName, usageName+" [short mode]")
	flag.StringVar(&region, "region", defaultRegion, usageRegion)
	flag.BoolVar(&displayVersion, "version", false, "Display app version")
	flag.StringVar(&username, "l", defaultUser, usageUser)

	flag.Parse()

	if displayVersion {
		fmt.Printf("sshaws version %s \n", cmd.Get())
		os.Exit(0)
	}

	if flag.NArg() != 0 {
		name = flag.Args()[0]
	}

	auth.NewLogin(name, region, username, silent, ssh, pushKey)
}
