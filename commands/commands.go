package commands

import (
	"github.com/satriahrh/winslow/commands/project"
	"github.com/satriahrh/winslow/ent"
	"github.com/satriahrh/winslow/machine"
	"github.com/urfave/cli/v2"
)

func NewCommands(client *ent.Client) []*cli.Command {
	m := machine.NewMachine(client)
	return []*cli.Command{
		project.NewCommands(m),
		//sprint.NewCommands(m),
		//story.NewCommands(m),
	}
}