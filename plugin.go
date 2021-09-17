package pluginlib

import (
	"plugin"
)

type Plugin interface {
	GetConfig() PluginConfig
	Setup()
	Deploy(project Project)
	Destroy(project Project)
}

func LoadPlugin(path string) (Plugin, error) {
	// load module
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	// 2. look up a symbol (an exported function or variable)
	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return nil, err
	}

	// 3. Assert that loaded symbol is of a desired type
	var plugin Plugin
	plugin, ok := symPlugin.(Plugin)
	if !ok {
		return nil, err
	}
	return plugin, nil
}
