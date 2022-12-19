package login

import (
	"strings"
	"time"
)

const subdomain string = "clidom.es"

// Instance struct
type Instance struct {
	Name       string
	DNS        string
	IP         string
	ID         string
	Size       string
	AZ         string
	LaunchTime time.Time
}

// NewInstance initialize struct Instance.
func NewInstance(name, ip, id, size, az string, launchTime time.Time) *Instance {
	newInstance := new(Instance)
	newInstance.Name = name
	dns := convertNameToDNS(name)
	newInstance.DNS = dns
	newInstance.IP = ip
	newInstance.ID = id
	newInstance.Size = size
	newInstance.AZ = az
	newInstance.LaunchTime = launchTime
	return newInstance
}

func convertNameToDNS(name string) string {
	if !strings.Contains(name, "bastion") {
		return name
	}
	replaceBastionV2 := strings.ReplaceAll(name, "-v2-", "")
	replaceBastion := strings.ReplaceAll(replaceBastionV2, "bastion", "bastion.")
	replaceProd := strings.ReplaceAll(replaceBastion, "-prod", ".prod")
	replaceStaging := strings.ReplaceAll(replaceProd, "-staging", ".staging")
	dnsSplit := strings.Split(replaceStaging, ".")
	dns := dnsSplit[1] + "." + dnsSplit[0] + "." + dnsSplit[2] + "." + subdomain
	return dns
}
