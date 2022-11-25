// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Album) By(ctx context.Context) (result []*Artist, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedBy(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.ByOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryBy().All(ctx)
	}
	return result, err
}

func (a *Album) IncludedIn(ctx context.Context) (result []*Topic, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedIncludedIn(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.IncludedInOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryIncludedIn().All(ctx)
	}
	return result, err
}

func (a *Artist) Wrote(ctx context.Context) (result []*Album, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedWrote(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.WroteOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryWrote().All(ctx)
	}
	return result, err
}

func (r *Review) Reviews(ctx context.Context) (result []*Topic, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedReviews(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.ReviewsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryReviews().All(ctx)
	}
	return result, err
}

func (t *Topic) ReviewedBy(ctx context.Context) (result []*Review, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedReviewedBy(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ReviewedByOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryReviewedBy().All(ctx)
	}
	return result, err
}

func (t *Topic) Includes(ctx context.Context) (result []*Album, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedIncludes(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.IncludesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryIncludes().All(ctx)
	}
	return result, err
}
