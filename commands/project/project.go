package project

import (
	"github.com/satriahrh/winslow/machine"

	"github.com/urfave/cli/v2"
)

type projectAction struct {
	Machine *machine.Machine
}

func NewCommands(m *machine.Machine) *cli.Command {
	pa := projectAction{Machine: m}
	return &cli.Command{
		Name:    "project",
		Aliases: []string{"cp"},
		Usage:   "Manage scrum project",
		Action:  pa.index,
		Subcommands: []*cli.Command{
			{
				Name:   "create",
				Usage:  "Create a new scrum project. You will be prompted to fulfill some form.",
				Action: pa.create,
			},
			{
				Name:   "use",
				Usage:  "Use this project for hereafter session, given SLUG of the project.",
				Action: pa.use,
			},
			{
				Name:  "config",
				Usage: "Configure this project",
			},
			{
				Name:   "list",
				Usage:  "List all project",
				Action: pa.list,
			},
		},
	}
}
