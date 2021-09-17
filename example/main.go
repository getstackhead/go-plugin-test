package main

import "fmt"
import "github.com/getstackhead/pluginlib"

type MyPlugin struct {
}

func (p MyPlugin) GetConfig() pluginlib.PluginConfig {
	return pluginlib.PluginConfig{
		Name:        "Test",
		Description: "This is an example StackHead Proxy plugin",
		Version:     "0.0.0-dev",
		Author:      "Mario Lubenka",
		PluginType:  pluginlib.PluginType.PROXY,
	}
}

func (p MyPlugin) Deploy(project pluginlib.Project) {
	fmt.Println("Hello " + project.Name)
}

func (p MyPlugin) Destroy(project pluginlib.Project) {
	fmt.Println("Bye " + project.Name)
}

func (p MyPlugin) Setup() {
	fmt.Println("Setup...")
}

var Plugin MyPlugin
