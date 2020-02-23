package logz

import (
	"fmt"
	"github.com/shawnwyckoff/commpkg/sys/machine_id"
	"github.com/shawnwyckoff/commpkg/sys/proc"
	"github.com/shawnwyckoff/commpkg/sys/sysinfo"
	"path"
)

type (
	Config struct {
		SaveDisk       bool
		SaveDir        string
		FileNameFormat string
		PrintScreen    bool

		MachId  string
		AppName string
	}
)

func DefaultConfig() (*Config, error) {
	res := &Config{}

	// Disk save filename format...
	res.FileNameFormat = "2006-01-02.log" // YEAR-MONTH-DAY.log
	res.SaveDisk = true
	res.PrintScreen = true
	// Get default logs directory
	pi, err := proc.GetProcInfo(proc.GetPidOfMyself())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	dir := sysinfo.GetAppLogFolder(pi.Name)
	res.SaveDir = dir
	// Get machine Id.
	id, err := machine_id.Get()
	if err != nil {
		return nil, err
	}
	res.MachId = id
	// Get app name.
	fn, err := proc.SelfPath()
	if err != nil {
		return nil, err
	}
	res.AppName = path.Base(fn)

	return res, nil
}
