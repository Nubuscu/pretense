// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"nubuscu/pretense/ent/predicate"
	"nubuscu/pretense/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AlbumUpdate is the builder for updating Album entities.
type AlbumUpdate struct {
	config
	hooks    []Hook
	mutation *AlbumMutation
}

// Where appends a list predicates to the AlbumUpdate builder.
func (au *AlbumUpdate) Where(ps ...predicate.Album) *AlbumUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AlbumUpdate) SetUpdatedAt(t time.Time) *AlbumUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetName sets the "name" field.
func (au *AlbumUpdate) SetName(s string) *AlbumUpdate {
	au.mutation.SetName(s)
	return au
}

// AddByIDs adds the "by" edge to the Artist entity by IDs.
func (au *AlbumUpdate) AddByIDs(ids ...int) *AlbumUpdate {
	au.mutation.AddByIDs(ids...)
	return au
}

// AddBy adds the "by" edges to the Artist entity.
func (au *AlbumUpdate) AddBy(a ...*Artist) *AlbumUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddByIDs(ids...)
}

// AddIncludedInIDs adds the "included_in" edge to the Topic entity by IDs.
func (au *AlbumUpdate) AddIncludedInIDs(ids ...int) *AlbumUpdate {
	au.mutation.AddIncludedInIDs(ids...)
	return au
}

// AddIncludedIn adds the "included_in" edges to the Topic entity.
func (au *AlbumUpdate) AddIncludedIn(t ...*Topic) *AlbumUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddIncludedInIDs(ids...)
}

// Mutation returns the AlbumMutation object of the builder.
func (au *AlbumUpdate) Mutation() *AlbumMutation {
	return au.mutation
}

// ClearBy clears all "by" edges to the Artist entity.
func (au *AlbumUpdate) ClearBy() *AlbumUpdate {
	au.mutation.ClearBy()
	return au
}

// RemoveByIDs removes the "by" edge to Artist entities by IDs.
func (au *AlbumUpdate) RemoveByIDs(ids ...int) *AlbumUpdate {
	au.mutation.RemoveByIDs(ids...)
	return au
}

// RemoveBy removes "by" edges to Artist entities.
func (au *AlbumUpdate) RemoveBy(a ...*Artist) *AlbumUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveByIDs(ids...)
}

// ClearIncludedIn clears all "included_in" edges to the Topic entity.
func (au *AlbumUpdate) ClearIncludedIn() *AlbumUpdate {
	au.mutation.ClearIncludedIn()
	return au
}

// RemoveIncludedInIDs removes the "included_in" edge to Topic entities by IDs.
func (au *AlbumUpdate) RemoveIncludedInIDs(ids ...int) *AlbumUpdate {
	au.mutation.RemoveIncludedInIDs(ids...)
	return au
}

// RemoveIncludedIn removes "included_in" edges to Topic entities.
func (au *AlbumUpdate) RemoveIncludedIn(t ...*Topic) *AlbumUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveIncludedInIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AlbumUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	au.defaults()
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AlbumUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AlbumUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AlbumUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AlbumUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := album.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AlbumUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := album.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Album.name": %w`, err)}
		}
	}
	return nil
}

func (au *AlbumUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   album.Table,
			Columns: album.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: album.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(album.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(album.FieldName, field.TypeString, value)
	}
	if au.mutation.ByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ByTable,
			Columns: album.ByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedByIDs(); len(nodes) > 0 && !au.mutation.ByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ByTable,
			Columns: album.ByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ByTable,
			Columns: album.ByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.IncludedInCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.IncludedInTable,
			Columns: album.IncludedInPrimaryKey,
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
	if nodes := au.mutation.RemovedIncludedInIDs(); len(nodes) > 0 && !au.mutation.IncludedInCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.IncludedInTable,
			Columns: album.IncludedInPrimaryKey,
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
	if nodes := au.mutation.IncludedInIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.IncludedInTable,
			Columns: album.IncludedInPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{album.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AlbumUpdateOne is the builder for updating a single Album entity.
type AlbumUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AlbumMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AlbumUpdateOne) SetUpdatedAt(t time.Time) *AlbumUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetName sets the "name" field.
func (auo *AlbumUpdateOne) SetName(s string) *AlbumUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// AddByIDs adds the "by" edge to the Artist entity by IDs.
func (auo *AlbumUpdateOne) AddByIDs(ids ...int) *AlbumUpdateOne {
	auo.mutation.AddByIDs(ids...)
	return auo
}

// AddBy adds the "by" edges to the Artist entity.
func (auo *AlbumUpdateOne) AddBy(a ...*Artist) *AlbumUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddByIDs(ids...)
}

// AddIncludedInIDs adds the "included_in" edge to the Topic entity by IDs.
func (auo *AlbumUpdateOne) AddIncludedInIDs(ids ...int) *AlbumUpdateOne {
	auo.mutation.AddIncludedInIDs(ids...)
	return auo
}

// AddIncludedIn adds the "included_in" edges to the Topic entity.
func (auo *AlbumUpdateOne) AddIncludedIn(t ...*Topic) *AlbumUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddIncludedInIDs(ids...)
}

// Mutation returns the AlbumMutation object of the builder.
func (auo *AlbumUpdateOne) Mutation() *AlbumMutation {
	return auo.mutation
}

// ClearBy clears all "by" edges to the Artist entity.
func (auo *AlbumUpdateOne) ClearBy() *AlbumUpdateOne {
	auo.mutation.ClearBy()
	return auo
}

// RemoveByIDs removes the "by" edge to Artist entities by IDs.
func (auo *AlbumUpdateOne) RemoveByIDs(ids ...int) *AlbumUpdateOne {
	auo.mutation.RemoveByIDs(ids...)
	return auo
}

// RemoveBy removes "by" edges to Artist entities.
func (auo *AlbumUpdateOne) RemoveBy(a ...*Artist) *AlbumUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveByIDs(ids...)
}

// ClearIncludedIn clears all "included_in" edges to the Topic entity.
func (auo *AlbumUpdateOne) ClearIncludedIn() *AlbumUpdateOne {
	auo.mutation.ClearIncludedIn()
	return auo
}

// RemoveIncludedInIDs removes the "included_in" edge to Topic entities by IDs.
func (auo *AlbumUpdateOne) RemoveIncludedInIDs(ids ...int) *AlbumUpdateOne {
	auo.mutation.RemoveIncludedInIDs(ids...)
	return auo
}

// RemoveIncludedIn removes "included_in" edges to Topic entities.
func (auo *AlbumUpdateOne) RemoveIncludedIn(t ...*Topic) *AlbumUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveIncludedInIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AlbumUpdateOne) Select(field string, fields ...string) *AlbumUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Album entity.
func (auo *AlbumUpdateOne) Save(ctx context.Context) (*Album, error) {
	var (
		err  error
		node *Album
	)
	auo.defaults()
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Album)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AlbumMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AlbumUpdateOne) SaveX(ctx context.Context) *Album {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AlbumUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AlbumUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AlbumUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := album.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AlbumUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := album.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Album.name": %w`, err)}
		}
	}
	return nil
}

func (auo *AlbumUpdateOne) sqlSave(ctx context.Context) (_node *Album, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   album.Table,
			Columns: album.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: album.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Album.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, album.FieldID)
		for _, f := range fields {
			if !album.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != album.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(album.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(album.FieldName, field.TypeString, value)
	}
	if auo.mutation.ByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ByTable,
			Columns: album.ByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedByIDs(); len(nodes) > 0 && !auo.mutation.ByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ByTable,
			Columns: album.ByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.ByTable,
			Columns: album.ByPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.IncludedInCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.IncludedInTable,
			Columns: album.IncludedInPrimaryKey,
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
	if nodes := auo.mutation.RemovedIncludedInIDs(); len(nodes) > 0 && !auo.mutation.IncludedInCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.IncludedInTable,
			Columns: album.IncludedInPrimaryKey,
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
	if nodes := auo.mutation.IncludedInIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   album.IncludedInTable,
			Columns: album.IncludedInPrimaryKey,
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
	_node = &Album{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{album.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}