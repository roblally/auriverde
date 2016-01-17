package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "basic-plugin-command" {
		fmt.Println("Running the basic-plugin-command")
	}
}

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "MyBasicPlugin",
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
				Name:     "basic-plugin-command",
				HelpText: "Basic plugin command's help text",
				UsageDetails: plugin.Usage{
					Usage: "basic-plugin-command\n   cf basic-plugin-command",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BasicPlugin))
}