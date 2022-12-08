// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"nubuscu/pretense/ent/migrate"

	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"nubuscu/pretense/ent/review"
	"nubuscu/pretense/ent/topic"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Album is the client for interacting with the Album builders.
	Album *AlbumClient
	// Artist is the client for interacting with the Artist builders.
	Artist *ArtistClient
	// Review is the client for interacting with the Review builders.
	Review *ReviewClient
	// Topic is the client for interacting with the Topic builders.
	Topic *TopicClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Album = NewAlbumClient(c.config)
	c.Artist = NewArtistClient(c.config)
	c.Review = NewReviewClient(c.config)
	c.Topic = NewTopicClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Album:  NewAlbumClient(cfg),
		Artist: NewArtistClient(cfg),
		Review: NewReviewClient(cfg),
		Topic:  NewTopicClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Album:  NewAlbumClient(cfg),
		Artist: NewArtistClient(cfg),
		Review: NewReviewClient(cfg),
		Topic:  NewTopicClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Album.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Album.Use(hooks...)
	c.Artist.Use(hooks...)
	c.Review.Use(hooks...)
	c.Topic.Use(hooks...)
}

// AlbumClient is a client for the Album schema.
type AlbumClient struct {
	config
}

// NewAlbumClient returns a client for the Album from the given config.
func NewAlbumClient(c config) *AlbumClient {
	return &AlbumClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `album.Hooks(f(g(h())))`.
func (c *AlbumClient) Use(hooks ...Hook) {
	c.hooks.Album = append(c.hooks.Album, hooks...)
}

// Create returns a builder for creating a Album entity.
func (c *AlbumClient) Create() *AlbumCreate {
	mutation := newAlbumMutation(c.config, OpCreate)
	return &AlbumCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Album entities.
func (c *AlbumClient) CreateBulk(builders ...*AlbumCreate) *AlbumCreateBulk {
	return &AlbumCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Album.
func (c *AlbumClient) Update() *AlbumUpdate {
	mutation := newAlbumMutation(c.config, OpUpdate)
	return &AlbumUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AlbumClient) UpdateOne(a *Album) *AlbumUpdateOne {
	mutation := newAlbumMutation(c.config, OpUpdateOne, withAlbum(a))
	return &AlbumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AlbumClient) UpdateOneID(id int) *AlbumUpdateOne {
	mutation := newAlbumMutation(c.config, OpUpdateOne, withAlbumID(id))
	return &AlbumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Album.
func (c *AlbumClient) Delete() *AlbumDelete {
	mutation := newAlbumMutation(c.config, OpDelete)
	return &AlbumDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AlbumClient) DeleteOne(a *Album) *AlbumDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AlbumClient) DeleteOneID(id int) *AlbumDeleteOne {
	builder := c.Delete().Where(album.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AlbumDeleteOne{builder}
}

// Query returns a query builder for Album.
func (c *AlbumClient) Query() *AlbumQuery {
	return &AlbumQuery{
		config: c.config,
	}
}

// Get returns a Album entity by its id.
func (c *AlbumClient) Get(ctx context.Context, id int) (*Album, error) {
	return c.Query().Where(album.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AlbumClient) GetX(ctx context.Context, id int) *Album {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBy queries the by edge of a Album.
func (c *AlbumClient) QueryBy(a *Album) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, album.ByTable, album.ByPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIncludedIn queries the included_in edge of a Album.
func (c *AlbumClient) QueryIncludedIn(a *Album) *TopicQuery {
	query := &TopicQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, album.IncludedInTable, album.IncludedInPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AlbumClient) Hooks() []Hook {
	return c.hooks.Album
}

// ArtistClient is a client for the Artist schema.
type ArtistClient struct {
	config
}

// NewArtistClient returns a client for the Artist from the given config.
func NewArtistClient(c config) *ArtistClient {
	return &ArtistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `artist.Hooks(f(g(h())))`.
func (c *ArtistClient) Use(hooks ...Hook) {
	c.hooks.Artist = append(c.hooks.Artist, hooks...)
}

// Create returns a builder for creating a Artist entity.
func (c *ArtistClient) Create() *ArtistCreate {
	mutation := newArtistMutation(c.config, OpCreate)
	return &ArtistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Artist entities.
func (c *ArtistClient) CreateBulk(builders ...*ArtistCreate) *ArtistCreateBulk {
	return &ArtistCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Artist.
func (c *ArtistClient) Update() *ArtistUpdate {
	mutation := newArtistMutation(c.config, OpUpdate)
	return &ArtistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArtistClient) UpdateOne(a *Artist) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtist(a))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArtistClient) UpdateOneID(id int) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtistID(id))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Artist.
func (c *ArtistClient) Delete() *ArtistDelete {
	mutation := newArtistMutation(c.config, OpDelete)
	return &ArtistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArtistClient) DeleteOne(a *Artist) *ArtistDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArtistClient) DeleteOneID(id int) *ArtistDeleteOne {
	builder := c.Delete().Where(artist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArtistDeleteOne{builder}
}

// Query returns a query builder for Artist.
func (c *ArtistClient) Query() *ArtistQuery {
	return &ArtistQuery{
		config: c.config,
	}
}

// Get returns a Artist entity by its id.
func (c *ArtistClient) Get(ctx context.Context, id int) (*Artist, error) {
	return c.Query().Where(artist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArtistClient) GetX(ctx context.Context, id int) *Artist {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryWrote queries the wrote edge of a Artist.
func (c *ArtistClient) QueryWrote(a *Artist) *AlbumQuery {
	query := &AlbumQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, artist.WroteTable, artist.WrotePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArtistClient) Hooks() []Hook {
	return c.hooks.Artist
}

// ReviewClient is a client for the Review schema.
type ReviewClient struct {
	config
}

// NewReviewClient returns a client for the Review from the given config.
func NewReviewClient(c config) *ReviewClient {
	return &ReviewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `review.Hooks(f(g(h())))`.
func (c *ReviewClient) Use(hooks ...Hook) {
	c.hooks.Review = append(c.hooks.Review, hooks...)
}

// Create returns a builder for creating a Review entity.
func (c *ReviewClient) Create() *ReviewCreate {
	mutation := newReviewMutation(c.config, OpCreate)
	return &ReviewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Review entities.
func (c *ReviewClient) CreateBulk(builders ...*ReviewCreate) *ReviewCreateBulk {
	return &ReviewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Review.
func (c *ReviewClient) Update() *ReviewUpdate {
	mutation := newReviewMutation(c.config, OpUpdate)
	return &ReviewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewClient) UpdateOne(r *Review) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReview(r))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewClient) UpdateOneID(id int) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReviewID(id))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Review.
func (c *ReviewClient) Delete() *ReviewDelete {
	mutation := newReviewMutation(c.config, OpDelete)
	return &ReviewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReviewClient) DeleteOne(r *Review) *ReviewDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ReviewClient) DeleteOneID(id int) *ReviewDeleteOne {
	builder := c.Delete().Where(review.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewDeleteOne{builder}
}

// Query returns a query builder for Review.
func (c *ReviewClient) Query() *ReviewQuery {
	return &ReviewQuery{
		config: c.config,
	}
}

// Get returns a Review entity by its id.
func (c *ReviewClient) Get(ctx context.Context, id int) (*Review, error) {
	return c.Query().Where(review.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewClient) GetX(ctx context.Context, id int) *Review {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReviews queries the reviews edge of a Review.
func (c *ReviewClient) QueryReviews(r *Review) *TopicQuery {
	query := &TopicQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, review.ReviewsTable, review.ReviewsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReviewClient) Hooks() []Hook {
	return c.hooks.Review
}

// TopicClient is a client for the Topic schema.
type TopicClient struct {
	config
}

// NewTopicClient returns a client for the Topic from the given config.
func NewTopicClient(c config) *TopicClient {
	return &TopicClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `topic.Hooks(f(g(h())))`.
func (c *TopicClient) Use(hooks ...Hook) {
	c.hooks.Topic = append(c.hooks.Topic, hooks...)
}

// Create returns a builder for creating a Topic entity.
func (c *TopicClient) Create() *TopicCreate {
	mutation := newTopicMutation(c.config, OpCreate)
	return &TopicCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Topic entities.
func (c *TopicClient) CreateBulk(builders ...*TopicCreate) *TopicCreateBulk {
	return &TopicCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Topic.
func (c *TopicClient) Update() *TopicUpdate {
	mutation := newTopicMutation(c.config, OpUpdate)
	return &TopicUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TopicClient) UpdateOne(t *Topic) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopic(t))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TopicClient) UpdateOneID(id int) *TopicUpdateOne {
	mutation := newTopicMutation(c.config, OpUpdateOne, withTopicID(id))
	return &TopicUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Topic.
func (c *TopicClient) Delete() *TopicDelete {
	mutation := newTopicMutation(c.config, OpDelete)
	return &TopicDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TopicClient) DeleteOne(t *Topic) *TopicDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TopicClient) DeleteOneID(id int) *TopicDeleteOne {
	builder := c.Delete().Where(topic.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TopicDeleteOne{builder}
}

// Query returns a query builder for Topic.
func (c *TopicClient) Query() *TopicQuery {
	return &TopicQuery{
		config: c.config,
	}
}

// Get returns a Topic entity by its id.
func (c *TopicClient) Get(ctx context.Context, id int) (*Topic, error) {
	return c.Query().Where(topic.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TopicClient) GetX(ctx context.Context, id int) *Topic {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReviewedBy queries the reviewed_by edge of a Topic.
func (c *TopicClient) QueryReviewedBy(t *Topic) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topic.Table, topic.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, topic.ReviewedByTable, topic.ReviewedByPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIncludes queries the includes edge of a Topic.
func (c *TopicClient) QueryIncludes(t *Topic) *AlbumQuery {
	query := &AlbumQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(topic.Table, topic.FieldID, id),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, topic.IncludesTable, topic.IncludesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TopicClient) Hooks() []Hook {
	return c.hooks.Topic
}