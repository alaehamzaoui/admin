// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bo-mathadventure/admin/ent/ban"
	"github.com/bo-mathadventure/admin/ent/predicate"
)

// BanUpdate is the builder for updating Ban entities.
type BanUpdate struct {
	config
	hooks    []Hook
	mutation *BanMutation
}

// Where appends a list predicates to the BanUpdate builder.
func (bu *BanUpdate) Where(ps ...predicate.Ban) *BanUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetCheck sets the "check" field.
func (bu *BanUpdate) SetCheck(s string) *BanUpdate {
	bu.mutation.SetCheck(s)
	return bu
}

// SetMessage sets the "message" field.
func (bu *BanUpdate) SetMessage(s string) *BanUpdate {
	bu.mutation.SetMessage(s)
	return bu
}

// SetValidUntil sets the "validUntil" field.
func (bu *BanUpdate) SetValidUntil(t time.Time) *BanUpdate {
	bu.mutation.SetValidUntil(t)
	return bu
}

// SetCreatedAt sets the "createdAt" field.
func (bu *BanUpdate) SetCreatedAt(t time.Time) *BanUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bu *BanUpdate) SetNillableCreatedAt(t *time.Time) *BanUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// Mutation returns the BanMutation object of the builder.
func (bu *BanUpdate) Mutation() *BanMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BanUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BanUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BanUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BanUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ban.Table, ban.Columns, sqlgraph.NewFieldSpec(ban.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Check(); ok {
		_spec.SetField(ban.FieldCheck, field.TypeString, value)
	}
	if value, ok := bu.mutation.Message(); ok {
		_spec.SetField(ban.FieldMessage, field.TypeString, value)
	}
	if value, ok := bu.mutation.ValidUntil(); ok {
		_spec.SetField(ban.FieldValidUntil, field.TypeTime, value)
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(ban.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ban.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BanUpdateOne is the builder for updating a single Ban entity.
type BanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BanMutation
}

// SetCheck sets the "check" field.
func (buo *BanUpdateOne) SetCheck(s string) *BanUpdateOne {
	buo.mutation.SetCheck(s)
	return buo
}

// SetMessage sets the "message" field.
func (buo *BanUpdateOne) SetMessage(s string) *BanUpdateOne {
	buo.mutation.SetMessage(s)
	return buo
}

// SetValidUntil sets the "validUntil" field.
func (buo *BanUpdateOne) SetValidUntil(t time.Time) *BanUpdateOne {
	buo.mutation.SetValidUntil(t)
	return buo
}

// SetCreatedAt sets the "createdAt" field.
func (buo *BanUpdateOne) SetCreatedAt(t time.Time) *BanUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (buo *BanUpdateOne) SetNillableCreatedAt(t *time.Time) *BanUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// Mutation returns the BanMutation object of the builder.
func (buo *BanUpdateOne) Mutation() *BanMutation {
	return buo.mutation
}

// Where appends a list predicates to the BanUpdate builder.
func (buo *BanUpdateOne) Where(ps ...predicate.Ban) *BanUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BanUpdateOne) Select(field string, fields ...string) *BanUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Ban entity.
func (buo *BanUpdateOne) Save(ctx context.Context) (*Ban, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BanUpdateOne) SaveX(ctx context.Context) *Ban {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BanUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BanUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BanUpdateOne) sqlSave(ctx context.Context) (_node *Ban, err error) {
	_spec := sqlgraph.NewUpdateSpec(ban.Table, ban.Columns, sqlgraph.NewFieldSpec(ban.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ban.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ban.FieldID)
		for _, f := range fields {
			if !ban.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ban.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Check(); ok {
		_spec.SetField(ban.FieldCheck, field.TypeString, value)
	}
	if value, ok := buo.mutation.Message(); ok {
		_spec.SetField(ban.FieldMessage, field.TypeString, value)
	}
	if value, ok := buo.mutation.ValidUntil(); ok {
		_spec.SetField(ban.FieldValidUntil, field.TypeTime, value)
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(ban.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Ban{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ban.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
