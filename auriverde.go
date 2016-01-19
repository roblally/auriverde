package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
//	if args[0] == "auriverde" {
		fmt.Println("Running the auriverde command")
//	}
}

const helpText = "auriverde is a tool for blue-green deployment"

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "auriverde",
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
				Name:     "auriverde",
				HelpText: helpText,
				UsageDetails: plugin.Usage{
					Usage: "cf auriverde",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BasicPlugin))
}