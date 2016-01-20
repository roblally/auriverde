package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Println("Executing command")
}

const pluginName = "auriverde"

const helpText = pluginName + " : zero downtime deploys and more"

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: pluginName,
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     pluginName,
				HelpText: helpText,
				UsageDetails: plugin.Usage{
					Usage: "cf " + pluginName,
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BasicPlugin))
}
