// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"restent/ent/blog"
	"restent/ent/cate"
	"restent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CateUpdate is the builder for updating Cate entities.
type CateUpdate struct {
	config
	hooks    []Hook
	mutation *CateMutation
}

// Where appends a list predicates to the CateUpdate builder.
func (cu *CateUpdate) Where(ps ...predicate.Cate) *CateUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CateUpdate) SetUpdatedAt(t time.Time) *CateUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cu *CateUpdate) ClearUpdatedAt() *CateUpdate {
	cu.mutation.ClearUpdatedAt()
	return cu
}

// SetName sets the "name" field.
func (cu *CateUpdate) SetName(s string) *CateUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddBlogIDs adds the "blogs" edge to the Blog entity by IDs.
func (cu *CateUpdate) AddBlogIDs(ids ...uuid.UUID) *CateUpdate {
	cu.mutation.AddBlogIDs(ids...)
	return cu
}

// AddBlogs adds the "blogs" edges to the Blog entity.
func (cu *CateUpdate) AddBlogs(b ...*Blog) *CateUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBlogIDs(ids...)
}

// Mutation returns the CateMutation object of the builder.
func (cu *CateUpdate) Mutation() *CateMutation {
	return cu.mutation
}

// ClearBlogs clears all "blogs" edges to the Blog entity.
func (cu *CateUpdate) ClearBlogs() *CateUpdate {
	cu.mutation.ClearBlogs()
	return cu
}

// RemoveBlogIDs removes the "blogs" edge to Blog entities by IDs.
func (cu *CateUpdate) RemoveBlogIDs(ids ...uuid.UUID) *CateUpdate {
	cu.mutation.RemoveBlogIDs(ids...)
	return cu
}

// RemoveBlogs removes "blogs" edges to Blog entities.
func (cu *CateUpdate) RemoveBlogs(b ...*Blog) *CateUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBlogIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CateUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CateUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CateUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CateUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok && !cu.mutation.UpdatedAtCleared() {
		v := cate.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cate.Table,
			Columns: cate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cate.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: cate.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cate.FieldUpdatedAt,
		})
	}
	if cu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: cate.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cate.FieldName,
		})
	}
	if cu.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cate.BlogsTable,
			Columns: cate.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedBlogsIDs(); len(nodes) > 0 && !cu.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cate.BlogsTable,
			Columns: cate.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BlogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cate.BlogsTable,
			Columns: cate.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CateUpdateOne is the builder for updating a single Cate entity.
type CateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CateMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CateUpdateOne) SetUpdatedAt(t time.Time) *CateUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cuo *CateUpdateOne) ClearUpdatedAt() *CateUpdateOne {
	cuo.mutation.ClearUpdatedAt()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CateUpdateOne) SetName(s string) *CateUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddBlogIDs adds the "blogs" edge to the Blog entity by IDs.
func (cuo *CateUpdateOne) AddBlogIDs(ids ...uuid.UUID) *CateUpdateOne {
	cuo.mutation.AddBlogIDs(ids...)
	return cuo
}

// AddBlogs adds the "blogs" edges to the Blog entity.
func (cuo *CateUpdateOne) AddBlogs(b ...*Blog) *CateUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBlogIDs(ids...)
}

// Mutation returns the CateMutation object of the builder.
func (cuo *CateUpdateOne) Mutation() *CateMutation {
	return cuo.mutation
}

// ClearBlogs clears all "blogs" edges to the Blog entity.
func (cuo *CateUpdateOne) ClearBlogs() *CateUpdateOne {
	cuo.mutation.ClearBlogs()
	return cuo
}

// RemoveBlogIDs removes the "blogs" edge to Blog entities by IDs.
func (cuo *CateUpdateOne) RemoveBlogIDs(ids ...uuid.UUID) *CateUpdateOne {
	cuo.mutation.RemoveBlogIDs(ids...)
	return cuo
}

// RemoveBlogs removes "blogs" edges to Blog entities.
func (cuo *CateUpdateOne) RemoveBlogs(b ...*Blog) *CateUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBlogIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CateUpdateOne) Select(field string, fields ...string) *CateUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cate entity.
func (cuo *CateUpdateOne) Save(ctx context.Context) (*Cate, error) {
	var (
		err  error
		node *Cate
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CateUpdateOne) SaveX(ctx context.Context) *Cate {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CateUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CateUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CateUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok && !cuo.mutation.UpdatedAtCleared() {
		v := cate.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CateUpdateOne) sqlSave(ctx context.Context) (_node *Cate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cate.Table,
			Columns: cate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cate.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cate.FieldID)
		for _, f := range fields {
			if !cate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: cate.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cate.FieldUpdatedAt,
		})
	}
	if cuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: cate.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cate.FieldName,
		})
	}
	if cuo.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cate.BlogsTable,
			Columns: cate.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blog.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedBlogsIDs(); len(nodes) > 0 && !cuo.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cate.BlogsTable,
			Columns: cate.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BlogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   cate.BlogsTable,
			Columns: cate.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blog.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cate{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
