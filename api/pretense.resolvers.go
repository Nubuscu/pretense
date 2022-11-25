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

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
