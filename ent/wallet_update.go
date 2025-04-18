// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gojeksrepo/ent/predicate"
	"gojeksrepo/ent/user"
	"gojeksrepo/ent/wallet"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WalletUpdate is the builder for updating Wallet entities.
type WalletUpdate struct {
	config
	hooks    []Hook
	mutation *WalletMutation
}

// Where appends a list predicates to the WalletUpdate builder.
func (wu *WalletUpdate) Where(ps ...predicate.Wallet) *WalletUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetUserID sets the "user_id" field.
func (wu *WalletUpdate) SetUserID(u uuid.UUID) *WalletUpdate {
	wu.mutation.SetUserID(u)
	return wu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableUserID(u *uuid.UUID) *WalletUpdate {
	if u != nil {
		wu.SetUserID(*u)
	}
	return wu
}

// SetBalance sets the "balance" field.
func (wu *WalletUpdate) SetBalance(f float64) *WalletUpdate {
	wu.mutation.ResetBalance()
	wu.mutation.SetBalance(f)
	return wu
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableBalance(f *float64) *WalletUpdate {
	if f != nil {
		wu.SetBalance(*f)
	}
	return wu
}

// AddBalance adds f to the "balance" field.
func (wu *WalletUpdate) AddBalance(f float64) *WalletUpdate {
	wu.mutation.AddBalance(f)
	return wu
}

// SetUser sets the "user" edge to the User entity.
func (wu *WalletUpdate) SetUser(u *User) *WalletUpdate {
	return wu.SetUserID(u.ID)
}

// Mutation returns the WalletMutation object of the builder.
func (wu *WalletUpdate) Mutation() *WalletMutation {
	return wu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wu *WalletUpdate) ClearUser() *WalletUpdate {
	wu.mutation.ClearUser()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WalletUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WalletUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WalletUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WalletUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WalletUpdate) check() error {
	if wu.mutation.UserCleared() && len(wu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Wallet.user"`)
	}
	return nil
}

func (wu *WalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(wallet.Table, wallet.Columns, sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeUUID))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Balance(); ok {
		_spec.SetField(wallet.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := wu.mutation.AddedBalance(); ok {
		_spec.AddField(wallet.FieldBalance, field.TypeFloat64, value)
	}
	if wu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.UserTable,
			Columns: []string{wallet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.UserTable,
			Columns: []string{wallet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WalletUpdateOne is the builder for updating a single Wallet entity.
type WalletUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WalletMutation
}

// SetUserID sets the "user_id" field.
func (wuo *WalletUpdateOne) SetUserID(u uuid.UUID) *WalletUpdateOne {
	wuo.mutation.SetUserID(u)
	return wuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableUserID(u *uuid.UUID) *WalletUpdateOne {
	if u != nil {
		wuo.SetUserID(*u)
	}
	return wuo
}

// SetBalance sets the "balance" field.
func (wuo *WalletUpdateOne) SetBalance(f float64) *WalletUpdateOne {
	wuo.mutation.ResetBalance()
	wuo.mutation.SetBalance(f)
	return wuo
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableBalance(f *float64) *WalletUpdateOne {
	if f != nil {
		wuo.SetBalance(*f)
	}
	return wuo
}

// AddBalance adds f to the "balance" field.
func (wuo *WalletUpdateOne) AddBalance(f float64) *WalletUpdateOne {
	wuo.mutation.AddBalance(f)
	return wuo
}

// SetUser sets the "user" edge to the User entity.
func (wuo *WalletUpdateOne) SetUser(u *User) *WalletUpdateOne {
	return wuo.SetUserID(u.ID)
}

// Mutation returns the WalletMutation object of the builder.
func (wuo *WalletUpdateOne) Mutation() *WalletMutation {
	return wuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wuo *WalletUpdateOne) ClearUser() *WalletUpdateOne {
	wuo.mutation.ClearUser()
	return wuo
}

// Where appends a list predicates to the WalletUpdate builder.
func (wuo *WalletUpdateOne) Where(ps ...predicate.Wallet) *WalletUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WalletUpdateOne) Select(field string, fields ...string) *WalletUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wallet entity.
func (wuo *WalletUpdateOne) Save(ctx context.Context) (*Wallet, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WalletUpdateOne) SaveX(ctx context.Context) *Wallet {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WalletUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WalletUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WalletUpdateOne) check() error {
	if wuo.mutation.UserCleared() && len(wuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Wallet.user"`)
	}
	return nil
}

func (wuo *WalletUpdateOne) sqlSave(ctx context.Context) (_node *Wallet, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(wallet.Table, wallet.Columns, sqlgraph.NewFieldSpec(wallet.FieldID, field.TypeUUID))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wallet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wallet.FieldID)
		for _, f := range fields {
			if !wallet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Balance(); ok {
		_spec.SetField(wallet.FieldBalance, field.TypeFloat64, value)
	}
	if value, ok := wuo.mutation.AddedBalance(); ok {
		_spec.AddField(wallet.FieldBalance, field.TypeFloat64, value)
	}
	if wuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.UserTable,
			Columns: []string{wallet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wallet.UserTable,
			Columns: []string{wallet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wallet{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
