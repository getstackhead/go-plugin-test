# StackHead Plugin API

## Expected structure

```go
package main

import "github.com/getstackhead/pluginlib"

type MyPlugin struct {
}

var PluginConfig = pluginlib.PluginConfig{
	Name:        "Test",
	Description: "This is an example StackHead Proxy plugin",
	Version:     "0.0.0-dev",
	Authors:     []string{"Your Name"},
	PluginType:  pluginlib.PluginType.PROXY,
}

func (p MyPlugin) Setup() {
	// implement software setup action
}

func (p MyPlugin) Deploy(project pluginlib.Project) {
	// implement project deploy action
}

func (p MyPlugin) Destroy(project pluginlib.Project) {
	// implement project destroy action
}

func (p MyPlugin) TriggerHook(hookName string, project pluginlib.Project) {
	// triggered by hooks
}

// Export plugin to StackHead. Must be named "Plugin"!
var Plugin MyPlugin
```

```shell
go build -buildmode=plugin -o plugin.so main.go
```
