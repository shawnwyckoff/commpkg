package gdeploy

import (
	"fmt"
	"github.com/shawnwyckoff/gopkg/net/ghttp/client"
	"github.com/shawnwyckoff/gopkg/spider/github"
	"github.com/shawnwyckoff/gopkg/sys/gfs"
	"github.com/shawnwyckoff/gopkg/sys/gproc"
	"os"
	"syscall"
)

// 还可以参考
// https://github.com/jpillora/overseer

// 自动检查更新，发现新版本后下载，替换，重启
// 要求正确填写参数，包括当前版本号currentVersion，否则不会更新。
// 不足之处，需要文件名后带版本号，否则会因为windows下无法删除正在运行的程序的文件导致删除旧版本失败
func CheckUpdate(currentVersion, githubUser, githubRepo, appName string) error {
	// Get current process folder.
	dir, err := gproc.SelfDir()
	if err != nil {
		return err
	}

	// Check whether old version file need to remove.
	removeConf := dir + gfs.DirSlash() + "update-to-remove"
	buf, err := gfs.FileToBytes(removeConf)
	if err == nil {
		if err := os.Remove(string(buf)); err == nil {
			os.Remove(removeConf)
		}
	}

	// Check new version.
	r, err := github.GetLatestRelease(githubUser, githubRepo)
	if err != nil {
		return err
	}
	if currentVersion == r.Version {
		return nil
	}
	a, err := r.ParseCurrentPlatform()
	if err != nil {
		return err
	}

	// Make a mark, remove old version after new version start.
	oldVersionFilepath, err := gproc.SelfPath()
	if err != nil {
		return err
	}
	if err := gfs.StringToFile(oldVersionFilepath, removeConf); err != nil {
		return err
	}
	newVersionFilepath := fmt.Sprintf("%s%s%s.%s", dir, gfs.DirSlash(), appName, r.Version)
	if err := client.DownloadBigFile(a.DownloadUrl, newVersionFilepath); err != nil {
		return err
	}

	// Start new version.
	files := make([]*os.File, syscall.Stderr+1)
	files[syscall.Stdin] = os.Stdin
	files[syscall.Stdout] = os.Stdout
	files[syscall.Stderr] = os.Stderr
	wd, err := os.Getwd()
	if nil != err {
		return err
	}
	_, err = os.StartProcess(newVersionFilepath, os.Args, &os.ProcAttr{
		Dir:   wd,
		Env:   os.Environ(),
		Files: files,
		Sys:   &syscall.SysProcAttr{},
	})
	if err != nil {
		return err
	}
	os.Exit(0)
	return nil
}
