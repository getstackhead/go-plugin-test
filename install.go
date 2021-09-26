package pluginlib

type PackageName struct {
	ApkPackageName string
}

func InstallPackage(packageName PackageName) <-chan error {
	return StackHeadMain.Execute(IntCmdInstallPkgApk, packageName.ApkPackageName)
}

func UninstallPackage(packageName PackageName) <-chan error {
	return StackHeadMain.Execute(IntCmdUninstallPkgApk, packageName.ApkPackageName)
}
