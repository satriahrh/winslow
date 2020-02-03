package sprint

import (
	"github.com/satriahrh/winslow/ent"
)

type Component struct {
	Client *ent.Client
}

func (c *Component) Current() *ent.Sprint {
	return nil
}

func (c *Component) Next() *ent.Sprint {
	return nil
}