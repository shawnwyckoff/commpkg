package gfs

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"runtime"
)

// Generate a new temp filename for cache
func NewTempFilename() (string, error) {
	if runtime.GOOS == "windows" {
		return "", errors.New("Unsupport windows for now")
	} else {
		return "/tmp/" + uuid.New().String() + ".temp", nil
	}
}
