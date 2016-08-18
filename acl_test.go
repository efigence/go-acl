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
	acl.AddBranch("test1.admin.some.very.long.nest")
	acl.AddBranch("test1.guest")
	err = acl.SetPerm("test1.admin.config","guest","read")
	Convey("SettingPermissions",t,func() {
		So(err,ShouldBeNil)
	})


	canRead, _ := 	acl.Role("guest").HasPermission("test1.admin.config","read")
	canAdmin, _ := acl.Role("guest").HasPermission("test1.admin.config","write")

	Convey("Creating acl", t, func() {
		So(acl, ShouldNotEqual, nil)
	})
	Convey("Check read permissions",t,func() {
		So(canRead, ShouldBeTrue)
	})
	Convey("Check admin permissions",t,func() {
		So(canAdmin, ShouldBeFalse)
	})

	canRecursiveRead, _ := acl.Role("guest").HasPermission("test1.admin.config.db","read")
	canRecursiveAdmin, _ := acl.Role("admin").HasPermission("test1.admin.config.db.some.thing","read")
	Convey("Recursively lookup read",t,func() {
		So(canRecursiveRead,ShouldBeTrue)
	})
	Convey("Recursively lookup admin",t,func() {
		So(canRecursiveAdmin,ShouldBeFalse)
	})

	acl.NewRole("author", RoleUser)
	err = acl.NewPermission("use")
	acl.SetPerm("test2.blog.editor", "author", "use")
	acl.SetPerm("test2.blog.editor.delete", "author", "use",false)
	canEditPost, _ := acl.Role("author").HasPermission("test2.blog.editor.zep","use")
	canDeletePost, _ := acl.Role("author").HasPermission("test2.blog.editor.delete","use")
	Convey("Deny specific leaf in tree",t,func() {
		So(canDeletePost,ShouldBeFalse)
	})
	Convey("Allow all but specific path",t,func() {
		So(canEditPost,ShouldBeTrue)
	})

//	Convey("DebugDump",t,func() {
//		So(acl.DebugDump(),ShouldEqual,nil)
//	})
	_ = err
}
