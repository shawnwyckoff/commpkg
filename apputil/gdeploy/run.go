package gdeploy

import (
	"github.com/VividCortex/godaemon"
	"github.com/takama/daemon"
)

// https://github.com/sevlyar/go-daemon
// https://github.com/VividCortex/godaemon

// 这些库非常有启发性 https://github.com/search?l=Go&o=desc&q=restart+process&s=stars&type=Repositories&utf8=%E2%9C%93

// Run options
// A) Self start // Start with OS http://www.cnblogs.com/nerxious/archive/2013/01/18/2866548.html
// B) Self supervise // Auto start if crashed, Supervisor
// C) No hangup // Independent of terminal
// D) Auto switch to sudo interface like sudo ./app

// Make current program a Unix/Linux daemon or Windows service
// name: daemon/service name
// desc: daemon/service description
func Daemonlize(name, desc string) error {
	service, err := daemon.New(name, desc)
	if err != nil {
		return err
	}
	_, err = service.Install()
	if err != nil {
		return err
	}
	service.Start()
	return nil
}

// Make current program NOT a Unix/Linux daemon or Windows service
func Undaemonlize() error {
	return nil
}

// https://github.com/vrecan/death
// https://github.com/tj/go-gracefully
// https://github.com/klauspost/shutdown
func GracefulShutdown() {

}

// Enter background, release terminal dependency
func NoHangup(onOff bool) error {
	_, _, err := godaemon.MakeDaemon(&godaemon.DaemonAttr{})
	return err
}
