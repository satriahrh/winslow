// Code generated by entc, DO NOT EDIT.

package activity

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/satriahrh/winslow/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Activity {
	return predicate.Activity(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Activity applies equality check predicate on the "activity" field. It's identical to ActivityEQ.
func Activity(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActivity), v))
	},
	)
}

// LoggedAt applies equality check predicate on the "logged_at" field. It's identical to LoggedAtEQ.
func LoggedAt(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLoggedAt), v))
	},
	)
}

// ActivityEQ applies the EQ predicate on the "activity" field.
func ActivityEQ(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldActivity), v))
	},
	)
}

// ActivityNEQ applies the NEQ predicate on the "activity" field.
func ActivityNEQ(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldActivity), v))
	},
	)
}

// ActivityIn applies the In predicate on the "activity" field.
func ActivityIn(vs ...string) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldActivity), v...))
	},
	)
}

// ActivityNotIn applies the NotIn predicate on the "activity" field.
func ActivityNotIn(vs ...string) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldActivity), v...))
	},
	)
}

// ActivityGT applies the GT predicate on the "activity" field.
func ActivityGT(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldActivity), v))
	},
	)
}

// ActivityGTE applies the GTE predicate on the "activity" field.
func ActivityGTE(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldActivity), v))
	},
	)
}

// ActivityLT applies the LT predicate on the "activity" field.
func ActivityLT(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldActivity), v))
	},
	)
}

// ActivityLTE applies the LTE predicate on the "activity" field.
func ActivityLTE(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldActivity), v))
	},
	)
}

// ActivityContains applies the Contains predicate on the "activity" field.
func ActivityContains(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldActivity), v))
	},
	)
}

// ActivityHasPrefix applies the HasPrefix predicate on the "activity" field.
func ActivityHasPrefix(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldActivity), v))
	},
	)
}

// ActivityHasSuffix applies the HasSuffix predicate on the "activity" field.
func ActivityHasSuffix(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldActivity), v))
	},
	)
}

// ActivityEqualFold applies the EqualFold predicate on the "activity" field.
func ActivityEqualFold(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldActivity), v))
	},
	)
}

// ActivityContainsFold applies the ContainsFold predicate on the "activity" field.
func ActivityContainsFold(v string) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldActivity), v))
	},
	)
}

// LoggedAtEQ applies the EQ predicate on the "logged_at" field.
func LoggedAtEQ(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLoggedAt), v))
	},
	)
}

// LoggedAtNEQ applies the NEQ predicate on the "logged_at" field.
func LoggedAtNEQ(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLoggedAt), v))
	},
	)
}

// LoggedAtIn applies the In predicate on the "logged_at" field.
func LoggedAtIn(vs ...time.Time) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLoggedAt), v...))
	},
	)
}

// LoggedAtNotIn applies the NotIn predicate on the "logged_at" field.
func LoggedAtNotIn(vs ...time.Time) predicate.Activity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Activity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLoggedAt), v...))
	},
	)
}

// LoggedAtGT applies the GT predicate on the "logged_at" field.
func LoggedAtGT(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLoggedAt), v))
	},
	)
}

// LoggedAtGTE applies the GTE predicate on the "logged_at" field.
func LoggedAtGTE(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLoggedAt), v))
	},
	)
}

// LoggedAtLT applies the LT predicate on the "logged_at" field.
func LoggedAtLT(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLoggedAt), v))
	},
	)
}

// LoggedAtLTE applies the LTE predicate on the "logged_at" field.
func LoggedAtLTE(v time.Time) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLoggedAt), v))
	},
	)
}

// HasStory applies the HasEdge predicate on the "story" edge.
func HasStory() predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StoryTable, StoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasStoryWith applies the HasEdge predicate on the "story" edge with a given conditions (other predicates).
func HasStoryWith(preds ...predicate.Story) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StoryTable, StoryPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Activity) predicate.Activity {
	return predicate.Activity(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Activity) predicate.Activity {
	return predicate.Activity(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Activity) predicate.Activity {
	return predicate.Activity(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
