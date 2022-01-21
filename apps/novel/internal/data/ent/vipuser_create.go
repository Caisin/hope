// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/socialuser"
	"hope/apps/novel/internal/data/ent/vipuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VipUserCreate is the builder for creating a VipUser entity.
type VipUserCreate struct {
	config
	mutation *VipUserMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (vuc *VipUserCreate) SetUserId(i int64) *VipUserCreate {
	vuc.mutation.SetUserId(i)
	return vuc
}

// SetVipType sets the "vipType" field.
func (vuc *VipUserCreate) SetVipType(i int64) *VipUserCreate {
	vuc.mutation.SetVipType(i)
	return vuc
}

// SetNillableVipType sets the "vipType" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableVipType(i *int64) *VipUserCreate {
	if i != nil {
		vuc.SetVipType(*i)
	}
	return vuc
}

// SetSvipType sets the "svipType" field.
func (vuc *VipUserCreate) SetSvipType(i int64) *VipUserCreate {
	vuc.mutation.SetSvipType(i)
	return vuc
}

// SetNillableSvipType sets the "svipType" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableSvipType(i *int64) *VipUserCreate {
	if i != nil {
		vuc.SetSvipType(*i)
	}
	return vuc
}

// SetSvipEffectTime sets the "svipEffectTime" field.
func (vuc *VipUserCreate) SetSvipEffectTime(t time.Time) *VipUserCreate {
	vuc.mutation.SetSvipEffectTime(t)
	return vuc
}

// SetNillableSvipEffectTime sets the "svipEffectTime" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableSvipEffectTime(t *time.Time) *VipUserCreate {
	if t != nil {
		vuc.SetSvipEffectTime(*t)
	}
	return vuc
}

// SetSvipExpiredTime sets the "svipExpiredTime" field.
func (vuc *VipUserCreate) SetSvipExpiredTime(t time.Time) *VipUserCreate {
	vuc.mutation.SetSvipExpiredTime(t)
	return vuc
}

// SetNillableSvipExpiredTime sets the "svipExpiredTime" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableSvipExpiredTime(t *time.Time) *VipUserCreate {
	if t != nil {
		vuc.SetSvipExpiredTime(*t)
	}
	return vuc
}

// SetRemark sets the "remark" field.
func (vuc *VipUserCreate) SetRemark(s string) *VipUserCreate {
	vuc.mutation.SetRemark(s)
	return vuc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableRemark(s *string) *VipUserCreate {
	if s != nil {
		vuc.SetRemark(*s)
	}
	return vuc
}

// SetEffectTime sets the "effectTime" field.
func (vuc *VipUserCreate) SetEffectTime(t time.Time) *VipUserCreate {
	vuc.mutation.SetEffectTime(t)
	return vuc
}

// SetNillableEffectTime sets the "effectTime" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableEffectTime(t *time.Time) *VipUserCreate {
	if t != nil {
		vuc.SetEffectTime(*t)
	}
	return vuc
}

// SetExpiredTime sets the "expiredTime" field.
func (vuc *VipUserCreate) SetExpiredTime(t time.Time) *VipUserCreate {
	vuc.mutation.SetExpiredTime(t)
	return vuc
}

// SetNillableExpiredTime sets the "expiredTime" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableExpiredTime(t *time.Time) *VipUserCreate {
	if t != nil {
		vuc.SetExpiredTime(*t)
	}
	return vuc
}

// SetCreatedAt sets the "createdAt" field.
func (vuc *VipUserCreate) SetCreatedAt(t time.Time) *VipUserCreate {
	vuc.mutation.SetCreatedAt(t)
	return vuc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableCreatedAt(t *time.Time) *VipUserCreate {
	if t != nil {
		vuc.SetCreatedAt(*t)
	}
	return vuc
}

// SetUpdatedAt sets the "updatedAt" field.
func (vuc *VipUserCreate) SetUpdatedAt(t time.Time) *VipUserCreate {
	vuc.mutation.SetUpdatedAt(t)
	return vuc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableUpdatedAt(t *time.Time) *VipUserCreate {
	if t != nil {
		vuc.SetUpdatedAt(*t)
	}
	return vuc
}

// SetCreateBy sets the "createBy" field.
func (vuc *VipUserCreate) SetCreateBy(i int64) *VipUserCreate {
	vuc.mutation.SetCreateBy(i)
	return vuc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableCreateBy(i *int64) *VipUserCreate {
	if i != nil {
		vuc.SetCreateBy(*i)
	}
	return vuc
}

// SetUpdateBy sets the "updateBy" field.
func (vuc *VipUserCreate) SetUpdateBy(i int64) *VipUserCreate {
	vuc.mutation.SetUpdateBy(i)
	return vuc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableUpdateBy(i *int64) *VipUserCreate {
	if i != nil {
		vuc.SetUpdateBy(*i)
	}
	return vuc
}

// SetTenantId sets the "tenantId" field.
func (vuc *VipUserCreate) SetTenantId(i int64) *VipUserCreate {
	vuc.mutation.SetTenantId(i)
	return vuc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (vuc *VipUserCreate) SetNillableTenantId(i *int64) *VipUserCreate {
	if i != nil {
		vuc.SetTenantId(*i)
	}
	return vuc
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (vuc *VipUserCreate) SetUserID(id int64) *VipUserCreate {
	vuc.mutation.SetUserID(id)
	return vuc
}

// SetUser sets the "user" edge to the SocialUser entity.
func (vuc *VipUserCreate) SetUser(s *SocialUser) *VipUserCreate {
	return vuc.SetUserID(s.ID)
}

// Mutation returns the VipUserMutation object of the builder.
func (vuc *VipUserCreate) Mutation() *VipUserMutation {
	return vuc.mutation
}

// Save creates the VipUser in the database.
func (vuc *VipUserCreate) Save(ctx context.Context) (*VipUser, error) {
	var (
		err  error
		node *VipUser
	)
	vuc.defaults()
	if len(vuc.hooks) == 0 {
		if err = vuc.check(); err != nil {
			return nil, err
		}
		node, err = vuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VipUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuc.check(); err != nil {
				return nil, err
			}
			vuc.mutation = mutation
			if node, err = vuc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vuc.hooks) - 1; i >= 0; i-- {
			if vuc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vuc *VipUserCreate) SaveX(ctx context.Context) *VipUser {
	v, err := vuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vuc *VipUserCreate) Exec(ctx context.Context) error {
	_, err := vuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuc *VipUserCreate) ExecX(ctx context.Context) {
	if err := vuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuc *VipUserCreate) defaults() {
	if _, ok := vuc.mutation.EffectTime(); !ok {
		v := vipuser.DefaultEffectTime()
		vuc.mutation.SetEffectTime(v)
	}
	if _, ok := vuc.mutation.ExpiredTime(); !ok {
		v := vipuser.DefaultExpiredTime()
		vuc.mutation.SetExpiredTime(v)
	}
	if _, ok := vuc.mutation.CreatedAt(); !ok {
		v := vipuser.DefaultCreatedAt()
		vuc.mutation.SetCreatedAt(v)
	}
	if _, ok := vuc.mutation.UpdatedAt(); !ok {
		v := vipuser.DefaultUpdatedAt()
		vuc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vuc.mutation.CreateBy(); !ok {
		v := vipuser.DefaultCreateBy
		vuc.mutation.SetCreateBy(v)
	}
	if _, ok := vuc.mutation.UpdateBy(); !ok {
		v := vipuser.DefaultUpdateBy
		vuc.mutation.SetUpdateBy(v)
	}
	if _, ok := vuc.mutation.TenantId(); !ok {
		v := vipuser.DefaultTenantId
		vuc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuc *VipUserCreate) check() error {
	if _, ok := vuc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "userId"`)}
	}
	if _, ok := vuc.mutation.EffectTime(); !ok {
		return &ValidationError{Name: "effectTime", err: errors.New(`ent: missing required field "effectTime"`)}
	}
	if _, ok := vuc.mutation.ExpiredTime(); !ok {
		return &ValidationError{Name: "expiredTime", err: errors.New(`ent: missing required field "expiredTime"`)}
	}
	if _, ok := vuc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := vuc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := vuc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := vuc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := vuc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	if _, ok := vuc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (vuc *VipUserCreate) sqlSave(ctx context.Context) (*VipUser, error) {
	_node, _spec := vuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (vuc *VipUserCreate) createSpec() (*VipUser, *sqlgraph.CreateSpec) {
	var (
		_node = &VipUser{config: vuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vipuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: vipuser.FieldID,
			},
		}
	)
	if value, ok := vuc.mutation.VipType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vipuser.FieldVipType,
		})
		_node.VipType = value
	}
	if value, ok := vuc.mutation.SvipType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vipuser.FieldSvipType,
		})
		_node.SvipType = value
	}
	if value, ok := vuc.mutation.SvipEffectTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vipuser.FieldSvipEffectTime,
		})
		_node.SvipEffectTime = value
	}
	if value, ok := vuc.mutation.SvipExpiredTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vipuser.FieldSvipExpiredTime,
		})
		_node.SvipExpiredTime = value
	}
	if value, ok := vuc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vipuser.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := vuc.mutation.EffectTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vipuser.FieldEffectTime,
		})
		_node.EffectTime = value
	}
	if value, ok := vuc.mutation.ExpiredTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vipuser.FieldExpiredTime,
		})
		_node.ExpiredTime = value
	}
	if value, ok := vuc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vipuser.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := vuc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: vipuser.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := vuc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vipuser.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := vuc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vipuser.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := vuc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: vipuser.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := vuc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vipuser.UserTable,
			Columns: []string{vipuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: socialuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VipUserCreateBulk is the builder for creating many VipUser entities in bulk.
type VipUserCreateBulk struct {
	config
	builders []*VipUserCreate
}

// Save creates the VipUser entities in the database.
func (vucb *VipUserCreateBulk) Save(ctx context.Context) ([]*VipUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vucb.builders))
	nodes := make([]*VipUser, len(vucb.builders))
	mutators := make([]Mutator, len(vucb.builders))
	for i := range vucb.builders {
		func(i int, root context.Context) {
			builder := vucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VipUserMutation)
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
					_, err = mutators[i+1].Mutate(root, vucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vucb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vucb *VipUserCreateBulk) SaveX(ctx context.Context) []*VipUser {
	v, err := vucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vucb *VipUserCreateBulk) Exec(ctx context.Context) error {
	_, err := vucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vucb *VipUserCreateBulk) ExecX(ctx context.Context) {
	if err := vucb.Exec(ctx); err != nil {
		panic(err)
	}
}
