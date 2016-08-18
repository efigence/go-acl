package acl

import (
	"encoding/json"
	"fmt"
	"sync"
)

// acl-s are multi-part names that use / as divider
// for example

type Instance struct {
	aclMap      *Branch
	roles       map[string]*Role
	permissions map[string]*Permission
	sync.RWMutex
}

func NewInstance() *Instance {
	var i Instance
	i.aclMap = newBranch("root")
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
	r.Name = name
	r.Type = roltype
	r.instance = acl // return ref so we can get to role's ACLs of user got role struct
	r.valid=true
	acl.roles[name] = &r
	acl.Unlock()
	return e
}

func (acl *Instance) AddBranch(name string) (err error) {
	splitP, err := splitPath(name)
	if err != nil {
		return err
	}
	currentPos := []string{}
	var currentPtr = acl.aclMap
	for _, val := range splitP {
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

func (acl *Instance) SetPerm(name string, role string, perm string) (err error) {
	path, err := splitPath(name)
	if err != nil { return err }
	branch, err := acl.getBranchPtr(path)
	fmt.Printf("---- %+v \n", branch)
	if err != nil { return err }
	if _, ok := branch.Perms[role]; !ok {
		branch.Perms[role] = make(map[string]bool)
	}
	branch.Perms[role][perm]=true
	return err
}

func (acl *Instance) getBranchPtr(path []string) (branch *Branch, err error) {
	currentPtr := acl.aclMap
	for _, val := range path {
		if child, ok := currentPtr.Branch[val]; ok {
			currentPtr = child
		} else {
			return branch, fmt.Errorf("No branch with that path [%+v], create it first",path)
		}
	}
	return currentPtr, err
}

func (acl *Instance) Role(name string) (z *Role) {
	if r, ok := acl.roles[name]; ok {
		return r
	} else {
		return &Role{
			valid: false,
		}
	}
}

func (acl *Instance) DebugDump() string {
	out, _ := json.MarshalIndent(struct{
		Map      *Branch
		Roles       map[string]*Role
		Permissions map[string]*Permission
	}{
		acl.aclMap,
		acl.roles,
		acl.permissions,
	}, "", "  ")
	return string(out)
}
