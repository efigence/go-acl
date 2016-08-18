package acl

import (
	"fmt"
)

type Role struct {
	Name string
	Type int
	instance *Instance
	valid bool
}

func (r *Role) HasPermission(path string, perm string) (ret bool, err error) {
	if (!r.valid) {return false, fmt.Errorf("No such role")}
	fmt.Printf("ffff: %+v", r.Name)
	pathSplit, err := splitPath(path)
	if err != nil { return false, err }
	branch, err := r.instance.getBranchPtr(pathSplit)
	if err != nil { return false, err }
	if _, ok := branch.Perms[r.Name]; ok {
		if ret, ok := branch.Perms[r.Name][perm]; ok {
			return ret, err
		} else {
			return false, err
		}

	} else {
		return false, err
	}
//	path, err
	return ret,err
}
