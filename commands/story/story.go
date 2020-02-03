package story

import (
	"bufio"
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/satriahrh/winslow/ent"
	prj "github.com/satriahrh/winslow/ent/project"
	str "github.com/satriahrh/winslow/ent/story"
	spr "github.com/satriahrh/winslow/ent/sprint"
	mcn "github.com/satriahrh/winslow/machine"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

func output(stories []*ent.Story) {
	data := [][]string{}
	for _, story := range stories {
		data = append(data, []string{
			story.Slug, story.Name, story.State.String(),
		})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Slug", "Name", "State"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func outputSingle(story *ent.Story) {
	sprint, err := story.QuerySprint().Only(context.Background())
	sprintCount := ""
	if err == nil {
		sprintCount = strconv.FormatUint(uint64(sprint.Counter), 10)
	}
	data := [][]string{
		{"Slug", story.Slug},
		{"State", story.State.String()},
		{"Sprint", sprintCount},
		{"Name", story.Name},
		{"Description", story.Description},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetCenterSeparator("|")
	table.SetRowSeparator("-")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func storyCommands(machine *mcn.Machine) *cli.Command {
	return &cli.Command{
		Name:    "story",
		Aliases: []string{"st"},
		Usage:   "Manage story on a project",
		Action: func(context *cli.Context) error {
			project, err := client.Project.Query().Where(prj.CurrentEQ(true)).First(context.Context)
			if err != nil {
				fmt.Println("Error querying current project.")
				return err
			}

			stories, err := project.QueryStories().Where(str.StateEQ(str.StateCreated)).All(context.Context)
			if err != nil {
				fmt.Println("Error querying stories")
				return err
			}

			output(stories)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create a new story, and put them in the backlog",
				Action: func(context *cli.Context) error {
					project, err := client.Project.Query().Where(prj.CurrentEQ(true)).First(context.Context)
					if err != nil {
						fmt.Println("Error querying current project.")
						return err
					}
					reader := bufio.NewReader(os.Stdin)
					story := client.Story.Create()

					story.SetState(str.StateCreated)

					fmt.Print("Name: ")
					text, _ := reader.ReadString('\n')
					text = text[:len(text)-1]
					story.SetName(text)

					fmt.Println("Description (long text, end it with `|`)")
					text, _ = reader.ReadString('|')
					text = text[:len(text)-1]
					story.SetDescription(text)

					numberOfStories, err := project.QueryStories().Count(context.Context)
					if err != nil {
						fmt.Println("Error counting the number of stories of current project.")
						return err
					}
					story.SetSlug(fmt.Sprintf("%v-%v", project.Slug, numberOfStories+1))

					story.AddProject(project)

					s, err := story.Save(context.Context)
					if err != nil {
						fmt.Println("Error creating the story.")
						return err
					}

					outputSingle(s)
					return nil
				},
			},
			{
				Name:  "read",
				Usage: "Read a specific of story",
				Action: func(context *cli.Context) error {
					story, err := client.Story.Query().Where(str.SlugEQ(context.Args().First())).First(context.Context)
					if err != nil {
						fmt.Println("Error querying the story.")
						return err
					}
					outputSingle(story)
					return nil
				},
			},
			{
				Name:  "config",
				Usage: "Config story point and details of this story",
			},
			{
				Name:  "backlog",
				Usage: "List created storis.",
				Action: func(context *cli.Context) error {
					project, err := client.Project.Query().Where(prj.CurrentEQ(true)).First(context.Context)
					if err != nil {
						fmt.Println("Error querying current project.")
						return err
					}

					stories, err := project.QueryStories().Where(str.StateEQ(str.StateCreated)).All(context.Context)
					if err != nil {
						fmt.Println("Error querying stories")
						return err
					}

					output(stories)
					return nil
				},
			},
			{
				Name:  "progress",
				Usage: "Update story progress",
				Action: func(context *cli.Context) error {

					return nil
				},
			},
			{
				Name:  "dispatch",
				Usage: "Dispatch a story to current sprint. BEWARE once you put them on a sprint, you can't take them down to the backlog",
				Subcommands: []*cli.Command{
					{
						Name:  "next",
						Usage: "Dispatch a story to next sprint.",
						Action: func(context *cli.Context) error {
							story, err := client.Story.Query().Where(str.SlugEQ(context.Args().First())).First(context.Context)
							if err != nil {
								fmt.Println("Error querying the story.")
								return err
							}

							project, err := client.Project.Query().Where(prj.CurrentEQ(true)).Only(context.Context)
							if err != nil {
								fmt.Println("Error querying current project.")
								return err
							}

							sprint, err := project.QuerySprints().Where(spr.StateEQ(spr.StateCreated)).Only(context.Context)
							if err != nil {
								if err.Error() == "ent: sprint not found" {
									sprintCounter, err := project.QuerySprints().Count(context.Context)
									if err != nil {
										fmt.Println("Error querying sprints.")
										return err
									}
									sprint, err = client.Sprint.Create().
										SetCounter(uint(sprintCounter + 1)).
										SetState(spr.StateStarted).
										Save(context.Context)
									if err != nil {
										fmt.Println("Error creating new next sprint.")
										return err
									}
								} else {
									fmt.Println("Error querying next sprint.")
									return err
								}
							}

							_, err = sprint.Update().AddStories(story).Save(context.Context)
							if err != nil {
								fmt.Println("Error dispatching story to next sprint.")
								return err
							}

							_, err = story.Update().SetState(str.StateTodo).Save(context.Context)
							if err != nil {
								fmt.Println("Error updating story state.")
								return err
							}

							outputSingle(story)
							return nil
						},
					},
					{
						Name:  "current",
						Usage: "Dispatch a story to current sprint.",
					},
				},
			},
			{
				Name:  "park",
				Usage: "Temporary park the story to parking lot only for this sprint",
			},
			{
				Name:  "forward",
				Usage: "Move the story progress forward",
			},
			{
				Name:  "backward",
				Usage: "Move the story progress backward",
			},
		},
	}
}