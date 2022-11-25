// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/vmihailenco/msgpack/v5"
)

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

// MarshalGQL implements graphql.Marshaler interface.
func (o OrderDirection) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(o.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (o *OrderDirection) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("order direction %T must be a string", val)
	}
	*o = OrderDirection(str)
	return o.Validate()
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

func (o OrderDirection) orderFunc(field string) OrderFunc {
	if o == OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

func cursorsToPredicates(direction OrderDirection, after, before *Cursor, field, idField string) []func(s *sql.Selector) {
	var predicates []func(s *sql.Selector)
	if after != nil {
		if after.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeGT
			} else {
				predicate = sql.CompositeLT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					after.Value, after.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.GT
			} else {
				predicate = sql.LT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					after.ID,
				))
			})
		}
	}
	if before != nil {
		if before.Value != nil {
			var predicate func([]string, ...interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.CompositeLT
			} else {
				predicate = sql.CompositeGT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.Columns(field, idField),
					before.Value, before.ID,
				))
			})
		} else {
			var predicate func(string, interface{}) *sql.Predicate
			if direction == OrderDirectionAsc {
				predicate = sql.LT
			} else {
				predicate = sql.GT
			}
			predicates = append(predicates, func(s *sql.Selector) {
				s.Where(predicate(
					s.C(idField),
					before.ID,
				))
			})
		}
	}
	return predicates
}

// PageInfo of a connection type.
type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *Cursor `json:"startCursor"`
	EndCursor       *Cursor `json:"endCursor"`
}

// Cursor of an edge type.
type Cursor struct {
	ID    int   `msgpack:"i"`
	Value Value `msgpack:"v,omitempty"`
}

// MarshalGQL implements graphql.Marshaler interface.
func (c Cursor) MarshalGQL(w io.Writer) {
	quote := []byte{'"'}
	w.Write(quote)
	defer w.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, w)
	defer wc.Close()
	_ = msgpack.NewEncoder(wc).Encode(c)
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (c *Cursor) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("%T is not a string", v)
	}
	if err := msgpack.NewDecoder(
		base64.NewDecoder(
			base64.RawStdEncoding,
			strings.NewReader(s),
		),
	).Decode(c); err != nil {
		return fmt.Errorf("cannot decode cursor: %w", err)
	}
	return nil
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// AlbumEdge is the edge representation of Album.
type AlbumEdge struct {
	Node   *Album `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// AlbumConnection is the connection containing edges to Album.
type AlbumConnection struct {
	Edges      []*AlbumEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *AlbumConnection) build(nodes []*Album, pager *albumPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Album
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Album {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Album {
			return nodes[i]
		}
	}
	c.Edges = make([]*AlbumEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &AlbumEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// AlbumPaginateOption enables pagination customization.
type AlbumPaginateOption func(*albumPager) error

// WithAlbumOrder configures pagination ordering.
func WithAlbumOrder(order *AlbumOrder) AlbumPaginateOption {
	if order == nil {
		order = DefaultAlbumOrder
	}
	o := *order
	return func(pager *albumPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultAlbumOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithAlbumFilter configures pagination filter.
func WithAlbumFilter(filter func(*AlbumQuery) (*AlbumQuery, error)) AlbumPaginateOption {
	return func(pager *albumPager) error {
		if filter == nil {
			return errors.New("AlbumQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type albumPager struct {
	order  *AlbumOrder
	filter func(*AlbumQuery) (*AlbumQuery, error)
}

func newAlbumPager(opts []AlbumPaginateOption) (*albumPager, error) {
	pager := &albumPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultAlbumOrder
	}
	return pager, nil
}

func (p *albumPager) applyFilter(query *AlbumQuery) (*AlbumQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *albumPager) toCursor(a *Album) Cursor {
	return p.order.Field.toCursor(a)
}

func (p *albumPager) applyCursors(query *AlbumQuery, after, before *Cursor) *AlbumQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultAlbumOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *albumPager) applyOrder(query *AlbumQuery, reverse bool) *AlbumQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultAlbumOrder.Field {
		query = query.Order(direction.orderFunc(DefaultAlbumOrder.Field.field))
	}
	return query
}

func (p *albumPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultAlbumOrder.Field {
			b.Comma().Ident(DefaultAlbumOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Album.
func (a *AlbumQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...AlbumPaginateOption,
) (*AlbumConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newAlbumPager(opts)
	if err != nil {
		return nil, err
	}
	if a, err = pager.applyFilter(a); err != nil {
		return nil, err
	}
	conn := &AlbumConnection{Edges: []*AlbumEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = a.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	a = pager.applyCursors(a, after, before)
	a = pager.applyOrder(a, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		a.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := a.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// AlbumOrderField defines the ordering field of Album.
type AlbumOrderField struct {
	field    string
	toCursor func(*Album) Cursor
}

// AlbumOrder defines the ordering of Album.
type AlbumOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *AlbumOrderField `json:"field"`
}

// DefaultAlbumOrder is the default ordering of Album.
var DefaultAlbumOrder = &AlbumOrder{
	Direction: OrderDirectionAsc,
	Field: &AlbumOrderField{
		field: album.FieldID,
		toCursor: func(a *Album) Cursor {
			return Cursor{ID: a.ID}
		},
	},
}

// ToEdge converts Album into AlbumEdge.
func (a *Album) ToEdge(order *AlbumOrder) *AlbumEdge {
	if order == nil {
		order = DefaultAlbumOrder
	}
	return &AlbumEdge{
		Node:   a,
		Cursor: order.Field.toCursor(a),
	}
}

// ArtistEdge is the edge representation of Artist.
type ArtistEdge struct {
	Node   *Artist `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// ArtistConnection is the connection containing edges to Artist.
type ArtistConnection struct {
	Edges      []*ArtistEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *ArtistConnection) build(nodes []*Artist, pager *artistPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Artist
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Artist {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Artist {
			return nodes[i]
		}
	}
	c.Edges = make([]*ArtistEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &ArtistEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// ArtistPaginateOption enables pagination customization.
type ArtistPaginateOption func(*artistPager) error

// WithArtistOrder configures pagination ordering.
func WithArtistOrder(order *ArtistOrder) ArtistPaginateOption {
	if order == nil {
		order = DefaultArtistOrder
	}
	o := *order
	return func(pager *artistPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultArtistOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithArtistFilter configures pagination filter.
func WithArtistFilter(filter func(*ArtistQuery) (*ArtistQuery, error)) ArtistPaginateOption {
	return func(pager *artistPager) error {
		if filter == nil {
			return errors.New("ArtistQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type artistPager struct {
	order  *ArtistOrder
	filter func(*ArtistQuery) (*ArtistQuery, error)
}

func newArtistPager(opts []ArtistPaginateOption) (*artistPager, error) {
	pager := &artistPager{}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultArtistOrder
	}
	return pager, nil
}

func (p *artistPager) applyFilter(query *ArtistQuery) (*ArtistQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *artistPager) toCursor(a *Artist) Cursor {
	return p.order.Field.toCursor(a)
}

func (p *artistPager) applyCursors(query *ArtistQuery, after, before *Cursor) *ArtistQuery {
	for _, predicate := range cursorsToPredicates(
		p.order.Direction, after, before,
		p.order.Field.field, DefaultArtistOrder.Field.field,
	) {
		query = query.Where(predicate)
	}
	return query
}

func (p *artistPager) applyOrder(query *ArtistQuery, reverse bool) *ArtistQuery {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	query = query.Order(direction.orderFunc(p.order.Field.field))
	if p.order.Field != DefaultArtistOrder.Field {
		query = query.Order(direction.orderFunc(DefaultArtistOrder.Field.field))
	}
	return query
}

func (p *artistPager) orderExpr(reverse bool) sql.Querier {
	direction := p.order.Direction
	if reverse {
		direction = direction.reverse()
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.field).Pad().WriteString(string(direction))
		if p.order.Field != DefaultArtistOrder.Field {
			b.Comma().Ident(DefaultArtistOrder.Field.field).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Artist.
func (a *ArtistQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...ArtistPaginateOption,
) (*ArtistConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newArtistPager(opts)
	if err != nil {
		return nil, err
	}
	if a, err = pager.applyFilter(a); err != nil {
		return nil, err
	}
	conn := &ArtistConnection{Edges: []*ArtistEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = a.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	a = pager.applyCursors(a, after, before)
	a = pager.applyOrder(a, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		a.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := a.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// ArtistOrderField defines the ordering field of Artist.
type ArtistOrderField struct {
	field    string
	toCursor func(*Artist) Cursor
}

// ArtistOrder defines the ordering of Artist.
type ArtistOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *ArtistOrderField `json:"field"`
}

// DefaultArtistOrder is the default ordering of Artist.
var DefaultArtistOrder = &ArtistOrder{
	Direction: OrderDirectionAsc,
	Field: &ArtistOrderField{
		field: artist.FieldID,
		toCursor: func(a *Artist) Cursor {
			return Cursor{ID: a.ID}
		},
	},
}

// ToEdge converts Artist into ArtistEdge.
func (a *Artist) ToEdge(order *ArtistOrder) *ArtistEdge {
	if order == nil {
		order = DefaultArtistOrder
	}
	return &ArtistEdge{
		Node:   a,
		Cursor: order.Field.toCursor(a),
	}
}
