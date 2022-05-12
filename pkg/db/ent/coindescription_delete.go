// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/coindescription"
	"github.com/NpoolPlatform/project-info-manager/pkg/db/ent/predicate"
)

// CoinDescriptionDelete is the builder for deleting a CoinDescription entity.
type CoinDescriptionDelete struct {
	config
	hooks    []Hook
	mutation *CoinDescriptionMutation
}

// Where appends a list predicates to the CoinDescriptionDelete builder.
func (cdd *CoinDescriptionDelete) Where(ps ...predicate.CoinDescription) *CoinDescriptionDelete {
	cdd.mutation.Where(ps...)
	return cdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cdd *CoinDescriptionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cdd.hooks) == 0 {
		affected, err = cdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cdd.mutation = mutation
			affected, err = cdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cdd.hooks) - 1; i >= 0; i-- {
			if cdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdd *CoinDescriptionDelete) ExecX(ctx context.Context) int {
	n, err := cdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cdd *CoinDescriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coindescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coindescription.FieldID,
			},
		},
	}
	if ps := cdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cdd.driver, _spec)
}

// CoinDescriptionDeleteOne is the builder for deleting a single CoinDescription entity.
type CoinDescriptionDeleteOne struct {
	cdd *CoinDescriptionDelete
}

// Exec executes the deletion query.
func (cddo *CoinDescriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := cddo.cdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coindescription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cddo *CoinDescriptionDeleteOne) ExecX(ctx context.Context) {
	cddo.cdd.ExecX(ctx)
}