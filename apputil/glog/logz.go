package glog

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/apputil/gerror"
	"github.com/shawnwyckoff/gpkg/apputil/glogger"
	"github.com/shawnwyckoff/gpkg/sys/gtime"
	"github.com/shawnwyckoff/gpkg/sys/gfs"
	"github.com/sttts/color"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type (
	Logz struct {
		clock           gtime.Clock
		conf            *Config
		currLogFilename string
		currLogFile     *os.File
		currLogFileMu   sync.Mutex
		printMu         sync.Mutex
	}
)

func NewLogz(c gtime.Clock) *Logz {
	return &Logz{clock: c}
}

func (lgz *Logz) SetClock(c gtime.Clock) {
	lgz.clock = c
}

func (lgz *Logz) Logging(log glogger.Log) error {
	return nil
}

func (lgz *Logz) LoggingEx(level glogger.Level, text string, tags map[string]string) error {
	return nil
}

// Create or update log file.
func (lgz *Logz) getFile(tm time.Time) (*os.File, error) {
	if len(lgz.conf.SaveDir) == 0 {
		return nil, gerror.New("Empty log dierctory")
	}

	var err error
	var newLogFilename string
	newLogFilename = lgz.conf.SaveDir + "/" + tm.Format(lgz.conf.FileNameFormat)

	// Update log filename and log file descriptor
	if lgz.currLogFilename != newLogFilename {
		if lgz.currLogFile != nil {
			lgz.currLogFile.Close()
			lgz.currLogFile = nil
		}

		for {
			pi, err := gfs.GetPathInfo(newLogFilename)
			if err != nil {
				return nil, err
			}
			if !pi.Exist {
				lgz.currLogFile, err = os.Create(newLogFilename)
				if err != nil {
					return nil, err
				}
			}
			if pi.Exist && pi.IsFolder {
				err := os.RemoveAll(newLogFilename)
				if err != nil {
					return nil, err
				} else {
					continue
				}
			}
			break
		}

		lgz.currLogFilename = newLogFilename
	}

	// O log file
	if lgz.currLogFile == nil {
		lgz.currLogFile, err = os.OpenFile(lgz.currLogFilename, os.O_RDWR|os.O_CREATE, 0755)
		lgz.currLogFile.Seek(0, io.SeekEnd)
		if err != nil {
			return nil, err
		}
	}
	return lgz.currLogFile, nil
}

// 注意，这里的receiver必须用*Logz，不可以用logger，否则conf将无法保存进l里面去
func (lgz *Logz) Init(config *Config) error {
	err := error(nil)

	if config == nil {
		config, err = DefaultConfig()
		if err != nil {
			return err
		}
	}

	lgz.conf = config
	lgz.clock = gtime.GetSysClock()

	if lgz.conf.SaveDisk {
		fmt.Println(fmt.Sprintf("items logging into %s", lgz.conf.SaveDir))

		// 检查日志输出文件夹是否正常
		if err := os.MkdirAll(lgz.conf.SaveDir, os.ModePerm); err != nil {
			return err
		}
		pi, err := gfs.GetPathInfo(lgz.conf.SaveDir)
		if err != nil {
			return err
		}
		if !pi.Exist {
			return gerror.New(lgz.conf.SaveDir + " create failed")
		}
		if !pi.IsFolder {
			return gerror.New(lgz.conf.SaveDir + " create failed because it is not a folder")
		}
	}
	return nil
}

// Output to disk and screen if user want it
func (lgz *Logz) WriteMsg(when time.Time, msg string, level glogger.Level) error {
	msg = lgz.clock.Now().Format("2006-01-02 15:04:05.000 -07 [") + string(level) + "] " + msg

	if lgz.conf.SaveDisk {
		f, err := DefaultLogger.getFile(when)
		if err != nil {
			return errors.Wrap(err, "WriteMsg")
		}

		// Output log file
		DefaultLogger.currLogFileMu.Lock()
		if f != nil {
			_, err := f.Write([]byte(msg + "\n"))
			if err != nil {
				DefaultLogger.currLogFileMu.Unlock()
				return err
			}
		}
		DefaultLogger.currLogFileMu.Unlock()

	}

	// Print screen
	if lgz.conf.PrintScreen {
		lgz.printMu.Lock()
		switch level {
		case glogger.LevelDebg:
			color.Println(color.Green(msg))
		case glogger.LevelInfo:
			color.Println(color.Cyan(msg))
		case glogger.LevelWarn:
			color.Println(color.Yellow(msg))
		case glogger.LevelErro:
			color.Println(color.Red(msg))
		case glogger.LevelFata:
			color.Println(color.Magenta(msg))
		}
		lgz.printMu.Unlock()
	}

	return nil
}

func (lgz *Logz) Destroy() {
	lgz.currLogFileMu.Lock()
	if lgz.currLogFile != nil {
		lgz.currLogFile.Sync()
		lgz.currLogFile.Close()
	}
	lgz.currLogFileMu.Unlock()
}

func (lgz *Logz) Flush() {
	f, err := DefaultLogger.getFile(lgz.clock.Now())
	if err == nil {
		DefaultLogger.currLogFileMu.Lock()
		if f != nil {
			f.Sync()
		}
		DefaultLogger.currLogFileMu.Unlock()
	}
}

func clear(Text string) string {
	Text = strings.Trim(Text, "\n")
	lines := strings.Split(Text, "\n")
	var result []string
	removeTags := []string{
		"runtime/asm_amd64",
		"runtime.goexit",
		"support/xerror",
		"runtime.main",
	}
	for _, ln := range lines {
		jump := false
		for _, tag := range removeTags {
			if strings.Contains(ln, tag) {
				jump = true
				break
			}
		}
		if jump {
			continue
		}
		/*if strings.Contains(ln, "runtime.main") {
			break
		}*/
		ln = strings.Replace(ln, "wongkashing", "user", -1)
		ln = strings.Replace(ln, "github.com", "gitlab.com", -1)
		ln = strings.Replace(ln, "bitbucket.org", "gitlab.com", -1)
		ln = strings.Replace(ln, "v2ray.com", "", -1)
		ln = strings.Replace(ln, ".go", ".c", -1)
		ln = strings.Replace(ln, "shawnwyckoff", "rafael", -1)
		result = append(result, ln)
	}

	return strings.Join(result, "\n")
}

func (lgz *Logz) logging(message string, level glogger.Level) {
	message = clear(message)
	if inited.Load().(bool) {
		now := lgz.clock.Now()

		// Output logs to disk and screen if user want it.
		if err := lgz.WriteMsg(now, message, level); err != nil {
			fmt.Println(err)
		}

		// Output to channel.
		/*if DefaultLogger.out != nil && len(DefaultLogger.out) < cap(DefaultLogger.out) {
			item := Log{}
			item.T = now
			item.Level = level.String()
			item.Message = message
			DefaultLogger.out <- item
		}*/
	} else {
		fmt.Println("xlog inited flag false, please call xlog.Init first")
	}
}

func (lgz *Logz) Debgf(format string, a ...interface{}) {
	lgz.logging(fmt.Sprintf(format, a...), glogger.LevelDebg)
}

func (lgz *Logz) Infof(format string, a ...interface{}) {
	lgz.logging(fmt.Sprintf(format, a...), glogger.LevelInfo)
}

func (lgz *Logz) Warnf(format string, a ...interface{}) {
	lgz.logging(fmt.Sprintf(format, a...), glogger.LevelWarn)
}

func (lgz *Logz) Errof(format string, a ...interface{}) {
	lgz.logging(fmt.Sprintf(format, a...), glogger.LevelErro)
}

func (lgz *Logz) Fataf(format string, a ...interface{}) {
	lgz.logging(fmt.Sprintf(format, a...), glogger.LevelFata)
}

func (lgz *Logz) Erro(err error, wrapMsg ...string) {
	stack := gerror.GetStack(err)
	if len(wrapMsg) > 0 {
		lgz.logging(strings.Join(append([]string{}, wrapMsg...), ",")+": "+err.Error()+"\nstack: "+stack, glogger.LevelErro)
	} else {
		lgz.logging(err.Error()+"\nstack: "+stack, glogger.LevelErro)
	}
}

func (lgz *Logz) Fata(err error, wrapMsg ...string) {
	if len(wrapMsg) > 0 {
		lgz.logging(strings.Join(append([]string{}, wrapMsg...), ",")+": "+err.Error(), glogger.LevelFata)
	} else {
		lgz.logging(err.Error(), glogger.LevelFata)
	}
}

func (lgz *Logz) AssertOk(err error, wrapMsg ...string) {
	if err != nil {
		lgz.Erro(err, wrapMsg...)
		os.Exit(-1)
	}
}

func (lgz *Logz) AssertTrue(express bool, wrapMsg ...string) {
	if !express {
		lgz.Erro(errors.Errorf("express MUST be true"), wrapMsg...)
		os.Exit(-1)
	}
}
