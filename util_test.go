package acl

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUtil(t *testing.T) {
	path, err := splitPath(`some.long.path`)
	Convey("splitLongPath", t, func() {
		So(err, ShouldEqual, nil)
		So(path, ShouldResemble, []string{`some`,`long`,`path`})
	})
	path, err = splitPath(`shortpath`)
	Convey("splitShortPath", t, func() {
		So(err, ShouldEqual, nil)
		So(path, ShouldResemble, []string{`shortpath`})
	})
}
