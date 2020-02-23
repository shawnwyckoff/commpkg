package ftp

//https://github.com/fclairamb/ftpserver/issues/73
import (
	//"github.com/andrewarrow/paradise_ftp/server"
	//"github.com/andrewarrow/paradise_ftp/paradise"
	"flag"
)

var (
	gracefulChild = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)

type FtpServer struct {
	localAddr string
	rootDir   string
}

func NewFtpServer(localAddr string, rootDir string) (*FtpServer, error) {
	fs := FtpServer{}
	fs.localAddr = localAddr
	fs.rootDir = rootDir
	return &fs, nil
}

func (fs *FtpServer) Start() {
	/*flag.Parse()
	go server.Monitor()
	fm := paradise.NewDefaultFileSystem()
	am := paradise.NewDefaultAuthSystem()
	server.Start(fm, am, *gracefulChild)*/
}
