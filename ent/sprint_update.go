// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/satriahrh/winslow/ent/predicate"
	"github.com/satriahrh/winslow/ent/project"
	"github.com/satriahrh/winslow/ent/sprint"
	"github.com/satriahrh/winslow/ent/story"
)

// SprintUpdate is the builder for updating Sprint entities.
type SprintUpdate struct {
	config
	counter          *uint
	addcounter       *uint
	sprint_goal      *string
	clearsprint_goal bool
	state            *sprint.State
	started_at       *time.Time
	clearstarted_at  bool
	finished_at      *time.Time
	clearfinished_at bool
	project          map[int]struct{}
	stories          map[int]struct{}
	removedProject   map[int]struct{}
	removedStories   map[int]struct{}
	predicates       []predicate.Sprint
}

// Where adds a new predicate for the builder.
func (su *SprintUpdate) Where(ps ...predicate.Sprint) *SprintUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetCounter sets the counter field.
func (su *SprintUpdate) SetCounter(u uint) *SprintUpdate {
	su.counter = &u
	su.addcounter = nil
	return su
}

// AddCounter adds u to counter.
func (su *SprintUpdate) AddCounter(u uint) *SprintUpdate {
	if su.addcounter == nil {
		su.addcounter = &u
	} else {
		*su.addcounter += u
	}
	return su
}

// SetSprintGoal sets the sprint_goal field.
func (su *SprintUpdate) SetSprintGoal(s string) *SprintUpdate {
	su.sprint_goal = &s
	return su
}

// SetNillableSprintGoal sets the sprint_goal field if the given value is not nil.
func (su *SprintUpdate) SetNillableSprintGoal(s *string) *SprintUpdate {
	if s != nil {
		su.SetSprintGoal(*s)
	}
	return su
}

// ClearSprintGoal clears the value of sprint_goal.
func (su *SprintUpdate) ClearSprintGoal() *SprintUpdate {
	su.sprint_goal = nil
	su.clearsprint_goal = true
	return su
}

// SetState sets the state field.
func (su *SprintUpdate) SetState(s sprint.State) *SprintUpdate {
	su.state = &s
	return su
}

// SetStartedAt sets the started_at field.
func (su *SprintUpdate) SetStartedAt(t time.Time) *SprintUpdate {
	su.started_at = &t
	return su
}

// SetNillableStartedAt sets the started_at field if the given value is not nil.
func (su *SprintUpdate) SetNillableStartedAt(t *time.Time) *SprintUpdate {
	if t != nil {
		su.SetStartedAt(*t)
	}
	return su
}

// ClearStartedAt clears the value of started_at.
func (su *SprintUpdate) ClearStartedAt() *SprintUpdate {
	su.started_at = nil
	su.clearstarted_at = true
	return su
}

// SetFinishedAt sets the finished_at field.
func (su *SprintUpdate) SetFinishedAt(t time.Time) *SprintUpdate {
	su.finished_at = &t
	return su
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (su *SprintUpdate) SetNillableFinishedAt(t *time.Time) *SprintUpdate {
	if t != nil {
		su.SetFinishedAt(*t)
	}
	return su
}

// ClearFinishedAt clears the value of finished_at.
func (su *SprintUpdate) ClearFinishedAt() *SprintUpdate {
	su.finished_at = nil
	su.clearfinished_at = true
	return su
}

// AddProjectIDs adds the project edge to Project by ids.
func (su *SprintUpdate) AddProjectIDs(ids ...int) *SprintUpdate {
	if su.project == nil {
		su.project = make(map[int]struct{})
	}
	for i := range ids {
		su.project[ids[i]] = struct{}{}
	}
	return su
}

// AddProject adds the project edges to Project.
func (su *SprintUpdate) AddProject(p ...*Project) *SprintUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddProjectIDs(ids...)
}

// AddStoryIDs adds the stories edge to Story by ids.
func (su *SprintUpdate) AddStoryIDs(ids ...int) *SprintUpdate {
	if su.stories == nil {
		su.stories = make(map[int]struct{})
	}
	for i := range ids {
		su.stories[ids[i]] = struct{}{}
	}
	return su
}

// AddStories adds the stories edges to Story.
func (su *SprintUpdate) AddStories(s ...*Story) *SprintUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddStoryIDs(ids...)
}

// RemoveProjectIDs removes the project edge to Project by ids.
func (su *SprintUpdate) RemoveProjectIDs(ids ...int) *SprintUpdate {
	if su.removedProject == nil {
		su.removedProject = make(map[int]struct{})
	}
	for i := range ids {
		su.removedProject[ids[i]] = struct{}{}
	}
	return su
}

// RemoveProject removes project edges to Project.
func (su *SprintUpdate) RemoveProject(p ...*Project) *SprintUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemoveProjectIDs(ids...)
}

// RemoveStoryIDs removes the stories edge to Story by ids.
func (su *SprintUpdate) RemoveStoryIDs(ids ...int) *SprintUpdate {
	if su.removedStories == nil {
		su.removedStories = make(map[int]struct{})
	}
	for i := range ids {
		su.removedStories[ids[i]] = struct{}{}
	}
	return su
}

// RemoveStories removes stories edges to Story.
func (su *SprintUpdate) RemoveStories(s ...*Story) *SprintUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveStoryIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SprintUpdate) Save(ctx context.Context) (int, error) {
	if su.state != nil {
		if err := sprint.StateValidator(*su.state); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
		}
	}
	return su.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SprintUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SprintUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SprintUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SprintUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sprint.Table,
			Columns: sprint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sprint.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := su.counter; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: sprint.FieldCounter,
		})
	}
	if value := su.addcounter; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: sprint.FieldCounter,
		})
	}
	if value := su.sprint_goal; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: sprint.FieldSprintGoal,
		})
	}
	if su.clearsprint_goal {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sprint.FieldSprintGoal,
		})
	}
	if value := su.state; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: sprint.FieldState,
		})
	}
	if value := su.started_at; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: sprint.FieldStartedAt,
		})
	}
	if su.clearstarted_at {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sprint.FieldStartedAt,
		})
	}
	if value := su.finished_at; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: sprint.FieldFinishedAt,
		})
	}
	if su.clearfinished_at {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sprint.FieldFinishedAt,
		})
	}
	if nodes := su.removedProject; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sprint.ProjectTable,
			Columns: sprint.ProjectPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := su.project; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sprint.ProjectTable,
			Columns: sprint.ProjectPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if nodes := su.removedStories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sprint.StoriesTable,
			Columns: sprint.StoriesPrimaryKey,
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
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := su.stories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sprint.StoriesTable,
			Columns: sprint.StoriesPrimaryKey,
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
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SprintUpdateOne is the builder for updating a single Sprint entity.
type SprintUpdateOne struct {
	config
	id               int
	counter          *uint
	addcounter       *uint
	sprint_goal      *string
	clearsprint_goal bool
	state            *sprint.State
	started_at       *time.Time
	clearstarted_at  bool
	finished_at      *time.Time
	clearfinished_at bool
	project          map[int]struct{}
	stories          map[int]struct{}
	removedProject   map[int]struct{}
	removedStories   map[int]struct{}
}

// SetCounter sets the counter field.
func (suo *SprintUpdateOne) SetCounter(u uint) *SprintUpdateOne {
	suo.counter = &u
	suo.addcounter = nil
	return suo
}

// AddCounter adds u to counter.
func (suo *SprintUpdateOne) AddCounter(u uint) *SprintUpdateOne {
	if suo.addcounter == nil {
		suo.addcounter = &u
	} else {
		*suo.addcounter += u
	}
	return suo
}

// SetSprintGoal sets the sprint_goal field.
func (suo *SprintUpdateOne) SetSprintGoal(s string) *SprintUpdateOne {
	suo.sprint_goal = &s
	return suo
}

// SetNillableSprintGoal sets the sprint_goal field if the given value is not nil.
func (suo *SprintUpdateOne) SetNillableSprintGoal(s *string) *SprintUpdateOne {
	if s != nil {
		suo.SetSprintGoal(*s)
	}
	return suo
}

// ClearSprintGoal clears the value of sprint_goal.
func (suo *SprintUpdateOne) ClearSprintGoal() *SprintUpdateOne {
	suo.sprint_goal = nil
	suo.clearsprint_goal = true
	return suo
}

// SetState sets the state field.
func (suo *SprintUpdateOne) SetState(s sprint.State) *SprintUpdateOne {
	suo.state = &s
	return suo
}

// SetStartedAt sets the started_at field.
func (suo *SprintUpdateOne) SetStartedAt(t time.Time) *SprintUpdateOne {
	suo.started_at = &t
	return suo
}

// SetNillableStartedAt sets the started_at field if the given value is not nil.
func (suo *SprintUpdateOne) SetNillableStartedAt(t *time.Time) *SprintUpdateOne {
	if t != nil {
		suo.SetStartedAt(*t)
	}
	return suo
}

// ClearStartedAt clears the value of started_at.
func (suo *SprintUpdateOne) ClearStartedAt() *SprintUpdateOne {
	suo.started_at = nil
	suo.clearstarted_at = true
	return suo
}

// SetFinishedAt sets the finished_at field.
func (suo *SprintUpdateOne) SetFinishedAt(t time.Time) *SprintUpdateOne {
	suo.finished_at = &t
	return suo
}

// SetNillableFinishedAt sets the finished_at field if the given value is not nil.
func (suo *SprintUpdateOne) SetNillableFinishedAt(t *time.Time) *SprintUpdateOne {
	if t != nil {
		suo.SetFinishedAt(*t)
	}
	return suo
}

// ClearFinishedAt clears the value of finished_at.
func (suo *SprintUpdateOne) ClearFinishedAt() *SprintUpdateOne {
	suo.finished_at = nil
	suo.clearfinished_at = true
	return suo
}

// AddProjectIDs adds the project edge to Project by ids.
func (suo *SprintUpdateOne) AddProjectIDs(ids ...int) *SprintUpdateOne {
	if suo.project == nil {
		suo.project = make(map[int]struct{})
	}
	for i := range ids {
		suo.project[ids[i]] = struct{}{}
	}
	return suo
}

// AddProject adds the project edges to Project.
func (suo *SprintUpdateOne) AddProject(p ...*Project) *SprintUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddProjectIDs(ids...)
}

// AddStoryIDs adds the stories edge to Story by ids.
func (suo *SprintUpdateOne) AddStoryIDs(ids ...int) *SprintUpdateOne {
	if suo.stories == nil {
		suo.stories = make(map[int]struct{})
	}
	for i := range ids {
		suo.stories[ids[i]] = struct{}{}
	}
	return suo
}

// AddStories adds the stories edges to Story.
func (suo *SprintUpdateOne) AddStories(s ...*Story) *SprintUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddStoryIDs(ids...)
}

// RemoveProjectIDs removes the project edge to Project by ids.
func (suo *SprintUpdateOne) RemoveProjectIDs(ids ...int) *SprintUpdateOne {
	if suo.removedProject == nil {
		suo.removedProject = make(map[int]struct{})
	}
	for i := range ids {
		suo.removedProject[ids[i]] = struct{}{}
	}
	return suo
}

// RemoveProject removes project edges to Project.
func (suo *SprintUpdateOne) RemoveProject(p ...*Project) *SprintUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemoveProjectIDs(ids...)
}

// RemoveStoryIDs removes the stories edge to Story by ids.
func (suo *SprintUpdateOne) RemoveStoryIDs(ids ...int) *SprintUpdateOne {
	if suo.removedStories == nil {
		suo.removedStories = make(map[int]struct{})
	}
	for i := range ids {
		suo.removedStories[ids[i]] = struct{}{}
	}
	return suo
}

// RemoveStories removes stories edges to Story.
func (suo *SprintUpdateOne) RemoveStories(s ...*Story) *SprintUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveStoryIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SprintUpdateOne) Save(ctx context.Context) (*Sprint, error) {
	if suo.state != nil {
		if err := sprint.StateValidator(*suo.state); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
		}
	}
	return suo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SprintUpdateOne) SaveX(ctx context.Context) *Sprint {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SprintUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SprintUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SprintUpdateOne) sqlSave(ctx context.Context) (s *Sprint, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sprint.Table,
			Columns: sprint.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  suo.id,
				Type:   field.TypeInt,
				Column: sprint.FieldID,
			},
		},
	}
	if value := suo.counter; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: sprint.FieldCounter,
		})
	}
	if value := suo.addcounter; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  *value,
			Column: sprint.FieldCounter,
		})
	}
	if value := suo.sprint_goal; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: sprint.FieldSprintGoal,
		})
	}
	if suo.clearsprint_goal {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sprint.FieldSprintGoal,
		})
	}
	if value := suo.state; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: sprint.FieldState,
		})
	}
	if value := suo.started_at; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: sprint.FieldStartedAt,
		})
	}
	if suo.clearstarted_at {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sprint.FieldStartedAt,
		})
	}
	if value := suo.finished_at; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: sprint.FieldFinishedAt,
		})
	}
	if suo.clearfinished_at {
		spec.Fields.Clear = append(spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sprint.FieldFinishedAt,
		})
	}
	if nodes := suo.removedProject; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sprint.ProjectTable,
			Columns: sprint.ProjectPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := suo.project; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   sprint.ProjectTable,
			Columns: sprint.ProjectPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if nodes := suo.removedStories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sprint.StoriesTable,
			Columns: sprint.StoriesPrimaryKey,
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
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := suo.stories; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   sprint.StoriesTable,
			Columns: sprint.StoriesPrimaryKey,
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
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	s = &Sprint{config: suo.config}
	spec.Assign = s.assignValues
	spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}