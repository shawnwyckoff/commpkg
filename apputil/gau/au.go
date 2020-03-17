package gau

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/net/ghttp"
	"github.com/shawnwyckoff/gpkg/spider/util/downloader/xhttpclient"
	"sync"
	"sync/atomic"
	"time"
)

const (
	AU_RELEASE = "https://raw.githubusercontent.com/shawnwyckoff/sbrew/master/release/%s"
	AU_CONF    = "https://raw.githubusercontent.com/shawnwyckoff/sbrew/master/conf/%s"
)

type (
	AU struct {
		proxy                    string
		currentPlatform          string          // windows/linux/darwin
		chNotifyAfterConfUpdated chan ConfUpdate // notify msg if new conf updated
		selfProcUpdate           atomic.Value    // whether update this process

		confToUpdateMtx sync.RWMutex
		confToUpdate    map[string]ConfUpdate // configs to update
	}

	ConfUpdate struct {
		Name     string
		SavePath string
	}
)

func GetLatestConf(name, proxy string, timeout time.Duration) (string, error) {
	path := fmt.Sprintf(AU_CONF, name)
	return ghttp.GetString(path, proxy, timeout)
}

func GetLatestConfTime(name, proxy string, timeout time.Duration) (time.Time, error) {
	//path := fmt.Sprintf(AU_CONF, name)
	return time.Now(), nil
}

func GetLatestRelease(name, savepath, proxy string, timeout time.Duration) error {
	return nil
}

func GetLatestReleaseTime(name, savepath, proxy string, timeout time.Duration) (time.Time, error) {
	return time.Now(), nil
}

func NewAU(chNotifyAfterConfUpdate chan ConfUpdate, proxy string) (*AU, error) {
	return &AU{chNotifyAfterConfUpdated: chNotifyAfterConfUpdate, proxy: proxy}, nil
}

func (au *AU) SetSelfProcUpdate(switcher bool) {
	au.selfProcUpdate.Store(switcher)
}

func (au *AU) SetConfUpdateParam(name, savePath string) {
	au.confToUpdateMtx.Lock()
	au.confToUpdate[name] = ConfUpdate{Name: name, SavePath: savePath}
	au.confToUpdateMtx.Unlock()
}

func (au *AU) LoopBackground() {
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}
