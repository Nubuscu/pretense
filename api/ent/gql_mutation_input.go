// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateAlbumInput represents a mutation input for creating albums.
type CreateAlbumInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	MetaLabels    []string
	SpotifyURL    string
	Name          string
	ByIDs         []int
	IncludedInIDs []int
	TaggedWithIDs []int
}

// Mutate applies the CreateAlbumInput on the AlbumMutation builder.
func (i *CreateAlbumInput) Mutate(m *AlbumMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	m.SetSpotifyURL(i.SpotifyURL)
	m.SetName(i.Name)
	if v := i.ByIDs; len(v) > 0 {
		m.AddByIDs(v...)
	}
	if v := i.IncludedInIDs; len(v) > 0 {
		m.AddIncludedInIDs(v...)
	}
	if v := i.TaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the CreateAlbumInput on the AlbumCreate builder.
func (c *AlbumCreate) SetInput(i CreateAlbumInput) *AlbumCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAlbumInput represents a mutation input for updating albums.
type UpdateAlbumInput struct {
	UpdatedAt           *time.Time
	MetaLabels          []string
	AppendMetaLabels    []string
	SpotifyURL          *string
	Name                *string
	AddByIDs            []int
	RemoveByIDs         []int
	AddIncludedInIDs    []int
	RemoveIncludedInIDs []int
	AddTaggedWithIDs    []int
	RemoveTaggedWithIDs []int
}

// Mutate applies the UpdateAlbumInput on the AlbumMutation builder.
func (i *UpdateAlbumInput) Mutate(m *AlbumMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	if i.AppendMetaLabels != nil {
		m.AppendMetaLabels(i.MetaLabels)
	}
	if v := i.SpotifyURL; v != nil {
		m.SetSpotifyURL(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.AddByIDs; len(v) > 0 {
		m.AddByIDs(v...)
	}
	if v := i.RemoveByIDs; len(v) > 0 {
		m.RemoveByIDs(v...)
	}
	if v := i.AddIncludedInIDs; len(v) > 0 {
		m.AddIncludedInIDs(v...)
	}
	if v := i.RemoveIncludedInIDs; len(v) > 0 {
		m.RemoveIncludedInIDs(v...)
	}
	if v := i.AddTaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
	if v := i.RemoveTaggedWithIDs; len(v) > 0 {
		m.RemoveTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateAlbumInput on the AlbumUpdate builder.
func (c *AlbumUpdate) SetInput(i UpdateAlbumInput) *AlbumUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAlbumInput on the AlbumUpdateOne builder.
func (c *AlbumUpdateOne) SetInput(i UpdateAlbumInput) *AlbumUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateArtistInput represents a mutation input for creating artists.
type CreateArtistInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	MetaLabels    []string
	SpotifyURL    string
	Name          string
	WroteIDs      []int
	TaggedWithIDs []int
}

// Mutate applies the CreateArtistInput on the ArtistMutation builder.
func (i *CreateArtistInput) Mutate(m *ArtistMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	m.SetSpotifyURL(i.SpotifyURL)
	m.SetName(i.Name)
	if v := i.WroteIDs; len(v) > 0 {
		m.AddWroteIDs(v...)
	}
	if v := i.TaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the CreateArtistInput on the ArtistCreate builder.
func (c *ArtistCreate) SetInput(i CreateArtistInput) *ArtistCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateArtistInput represents a mutation input for updating artists.
type UpdateArtistInput struct {
	UpdatedAt           *time.Time
	MetaLabels          []string
	AppendMetaLabels    []string
	SpotifyURL          *string
	Name                *string
	AddWroteIDs         []int
	RemoveWroteIDs      []int
	AddTaggedWithIDs    []int
	RemoveTaggedWithIDs []int
}

// Mutate applies the UpdateArtistInput on the ArtistMutation builder.
func (i *UpdateArtistInput) Mutate(m *ArtistMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	if i.AppendMetaLabels != nil {
		m.AppendMetaLabels(i.MetaLabels)
	}
	if v := i.SpotifyURL; v != nil {
		m.SetSpotifyURL(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.AddWroteIDs; len(v) > 0 {
		m.AddWroteIDs(v...)
	}
	if v := i.RemoveWroteIDs; len(v) > 0 {
		m.RemoveWroteIDs(v...)
	}
	if v := i.AddTaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
	if v := i.RemoveTaggedWithIDs; len(v) > 0 {
		m.RemoveTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateArtistInput on the ArtistUpdate builder.
func (c *ArtistUpdate) SetInput(i UpdateArtistInput) *ArtistUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateArtistInput on the ArtistUpdateOne builder.
func (c *ArtistUpdateOne) SetInput(i UpdateArtistInput) *ArtistUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateReviewInput represents a mutation input for creating reviews.
type CreateReviewInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	MetaLabels    []string
	Name          string
	Body          string
	ReviewIDs     []int
	TaggedWithIDs []int
}

// Mutate applies the CreateReviewInput on the ReviewMutation builder.
func (i *CreateReviewInput) Mutate(m *ReviewMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	m.SetName(i.Name)
	m.SetBody(i.Body)
	if v := i.ReviewIDs; len(v) > 0 {
		m.AddReviewIDs(v...)
	}
	if v := i.TaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the CreateReviewInput on the ReviewCreate builder.
func (c *ReviewCreate) SetInput(i CreateReviewInput) *ReviewCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateReviewInput represents a mutation input for updating reviews.
type UpdateReviewInput struct {
	UpdatedAt           *time.Time
	MetaLabels          []string
	AppendMetaLabels    []string
	Name                *string
	Body                *string
	AddReviewIDs        []int
	RemoveReviewIDs     []int
	AddTaggedWithIDs    []int
	RemoveTaggedWithIDs []int
}

// Mutate applies the UpdateReviewInput on the ReviewMutation builder.
func (i *UpdateReviewInput) Mutate(m *ReviewMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	if i.AppendMetaLabels != nil {
		m.AppendMetaLabels(i.MetaLabels)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
	if v := i.AddReviewIDs; len(v) > 0 {
		m.AddReviewIDs(v...)
	}
	if v := i.RemoveReviewIDs; len(v) > 0 {
		m.RemoveReviewIDs(v...)
	}
	if v := i.AddTaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
	if v := i.RemoveTaggedWithIDs; len(v) > 0 {
		m.RemoveTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateReviewInput on the ReviewUpdate builder.
func (c *ReviewUpdate) SetInput(i UpdateReviewInput) *ReviewUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateReviewInput on the ReviewUpdateOne builder.
func (c *ReviewUpdateOne) SetInput(i UpdateReviewInput) *ReviewUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateTopicInput represents a mutation input for creating topics.
type CreateTopicInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	MetaLabels    []string
	Name          string
	ReviewedByIDs []int
	IncludeIDs    []int
	TaggedWithIDs []int
}

// Mutate applies the CreateTopicInput on the TopicMutation builder.
func (i *CreateTopicInput) Mutate(m *TopicMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	m.SetName(i.Name)
	if v := i.ReviewedByIDs; len(v) > 0 {
		m.AddReviewedByIDs(v...)
	}
	if v := i.IncludeIDs; len(v) > 0 {
		m.AddIncludeIDs(v...)
	}
	if v := i.TaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTopicInput on the TopicCreate builder.
func (c *TopicCreate) SetInput(i CreateTopicInput) *TopicCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTopicInput represents a mutation input for updating topics.
type UpdateTopicInput struct {
	UpdatedAt           *time.Time
	MetaLabels          []string
	AppendMetaLabels    []string
	Name                *string
	AddReviewedByIDs    []int
	RemoveReviewedByIDs []int
	AddIncludeIDs       []int
	RemoveIncludeIDs    []int
	AddTaggedWithIDs    []int
	RemoveTaggedWithIDs []int
}

// Mutate applies the UpdateTopicInput on the TopicMutation builder.
func (i *UpdateTopicInput) Mutate(m *TopicMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.MetaLabels; v != nil {
		m.SetMetaLabels(v)
	}
	if i.AppendMetaLabels != nil {
		m.AppendMetaLabels(i.MetaLabels)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.AddReviewedByIDs; len(v) > 0 {
		m.AddReviewedByIDs(v...)
	}
	if v := i.RemoveReviewedByIDs; len(v) > 0 {
		m.RemoveReviewedByIDs(v...)
	}
	if v := i.AddIncludeIDs; len(v) > 0 {
		m.AddIncludeIDs(v...)
	}
	if v := i.RemoveIncludeIDs; len(v) > 0 {
		m.RemoveIncludeIDs(v...)
	}
	if v := i.AddTaggedWithIDs; len(v) > 0 {
		m.AddTaggedWithIDs(v...)
	}
	if v := i.RemoveTaggedWithIDs; len(v) > 0 {
		m.RemoveTaggedWithIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateTopicInput on the TopicUpdate builder.
func (c *TopicUpdate) SetInput(i UpdateTopicInput) *TopicUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTopicInput on the TopicUpdateOne builder.
func (c *TopicUpdateOne) SetInput(i UpdateTopicInput) *TopicUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
