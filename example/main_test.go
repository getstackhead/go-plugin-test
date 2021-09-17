package main_test

import (
	"testing"

	"github.com/getstackhead/pluginlib"
)

func TestPluginLoading(t *testing.T) {
	loadedPlugin, err := pluginlib.LoadPlugin("./plugin_myplugin.so")
	if err != nil {
		t.Fatal("Unable to load plugin. Error occurred: " + err.Error())
		return
	}
	if loadedPlugin == nil {
		t.Fatal("Unable to load plugin. Plugin is nil")
		return
	}
}
