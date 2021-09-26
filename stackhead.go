package pluginlib

type IStackHeadMain interface {
	Execute(command string, args ...interface{}) <-chan error
}

const (
	IntCmdInstallPkgApk   string = "stackhead:install:package:apk"
	IntCmdUninstallPkgApk string = "stackhead:uninstall:package:apk"
)

var StackHeadMain IStackHeadMain = nil
