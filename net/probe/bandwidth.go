package probe

import (
	"github.com/shawnwyckoff/gpkg/dsa/speed"
	"github.com/shawnwyckoff/gpkg/net/probe/util/speedtest"
)

// github.com/showwin/speedtest-go 测试下来功能正常，但代码较乱 另外，代码比较清爽但是star很少且没有验证的库 https://github.com/sivel/speedtest/blob/master/speedtest.go

func WanDownloadSpeedTest() (speed.Speed, error) {
	return speedtest.TestDownloadSpeed()
}

func WanUploadSpeedTest() (speed.Speed, error) {
	return speedtest.TestUploadSpeed()
}

// 测试端到端的网速
// https://github.com/blang/speedtest
// https://github.com/DhruvKalaria/SpeedTest
// https://github.com/itimofeev/netspeed
func End2EndUploadSpeedTest() (*speed.Speed, error) {
	return nil, nil
}

func End2EndDownloadSpeedTest() (*speed.Speed, error) {
	return nil, nil
}
