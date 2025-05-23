// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gojeksrepo/ent/predicate"
	"gojeksrepo/ent/usersadmin"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersAdminUpdate is the builder for updating UsersAdmin entities.
type UsersAdminUpdate struct {
	config
	hooks    []Hook
	mutation *UsersAdminMutation
}

// Where appends a list predicates to the UsersAdminUpdate builder.
func (uau *UsersAdminUpdate) Where(ps ...predicate.UsersAdmin) *UsersAdminUpdate {
	uau.mutation.Where(ps...)
	return uau
}

// SetName sets the "name" field.
func (uau *UsersAdminUpdate) SetName(s string) *UsersAdminUpdate {
	uau.mutation.SetName(s)
	return uau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uau *UsersAdminUpdate) SetNillableName(s *string) *UsersAdminUpdate {
	if s != nil {
		uau.SetName(*s)
	}
	return uau
}

// SetUsername sets the "username" field.
func (uau *UsersAdminUpdate) SetUsername(s string) *UsersAdminUpdate {
	uau.mutation.SetUsername(s)
	return uau
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uau *UsersAdminUpdate) SetNillableUsername(s *string) *UsersAdminUpdate {
	if s != nil {
		uau.SetUsername(*s)
	}
	return uau
}

// SetPassword sets the "password" field.
func (uau *UsersAdminUpdate) SetPassword(s string) *UsersAdminUpdate {
	uau.mutation.SetPassword(s)
	return uau
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uau *UsersAdminUpdate) SetNillablePassword(s *string) *UsersAdminUpdate {
	if s != nil {
		uau.SetPassword(*s)
	}
	return uau
}

// SetStatusAdmin sets the "status_admin" field.
func (uau *UsersAdminUpdate) SetStatusAdmin(i int) *UsersAdminUpdate {
	uau.mutation.ResetStatusAdmin()
	uau.mutation.SetStatusAdmin(i)
	return uau
}

// SetNillableStatusAdmin sets the "status_admin" field if the given value is not nil.
func (uau *UsersAdminUpdate) SetNillableStatusAdmin(i *int) *UsersAdminUpdate {
	if i != nil {
		uau.SetStatusAdmin(*i)
	}
	return uau
}

// AddStatusAdmin adds i to the "status_admin" field.
func (uau *UsersAdminUpdate) AddStatusAdmin(i int) *UsersAdminUpdate {
	uau.mutation.AddStatusAdmin(i)
	return uau
}

// Mutation returns the UsersAdminMutation object of the builder.
func (uau *UsersAdminUpdate) Mutation() *UsersAdminMutation {
	return uau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uau *UsersAdminUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uau.sqlSave, uau.mutation, uau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uau *UsersAdminUpdate) SaveX(ctx context.Context) int {
	affected, err := uau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uau *UsersAdminUpdate) Exec(ctx context.Context) error {
	_, err := uau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uau *UsersAdminUpdate) ExecX(ctx context.Context) {
	if err := uau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uau *UsersAdminUpdate) check() error {
	if v, ok := uau.mutation.Name(); ok {
		if err := usersadmin.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UsersAdmin.name": %w`, err)}
		}
	}
	if v, ok := uau.mutation.Username(); ok {
		if err := usersadmin.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "UsersAdmin.username": %w`, err)}
		}
	}
	if v, ok := uau.mutation.Password(); ok {
		if err := usersadmin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "UsersAdmin.password": %w`, err)}
		}
	}
	return nil
}

func (uau *UsersAdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersadmin.Table, usersadmin.Columns, sqlgraph.NewFieldSpec(usersadmin.FieldID, field.TypeUUID))
	if ps := uau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uau.mutation.Name(); ok {
		_spec.SetField(usersadmin.FieldName, field.TypeString, value)
	}
	if value, ok := uau.mutation.Username(); ok {
		_spec.SetField(usersadmin.FieldUsername, field.TypeString, value)
	}
	if value, ok := uau.mutation.Password(); ok {
		_spec.SetField(usersadmin.FieldPassword, field.TypeString, value)
	}
	if value, ok := uau.mutation.StatusAdmin(); ok {
		_spec.SetField(usersadmin.FieldStatusAdmin, field.TypeInt, value)
	}
	if value, ok := uau.mutation.AddedStatusAdmin(); ok {
		_spec.AddField(usersadmin.FieldStatusAdmin, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uau.mutation.done = true
	return n, nil
}

// UsersAdminUpdateOne is the builder for updating a single UsersAdmin entity.
type UsersAdminUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersAdminMutation
}

// SetName sets the "name" field.
func (uauo *UsersAdminUpdateOne) SetName(s string) *UsersAdminUpdateOne {
	uauo.mutation.SetName(s)
	return uauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uauo *UsersAdminUpdateOne) SetNillableName(s *string) *UsersAdminUpdateOne {
	if s != nil {
		uauo.SetName(*s)
	}
	return uauo
}

// SetUsername sets the "username" field.
func (uauo *UsersAdminUpdateOne) SetUsername(s string) *UsersAdminUpdateOne {
	uauo.mutation.SetUsername(s)
	return uauo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uauo *UsersAdminUpdateOne) SetNillableUsername(s *string) *UsersAdminUpdateOne {
	if s != nil {
		uauo.SetUsername(*s)
	}
	return uauo
}

// SetPassword sets the "password" field.
func (uauo *UsersAdminUpdateOne) SetPassword(s string) *UsersAdminUpdateOne {
	uauo.mutation.SetPassword(s)
	return uauo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uauo *UsersAdminUpdateOne) SetNillablePassword(s *string) *UsersAdminUpdateOne {
	if s != nil {
		uauo.SetPassword(*s)
	}
	return uauo
}

// SetStatusAdmin sets the "status_admin" field.
func (uauo *UsersAdminUpdateOne) SetStatusAdmin(i int) *UsersAdminUpdateOne {
	uauo.mutation.ResetStatusAdmin()
	uauo.mutation.SetStatusAdmin(i)
	return uauo
}

// SetNillableStatusAdmin sets the "status_admin" field if the given value is not nil.
func (uauo *UsersAdminUpdateOne) SetNillableStatusAdmin(i *int) *UsersAdminUpdateOne {
	if i != nil {
		uauo.SetStatusAdmin(*i)
	}
	return uauo
}

// AddStatusAdmin adds i to the "status_admin" field.
func (uauo *UsersAdminUpdateOne) AddStatusAdmin(i int) *UsersAdminUpdateOne {
	uauo.mutation.AddStatusAdmin(i)
	return uauo
}

// Mutation returns the UsersAdminMutation object of the builder.
func (uauo *UsersAdminUpdateOne) Mutation() *UsersAdminMutation {
	return uauo.mutation
}

// Where appends a list predicates to the UsersAdminUpdate builder.
func (uauo *UsersAdminUpdateOne) Where(ps ...predicate.UsersAdmin) *UsersAdminUpdateOne {
	uauo.mutation.Where(ps...)
	return uauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uauo *UsersAdminUpdateOne) Select(field string, fields ...string) *UsersAdminUpdateOne {
	uauo.fields = append([]string{field}, fields...)
	return uauo
}

// Save executes the query and returns the updated UsersAdmin entity.
func (uauo *UsersAdminUpdateOne) Save(ctx context.Context) (*UsersAdmin, error) {
	return withHooks(ctx, uauo.sqlSave, uauo.mutation, uauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uauo *UsersAdminUpdateOne) SaveX(ctx context.Context) *UsersAdmin {
	node, err := uauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uauo *UsersAdminUpdateOne) Exec(ctx context.Context) error {
	_, err := uauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uauo *UsersAdminUpdateOne) ExecX(ctx context.Context) {
	if err := uauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uauo *UsersAdminUpdateOne) check() error {
	if v, ok := uauo.mutation.Name(); ok {
		if err := usersadmin.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "UsersAdmin.name": %w`, err)}
		}
	}
	if v, ok := uauo.mutation.Username(); ok {
		if err := usersadmin.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "UsersAdmin.username": %w`, err)}
		}
	}
	if v, ok := uauo.mutation.Password(); ok {
		if err := usersadmin.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "UsersAdmin.password": %w`, err)}
		}
	}
	return nil
}

func (uauo *UsersAdminUpdateOne) sqlSave(ctx context.Context) (_node *UsersAdmin, err error) {
	if err := uauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersadmin.Table, usersadmin.Columns, sqlgraph.NewFieldSpec(usersadmin.FieldID, field.TypeUUID))
	id, ok := uauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UsersAdmin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersadmin.FieldID)
		for _, f := range fields {
			if !usersadmin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usersadmin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uauo.mutation.Name(); ok {
		_spec.SetField(usersadmin.FieldName, field.TypeString, value)
	}
	if value, ok := uauo.mutation.Username(); ok {
		_spec.SetField(usersadmin.FieldUsername, field.TypeString, value)
	}
	if value, ok := uauo.mutation.Password(); ok {
		_spec.SetField(usersadmin.FieldPassword, field.TypeString, value)
	}
	if value, ok := uauo.mutation.StatusAdmin(); ok {
		_spec.SetField(usersadmin.FieldStatusAdmin, field.TypeInt, value)
	}
	if value, ok := uauo.mutation.AddedStatusAdmin(); ok {
		_spec.AddField(usersadmin.FieldStatusAdmin, field.TypeInt, value)
	}
	_node = &UsersAdmin{config: uauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uauo.mutation.done = true
	return _node, nil
}
