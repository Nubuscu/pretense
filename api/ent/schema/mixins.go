package schema

import (
	"time"

	"entgo.io/contrib/entgql"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Adds default created/updated values
// and graphql queryability
type DefaultsMixin struct {
	mixin.Schema
}

func (DefaultsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
func (DefaultsMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

type ReviewableMixin struct {
	mixin.Schema
}

func (ReviewableMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("reviewed_by", Review.Type).Ref("reviews"),
	}
}
