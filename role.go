package acl

import (
//	"fmt"
)

type Role struct {
	Name string
	Type int
	instance *Instance
	valid bool
}

func (r *Role) HasPermission(path string, perm string) (ret bool, err error) {
	return r.instance.hasPermission(r, path, perm)
}
