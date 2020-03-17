package gmachineid

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/crypto/hash"
	"github.com/shawnwyckoff/gpkg/dsa/stringz"
	"github.com/shawnwyckoff/gpkg/net/addr"
	"os/exec"
	"runtime"
	"strings"
)

// Get hardware UUID of MacOS
func MacosHardwareUUID() (string, error) {
	if runtime.GOOS != "darwin" {
		return "", errors.New("MacosHardwareUUID does not support " + runtime.GOOS)
	}
	output, err := exec.Command("system_profiler", "SPHardwareDataType").CombinedOutput()
	if err != nil {
		return "", err
	}
	uuid, err := stringz.SubstrBetween(string(output), "Hardware UUID:", "\n", true, true, false, false)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(uuid), nil
}

func nonMacosPhysicalMACs() (string, error) {
	ns, err := addr.GetAllNicNames()
	if err != nil {
		return "", err
	}

	var macs string

	for _, s := range ns {
		ni, _ := addr.GetNicInfo(s)
		if !ni.IsPhysical {
			continue
		}
		macs += ni.MAC
	}

	return hash.Md5Str(macs)
}

func Get() (string, error) {
	var str string
	var err error

	if runtime.GOOS == "darwin" {
		str, err = MacosHardwareUUID()
	} else if runtime.GOOS == "linux" || runtime.GOOS == "windows" {
		str, err = nonMacosPhysicalMACs()
	} else {
		return "", errors.New("Unsupported OS " + runtime.GOOS)
	}

	if err != nil {
		return "", err
	}
	md5, err := hash.Md5Str(str + "salt-duck-machid")
	if err != nil {
		return "", err
	}
	return md5, nil
}
