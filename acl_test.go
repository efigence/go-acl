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
	err = acl.NewPermission(`write`)
	acl.AddBranch(`root.admin.some.very.long.nest`)
	acl.AddBranch(`root.guest`)
	err = acl.SetPerm(`root.admin.config`,`guest`,`read`)
	Convey("SettingPermissions",t,func() {
		So(err,ShouldBeNil)
	})


//	acl.SetPerm(`root.admin`,`admin`,`read`)
//	acl.SetPerm(`root.admin`,`gadmin`,`write`)
	canRead, _ := 	acl.Role("guest").HasPermission("root.admin.config","read")
	canAdmin, _ := acl.Role("guest").HasPermission("root.admin.config","write")

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
		So(acl.DebugDump(),ShouldNotEqual,nil)
	})

	canRecursiveRead, _ := acl.Role("guest").HasPermission("root.admin.config.db","read")
	canRecursiveAdmin, _ := acl.Role("admin").HasPermission("root.admin.config.db.some.thing","read")
	Convey("Recursively lookup read",t,func() {
		So(canRecursiveRead,ShouldBeTrue)
	})
	Convey("Recursively lookup admin",t,func() {
		So(canRecursiveAdmin,ShouldBeFalse)
	})




	_ = err
}
