package pluginlib

type PackageVendor string

var PackageVendorApt PackageVendor = "apt"
var PackageVendorApk PackageVendor = "apk"

type Package struct {
	Name   string
	Vendor PackageVendor
}

func InstallPackage(packages []Package) error {
	// mock function
	// behaviour is implemented in StackHead main repository
	return nil
}
