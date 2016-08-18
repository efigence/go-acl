[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/efigence/go-acl)

## go-acl

Simple lib to manage (add roles/check permissions) ACLs. WiP


### Format

ACL consist of path, action, role, and allow/deny (true/false flags)

#### Path

Path consist of UTF8 elements separated with `.` character, up to `MaxPathDeph` (currently 64). Elements itself can't contain `.`

#### Role Name

Just a golang string for now.

#### Action

Golang string. If the "default" action is needed, the `#` should be used. It is the default if argument is not specified.

### ACL hierarchy

Default policy is to deny e.g. if there is no ACL in the path to root, the result will be deny/false

TO change that just insert allow rule at top of the tree via `acl.SetPerm("myroot","guest","read")`

ACLs are checked from *most specific* to *least specific* path, regardless of if it is deny or allow

So if `el1.el2.el3` is set to "deny" but `el1.el2` to allow, the query for `el1.el2.el10` will return allow but query for `el1.el2.el3` will return deny


ACLs are structured in hierarchical tree wit

To create ACL object:

    ```go
    acl := NewInstance()
    acl.newRole("guest",acl.RoleUser)
    acl.newRole("editor",acl.RoleUser)
    acl.newRole("admin",acl.RoleUser)
    acl.SetPerm("blog.post.read", "guest", "use") // allow guest to read posts
    acl.SetPerm("blog.post", "editor", "use", false) // allow editor to change them
    acl.SetPerm("blog.post.delete", "editor", "use", false) // but not to remove them
    acl.SetPerm("blog.post", "admin", "read") // allow all to admin


    if acl.Role("guest").HasPermissions("blog.post.read","use") {
    fmt.Println("I can read stuff")

    ```
