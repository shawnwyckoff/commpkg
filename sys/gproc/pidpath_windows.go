package gproc

import "github.com/shawnwyckoff/gpkg/apputil/gerror"

// TODO
func GetExePathFromPid(pid int) (path string, err error) {
	return "", gerror.Errorf("windows unsupported for now")
}
