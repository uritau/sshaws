package login

//Instance struct
type Instance struct {
	Name string
	IP   string
	ID   string
	Size string
}

//NewInstance initialize struct Instnace.
func NewInstance(name, ip, id, size string) *Instance {
	newInstance := new(Instance)
	newInstance.Name = name
	newInstance.IP = ip
	newInstance.ID = id
	newInstance.Size = size
	return newInstance
}
