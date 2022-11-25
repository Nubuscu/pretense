// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/predicate"
	"nubuscu/pretense/ent/review"
	"nubuscu/pretense/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// Where appends a list predicates to the TopicUpdate builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TopicUpdate) SetUpdatedAt(t time.Time) *TopicUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetName sets the "name" field.
func (tu *TopicUpdate) SetName(s string) *TopicUpdate {
	tu.mutation.SetName(s)
	return tu
}

// AddReviewedByIDs adds the "reviewed_by" edge to the Review entity by IDs.
func (tu *TopicUpdate) AddReviewedByIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddReviewedByIDs(ids...)
	return tu
}

// AddReviewedBy adds the "reviewed_by" edges to the Review entity.
func (tu *TopicUpdate) AddReviewedBy(r ...*Review) *TopicUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddReviewedByIDs(ids...)
}

// AddIncludeIDs adds the "includes" edge to the Album entity by IDs.
func (tu *TopicUpdate) AddIncludeIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddIncludeIDs(ids...)
	return tu
}

// AddIncludes adds the "includes" edges to the Album entity.
func (tu *TopicUpdate) AddIncludes(a ...*Album) *TopicUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.AddIncludeIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// ClearReviewedBy clears all "reviewed_by" edges to the Review entity.
func (tu *TopicUpdate) ClearReviewedBy() *TopicUpdate {
	tu.mutation.ClearReviewedBy()
	return tu
}

// RemoveReviewedByIDs removes the "reviewed_by" edge to Review entities by IDs.
func (tu *TopicUpdate) RemoveReviewedByIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveReviewedByIDs(ids...)
	return tu
}

// RemoveReviewedBy removes "reviewed_by" edges to Review entities.
func (tu *TopicUpdate) RemoveReviewedBy(r ...*Review) *TopicUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveReviewedByIDs(ids...)
}

// ClearIncludes clears all "includes" edges to the Album entity.
func (tu *TopicUpdate) ClearIncludes() *TopicUpdate {
	tu.mutation.ClearIncludes()
	return tu
}

// RemoveIncludeIDs removes the "includes" edge to Album entities by IDs.
func (tu *TopicUpdate) RemoveIncludeIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveIncludeIDs(ids...)
	return tu
}

// RemoveIncludes removes "includes" edges to Album entities.
func (tu *TopicUpdate) RemoveIncludes(a ...*Album) *TopicUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.RemoveIncludeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TopicUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := topic.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(topic.FieldName, field.TypeString, value)
	}
	if tu.mutation.ReviewedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   topic.ReviewedByTable,
			Columns: topic.ReviewedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReviewedByIDs(); len(nodes) > 0 && !tu.mutation.ReviewedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   topic.ReviewedByTable,
			Columns: topic.ReviewedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReviewedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   topic.ReviewedByTable,
			Columns: topic.ReviewedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.IncludesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.IncludesTable,
			Columns: topic.IncludesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedIncludesIDs(); len(nodes) > 0 && !tu.mutation.IncludesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.IncludesTable,
			Columns: topic.IncludesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.IncludesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.IncludesTable,
			Columns: topic.IncludesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TopicMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TopicUpdateOne) SetUpdatedAt(t time.Time) *TopicUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TopicUpdateOne) SetName(s string) *TopicUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// AddReviewedByIDs adds the "reviewed_by" edge to the Review entity by IDs.
func (tuo *TopicUpdateOne) AddReviewedByIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddReviewedByIDs(ids...)
	return tuo
}

// AddReviewedBy adds the "reviewed_by" edges to the Review entity.
func (tuo *TopicUpdateOne) AddReviewedBy(r ...*Review) *TopicUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddReviewedByIDs(ids...)
}

// AddIncludeIDs adds the "includes" edge to the Album entity by IDs.
func (tuo *TopicUpdateOne) AddIncludeIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddIncludeIDs(ids...)
	return tuo
}

// AddIncludes adds the "includes" edges to the Album entity.
func (tuo *TopicUpdateOne) AddIncludes(a ...*Album) *TopicUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.AddIncludeIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// ClearReviewedBy clears all "reviewed_by" edges to the Review entity.
func (tuo *TopicUpdateOne) ClearReviewedBy() *TopicUpdateOne {
	tuo.mutation.ClearReviewedBy()
	return tuo
}

// RemoveReviewedByIDs removes the "reviewed_by" edge to Review entities by IDs.
func (tuo *TopicUpdateOne) RemoveReviewedByIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveReviewedByIDs(ids...)
	return tuo
}

// RemoveReviewedBy removes "reviewed_by" edges to Review entities.
func (tuo *TopicUpdateOne) RemoveReviewedBy(r ...*Review) *TopicUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveReviewedByIDs(ids...)
}

// ClearIncludes clears all "includes" edges to the Album entity.
func (tuo *TopicUpdateOne) ClearIncludes() *TopicUpdateOne {
	tuo.mutation.ClearIncludes()
	return tuo
}

// RemoveIncludeIDs removes the "includes" edge to Album entities by IDs.
func (tuo *TopicUpdateOne) RemoveIncludeIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveIncludeIDs(ids...)
	return tuo
}

// RemoveIncludes removes "includes" edges to Album entities.
func (tuo *TopicUpdateOne) RemoveIncludes(a ...*Album) *TopicUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.RemoveIncludeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TopicUpdateOne) Select(field string, fields ...string) *TopicUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Topic entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	var (
		err  error
		node *Topic
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Topic)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TopicMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TopicUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := topic.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (_node *Topic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Topic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topic.FieldID)
		for _, f := range fields {
			if !topic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(topic.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(topic.FieldName, field.TypeString, value)
	}
	if tuo.mutation.ReviewedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   topic.ReviewedByTable,
			Columns: topic.ReviewedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReviewedByIDs(); len(nodes) > 0 && !tuo.mutation.ReviewedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   topic.ReviewedByTable,
			Columns: topic.ReviewedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReviewedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   topic.ReviewedByTable,
			Columns: topic.ReviewedByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.IncludesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.IncludesTable,
			Columns: topic.IncludesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedIncludesIDs(); len(nodes) > 0 && !tuo.mutation.IncludesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.IncludesTable,
			Columns: topic.IncludesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.IncludesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.IncludesTable,
			Columns: topic.IncludesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: album.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Topic{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
