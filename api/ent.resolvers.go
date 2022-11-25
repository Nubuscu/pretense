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
func (r *queryResolver) Albums(ctx context.Context) ([]*ent.Album, error) {
	return r.client.Album.Query().All(ctx)
}

// Artists is the resolver for the artists field.
func (r *queryResolver) Artists(ctx context.Context) ([]*ent.Artist, error) {
	return r.client.Artist.Query().All(ctx)
}

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context) ([]*ent.Review, error) {
	return r.client.Review.Query().All(ctx)
}

// Topics is the resolver for the topics field.
func (r *queryResolver) Topics(ctx context.Context) ([]*ent.Topic, error) {
	return r.client.Topic.Query().All(ctx)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
