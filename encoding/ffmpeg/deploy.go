package ffmpeg

import (
	"github.com/shawnwyckoff/gpkg/sys/gcmd"
	"strings"
)

// make ffmpeg a netx service, and get a highly available live cluster

func IsInstalled() bool {
	result, _ := gcmd.ExecWaitReturn("ffmpeg")
	if strings.Contains(string(result), "version") {
		return true
	}
	return false
}
