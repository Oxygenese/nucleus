// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/basemodel"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/predicate"
)

// BaseModelDelete is the builder for deleting a BaseModel entity.
type BaseModelDelete struct {
	config
	hooks    []Hook
	mutation *BaseModelMutation
}

// Where appends a list predicates to the BaseModelDelete builder.
func (bmd *BaseModelDelete) Where(ps ...predicate.BaseModel) *BaseModelDelete {
	bmd.mutation.Where(ps...)
	return bmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bmd *BaseModelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bmd.hooks) == 0 {
		affected, err = bmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BaseModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bmd.mutation = mutation
			affected, err = bmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bmd.hooks) - 1; i >= 0; i-- {
			if bmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bmd *BaseModelDelete) ExecX(ctx context.Context) int {
	n, err := bmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bmd *BaseModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: basemodel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: basemodel.FieldID,
			},
		},
	}
	if ps := bmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// BaseModelDeleteOne is the builder for deleting a single BaseModel entity.
type BaseModelDeleteOne struct {
	bmd *BaseModelDelete
}

// Exec executes the deletion query.
func (bmdo *BaseModelDeleteOne) Exec(ctx context.Context) error {
	n, err := bmdo.bmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{basemodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bmdo *BaseModelDeleteOne) ExecX(ctx context.Context) {
	bmdo.bmd.ExecX(ctx)
}
