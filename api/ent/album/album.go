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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeBy holds the string denoting the by edge name in mutations.
	EdgeBy = "by"
	// Table holds the table name of the album in the database.
	Table = "albums"
	// ByTable is the table that holds the by relation/edge. The primary key declared below.
	ByTable = "artist_wrote"
	// ByInverseTable is the table name for the Artist entity.
	// It exists in this package in order to avoid circular dependency with the "artist" package.
	ByInverseTable = "artists"
)

// Columns holds all SQL columns for album fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
}

var (
	// ByPrimaryKey and ByColumn2 are the table columns denoting the
	// primary key for the by relation (M2M).
	ByPrimaryKey = []string{"artist_id", "album_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
