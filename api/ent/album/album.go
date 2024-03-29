// Code generated by ent, DO NOT EDIT.

package album

import (
	"time"
)

const (
	// Label holds the string label denoting the album type in the database.
	Label = "album"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMetaLabels holds the string denoting the meta_labels field in the database.
	FieldMetaLabels = "meta_labels"
	// FieldSpotifyURL holds the string denoting the spotify_url field in the database.
	FieldSpotifyURL = "spotify_url"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeBy holds the string denoting the by edge name in mutations.
	EdgeBy = "by"
	// EdgeIncludedIn holds the string denoting the included_in edge name in mutations.
	EdgeIncludedIn = "included_in"
	// EdgeTaggedWith holds the string denoting the tagged_with edge name in mutations.
	EdgeTaggedWith = "tagged_with"
	// Table holds the table name of the album in the database.
	Table = "albums"
	// ByTable is the table that holds the by relation/edge. The primary key declared below.
	ByTable = "artist_wrote"
	// ByInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	ByInverseTable = "artists"
	// IncludedInTable is the table that holds the included_in relation/edge. The primary key declared below.
	IncludedInTable = "topic_includes"
	// IncludedInInverseTable is the table name for the Topic entity.
	// It exists in this package in order to avoid circular dependency with the "topic" package.
	IncludedInInverseTable = "topics"
	// TaggedWithTable is the table that holds the tagged_with relation/edge. The primary key declared below.
	TaggedWithTable = "album_tagged_with"
	// TaggedWithInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TaggedWithInverseTable = "tags"
)

// Columns holds all SQL columns for album fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMetaLabels,
	FieldSpotifyURL,
	FieldName,
}

var (
	// ByPrimaryKey and ByColumn2 are the table columns denoting the
	// primary key for the by relation (M2M).
	ByPrimaryKey = []string{"artist_id", "album_id"}
	// IncludedInPrimaryKey and IncludedInColumn2 are the table columns denoting the
	// primary key for the included_in relation (M2M).
	IncludedInPrimaryKey = []string{"topic_id", "album_id"}
	// TaggedWithPrimaryKey and TaggedWithColumn2 are the table columns denoting the
	// primary key for the tagged_with relation (M2M).
	TaggedWithPrimaryKey = []string{"album_id", "tag_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMetaLabels holds the default value on creation for the "meta_labels" field.
	DefaultMetaLabels []string
	// SpotifyURLValidator is a validator for the "spotify_url" field. It is called by the builders before save.
	SpotifyURLValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
