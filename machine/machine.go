package machine

import (
	"context"
	"github.com/satriahrh/winslow/ent"
	"github.com/satriahrh/winslow/ent/story"
	"github.com/satriahrh/winslow/machine/project"
)

func NewMachine(client *ent.Client) *Machine {
	return &Machine{
		Project: &project.Component{Client: client},
		//Sprint: &sprint.Component{Client: client},
		//Story: &story.Component{Client: client},
	}
}

type Machine struct {
	Project ProjectMachine
	Sprint SprintMachine
	Story StoryMachine
}

type ProjectMachine interface {
	Creation(ctx context.Context, projectName, projectSlug string) (*ent.Project, error)
	SetCurrent(ctx context.Context, slug string) error
	GetCurrent(ctx context.Context) (*ent.Project, error)
	Get(ctx context.Context, slug string) (*ent.Project, error)
	GetAll(ctx context.Context) ([]*ent.Project, error)
	Outputs(projects []*ent.Project)
}

type SprintMachine interface {
	Current() *ent.Sprint
	Next() *ent.Sprint
}

type StoryMachine interface {
	Get(slug string) *ent.Story
	GetBySprint(sprint *ent.Sprint) []*ent.Story
	GetByState(project *ent.Project, state story.State) []*ent.Story
}

