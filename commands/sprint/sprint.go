package sprint

import (
	"fmt"
	"github.com/urfave/cli/v2"
	prj "github.com/satriahrh/winslow/ent/project"
	spr "github.com/satriahrh/winslow/ent/sprint"
	"time"
	"github.com/satriahrh/winslow/machine"

)

type sprinttAction struct {
	Machine *machine.Machine
}

func sprintCommands(m *machine.Machine) *cli.Command {
	sa := sprinttAction{Machine: m}
	return &cli.Command{
		Name:    "sprint",
		Aliases: []string{"sp"},
		Usage:   "Manage sprint of a project",
		Subcommands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start a new sprint with attached sprint",
				Action: func(context *cli.Context) error {
					project, err := client.Project.Query().Where(prj.CurrentEQ(true)).First(context.Context)
					if err != nil {
						fmt.Println("Error querying current project.")
						return err
					}

					sprint, err := project.QuerySprints().Where(spr.StateEQ(spr.StateCreated)).First(context.Context)
					if err != nil {
						fmt.Println("Error querying next sprint.")
						return err
					}

					current := time.Now()
					s, err := sprint.Update().
						SetState(spr.StateStarted).
						SetStartedAt(current).
						SetFinishedAt(current.AddDate(0, 0, 14)).
						Save(context.Context)
					if err != nil {
						fmt.Println("Error starting new sprint.")
						return err
					}

					fmt.Printf("Sprint %v\n", s.Counter)
					fmt.Printf("Started at %v and scheduled to finished at %v.\n", s.StartedAt.String(), s.FinishedAt.String())
					fmt.Println("Sprint Goals:")
					fmt.Printf("\t%v\n", s.SprintGoal)
					return nil
				},
			},
			{
				Name:  "finish",
				Usage: "Finish the current running sprint",
			},
		},
	}
}
