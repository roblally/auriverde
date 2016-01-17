package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "auriverde" {
		fmt.Println("Running the auriverde command")
	}
}

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Auriverde",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "auriverde",
				HelpText: "Some help...",
				UsageDetails: plugin.Usage{
					Usage: "auriverde\n   cf auriverde",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BasicPlugin))
}