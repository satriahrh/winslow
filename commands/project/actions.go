package project

import (
	"bufio"
	"github.com/satriahrh/winslow/ent"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"fmt"
)

func (pa *projectAction) index(ctx *cli.Context) error {
	project, err := pa.Machine.Project.GetCurrent(ctx.Context)
	if err != nil {
		return err
	}

	pa.Machine.Project.Outputs([]*ent.Project{project})
	return nil
}

func (pa *projectAction) create(ctx *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name: ")
	projectName, _ := reader.ReadString('\n')
	projectName = projectName[:len(projectName)-1]

	fmt.Print("Slug: ")
	projectSlug, _ := reader.ReadString('\n')
	projectSlug = strings.ToUpper(projectSlug[:len(projectSlug)-1])

	project, err := pa.Machine.Project.Creation(ctx.Context, projectName, projectSlug)
	if err != nil {
		return err
	}

	pa.Machine.Project.Outputs([]*ent.Project{project})
	return nil
}

func (pa *projectAction) use(ctx *cli.Context) error {
	slug := strings.ToUpper(ctx.Args().First())

	err := pa.Machine.Project.SetCurrent(ctx.Context, slug)
	if err != nil {
		return err
	}

	project, err := pa.Machine.Project.GetCurrent(ctx.Context)
	if err != nil {
		return err
	}

	pa.Machine.Project.Outputs([]*ent.Project{project})
	return nil
}

func (pa *projectAction) list(ctx *cli.Context) error {
	projects, err := pa.Machine.Project.GetAll(ctx.Context)
	if err != nil {
		return err
	}

	pa.Machine.Project.Outputs(projects)
	return err
}