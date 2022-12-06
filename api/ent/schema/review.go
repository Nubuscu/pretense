package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

func (Review) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultsMixin{},
	}
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.Text("body").NotEmpty(),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reviews", Topic.Type),
		// edge.To("reviews", Album.Type),
		// edge.To("reviews", Artist.Type),
	}
}
func (Review) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationUpdate()),
	}
}
