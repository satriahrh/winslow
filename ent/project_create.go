// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/satriahrh/winslow/ent/project"
	"github.com/satriahrh/winslow/ent/sprint"
	"github.com/satriahrh/winslow/ent/story"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	name    *string
	slug    *string
	current *bool
	sprints map[int]struct{}
	stories map[int]struct{}
}

// SetName sets the name field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.name = &s
	return pc
}

// SetSlug sets the slug field.
func (pc *ProjectCreate) SetSlug(s string) *ProjectCreate {
	pc.slug = &s
	return pc
}

// SetCurrent sets the current field.
func (pc *ProjectCreate) SetCurrent(b bool) *ProjectCreate {
	pc.current = &b
	return pc
}

// SetNillableCurrent sets the current field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCurrent(b *bool) *ProjectCreate {
	if b != nil {
		pc.SetCurrent(*b)
	}
	return pc
}

// AddSprintIDs adds the sprints edge to Sprint by ids.
func (pc *ProjectCreate) AddSprintIDs(ids ...int) *ProjectCreate {
	if pc.sprints == nil {
		pc.sprints = make(map[int]struct{})
	}
	for i := range ids {
		pc.sprints[ids[i]] = struct{}{}
	}
	return pc
}

// AddSprints adds the sprints edges to Sprint.
func (pc *ProjectCreate) AddSprints(s ...*Sprint) *ProjectCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSprintIDs(ids...)
}

// AddStoryIDs adds the stories edge to Story by ids.
func (pc *ProjectCreate) AddStoryIDs(ids ...int) *ProjectCreate {
	if pc.stories == nil {
		pc.stories = make(map[int]struct{})
	}
	for i := range ids {
		pc.stories[ids[i]] = struct{}{}
	}
	return pc
}

// AddStories adds the stories edges to Story.
func (pc *ProjectCreate) AddStories(s ...*Story) *ProjectCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddStoryIDs(ids...)
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	if pc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := project.NameValidator(*pc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	if pc.slug == nil {
		return nil, errors.New("ent: missing required field \"slug\"")
	}
	if err := project.SlugValidator(*pc.slug); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"slug\": %v", err)
	}
	if pc.current == nil {
		v := project.DefaultCurrent
		pc.current = &v
	}
	return pc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	var (
		pr   = &Project{config: pc.config}
		spec = &sqlgraph.CreateSpec{
			Table: project.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: project.FieldID,
			},
		}
	)
	if value := pc.name; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: project.FieldName,
		})
		pr.Name = *value
	}
	if value := pc.slug; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: project.FieldSlug,
		})
		pr.Slug = *value
	}
	if value := pc.current; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: project.FieldCurrent,
		})
		pr.Current = *value
	}
	if nodes := pc.sprints; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.SprintsTable,
			Columns: project.SprintsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sprint.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := pc.stories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.StoriesTable,
			Columns: project.StoriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: story.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, pc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}