package acl

import (
	"fmt"
)

type Role struct {
	Name     string
	Type     int
	instance *Instance
	valid    bool
}

// Check if path have a permission for role. Omit perm for "default" (#) one
func (r *Role) HasPerm(path string, perm ...string) (ret bool, err error) {
	if len(perm) == 0 {
		return r.instance.hasPermission(r, path, "#")
	} else if len(perm) == 1 {
		return r.instance.hasPermission(r, path, perm[0])
	} else {
		return false, fmt.Errorf("Wrong number of arguments [%d], should be 0 or 1", len(perm))
	}

}
