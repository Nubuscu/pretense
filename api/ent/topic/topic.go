// Code generated by ent, DO NOT EDIT.

package topic

import (
	"time"
)

const (
	// Label holds the string label denoting the topic type in the database.
	Label = "topic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMetaLabels holds the string denoting the meta_labels field in the database.
	FieldMetaLabels = "meta_labels"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeReviewedBy holds the string denoting the reviewed_by edge name in mutations.
	EdgeReviewedBy = "reviewed_by"
	// EdgeIncludes holds the string denoting the includes edge name in mutations.
	EdgeIncludes = "includes"
	// EdgeTaggedWith holds the string denoting the tagged_with edge name in mutations.
	EdgeTaggedWith = "tagged_with"
	// Table holds the table name of the topic in the database.
	Table = "topics"
	// ReviewedByTable is the table that holds the reviewed_by relation/edge. The primary key declared below.
	ReviewedByTable = "review_reviews"
	// ReviewedByInverseTable is the table name for the Review entity.
	// It exists in this package in order to avoid circular dependency with the "review" package.
	ReviewedByInverseTable = "reviews"
	// IncludesTable is the table that holds the includes relation/edge. The primary key declared below.
	IncludesTable = "topic_includes"
	// IncludesInverseTable is the table name for the Album entity.
	// It exists in this package in order to avoid circular dependency with the "album" package.
	IncludesInverseTable = "albums"
	// TaggedWithTable is the table that holds the tagged_with relation/edge. The primary key declared below.
	TaggedWithTable = "topic_tagged_with"
	// TaggedWithInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TaggedWithInverseTable = "tags"
)

// Columns holds all SQL columns for topic fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMetaLabels,
	FieldName,
}

var (
	// ReviewedByPrimaryKey and ReviewedByColumn2 are the table columns denoting the
	// primary key for the reviewed_by relation (M2M).
	ReviewedByPrimaryKey = []string{"review_id", "topic_id"}
	// IncludesPrimaryKey and IncludesColumn2 are the table columns denoting the
	// primary key for the includes relation (M2M).
	IncludesPrimaryKey = []string{"topic_id", "album_id"}
	// TaggedWithPrimaryKey and TaggedWithColumn2 are the table columns denoting the
	// primary key for the tagged_with relation (M2M).
	TaggedWithPrimaryKey = []string{"topic_id", "tag_id"}
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
)
