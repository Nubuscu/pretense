// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"nubuscu/pretense/ent/review"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Review is the model entity for the Review schema.
type Review struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReviewQuery when eager-loading is set.
	Edges ReviewEdges `json:"edges"`
}

// ReviewEdges holds the relations/edges for other nodes in the graph.
type ReviewEdges struct {
	// Reviews holds the value of the reviews edge.
	Reviews []*Topic `json:"reviews,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedReviews map[string][]*Topic
}

// ReviewsOrErr returns the Reviews value or an error if the edge
// was not loaded in eager-loading.
func (e ReviewEdges) ReviewsOrErr() ([]*Topic, error) {
	if e.loadedTypes[0] {
		return e.Reviews, nil
	}
	return nil, &NotLoadedError{edge: "reviews"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Review) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case review.FieldID:
			values[i] = new(sql.NullInt64)
		case review.FieldBody:
			values[i] = new(sql.NullString)
		case review.FieldCreatedAt, review.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Review", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Review fields.
func (r *Review) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case review.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case review.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case review.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case review.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				r.Body = value.String
			}
		}
	}
	return nil
}

// QueryReviews queries the "reviews" edge of the Review entity.
func (r *Review) QueryReviews() *TopicQuery {
	return (&ReviewClient{config: r.config}).QueryReviews(r)
}

// Update returns a builder for updating this Review.
// Note that you need to call Review.Unwrap() before calling this method if this Review
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Review) Update() *ReviewUpdateOne {
	return (&ReviewClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Review entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Review) Unwrap() *Review {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Review is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Review) String() string {
	var builder strings.Builder
	builder.WriteString("Review(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(r.Body)
	builder.WriteByte(')')
	return builder.String()
}

// NamedReviews returns the Reviews named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Review) NamedReviews(name string) ([]*Topic, error) {
	if r.Edges.namedReviews == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedReviews[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Review) appendNamedReviews(name string, edges ...*Topic) {
	if r.Edges.namedReviews == nil {
		r.Edges.namedReviews = make(map[string][]*Topic)
	}
	if len(edges) == 0 {
		r.Edges.namedReviews[name] = []*Topic{}
	} else {
		r.Edges.namedReviews[name] = append(r.Edges.namedReviews[name], edges...)
	}
}

// Reviews is a parsable slice of Review.
type Reviews []*Review

func (r Reviews) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}