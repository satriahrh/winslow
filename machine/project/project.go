package project

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/satriahrh/winslow/ent"
	prj "github.com/satriahrh/winslow/ent/project"
	"os"
)

type Component struct {
	Client *ent.Client
}

func (c *Component) Creation(ctx context.Context, projectName, projectSlug string) (*ent.Project, error) {
	project := c.Client.Project.Create()
	project.SetName(projectName)
	project.SetSlug(projectSlug)

	_, err := c.Client.Project.Update().SetCurrent(false).Save(ctx)
	if err != nil {
		fmt.Println("Error updating the project.")
		return nil, err
	}
	project.SetCurrent(true)

	projectSaved, err := project.Save(ctx)
	if err != nil {
		fmt.Println("Error creating the project.")
		return nil, err
	}

	return projectSaved, nil
}

func (c* Component) SetCurrent(ctx context.Context, slug string) (error) {
	_, err := c.Client.Project.Update().SetCurrent(false).Where(prj.CurrentEQ(true)).Save(ctx)
	if err != nil {
		fmt.Println("Error unset the project.")
		return err
	}

	_ , err = c.Client.Project.Update().SetCurrent(true).Where(prj.SlugEQ(slug)).Save(ctx)
	if err != nil {
		fmt.Println("Error setting new current project.")
		return err
	}

	return nil
}

func (c *Component) GetCurrent(ctx context.Context) (*ent.Project, error) {
	project, err := c.Client.Project.Query().Where(prj.CurrentEQ(true)).First(ctx)
	if err != nil {
		fmt.Println("Error querying current project.")
		return nil, err
	}
	return project, nil
}

func (c *Component) Get(ctx context.Context, slug string) (*ent.Project, error) {
	project, err := c.Client.Project.Query().Where(prj.SlugEQ(slug)).Only(ctx)
	if err !=  nil {
		fmt.Println("Error querying project.")
		return nil, err
	}
	return project, nil
}

func (c *Component) GetAll(ctx context.Context) ([]*ent.Project, error) {
	projects, err := c.Client.Project.Query().All(ctx)
	if err != nil {
		fmt.Println("Error querying project")
		return []*ent.Project{}, err
	}
	return projects, nil
}

func (c *Component) Outputs(projects []*ent.Project) {
	if len(projects) == 0 {
		fmt.Println("There is no projects.")
		return
	}

	data := [][]string{}
	for _, project := range projects {
		current := "0"
		if project.Current {
			current = "1"
		}
		data = append(data, []string{
			project.Slug, project.Name, current,
		})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Slug", "Name", "Current"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator("|")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}