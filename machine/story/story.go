package story

import (
	"github.com/satriahrh/winslow/ent"
	"github.com/satriahrh/winslow/ent/story"
)

type Component struct {
	Client *ent.Client
}

func (c *Component) Get(slug string) *ent.Story {
	return nil
}

func (c *Component) GetBySprint(sprint *ent.Sprint) []*ent.Story {
	return nil
}

func (c *Component) GetByState(project *ent.Project, state story.State) []*ent.Story {
	return nil
}