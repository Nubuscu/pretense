package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

func (Artist) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultsMixin{},
		SpotifyMixin{},
	}
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("wrote", Album.Type),
		edge.To("tagged_with", Tag.Type),
	}
}

func (Artist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationUpdate()),
	}
}
