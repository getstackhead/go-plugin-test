package pluginlib

type PackageName struct {
	ApkPackageName string
}

func InstallPackage(packageName PackageName) error {
	return ExecCmd(IntCmdInstallPkgApk, packageName.ApkPackageName)
}

func UninstallPackage(packageName PackageName) error {
	return ExecCmd(IntCmdUninstallPkgApk, packageName.ApkPackageName)
}
