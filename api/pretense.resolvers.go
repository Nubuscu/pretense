package pretense

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"nubuscu/pretense/ent"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"nubuscu/pretense/ent/review"
	"nubuscu/pretense/ent/topic"
)

// UpsertAlbum is the resolver for the upsertAlbum field.
func (r *mutationResolver) UpsertAlbum(ctx context.Context, input ent.CreateAlbumInput) (*ent.Album, error) {
	id, err := r.client.Album.Create().SetInput(input).OnConflictColumns(album.FieldName).UpdateNewValues().ID(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to upsert Album %w", err))
	}
	return r.client.Album.Get(ctx, id)
}

// UpsertArtist is the resolver for the upsertArtist field.
func (r *mutationResolver) UpsertArtist(ctx context.Context, input ent.CreateArtistInput) (*ent.Artist, error) {
	id, err := r.client.Artist.Create().SetInput(input).OnConflictColumns(artist.FieldName).UpdateNewValues().ID(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to upsert Artist %w", err))
	}
	return r.client.Artist.Get(ctx, id)
}

// UpsertReview is the resolver for the upsertReview field.
func (r *mutationResolver) UpsertReview(ctx context.Context, input ent.CreateReviewInput) (*ent.Review, error) {
	id, err := r.client.Review.Create().SetInput(input).OnConflictColumns(review.FieldName).UpdateNewValues().ID(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to upsert Review %w", err))
	}
	return r.client.Review.Get(ctx, id)
}

// UpsertTopic is the resolver for the upsertTopic field.
func (r *mutationResolver) UpsertTopic(ctx context.Context, input ent.CreateTopicInput) (*ent.Topic, error) {
	id, err := r.client.Topic.Create().SetInput(input).OnConflictColumns(topic.FieldName).UpdateNewValues().ID(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to upsert Topic %w", err))
	}
	return r.client.Topic.Get(ctx, id)
}

// CreateAlbumAndArtists is the resolver for the createAlbumAndArtists field.
func (r *mutationResolver) CreateAlbumAndArtists(ctx context.Context, album ent.CreateAlbumInput, artists []*ent.CreateArtistInput) (*ent.Album, error) {
	createdAlbum, err := r.UpsertAlbum(ctx, album)
	if err != nil {
		return nil, err
	}
	var artistIds []int
	for _, artistInput := range artists {
		createdArtist, err := r.UpsertArtist(ctx, *artistInput)
		if err != nil {
			return nil, err
		}
		artistIds = append(artistIds, createdArtist.ID)
	}
	return r.client.Album.UpdateOne(createdAlbum).AddByIDs(artistIds...).Save(ctx)
}

// CreateTopicWithReview is the resolver for the createTopicWithReview field.
func (r *mutationResolver) CreateTopicWithReview(ctx context.Context, topicName string, reviewName string, reviewBody string, albums []*ent.CreateAlbumInput) (*ent.Topic, error) {
	var albumIds []int
	for _, album := range albums {
		createdAlbum, _ := r.UpsertAlbum(ctx, *album)
		albumIds = append(albumIds, createdAlbum.ID)
	}
	topic, _ := r.UpsertTopic(ctx, ent.CreateTopicInput{
		Name:       topicName,
		IncludeIDs: albumIds,
	})
	r.UpsertReview(ctx, ent.CreateReviewInput{
		Name:      reviewName,
		Body:      reviewBody,
		ReviewIDs: []int{topic.ID},
	})
	return topic, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
