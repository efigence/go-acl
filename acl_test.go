package acl

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService(t *testing.T) {
	acl := NewInstance()
	err := acl.NewRole("guest", RoleUser)
	err = acl.NewRole("admin", RoleUser)
	err = acl.NewPermission(`read`)
	acl.AddBranch(`root.admin.some.very.long.nest`)
	acl.AddBranch(`root.guest`)
	acl.AddPerm(`root.admin`,`read`)
	canRead, _ := 	acl.Role("guest").HasPermission("read")
	canAdmin, _ := acl.Role("guest").HasPermission("admin")

	Convey("Creating acl", t, func() {
		So(acl, ShouldNotEqual, nil)
	})
	Convey("Check read permissions",t,func() {
		So(canRead, ShouldBeTrue)
	})
	Convey("Check admin permissions",t,func() {
		So(canAdmin, ShouldBeFalse)
	})
	Convey("DebugDump",t,func() {
		So(acl.DebugDump(),ShouldEqual,nil)
	})


	_ = err
}
