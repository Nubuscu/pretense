package pretense

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"nubuscu/pretense/ent"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"nubuscu/pretense/ent/review"
	"nubuscu/pretense/ent/topic"
)

// UpsertAlbum is the resolver for the upsertAlbum field.
func (r *mutationResolver) UpsertAlbum(ctx context.Context, input ent.CreateAlbumInput) (*ent.Album, error) {
	id, err := r.client.Album.Create().SetInput(input).OnConflictColumns(album.FieldName).Ignore().ID(ctx)
	if err != nil {
		log.Println("failed to upsert Album", input.Name, err)
		return nil, err
	}
	return r.client.Album.Get(ctx, id)
}

// UpsertArtist is the resolver for the upsertArtist field.
func (r *mutationResolver) UpsertArtist(ctx context.Context, input ent.CreateArtistInput) (*ent.Artist, error) {
	id, err := r.client.Artist.Create().SetInput(input).OnConflictColumns(artist.FieldName).Ignore().ID(ctx)
	if err != nil {
		log.Println("failed to upsert Artist", input.Name, err)
		return nil, err
	}
	return r.client.Artist.Get(ctx, id)
}

// UpsertReview is the resolver for the upsertReview field.
func (r *mutationResolver) UpsertReview(ctx context.Context, input ent.CreateReviewInput) (*ent.Review, error) {
	// should still be able to update review body for the same name
	id, err := r.client.Review.Create().SetInput(input).OnConflictColumns(review.FieldName).UpdateNewValues().ID(ctx)
	if err != nil {
		log.Println("failed to upsert Review", input.Name, err)
		return nil, err
	}
	return r.client.Review.Get(ctx, id)
}

// UpsertTopic is the resolver for the upsertTopic field.
func (r *mutationResolver) UpsertTopic(ctx context.Context, input ent.CreateTopicInput) (*ent.Topic, error) {
	id, err := r.client.Topic.Create().SetName(input.Name).OnConflictColumns(topic.FieldName).Ignore().ID(ctx)
	if err != nil {
		log.Println("failed to upsert Topic", input.Name, err)
		return nil, err
	}
	created, err := r.client.Topic.UpdateOneID(id).ClearIncludes().AddIncludeIDs(input.IncludeIDs...).Save(ctx)
	if err != nil {
		log.Println("failed to upsert Topic includes", input.Name, err)
		return nil, err
	}
	return created, nil
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
func (r *mutationResolver) CreateTopicWithReview(ctx context.Context, topicName string, reviewName string, reviewBody string, albumNames []string) (*ent.Topic, error) {
	var albumIds []int
	for _, albumName := range albumNames {

		id, err := r.client.Album.Query().Where(album.NameEQ(albumName)).FirstID(ctx)
		if err != nil {
			return nil, err
		}
		albumIds = append(albumIds, id)
	}
	topic, err := r.UpsertTopic(ctx, ent.CreateTopicInput{
		Name:       topicName,
		IncludeIDs: albumIds,
	})
	if err != nil {
		return nil, err
	}
	review, err := r.UpsertReview(ctx, ent.CreateReviewInput{
		Name: reviewName,
		Body: reviewBody,
	})
	if err != nil {
		return nil, err
	}
	_, revErr := r.client.Review.UpdateOne(review).AddReviews(topic).Save(ctx)
	if revErr != nil {
		log.Println("upserting review - topic edge failed, may already exist.", revErr)
	}
	return topic, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
