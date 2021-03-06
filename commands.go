package main

import (
	"github.com/codegangsta/cli"
)

// Commands defines our top level commands.
var Commands = []cli.Command{
	commandTags,
	commandSearch,
	commandList,
}

// commandTags defines the sub-commands for all tag operations
var commandTags = cli.Command{
	Name:        "tags",
	Usage:       "manipulate tags for a specific repository",
	Description: `...`,
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "list tags for a specific repository",
			Action: doListTags,
			Flags: []cli.Flag{
				QuietFlag(),
			},
		},
		{
			Name:   "info",
			Usage:  "show detailed tag info",
			Action: doTagInfo,
			Flags: []cli.Flag{
				QuietFlag(),
			},
		},
		{
			Name:   "delete",
			Usage:  "delete a tag in a repository",
			Action: doDeleteTag,
		},
		{
			Name:   "create",
			Usage:  "create a tag for an image",
			Action: doCreateTag,
		},
	},
}

// commandSearch defines the search command.
var commandSearch = cli.Command{
	Name:        "search",
	Usage:       "search for a repository in the registry",
	Description: `...`,
	Action:      doSearch,
}

// commandList defines the list command for listing all repositories.
var commandList = cli.Command{
	Name:        "list",
	Usage:       "list all repositories in the registry",
	Description: `...`,
	Action:      doSearchAll,
}
