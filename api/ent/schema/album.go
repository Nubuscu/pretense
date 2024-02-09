package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

func (Album) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultsMixin{},
		SpotifyMixin{},
	}
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("by", Artist.Type).Ref("wrote"),
		edge.From("included_in", Topic.Type).Ref("includes"),
		edge.To("tagged_with", Tag.Type),
	}
}

func (Album) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationUpdate()),
	}
}
