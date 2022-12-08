package pretense

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"nubuscu/pretense/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Albums is the resolver for the albums field.
func (r *queryResolver) Albums(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.AlbumWhereInput) (*ent.AlbumConnection, error) {
	return r.client.Album.Query().
		Paginate(ctx, after, first, before, last, ent.WithAlbumOrder(ent.DefaultAlbumOrder), ent.WithAlbumFilter(where.Filter))
}

// Artists is the resolver for the artists field.
func (r *queryResolver) Artists(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ArtistWhereInput) (*ent.ArtistConnection, error) {
	return r.client.Artist.Query().
		Paginate(ctx, after, first, before, last, ent.WithArtistOrder(ent.DefaultArtistOrder), ent.WithArtistFilter(where.Filter))
}

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.ReviewWhereInput) (*ent.ReviewConnection, error) {
	return r.client.Review.Query().
		Paginate(ctx, after, first, before, last, ent.WithReviewOrder(ent.DefaultReviewOrder), ent.WithReviewFilter(where.Filter))
}

// Topics is the resolver for the topics field.
func (r *queryResolver) Topics(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TopicWhereInput) (*ent.TopicConnection, error) {
	return r.client.Topic.Query().
		Paginate(ctx, after, first, before, last, ent.WithTopicOrder(ent.DefaultTopicOrder), ent.WithTopicFilter(where.Filter))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
