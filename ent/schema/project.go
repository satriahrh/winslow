package schema

import (
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Match(
				regexp.MustCompile("[a-z0-9-]+$"),
			),
		field.String("slug").
			Unique().
			Match(
				regexp.MustCompile("[A-Z]+$"),
			),
		field.Bool("current").
			Default(false),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sprints", Sprint.Type),
		edge.To("stories", Story.Type),
	}
}