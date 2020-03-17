package users

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shawnwyckoff/gpkg/sys/fs"
	"github.com/shawnwyckoff/gpkg/sys/sysinfo"
	"os"
)

func IsRunAsAdmin() (bool, error) {
	rootDir, err := sysinfo.SysRootDir()
	if err != nil {
		return false, err
	}

	for {
		uuid, err := uuid.NewUUID()
		if err != nil {
			fmt.Println(err.Error())
			return false, err
		}

		testDir := rootDir + uuid.String()
		pi, err := fs.GetPathInfo(testDir)
		if err != nil {
			fmt.Println(err.Error())
			return false, err
		}

		if pi.Exist {
			continue
		}
		err = os.Mkdir(testDir, os.ModePerm)
		os.Remove(testDir)
		if err == nil {
			return true, nil
		} else {
			return false, nil
		}
	}
}

// untested
func IsRunAsAdmin2() (bool, error) {
	fmt.Println(os.Geteuid())
	if os.Geteuid() == 0 {
		return true, nil
	}
	return false, nil
}

// if run as normal user, this function will restart current app in sudo/administrator
func AutoSwitchRoot() error {
	return nil
}
