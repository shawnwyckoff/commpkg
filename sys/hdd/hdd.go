package hdd

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/dsa/volume"
	"github.com/shirou/gopsutil/disk"
)

// https://github.com/cydev/du
// https://github.com/ricochet2200/go-disk-usage
// https://gist.github.com/lunny/9828326
// http://wendal.net/2012/1224.html
// https://github.com/lxn/win
// https://github.com/AllenDang/w32

type VolumeInfo struct {
	FilesystemType string
	Available      volume.Volume
	Free           volume.Volume
	Total          volume.Volume
}

func GetVolumeInfo(volumePath string) (*VolumeInfo, error) {
	var vi VolumeInfo
	du, err := disk.Usage(volumePath)
	if err != nil {
		return nil, err
	}
	vi.Available, err = volume.FromByteSize(float64(du.Total))
	if err != nil {
		return nil, err
	}
	vi.Free, err = volume.FromByteSize(float64(du.Total))
	if err != nil {
		return nil, err
	}
	vi.Total, err = volume.FromByteSize(float64(du.Total))
	if err != nil {
		return nil, err
	}
	vi.FilesystemType = du.Fstype
	return &vi, nil
}

// returns partition path, same as mountpoint in Unix, or logical drive in Windows
func ListVolumes() (partitionPath []string, err error) {
	ps, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}
	if len(ps) == 0 {
		return nil, errors.New("Get disk partitions error")
	}

	var result []string
	for _, item := range ps {
		result = append(result, item.Mountpoint)
	}
	return result, nil
}
