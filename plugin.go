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

func LoadPlugin(path string) (Plugin, *PluginConfig, error) {
	// load module
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, nil, err
	}

	// 2. look up a symbol (an exported function or variable)
	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return nil, nil, err
	}
	var ok bool
	var pluginObj Plugin
	pluginObj, ok = symPlugin.(Plugin)
	if !ok {
		return nil, nil, err
	}

	// 3. look up a symbol (an exported function or variable)
	symPluginConfig, err := plug.Lookup("PluginConfig")
	if err != nil {
		return nil, nil, err
	}
	var pluginConfig *PluginConfig
	pluginConfig, ok = symPluginConfig.(*PluginConfig)
	if !ok {
		return nil, nil, err
	}

	// 4. Assert that loaded symbol is of a desired type
	return pluginObj, pluginConfig, nil
}
