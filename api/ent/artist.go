// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"nubuscu/pretense/ent/artist"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Artist is the model entity for the Artist schema.
type Artist struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArtistQuery when eager-loading is set.
	Edges ArtistEdges `json:"edges"`
}

// ArtistEdges holds the relations/edges for other nodes in the graph.
type ArtistEdges struct {
	// Wrote holds the value of the wrote edge.
	Wrote []*Album `json:"wrote,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedWrote map[string][]*Album
}

// WroteOrErr returns the Wrote value or an error if the edge
// was not loaded in eager-loading.
func (e ArtistEdges) WroteOrErr() ([]*Album, error) {
	if e.loadedTypes[0] {
		return e.Wrote, nil
	}
	return nil, &NotLoadedError{edge: "wrote"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Artist) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case artist.FieldID:
			values[i] = new(sql.NullInt64)
		case artist.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Artist", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Artist fields.
func (a *Artist) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case artist.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case artist.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		}
	}
	return nil
}

// QueryWrote queries the "wrote" edge of the Artist entity.
func (a *Artist) QueryWrote() *AlbumQuery {
	return (&ArtistClient{config: a.config}).QueryWrote(a)
}

// Update returns a builder for updating this Artist.
// Note that you need to call Artist.Unwrap() before calling this method if this Artist
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Artist) Update() *ArtistUpdateOne {
	return (&ArtistClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Artist entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Artist) Unwrap() *Artist {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Artist is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Artist) String() string {
	var builder strings.Builder
	builder.WriteString("Artist(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedWrote returns the Wrote named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Artist) NamedWrote(name string) ([]*Album, error) {
	if a.Edges.namedWrote == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedWrote[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Artist) appendNamedWrote(name string, edges ...*Album) {
	if a.Edges.namedWrote == nil {
		a.Edges.namedWrote = make(map[string][]*Album)
	}
	if len(edges) == 0 {
		a.Edges.namedWrote[name] = []*Album{}
	} else {
		a.Edges.namedWrote[name] = append(a.Edges.namedWrote[name], edges...)
	}
}

// Artists is a parsable slice of Artist.
type Artists []*Artist

func (a Artists) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
