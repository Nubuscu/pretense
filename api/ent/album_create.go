// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nubuscu/pretense/ent/album"
	"nubuscu/pretense/ent/artist"
	"nubuscu/pretense/ent/topic"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AlbumCreate is the builder for creating a Album entity.
type AlbumCreate struct {
	config
	mutation *AlbumMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ac *AlbumCreate) SetCreatedAt(t time.Time) *AlbumCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AlbumCreate) SetNillableCreatedAt(t *time.Time) *AlbumCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AlbumCreate) SetUpdatedAt(t time.Time) *AlbumCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AlbumCreate) SetNillableUpdatedAt(t *time.Time) *AlbumCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AlbumCreate) SetName(s string) *AlbumCreate {
	ac.mutation.SetName(s)
	return ac
}

// AddByIDs adds the "by" edge to the Artist entity by IDs.
func (ac *AlbumCreate) AddByIDs(ids ...int) *AlbumCreate {
	ac.mutation.AddByIDs(ids...)
	return ac
}

// AddBy adds the "by" edges to the Artist entity.
func (ac *AlbumCreate) AddBy(a ...*Artist) *AlbumCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddByIDs(ids...)
}

// AddIncludedInIDs adds the "included_in" edge to the Topic entity by IDs.
func (ac *AlbumCreate) AddIncludedInIDs(ids ...int) *AlbumCreate {
	ac.mutation.AddIncludedInIDs(ids...)
	return ac
}

// AddIncludedIn adds the "included_in" edges to the Topic entity.
func (ac *AlbumCreate) AddIncludedIn(t ...*Topic) *AlbumCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddIncludedInIDs(ids...)
}

// Mutation returns the AlbumMutation object of the builder.
func (ac *AlbumCreate) Mutation() *AlbumMutation {
	return ac.mutation
}

// Save creates the Album in the database.
func (ac *AlbumCreate) Save(ctx context.Context) (*Album, error) {
	var (
		err  error
		node *Album
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AlbumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ac *AlbumCreate) SaveX(ctx context.Context) *Album {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AlbumCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AlbumCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AlbumCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := album.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := album.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AlbumCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Album.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Album.updated_at"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Album.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := album.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Album.name": %w`, err)}
		}
	}
	return nil
}

func (ac *AlbumCreate) sqlSave(ctx context.Context) (*Album, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AlbumCreate) createSpec() (*Album, *sqlgraph.CreateSpec) {
	var (
		_node = &Album{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: album.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: album.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(album.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(album.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(album.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := ac.mutation.ByIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.IncludedInIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Album.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AlbumUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ac *AlbumCreate) OnConflict(opts ...sql.ConflictOption) *AlbumUpsertOne {
	ac.conflict = opts
	return &AlbumUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Album.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AlbumCreate) OnConflictColumns(columns ...string) *AlbumUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AlbumUpsertOne{
		create: ac,
	}
}

type (
	// AlbumUpsertOne is the builder for "upsert"-ing
	//  one Album node.
	AlbumUpsertOne struct {
		create *AlbumCreate
	}

	// AlbumUpsert is the "OnConflict" setter.
	AlbumUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AlbumUpsert) SetUpdatedAt(v time.Time) *AlbumUpsert {
	u.Set(album.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AlbumUpsert) UpdateUpdatedAt() *AlbumUpsert {
	u.SetExcluded(album.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *AlbumUpsert) SetName(v string) *AlbumUpsert {
	u.Set(album.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AlbumUpsert) UpdateName() *AlbumUpsert {
	u.SetExcluded(album.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Album.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AlbumUpsertOne) UpdateNewValues() *AlbumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(album.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Album.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AlbumUpsertOne) Ignore() *AlbumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AlbumUpsertOne) DoNothing() *AlbumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AlbumCreate.OnConflict
// documentation for more info.
func (u *AlbumUpsertOne) Update(set func(*AlbumUpsert)) *AlbumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AlbumUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AlbumUpsertOne) SetUpdatedAt(v time.Time) *AlbumUpsertOne {
	return u.Update(func(s *AlbumUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AlbumUpsertOne) UpdateUpdatedAt() *AlbumUpsertOne {
	return u.Update(func(s *AlbumUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AlbumUpsertOne) SetName(v string) *AlbumUpsertOne {
	return u.Update(func(s *AlbumUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AlbumUpsertOne) UpdateName() *AlbumUpsertOne {
	return u.Update(func(s *AlbumUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *AlbumUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AlbumCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AlbumUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AlbumUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AlbumUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AlbumCreateBulk is the builder for creating many Album entities in bulk.
type AlbumCreateBulk struct {
	config
	builders []*AlbumCreate
	conflict []sql.ConflictOption
}

// Save creates the Album entities in the database.
func (acb *AlbumCreateBulk) Save(ctx context.Context) ([]*Album, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Album, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AlbumMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AlbumCreateBulk) SaveX(ctx context.Context) []*Album {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AlbumCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AlbumCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Album.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AlbumUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (acb *AlbumCreateBulk) OnConflict(opts ...sql.ConflictOption) *AlbumUpsertBulk {
	acb.conflict = opts
	return &AlbumUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Album.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AlbumCreateBulk) OnConflictColumns(columns ...string) *AlbumUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AlbumUpsertBulk{
		create: acb,
	}
}

// AlbumUpsertBulk is the builder for "upsert"-ing
// a bulk of Album nodes.
type AlbumUpsertBulk struct {
	create *AlbumCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Album.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AlbumUpsertBulk) UpdateNewValues() *AlbumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(album.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Album.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AlbumUpsertBulk) Ignore() *AlbumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AlbumUpsertBulk) DoNothing() *AlbumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AlbumCreateBulk.OnConflict
// documentation for more info.
func (u *AlbumUpsertBulk) Update(set func(*AlbumUpsert)) *AlbumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AlbumUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AlbumUpsertBulk) SetUpdatedAt(v time.Time) *AlbumUpsertBulk {
	return u.Update(func(s *AlbumUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AlbumUpsertBulk) UpdateUpdatedAt() *AlbumUpsertBulk {
	return u.Update(func(s *AlbumUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AlbumUpsertBulk) SetName(v string) *AlbumUpsertBulk {
	return u.Update(func(s *AlbumUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AlbumUpsertBulk) UpdateName() *AlbumUpsertBulk {
	return u.Update(func(s *AlbumUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *AlbumUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AlbumCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AlbumCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AlbumUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
