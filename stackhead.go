package pluginlib

import "github.com/valyala/gorpc"

const (
	IntCmdInstallPkgApk   string = "stackhead:install:package:apk"
	IntCmdUninstallPkgApk string = "stackhead:uninstall:package:apk"
)

type StackHeadRpcRequest struct {
	Command string
	Data    interface{}
}

func ExecCmd(command string, args ...interface{}) error {
	c := &gorpc.Client{Addr: "localhost:1412"}
	c.Start()
	defer c.Stop()
	// All client methods issuing RPCs are thread-safe and goroutine-safe,
	// i.e. it is safe to call them from multiple concurrently running goroutines.
	_, err := c.Call(StackHeadRpcRequest{Command: command, Data: args})
	return err
}
