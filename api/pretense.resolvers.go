package pretense

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"nubuscu/pretense/ent"
)

// CreateAlbum is the resolver for the createAlbum field.
func (r *mutationResolver) CreateAlbum(ctx context.Context, input ent.CreateAlbumInput) (*ent.Album, error) {
	return r.client.Album.Create().SetInput(input).Save(ctx)
}

// CreateArtist is the resolver for the createArtist field.
func (r *mutationResolver) CreateArtist(ctx context.Context, input ent.CreateArtistInput) (*ent.Artist, error) {
	return r.client.Artist.Create().SetInput(input).Save(ctx)
}

// CreateReview is the resolver for the CreateReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, input ent.CreateReviewInput) (*ent.Review, error) {
	return r.client.Review.Create().SetInput(input).Save(ctx)
}

// CreateTopic is the resolver for the CreateTopic field.
func (r *mutationResolver) CreateTopic(ctx context.Context, input ent.CreateTopicInput) (*ent.Topic, error) {
	return r.client.Topic.Create().SetInput(input).Save(ctx)
}

// UpdateAlbum is the resolver for the updateAlbum field.
func (r *mutationResolver) UpdateAlbum(ctx context.Context, id int, input ent.UpdateAlbumInput) (*ent.Album, error) {
	return r.client.Album.UpdateOneID(id).SetInput(input).Save(ctx)
}

// UpdateArtist is the resolver for the updateArtist field.
func (r *mutationResolver) UpdateArtist(ctx context.Context, id int, input ent.UpdateArtistInput) (*ent.Artist, error) {
	return r.client.Artist.UpdateOneID(id).SetInput(input).Save(ctx)
}

// UpdateReview is the resolver for the updateReview field.
func (r *mutationResolver) UpdateReview(ctx context.Context, id int, input ent.UpdateReviewInput) (*ent.Review, error) {
	return r.client.Review.UpdateOneID(id).SetInput(input).Save(ctx)
}

// UpdateTopic is the resolver for the updateTopic field.
func (r *mutationResolver) UpdateTopic(ctx context.Context, id int, input ent.UpdateTopicInput) (*ent.Topic, error) {
	return r.client.Topic.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
