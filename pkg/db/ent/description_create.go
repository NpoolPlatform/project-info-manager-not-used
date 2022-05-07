// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/description"
	"github.com/google/uuid"
)

// DescriptionCreate is the builder for creating a Description entity.
type DescriptionCreate struct {
	config
	mutation *DescriptionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DescriptionCreate) SetCreatedAt(u uint32) *DescriptionCreate {
	dc.mutation.SetCreatedAt(u)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableCreatedAt(u *uint32) *DescriptionCreate {
	if u != nil {
		dc.SetCreatedAt(*u)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DescriptionCreate) SetUpdatedAt(u uint32) *DescriptionCreate {
	dc.mutation.SetUpdatedAt(u)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableUpdatedAt(u *uint32) *DescriptionCreate {
	if u != nil {
		dc.SetUpdatedAt(*u)
	}
	return dc
}

// SetDeletedAt sets the "deleted_at" field.
func (dc *DescriptionCreate) SetDeletedAt(u uint32) *DescriptionCreate {
	dc.mutation.SetDeletedAt(u)
	return dc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableDeletedAt(u *uint32) *DescriptionCreate {
	if u != nil {
		dc.SetDeletedAt(*u)
	}
	return dc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (dc *DescriptionCreate) SetCoinTypeID(u uuid.UUID) *DescriptionCreate {
	dc.mutation.SetCoinTypeID(u)
	return dc
}

// SetTitle sets the "title" field.
func (dc *DescriptionCreate) SetTitle(s string) *DescriptionCreate {
	dc.mutation.SetTitle(s)
	return dc
}

// SetMessage sets the "message" field.
func (dc *DescriptionCreate) SetMessage(s string) *DescriptionCreate {
	dc.mutation.SetMessage(s)
	return dc
}

// SetUsedFor sets the "used_for" field.
func (dc *DescriptionCreate) SetUsedFor(s string) *DescriptionCreate {
	dc.mutation.SetUsedFor(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DescriptionCreate) SetID(u uuid.UUID) *DescriptionCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dc *DescriptionCreate) SetNillableID(u *uuid.UUID) *DescriptionCreate {
	if u != nil {
		dc.SetID(*u)
	}
	return dc
}

// Mutation returns the DescriptionMutation object of the builder.
func (dc *DescriptionCreate) Mutation() *DescriptionMutation {
	return dc.mutation
}

// Save creates the Description in the database.
func (dc *DescriptionCreate) Save(ctx context.Context) (*Description, error) {
	var (
		err  error
		node *Description
	)
	if err := dc.defaults(); err != nil {
		return nil, err
	}
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			if node, err = dc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			if dc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DescriptionCreate) SaveX(ctx context.Context) *Description {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DescriptionCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DescriptionCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DescriptionCreate) defaults() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		if description.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized description.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := description.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		if description.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized description.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := description.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		if description.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized description.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := description.DefaultDeletedAt()
		dc.mutation.SetDeletedAt(v)
	}
	if _, ok := dc.mutation.ID(); !ok {
		if description.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized description.DefaultID (forgotten import ent/runtime?)")
		}
		v := description.DefaultID()
		dc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dc *DescriptionCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Description.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Description.updated_at"`)}
	}
	if _, ok := dc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Description.deleted_at"`)}
	}
	if _, ok := dc.mutation.CoinTypeID(); !ok {
		return &ValidationError{Name: "coin_type_id", err: errors.New(`ent: missing required field "Description.coin_type_id"`)}
	}
	if _, ok := dc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Description.title"`)}
	}
	if _, ok := dc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "Description.message"`)}
	}
	if v, ok := dc.mutation.Message(); ok {
		if err := description.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Description.message": %w`, err)}
		}
	}
	if _, ok := dc.mutation.UsedFor(); !ok {
		return &ValidationError{Name: "used_for", err: errors.New(`ent: missing required field "Description.used_for"`)}
	}
	return nil
}

func (dc *DescriptionCreate) sqlSave(ctx context.Context) (*Description, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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
	return _node, nil
}

func (dc *DescriptionCreate) createSpec() (*Description, *sqlgraph.CreateSpec) {
	var (
		_node = &Description{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: description.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: description.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: description.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: description.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: description.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := dc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: description.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := dc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: description.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := dc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: description.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := dc.mutation.UsedFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: description.FieldUsedFor,
		})
		_node.UsedFor = value
	}
	return _node, _spec
}

// DescriptionCreateBulk is the builder for creating many Description entities in bulk.
type DescriptionCreateBulk struct {
	config
	builders []*DescriptionCreate
}

// Save creates the Description entities in the database.
func (dcb *DescriptionCreateBulk) Save(ctx context.Context) ([]*Description, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Description, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DescriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DescriptionCreateBulk) SaveX(ctx context.Context) []*Description {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DescriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DescriptionCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
