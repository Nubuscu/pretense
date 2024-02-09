// Code generated by ent, DO NOT EDIT.

package album

import (
	"nubuscu/pretense/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// SpotifyURL applies equality check predicate on the "spotify_url" field. It's identical to SpotifyURLEQ.
func SpotifyURL(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpotifyURL), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// SpotifyURLEQ applies the EQ predicate on the "spotify_url" field.
func SpotifyURLEQ(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLNEQ applies the NEQ predicate on the "spotify_url" field.
func SpotifyURLNEQ(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLIn applies the In predicate on the "spotify_url" field.
func SpotifyURLIn(vs ...string) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSpotifyURL), v...))
	})
}

// SpotifyURLNotIn applies the NotIn predicate on the "spotify_url" field.
func SpotifyURLNotIn(vs ...string) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSpotifyURL), v...))
	})
}

// SpotifyURLGT applies the GT predicate on the "spotify_url" field.
func SpotifyURLGT(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLGTE applies the GTE predicate on the "spotify_url" field.
func SpotifyURLGTE(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLLT applies the LT predicate on the "spotify_url" field.
func SpotifyURLLT(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLLTE applies the LTE predicate on the "spotify_url" field.
func SpotifyURLLTE(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLContains applies the Contains predicate on the "spotify_url" field.
func SpotifyURLContains(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLHasPrefix applies the HasPrefix predicate on the "spotify_url" field.
func SpotifyURLHasPrefix(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLHasSuffix applies the HasSuffix predicate on the "spotify_url" field.
func SpotifyURLHasSuffix(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLEqualFold applies the EqualFold predicate on the "spotify_url" field.
func SpotifyURLEqualFold(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSpotifyURL), v))
	})
}

// SpotifyURLContainsFold applies the ContainsFold predicate on the "spotify_url" field.
func SpotifyURLContainsFold(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSpotifyURL), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Album {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// HasBy applies the HasEdge predicate on the "by" edge.
func HasBy() predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ByTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ByTable, ByPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasByWith applies the HasEdge predicate on the "by" edge with a given conditions (other predicates).
func HasByWith(preds ...predicate.Artist) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ByInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ByTable, ByPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIncludedIn applies the HasEdge predicate on the "included_in" edge.
func HasIncludedIn() predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IncludedInTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, IncludedInTable, IncludedInPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIncludedInWith applies the HasEdge predicate on the "included_in" edge with a given conditions (other predicates).
func HasIncludedInWith(preds ...predicate.Topic) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IncludedInInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, IncludedInTable, IncludedInPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTaggedWith applies the HasEdge predicate on the "tagged_with" edge.
func HasTaggedWith() predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaggedWithTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TaggedWithTable, TaggedWithPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaggedWithWith applies the HasEdge predicate on the "tagged_with" edge with a given conditions (other predicates).
func HasTaggedWithWith(preds ...predicate.Tag) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TaggedWithInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TaggedWithTable, TaggedWithPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Album) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Album) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Album) predicate.Album {
	return predicate.Album(func(s *sql.Selector) {
		p(s.Not())
	})
}
