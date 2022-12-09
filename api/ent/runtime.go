// Code generated by ent, DO NOT EDIT.

package ent

import (
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"nubuscu/pretense/ent/review"
	"nubuscu/pretense/ent/schema"
	"nubuscu/pretense/ent/topic"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	albumMixin := schema.Album{}.Mixin()
	albumMixinFields0 := albumMixin[0].Fields()
	_ = albumMixinFields0
	albumFields := schema.Album{}.Fields()
	_ = albumFields
	// albumDescCreatedAt is the schema descriptor for created_at field.
	albumDescCreatedAt := albumMixinFields0[0].Descriptor()
	// album.DefaultCreatedAt holds the default value on creation for the created_at field.
	album.DefaultCreatedAt = albumDescCreatedAt.Default.(func() time.Time)
	// albumDescUpdatedAt is the schema descriptor for updated_at field.
	albumDescUpdatedAt := albumMixinFields0[1].Descriptor()
	// album.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	album.DefaultUpdatedAt = albumDescUpdatedAt.Default.(func() time.Time)
	// album.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	album.UpdateDefaultUpdatedAt = albumDescUpdatedAt.UpdateDefault.(func() time.Time)
	// albumDescName is the schema descriptor for name field.
	albumDescName := albumFields[0].Descriptor()
	// album.NameValidator is a validator for the "name" field. It is called by the builders before save.
	album.NameValidator = albumDescName.Validators[0].(func(string) error)
	artistMixin := schema.Artist{}.Mixin()
	artistMixinFields0 := artistMixin[0].Fields()
	_ = artistMixinFields0
	artistFields := schema.Artist{}.Fields()
	_ = artistFields
	// artistDescCreatedAt is the schema descriptor for created_at field.
	artistDescCreatedAt := artistMixinFields0[0].Descriptor()
	// artist.DefaultCreatedAt holds the default value on creation for the created_at field.
	artist.DefaultCreatedAt = artistDescCreatedAt.Default.(func() time.Time)
	// artistDescUpdatedAt is the schema descriptor for updated_at field.
	artistDescUpdatedAt := artistMixinFields0[1].Descriptor()
	// artist.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	artist.DefaultUpdatedAt = artistDescUpdatedAt.Default.(func() time.Time)
	// artist.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	artist.UpdateDefaultUpdatedAt = artistDescUpdatedAt.UpdateDefault.(func() time.Time)
	reviewMixin := schema.Review{}.Mixin()
	reviewMixinFields0 := reviewMixin[0].Fields()
	_ = reviewMixinFields0
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescCreatedAt is the schema descriptor for created_at field.
	reviewDescCreatedAt := reviewMixinFields0[0].Descriptor()
	// review.DefaultCreatedAt holds the default value on creation for the created_at field.
	review.DefaultCreatedAt = reviewDescCreatedAt.Default.(func() time.Time)
	// reviewDescUpdatedAt is the schema descriptor for updated_at field.
	reviewDescUpdatedAt := reviewMixinFields0[1].Descriptor()
	// review.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	review.DefaultUpdatedAt = reviewDescUpdatedAt.Default.(func() time.Time)
	// review.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	review.UpdateDefaultUpdatedAt = reviewDescUpdatedAt.UpdateDefault.(func() time.Time)
	// reviewDescName is the schema descriptor for name field.
	reviewDescName := reviewFields[0].Descriptor()
	// review.NameValidator is a validator for the "name" field. It is called by the builders before save.
	review.NameValidator = reviewDescName.Validators[0].(func(string) error)
	// reviewDescBody is the schema descriptor for body field.
	reviewDescBody := reviewFields[1].Descriptor()
	// review.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	review.BodyValidator = reviewDescBody.Validators[0].(func(string) error)
	topicMixin := schema.Topic{}.Mixin()
	topicMixinFields0 := topicMixin[0].Fields()
	_ = topicMixinFields0
	topicFields := schema.Topic{}.Fields()
	_ = topicFields
	// topicDescCreatedAt is the schema descriptor for created_at field.
	topicDescCreatedAt := topicMixinFields0[0].Descriptor()
	// topic.DefaultCreatedAt holds the default value on creation for the created_at field.
	topic.DefaultCreatedAt = topicDescCreatedAt.Default.(func() time.Time)
	// topicDescUpdatedAt is the schema descriptor for updated_at field.
	topicDescUpdatedAt := topicMixinFields0[1].Descriptor()
	// topic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	topic.DefaultUpdatedAt = topicDescUpdatedAt.Default.(func() time.Time)
	// topic.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	topic.UpdateDefaultUpdatedAt = topicDescUpdatedAt.UpdateDefault.(func() time.Time)
}
