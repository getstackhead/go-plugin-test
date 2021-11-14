package pluginlib

type PluginTerraformConfigProvider struct {
	Vendor             string
	Name               string
	Version            string
	ResourceName       string
	Init               string `yaml:"init,omitempty"`
	ProviderPerProject bool   `yaml:"provider_per_project,omitempty"` // todo
	NameSuffix         string `yaml:"name_suffix,omitempty"`          // only used internally for upgrading Terraform providers
}

type PluginTerraformConfig struct {
	Provider PluginTerraformConfigProvider
}

type PluginConfig struct {
	Type      string
	Terraform PluginTerraformConfig `yaml:"terraform,omitempty"`
}
