// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gojeksrepo/ent/trip"
	"gojeksrepo/ent/triprating"
	"gojeksrepo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TripRatingCreate is the builder for creating a TripRating entity.
type TripRatingCreate struct {
	config
	mutation *TripRatingMutation
	hooks    []Hook
}

// SetTripID sets the "trip_id" field.
func (trc *TripRatingCreate) SetTripID(u uuid.UUID) *TripRatingCreate {
	trc.mutation.SetTripID(u)
	return trc
}

// SetFromUserID sets the "from_user_id" field.
func (trc *TripRatingCreate) SetFromUserID(u uuid.UUID) *TripRatingCreate {
	trc.mutation.SetFromUserID(u)
	return trc
}

// SetToUserID sets the "to_user_id" field.
func (trc *TripRatingCreate) SetToUserID(u uuid.UUID) *TripRatingCreate {
	trc.mutation.SetToUserID(u)
	return trc
}

// SetRating sets the "rating" field.
func (trc *TripRatingCreate) SetRating(i int) *TripRatingCreate {
	trc.mutation.SetRating(i)
	return trc
}

// SetComment sets the "comment" field.
func (trc *TripRatingCreate) SetComment(s string) *TripRatingCreate {
	trc.mutation.SetComment(s)
	return trc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (trc *TripRatingCreate) SetNillableComment(s *string) *TripRatingCreate {
	if s != nil {
		trc.SetComment(*s)
	}
	return trc
}

// SetCreatedAt sets the "created_at" field.
func (trc *TripRatingCreate) SetCreatedAt(t time.Time) *TripRatingCreate {
	trc.mutation.SetCreatedAt(t)
	return trc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (trc *TripRatingCreate) SetNillableCreatedAt(t *time.Time) *TripRatingCreate {
	if t != nil {
		trc.SetCreatedAt(*t)
	}
	return trc
}

// SetID sets the "id" field.
func (trc *TripRatingCreate) SetID(u uuid.UUID) *TripRatingCreate {
	trc.mutation.SetID(u)
	return trc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (trc *TripRatingCreate) SetNillableID(u *uuid.UUID) *TripRatingCreate {
	if u != nil {
		trc.SetID(*u)
	}
	return trc
}

// SetTrip sets the "trip" edge to the Trip entity.
func (trc *TripRatingCreate) SetTrip(t *Trip) *TripRatingCreate {
	return trc.SetTripID(t.ID)
}

// SetFromUser sets the "from_user" edge to the User entity.
func (trc *TripRatingCreate) SetFromUser(u *User) *TripRatingCreate {
	return trc.SetFromUserID(u.ID)
}

// SetToUser sets the "to_user" edge to the User entity.
func (trc *TripRatingCreate) SetToUser(u *User) *TripRatingCreate {
	return trc.SetToUserID(u.ID)
}

// Mutation returns the TripRatingMutation object of the builder.
func (trc *TripRatingCreate) Mutation() *TripRatingMutation {
	return trc.mutation
}

// Save creates the TripRating in the database.
func (trc *TripRatingCreate) Save(ctx context.Context) (*TripRating, error) {
	trc.defaults()
	return withHooks(ctx, trc.sqlSave, trc.mutation, trc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (trc *TripRatingCreate) SaveX(ctx context.Context) *TripRating {
	v, err := trc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trc *TripRatingCreate) Exec(ctx context.Context) error {
	_, err := trc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trc *TripRatingCreate) ExecX(ctx context.Context) {
	if err := trc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (trc *TripRatingCreate) defaults() {
	if _, ok := trc.mutation.Comment(); !ok {
		v := triprating.DefaultComment
		trc.mutation.SetComment(v)
	}
	if _, ok := trc.mutation.CreatedAt(); !ok {
		v := triprating.DefaultCreatedAt
		trc.mutation.SetCreatedAt(v)
	}
	if _, ok := trc.mutation.ID(); !ok {
		v := triprating.DefaultID()
		trc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trc *TripRatingCreate) check() error {
	if _, ok := trc.mutation.TripID(); !ok {
		return &ValidationError{Name: "trip_id", err: errors.New(`ent: missing required field "TripRating.trip_id"`)}
	}
	if _, ok := trc.mutation.FromUserID(); !ok {
		return &ValidationError{Name: "from_user_id", err: errors.New(`ent: missing required field "TripRating.from_user_id"`)}
	}
	if _, ok := trc.mutation.ToUserID(); !ok {
		return &ValidationError{Name: "to_user_id", err: errors.New(`ent: missing required field "TripRating.to_user_id"`)}
	}
	if _, ok := trc.mutation.Rating(); !ok {
		return &ValidationError{Name: "rating", err: errors.New(`ent: missing required field "TripRating.rating"`)}
	}
	if _, ok := trc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "TripRating.comment"`)}
	}
	if _, ok := trc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TripRating.created_at"`)}
	}
	if len(trc.mutation.TripIDs()) == 0 {
		return &ValidationError{Name: "trip", err: errors.New(`ent: missing required edge "TripRating.trip"`)}
	}
	if len(trc.mutation.FromUserIDs()) == 0 {
		return &ValidationError{Name: "from_user", err: errors.New(`ent: missing required edge "TripRating.from_user"`)}
	}
	if len(trc.mutation.ToUserIDs()) == 0 {
		return &ValidationError{Name: "to_user", err: errors.New(`ent: missing required edge "TripRating.to_user"`)}
	}
	return nil
}

func (trc *TripRatingCreate) sqlSave(ctx context.Context) (*TripRating, error) {
	if err := trc.check(); err != nil {
		return nil, err
	}
	_node, _spec := trc.createSpec()
	if err := sqlgraph.CreateNode(ctx, trc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	trc.mutation.id = &_node.ID
	trc.mutation.done = true
	return _node, nil
}

func (trc *TripRatingCreate) createSpec() (*TripRating, *sqlgraph.CreateSpec) {
	var (
		_node = &TripRating{config: trc.config}
		_spec = sqlgraph.NewCreateSpec(triprating.Table, sqlgraph.NewFieldSpec(triprating.FieldID, field.TypeUUID))
	)
	if id, ok := trc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := trc.mutation.Rating(); ok {
		_spec.SetField(triprating.FieldRating, field.TypeInt, value)
		_node.Rating = value
	}
	if value, ok := trc.mutation.Comment(); ok {
		_spec.SetField(triprating.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if value, ok := trc.mutation.CreatedAt(); ok {
		_spec.SetField(triprating.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := trc.mutation.TripIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triprating.TripTable,
			Columns: []string{triprating.TripColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trip.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TripID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.FromUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triprating.FromUserTable,
			Columns: []string{triprating.FromUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FromUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.ToUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   triprating.ToUserTable,
			Columns: []string{triprating.ToUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ToUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TripRatingCreateBulk is the builder for creating many TripRating entities in bulk.
type TripRatingCreateBulk struct {
	config
	err      error
	builders []*TripRatingCreate
}

// Save creates the TripRating entities in the database.
func (trcb *TripRatingCreateBulk) Save(ctx context.Context) ([]*TripRating, error) {
	if trcb.err != nil {
		return nil, trcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(trcb.builders))
	nodes := make([]*TripRating, len(trcb.builders))
	mutators := make([]Mutator, len(trcb.builders))
	for i := range trcb.builders {
		func(i int, root context.Context) {
			builder := trcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TripRatingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, trcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, trcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, trcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (trcb *TripRatingCreateBulk) SaveX(ctx context.Context) []*TripRating {
	v, err := trcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trcb *TripRatingCreateBulk) Exec(ctx context.Context) error {
	_, err := trcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trcb *TripRatingCreateBulk) ExecX(ctx context.Context) {
	if err := trcb.Exec(ctx); err != nil {
		panic(err)
	}
}
