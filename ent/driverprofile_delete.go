// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"gojeksrepo/ent/driverprofile"
	"gojeksrepo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DriverProfileDelete is the builder for deleting a DriverProfile entity.
type DriverProfileDelete struct {
	config
	hooks    []Hook
	mutation *DriverProfileMutation
}

// Where appends a list predicates to the DriverProfileDelete builder.
func (dpd *DriverProfileDelete) Where(ps ...predicate.DriverProfile) *DriverProfileDelete {
	dpd.mutation.Where(ps...)
	return dpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dpd *DriverProfileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dpd.sqlExec, dpd.mutation, dpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dpd *DriverProfileDelete) ExecX(ctx context.Context) int {
	n, err := dpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dpd *DriverProfileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(driverprofile.Table, sqlgraph.NewFieldSpec(driverprofile.FieldID, field.TypeUUID))
	if ps := dpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dpd.mutation.done = true
	return affected, err
}

// DriverProfileDeleteOne is the builder for deleting a single DriverProfile entity.
type DriverProfileDeleteOne struct {
	dpd *DriverProfileDelete
}

// Where appends a list predicates to the DriverProfileDelete builder.
func (dpdo *DriverProfileDeleteOne) Where(ps ...predicate.DriverProfile) *DriverProfileDeleteOne {
	dpdo.dpd.mutation.Where(ps...)
	return dpdo
}

// Exec executes the deletion query.
func (dpdo *DriverProfileDeleteOne) Exec(ctx context.Context) error {
	n, err := dpdo.dpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{driverprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dpdo *DriverProfileDeleteOne) ExecX(ctx context.Context) {
	if err := dpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
