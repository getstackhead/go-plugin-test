package main

import (
	"fmt"

	"github.com/getstackhead/pluginlib"
)

var Plugin MyPlugin
var PluginConfig = pluginlib.PluginConfig{
	Name:        "Test",
	Description: "This is an example StackHead Proxy plugin",
	Version:     "0.0.0-dev",
	Authors:     []string{"Mario Lubenka"},
	PluginType:  pluginlib.PluginType.PROXY,
}

type MyPlugin struct {
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
