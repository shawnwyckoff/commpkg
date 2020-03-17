package gffmpeg

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/apputil/std_logger"
	"github.com/shawnwyckoff/gpkg/sys/gcmd"
	"github.com/shawnwyckoff/gpkg/sys/gfs"
	"github.com/shawnwyckoff/gpkg/sys/gproc"
	"strconv"
	"strings"
)

// ffmpeg -re -i mymp4.mp4 -stimeout 10000 -c copy -f flv -r 25 -an rtmp://10.78.101.74:1935/live/test100
// rtmp://10.78.101.74:1935/live/test100

// ffmpeg -rtsp_transport tcp -i rtsp://admin:11111111@10.1.3.160 -stimeout 10000 -c copy -f flv -r 25 -an rtmp://221.6.110.211:1934/live/test
// cmd := "ffmpeg -rtsp_transport tcp -i " + rtspUrl + " -stimeout 10000 -c copy -f flv -an rtmp://" + rtmpUrl

/*
func Rtsp2RtmpWait(rtspUrl, rtmpUrl string) {
	cmd := "ffmpeg -rtsp_transport tcp -i " + rtspUrl + " -stimeout 100000 -vcodec copy -acodec libfdk_aac -b:a 64k -bt 64k -ac 2 -ar 44100 -f flv rtmp://" + rtmpUrl
	xlog.Info("ffmpeg command:" + cmd)
	xcmd.ExecWait(cmd, true)
}*/

type Transmitter struct {
	cmder        *gcmd.Cmder
	cmds         []string
	transmitType TransmitType
}

const (
	TransmitRtsp2Rtmp = iota
	TransmitFile2Rtmp
)

type TransmitType int

func NewTransmitter(inputUrl, outputUrl string, transmitType TransmitType, logger std_logger.StdLogger) (*Transmitter, error) {
	var t Transmitter

	if transmitType == TransmitRtsp2Rtmp {
		t.cmds = []string{"ffmpeg", "-rtsp_transport", "tcp", "-i", inputUrl, "-stimeout", "100000", "-c:v", "copy", "-c:a libfdk_aac", "-b:a", "64k", "-bt", "64k", "-ac", "2", "-ar", "44100", "-f", "flv", "rtmp://" + outputUrl}
	} else if transmitType == TransmitFile2Rtmp {
		pi, err := gfs.GetPathInfo(inputUrl)
		if err == nil && !pi.Exist {
			return nil, errors.New("Not exist input file " + inputUrl)
		}
		// -vcodec h264 -b:v 500k
		// t.cmd = "ffmpeg -re -i " + inputUrl + " -c:v copy -c:a libfdk_aac -b:a 64k -bt 64k -ac 2 -ar 44100 -f flv rtmp://" + outputUrl
		t.cmds = []string{"ffmpeg", "-re", "-i", inputUrl, "-c:v", "copy", "-c:a", "aac", "-f", "flv", "rtmp://" + outputUrl}
	} else {
		return nil, errors.New("invalid transmit type " + strconv.FormatInt(int64(transmitType), 10))
	}
	logger.Printf("ffmpeg command:" + strings.Join(t.cmds, " "))

	return &t, nil
}

func (t *Transmitter) RunNowait(retry bool) {
	for {
		t.cmder = gcmd.ExecNowait(true, t.cmds[0], t.cmds[1:]...)
		if !retry {
			break
		}
	}
}

func (t *Transmitter) GetPid() gproc.ProcId {
	if t.cmder == nil {
		return gproc.InvalidProcId
	} else {
		return gproc.ProcId(t.cmder.GetPid())
	}
}

func (t *Transmitter) Wait() {
	if t.cmder != nil {
		t.cmder.Wait()
	}
}

func (t *Transmitter) Close() {
	pid := t.GetPid()
	if pid != gproc.InvalidProcId {
		gproc.Terminate(pid)
		t.cmder = nil
	}
}
