package acl

type Branch struct {
	Name string
	Branch map[string]*Branch
	Perms map[string]map[string]bool
}


func newBranch(name string) (*Branch) {
	var a Branch
	a.Name = name
	a.Branch = make(map[string]*Branch)
	a.Perms = make(map[string]map[string]bool)
	return &a
}


//func (s *Subtree)
