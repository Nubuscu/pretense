package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

func (Topic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultsMixin{},
		ReviewableMixin{},
	}
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("includes", Album.Type),
	}
}
