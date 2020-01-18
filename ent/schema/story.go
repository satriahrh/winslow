package schema

import (
	"github.com/facebookincubator/ent/schema/edge"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Story holds the schema definition for the Story entity.
type Story struct {
	ent.Schema
}

// Fields of the Story.
func (Story) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			Unique(),
		field.String("name").
			Match(
				regexp.MustCompile("[A-Za-z0-9-]+$"),
			),
		field.String("description"),
		field.Enum("state").
			Values("created", "parked", "todo", "on_progress", "on_review", "done"),
	}
}

// Edges of the Story.
func (Story) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("stories"),
		edge.From("sprint", Sprint.Type).
			Ref("stories"),
		edge.To("activities", Activity.Type),
	}
}
