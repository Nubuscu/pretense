// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"nubuscu/pretense/ent/album"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Album is the model entity for the Album schema.
type Album struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AlbumQuery when eager-loading is set.
	Edges AlbumEdges `json:"edges"`
}

// AlbumEdges holds the relations/edges for other nodes in the graph.
type AlbumEdges struct {
	// By holds the value of the by edge.
	By []*Artist `json:"by,omitempty"`
	// IncludedIn holds the value of the included_in edge.
	IncludedIn []*Topic `json:"included_in,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedBy         map[string][]*Artist
	namedIncludedIn map[string][]*Topic
}

// ByOrErr returns the By value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumEdges) ByOrErr() ([]*Artist, error) {
	if e.loadedTypes[0] {
		return e.By, nil
	}
	return nil, &NotLoadedError{edge: "by"}
}

// IncludedInOrErr returns the IncludedIn value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumEdges) IncludedInOrErr() ([]*Topic, error) {
	if e.loadedTypes[1] {
		return e.IncludedIn, nil
	}
	return nil, &NotLoadedError{edge: "included_in"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Album) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case album.FieldID:
			values[i] = new(sql.NullInt64)
		case album.FieldName:
			values[i] = new(sql.NullString)
		case album.FieldCreatedAt, album.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Album", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Album fields.
func (a *Album) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case album.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case album.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case album.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case album.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		}
	}
	return nil
}

// QueryBy queries the "by" edge of the Album entity.
func (a *Album) QueryBy() *ArtistQuery {
	return (&AlbumClient{config: a.config}).QueryBy(a)
}

// QueryIncludedIn queries the "included_in" edge of the Album entity.
func (a *Album) QueryIncludedIn() *TopicQuery {
	return (&AlbumClient{config: a.config}).QueryIncludedIn(a)
}

// Update returns a builder for updating this Album.
// Note that you need to call Album.Unwrap() before calling this method if this Album
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Album) Update() *AlbumUpdateOne {
	return (&AlbumClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Album entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Album) Unwrap() *Album {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Album is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Album) String() string {
	var builder strings.Builder
	builder.WriteString("Album(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBy returns the By named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Album) NamedBy(name string) ([]*Artist, error) {
	if a.Edges.namedBy == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedBy[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Album) appendNamedBy(name string, edges ...*Artist) {
	if a.Edges.namedBy == nil {
		a.Edges.namedBy = make(map[string][]*Artist)
	}
	if len(edges) == 0 {
		a.Edges.namedBy[name] = []*Artist{}
	} else {
		a.Edges.namedBy[name] = append(a.Edges.namedBy[name], edges...)
	}
}

// NamedIncludedIn returns the IncludedIn named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Album) NamedIncludedIn(name string) ([]*Topic, error) {
	if a.Edges.namedIncludedIn == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedIncludedIn[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Album) appendNamedIncludedIn(name string, edges ...*Topic) {
	if a.Edges.namedIncludedIn == nil {
		a.Edges.namedIncludedIn = make(map[string][]*Topic)
	}
	if len(edges) == 0 {
		a.Edges.namedIncludedIn[name] = []*Topic{}
	} else {
		a.Edges.namedIncludedIn[name] = append(a.Edges.namedIncludedIn[name], edges...)
	}
}

// Albums is a parsable slice of Album.
type Albums []*Album

func (a Albums) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
