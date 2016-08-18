package acl

import (
	"encoding/json"
	"fmt"
	"sync"
)

// acl-s are multi-part names that use / as divider
// for example

type Instance struct {
	aclMap      map[string]*Branch
	roles       map[string]*Role
	permissions map[string]*Permission
	sync.RWMutex
}

func NewInstance() *Instance {
	var i Instance
	i.aclMap = make(map[string]*Branch)
	i.roles = make(map[string]*Role)
	i.permissions = make(map[string]*Permission)
	return &i
}

const (
	RoleUser  = iota
	RoleGroup = iota
	// <- insert new roles here
	RoleLast = iota //end of enum pointer, used by validator, must be last
)

type Permission struct {
	Name       string
	nameParsed []string
}

type Acl struct {
	Name   string
	Grants []Permission
}

func (acl *Instance) NewRole(name string, roltype int) (e error) {
	if roltype >= RoleLast {
		return fmt.Errorf("Invalid role ID, either you did not use package constants or are trying to import new data into old acl lib!")
	}
	var r Role
	acl.Lock()
	r.Type = roltype
	r.instance = acl // return ref so we can get to role's ACLs of user got role struct
	acl.Unlock()
	return e
}

func (acl *Instance) AddBranch(name string) (err error) {
	splitP, err := splitPath(name)
	if err != nil {
		return err
	}
	currentPos := []string{}
	var currentPtr *Branch
	for idx, val := range splitP {
		if idx == 0 {
			if _, ok := acl.aclMap[val]; !ok {
				acl.aclMap[val] = newBranch(val)
			}
			currentPtr = acl.aclMap[val]
			continue

		}
		currentPos = append(currentPos, val)
		if _, ok := currentPtr.Branch[val]; !ok {
			currentPtr.Branch[val] = newBranch(val)
		}
		currentPtr = currentPtr.Branch[val]

	}
	_ = currentPtr
	return err
}

func (acl *Instance) NewPermission(name string) (err error) {
	var perm Permission
	perm.Name = name
	acl.Lock()
	acl.permissions[name] = &perm
	acl.Unlock()
	return err
}

func (acl *Instance) AddPerm(name string, perm string) (err error) {
	return err
}

func (acl *Instance) Role(name string) *Role {
	return acl.roles[name]
}

func (acl *Instance) DebugDump() string {
	out, _ := json.MarshalIndent(struct{
		Map      map[string]*Branch
		Roles       map[string]*Role
		Permissions map[string]*Permission
	}{
		acl.aclMap,
		acl.roles,
		acl.permissions,
	}, "", "    ")
	return string(out)
}
