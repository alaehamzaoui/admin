// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bo-mathadventure/admin/ent/maps"
)

// MapsCreate is the builder for creating a Maps entity.
type MapsCreate struct {
	config
	mutation *MapsMutation
	hooks    []Hook
}

// SetRoomName sets the "roomName" field.
func (mc *MapsCreate) SetRoomName(s string) *MapsCreate {
	mc.mutation.SetRoomName(s)
	return mc
}

// SetMapUrl sets the "mapUrl" field.
func (mc *MapsCreate) SetMapUrl(s string) *MapsCreate {
	mc.mutation.SetMapUrl(s)
	return mc
}

// SetPolicyNumber sets the "policyNumber" field.
func (mc *MapsCreate) SetPolicyNumber(i int) *MapsCreate {
	mc.mutation.SetPolicyNumber(i)
	return mc
}

// SetContactPage sets the "contactPage" field.
func (mc *MapsCreate) SetContactPage(s string) *MapsCreate {
	mc.mutation.SetContactPage(s)
	return mc
}

// SetTags sets the "tags" field.
func (mc *MapsCreate) SetTags(s []string) *MapsCreate {
	mc.mutation.SetTags(s)
	return mc
}

// SetEnableChat sets the "enableChat" field.
func (mc *MapsCreate) SetEnableChat(b bool) *MapsCreate {
	mc.mutation.SetEnableChat(b)
	return mc
}

// SetEnableChatUpload sets the "enableChatUpload" field.
func (mc *MapsCreate) SetEnableChatUpload(b bool) *MapsCreate {
	mc.mutation.SetEnableChatUpload(b)
	return mc
}

// SetEnableChatOnlineList sets the "enableChatOnlineList" field.
func (mc *MapsCreate) SetEnableChatOnlineList(b bool) *MapsCreate {
	mc.mutation.SetEnableChatOnlineList(b)
	return mc
}

// SetEnableChatDisconnectedList sets the "enableChatDisconnectedList" field.
func (mc *MapsCreate) SetEnableChatDisconnectedList(b bool) *MapsCreate {
	mc.mutation.SetEnableChatDisconnectedList(b)
	return mc
}

// SetCanReport sets the "canReport" field.
func (mc *MapsCreate) SetCanReport(b bool) *MapsCreate {
	mc.mutation.SetCanReport(b)
	return mc
}

// SetExpireOn sets the "expireOn" field.
func (mc *MapsCreate) SetExpireOn(t time.Time) *MapsCreate {
	mc.mutation.SetExpireOn(t)
	return mc
}

// SetCreatedAt sets the "createdAt" field.
func (mc *MapsCreate) SetCreatedAt(t time.Time) *MapsCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (mc *MapsCreate) SetNillableCreatedAt(t *time.Time) *MapsCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// Mutation returns the MapsMutation object of the builder.
func (mc *MapsCreate) Mutation() *MapsMutation {
	return mc.mutation
}

// Save creates the Maps in the database.
func (mc *MapsCreate) Save(ctx context.Context) (*Maps, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MapsCreate) SaveX(ctx context.Context) *Maps {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MapsCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MapsCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MapsCreate) defaults() {
	if _, ok := mc.mutation.Tags(); !ok {
		v := maps.DefaultTags
		mc.mutation.SetTags(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := maps.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MapsCreate) check() error {
	if _, ok := mc.mutation.RoomName(); !ok {
		return &ValidationError{Name: "roomName", err: errors.New(`ent: missing required field "Maps.roomName"`)}
	}
	if _, ok := mc.mutation.MapUrl(); !ok {
		return &ValidationError{Name: "mapUrl", err: errors.New(`ent: missing required field "Maps.mapUrl"`)}
	}
	if _, ok := mc.mutation.PolicyNumber(); !ok {
		return &ValidationError{Name: "policyNumber", err: errors.New(`ent: missing required field "Maps.policyNumber"`)}
	}
	if _, ok := mc.mutation.ContactPage(); !ok {
		return &ValidationError{Name: "contactPage", err: errors.New(`ent: missing required field "Maps.contactPage"`)}
	}
	if _, ok := mc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "Maps.tags"`)}
	}
	if _, ok := mc.mutation.EnableChat(); !ok {
		return &ValidationError{Name: "enableChat", err: errors.New(`ent: missing required field "Maps.enableChat"`)}
	}
	if _, ok := mc.mutation.EnableChatUpload(); !ok {
		return &ValidationError{Name: "enableChatUpload", err: errors.New(`ent: missing required field "Maps.enableChatUpload"`)}
	}
	if _, ok := mc.mutation.EnableChatOnlineList(); !ok {
		return &ValidationError{Name: "enableChatOnlineList", err: errors.New(`ent: missing required field "Maps.enableChatOnlineList"`)}
	}
	if _, ok := mc.mutation.EnableChatDisconnectedList(); !ok {
		return &ValidationError{Name: "enableChatDisconnectedList", err: errors.New(`ent: missing required field "Maps.enableChatDisconnectedList"`)}
	}
	if _, ok := mc.mutation.CanReport(); !ok {
		return &ValidationError{Name: "canReport", err: errors.New(`ent: missing required field "Maps.canReport"`)}
	}
	if _, ok := mc.mutation.ExpireOn(); !ok {
		return &ValidationError{Name: "expireOn", err: errors.New(`ent: missing required field "Maps.expireOn"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Maps.createdAt"`)}
	}
	return nil
}

func (mc *MapsCreate) sqlSave(ctx context.Context) (*Maps, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MapsCreate) createSpec() (*Maps, *sqlgraph.CreateSpec) {
	var (
		_node = &Maps{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(maps.Table, sqlgraph.NewFieldSpec(maps.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.RoomName(); ok {
		_spec.SetField(maps.FieldRoomName, field.TypeString, value)
		_node.RoomName = value
	}
	if value, ok := mc.mutation.MapUrl(); ok {
		_spec.SetField(maps.FieldMapUrl, field.TypeString, value)
		_node.MapUrl = value
	}
	if value, ok := mc.mutation.PolicyNumber(); ok {
		_spec.SetField(maps.FieldPolicyNumber, field.TypeInt, value)
		_node.PolicyNumber = value
	}
	if value, ok := mc.mutation.ContactPage(); ok {
		_spec.SetField(maps.FieldContactPage, field.TypeString, value)
		_node.ContactPage = value
	}
	if value, ok := mc.mutation.Tags(); ok {
		_spec.SetField(maps.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := mc.mutation.EnableChat(); ok {
		_spec.SetField(maps.FieldEnableChat, field.TypeBool, value)
		_node.EnableChat = value
	}
	if value, ok := mc.mutation.EnableChatUpload(); ok {
		_spec.SetField(maps.FieldEnableChatUpload, field.TypeBool, value)
		_node.EnableChatUpload = value
	}
	if value, ok := mc.mutation.EnableChatOnlineList(); ok {
		_spec.SetField(maps.FieldEnableChatOnlineList, field.TypeBool, value)
		_node.EnableChatOnlineList = value
	}
	if value, ok := mc.mutation.EnableChatDisconnectedList(); ok {
		_spec.SetField(maps.FieldEnableChatDisconnectedList, field.TypeBool, value)
		_node.EnableChatDisconnectedList = value
	}
	if value, ok := mc.mutation.CanReport(); ok {
		_spec.SetField(maps.FieldCanReport, field.TypeBool, value)
		_node.CanReport = value
	}
	if value, ok := mc.mutation.ExpireOn(); ok {
		_spec.SetField(maps.FieldExpireOn, field.TypeTime, value)
		_node.ExpireOn = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(maps.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// MapsCreateBulk is the builder for creating many Maps entities in bulk.
type MapsCreateBulk struct {
	config
	builders []*MapsCreate
}

// Save creates the Maps entities in the database.
func (mcb *MapsCreateBulk) Save(ctx context.Context) ([]*Maps, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Maps, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MapsMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MapsCreateBulk) SaveX(ctx context.Context) []*Maps {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MapsCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MapsCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
