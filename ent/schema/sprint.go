package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Sprint holds the schema definition for the Sprint entity.
type Sprint struct {
	ent.Schema
}

// Fields of the Sprint.
func (Sprint) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("counter"),
		field.Text("sprint_goal").
			Optional(),
		field.Enum("state").
			Values("created", "started", "finished"),
		field.Time("started_at").
			Optional().Nillable(),
		field.Time("finished_at").
			Optional().Nillable(),
	}
}

// Edges of the Sprint.
func (Sprint) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("sprints"),
		edge.To("stories", Story.Type),
	}
}
