# StackHead Plugin API

## Expected structure

```go
package main

import "github.com/getstackhead/pluginlib"

type MyPlugin struct {
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

func (p MyPlugin) HookPreTerraformPlan(project pluginlib.Project) {
	// pre terraform plan hook (to be implemented)
}

// Export plugin to StackHead. Must be named "Plugin"!
var Plugin MyPlugin
```

```shell
go build -buildmode=plugin -o plugin_myplugin.so main.go
```
