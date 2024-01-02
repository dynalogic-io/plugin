package v1

import (
	"fmt"
	"github.com/dynalogic-io/plugin/v1/core"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

type Plugin interface {
	New(core core.Core)
	Info() core.Manifest
	Process(value interface{}) interface{}
}

type Plugins map[string]Plugin

// Load loads plugins from the specified path and returns a map of loaded plugins.
//
// The plugins are loaded using the `plugin.Open` function. If loading a plugin
// fails, an error is printed to the console and the application exits with code 1.
// For each successfully loaded plugin, the `New` method of the interface type
// `Plugin` is called with a reference to the `Instance` created by `core.New`.
// The loaded plugin is then added to the `plugins` map using the package name
// from the plugin's `Info()` method as the key.
//
// It returns a map of loaded plugins with the package name as the key and the plugin as
// the value.
func Load(path string) (plugins Plugins) {
	plugins = make(Plugins)
	instance := core.New()

	if pluginsList, err := filepath.Glob(path); err == nil {
		for _, cursor := range pluginsList {
			// load module
			plug, err := plugin.Open(cursor)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			lookup, err := plug.Lookup("Addon")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if currentPlugin, ok := lookup.(Plugin); ok {
				currentPlugin.New(&instance)
				plugins[currentPlugin.Info().Package] = currentPlugin
				log.Println(fmt.Sprintf("package [%s] %s - loaded", currentPlugin.Info().Package, currentPlugin.Info().Name))
			}
		}
	}

	return
}
