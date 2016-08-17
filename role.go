package acl

type Role struct {
	Type int
	instance *Instance
}

func (r *Role) HasPermission(name string) (ret bool, err error) {
	return ret,err
}
