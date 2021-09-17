package pluginlib

var PluginType = struct {
	PROXY     int
	CONTAINER int
	DNS       int
	MISC      int
}{
	PROXY:     0,
	CONTAINER: 1,
	DNS:       2,
	MISC:      3,
}

type TerraformProvider struct {
	name               string
	vendor             string
	version            string
	resourceName       string
	init               string
	providerPerProject bool
}

type PluginConfig struct {
	Name        string
	Description string
	Version     string
	Author      string
	PluginType  int
	Terraform   TerraformProvider
}
