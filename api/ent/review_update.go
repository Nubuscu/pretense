// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nubuscu/pretense/ent/predicate"
	"nubuscu/pretense/ent/review"
	"nubuscu/pretense/ent/tag"
	"nubuscu/pretense/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ReviewUpdate) SetUpdatedAt(t time.Time) *ReviewUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetMetaLabels sets the "meta_labels" field.
func (ru *ReviewUpdate) SetMetaLabels(s []string) *ReviewUpdate {
	ru.mutation.SetMetaLabels(s)
	return ru
}

// AppendMetaLabels appends s to the "meta_labels" field.
func (ru *ReviewUpdate) AppendMetaLabels(s []string) *ReviewUpdate {
	ru.mutation.AppendMetaLabels(s)
	return ru
}

// SetName sets the "name" field.
func (ru *ReviewUpdate) SetName(s string) *ReviewUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetBody sets the "body" field.
func (ru *ReviewUpdate) SetBody(s string) *ReviewUpdate {
	ru.mutation.SetBody(s)
	return ru
}

// AddReviewIDs adds the "reviews" edge to the Topic entity by IDs.
func (ru *ReviewUpdate) AddReviewIDs(ids ...int) *ReviewUpdate {
	ru.mutation.AddReviewIDs(ids...)
	return ru
}

// AddReviews adds the "reviews" edges to the Topic entity.
func (ru *ReviewUpdate) AddReviews(t ...*Topic) *ReviewUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.AddReviewIDs(ids...)
}

// AddTaggedWithIDs adds the "tagged_with" edge to the Tag entity by IDs.
func (ru *ReviewUpdate) AddTaggedWithIDs(ids ...int) *ReviewUpdate {
	ru.mutation.AddTaggedWithIDs(ids...)
	return ru
}

// AddTaggedWith adds the "tagged_with" edges to the Tag entity.
func (ru *ReviewUpdate) AddTaggedWith(t ...*Tag) *ReviewUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.AddTaggedWithIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// ClearReviews clears all "reviews" edges to the Topic entity.
func (ru *ReviewUpdate) ClearReviews() *ReviewUpdate {
	ru.mutation.ClearReviews()
	return ru
}

// RemoveReviewIDs removes the "reviews" edge to Topic entities by IDs.
func (ru *ReviewUpdate) RemoveReviewIDs(ids ...int) *ReviewUpdate {
	ru.mutation.RemoveReviewIDs(ids...)
	return ru
}

// RemoveReviews removes "reviews" edges to Topic entities.
func (ru *ReviewUpdate) RemoveReviews(t ...*Topic) *ReviewUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.RemoveReviewIDs(ids...)
}

// ClearTaggedWith clears all "tagged_with" edges to the Tag entity.
func (ru *ReviewUpdate) ClearTaggedWith() *ReviewUpdate {
	ru.mutation.ClearTaggedWith()
	return ru
}

// RemoveTaggedWithIDs removes the "tagged_with" edge to Tag entities by IDs.
func (ru *ReviewUpdate) RemoveTaggedWithIDs(ids ...int) *ReviewUpdate {
	ru.mutation.RemoveTaggedWithIDs(ids...)
	return ru
}

// RemoveTaggedWith removes "tagged_with" edges to Tag entities.
func (ru *ReviewUpdate) RemoveTaggedWith(t ...*Tag) *ReviewUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.RemoveTaggedWithIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReviewUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := review.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReviewUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := review.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Review.name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Body(); ok {
		if err := review.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Review.body": %w`, err)}
		}
	}
	return nil
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(review.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.MetaLabels(); ok {
		_spec.SetField(review.FieldMetaLabels, field.TypeJSON, value)
	}
	if value, ok := ru.mutation.AppendedMetaLabels(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, review.FieldMetaLabels, value)
		})
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(review.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Body(); ok {
		_spec.SetField(review.FieldBody, field.TypeString, value)
	}
	if ru.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.ReviewsTable,
			Columns: review.ReviewsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ru.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.ReviewsTable,
			Columns: review.ReviewsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.ReviewsTable,
			Columns: review.ReviewsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.TaggedWithCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.TaggedWithTable,
			Columns: review.TaggedWithPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedTaggedWithIDs(); len(nodes) > 0 && !ru.mutation.TaggedWithCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.TaggedWithTable,
			Columns: review.TaggedWithPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TaggedWithIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.TaggedWithTable,
			Columns: review.TaggedWithPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ReviewUpdateOne) SetUpdatedAt(t time.Time) *ReviewUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetMetaLabels sets the "meta_labels" field.
func (ruo *ReviewUpdateOne) SetMetaLabels(s []string) *ReviewUpdateOne {
	ruo.mutation.SetMetaLabels(s)
	return ruo
}

// AppendMetaLabels appends s to the "meta_labels" field.
func (ruo *ReviewUpdateOne) AppendMetaLabels(s []string) *ReviewUpdateOne {
	ruo.mutation.AppendMetaLabels(s)
	return ruo
}

// SetName sets the "name" field.
func (ruo *ReviewUpdateOne) SetName(s string) *ReviewUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetBody sets the "body" field.
func (ruo *ReviewUpdateOne) SetBody(s string) *ReviewUpdateOne {
	ruo.mutation.SetBody(s)
	return ruo
}

// AddReviewIDs adds the "reviews" edge to the Topic entity by IDs.
func (ruo *ReviewUpdateOne) AddReviewIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.AddReviewIDs(ids...)
	return ruo
}

// AddReviews adds the "reviews" edges to the Topic entity.
func (ruo *ReviewUpdateOne) AddReviews(t ...*Topic) *ReviewUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.AddReviewIDs(ids...)
}

// AddTaggedWithIDs adds the "tagged_with" edge to the Tag entity by IDs.
func (ruo *ReviewUpdateOne) AddTaggedWithIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.AddTaggedWithIDs(ids...)
	return ruo
}

// AddTaggedWith adds the "tagged_with" edges to the Tag entity.
func (ruo *ReviewUpdateOne) AddTaggedWith(t ...*Tag) *ReviewUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.AddTaggedWithIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// ClearReviews clears all "reviews" edges to the Topic entity.
func (ruo *ReviewUpdateOne) ClearReviews() *ReviewUpdateOne {
	ruo.mutation.ClearReviews()
	return ruo
}

// RemoveReviewIDs removes the "reviews" edge to Topic entities by IDs.
func (ruo *ReviewUpdateOne) RemoveReviewIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.RemoveReviewIDs(ids...)
	return ruo
}

// RemoveReviews removes "reviews" edges to Topic entities.
func (ruo *ReviewUpdateOne) RemoveReviews(t ...*Topic) *ReviewUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.RemoveReviewIDs(ids...)
}

// ClearTaggedWith clears all "tagged_with" edges to the Tag entity.
func (ruo *ReviewUpdateOne) ClearTaggedWith() *ReviewUpdateOne {
	ruo.mutation.ClearTaggedWith()
	return ruo
}

// RemoveTaggedWithIDs removes the "tagged_with" edge to Tag entities by IDs.
func (ruo *ReviewUpdateOne) RemoveTaggedWithIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.RemoveTaggedWithIDs(ids...)
	return ruo
}

// RemoveTaggedWith removes "tagged_with" edges to Tag entities.
func (ruo *ReviewUpdateOne) RemoveTaggedWith(t ...*Tag) *ReviewUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.RemoveTaggedWithIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Review)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ReviewMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReviewUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := review.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReviewUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := review.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Review.name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Body(); ok {
		if err := review.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Review.body": %w`, err)}
		}
	}
	return nil
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Review.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(review.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.MetaLabels(); ok {
		_spec.SetField(review.FieldMetaLabels, field.TypeJSON, value)
	}
	if value, ok := ruo.mutation.AppendedMetaLabels(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, review.FieldMetaLabels, value)
		})
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(review.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Body(); ok {
		_spec.SetField(review.FieldBody, field.TypeString, value)
	}
	if ruo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.ReviewsTable,
			Columns: review.ReviewsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedReviewsIDs(); len(nodes) > 0 && !ruo.mutation.ReviewsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.ReviewsTable,
			Columns: review.ReviewsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.ReviewsTable,
			Columns: review.ReviewsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.TaggedWithCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.TaggedWithTable,
			Columns: review.TaggedWithPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedTaggedWithIDs(); len(nodes) > 0 && !ruo.mutation.TaggedWithCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.TaggedWithTable,
			Columns: review.TaggedWithPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TaggedWithIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.TaggedWithTable,
			Columns: review.TaggedWithPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
