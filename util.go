package acl

import(
	"strings"
)


func splitPath(path string) (s []string, err error) {
	return strings.Split(path, "."), err

}
