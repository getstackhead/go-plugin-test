package pluginlib

var PluginType = struct {
	PROXY       int
	CONTAINER   int
	DNS         int
	APPLICATION int
}{
	PROXY:       0,
	CONTAINER:   1,
	DNS:         2,
	APPLICATION: 3,
}

type TerraformProvider struct {
	name               string `yaml:"name"`
	vendor             string `yaml:"vendor"`
	version            string `yaml:"version"`
	resourceName       string `yaml:"resource_name"`
	init               string `yaml:"init"`
	providerPerProject bool   `yaml:"provider_per_project"`
}

type PluginConfig struct {
	Name        string 	 `yaml:"name"`
	Description string 	 `yaml:"description,omitempty"`
	Version     string 	 `yaml:"version,omitempty"`
	Authors     []string `yaml:"authors,omitempty"`
	PluginType  int      `yaml:"plugin_type"`
	Terraform   TerraformProvider
}
