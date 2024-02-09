package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty(),
		field.String("value").NotEmpty(),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tags_album", Album.Type).Ref("tagged_with"),
		edge.From("tags_artist", Artist.Type).Ref("tagged_with"),
		edge.From("tags_review", Review.Type).Ref("tagged_with"),
		edge.From("tags_topic", Topic.Type).Ref("tagged_with"),
	}
}
