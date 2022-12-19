package login

import (
	"time"
)

// Instance struct
type Instance struct {
	Name       string
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
	newInstance.IP = ip
	newInstance.ID = id
	newInstance.Size = size
	newInstance.AZ = az
	newInstance.LaunchTime = launchTime
	return newInstance
}
