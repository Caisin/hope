// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/agreementlog"
	"hope/apps/novel/internal/data/ent/payorder"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgreementLogCreate is the builder for creating a AgreementLog entity.
type AgreementLogCreate struct {
	config
	mutation *AgreementLogMutation
	hooks    []Hook
}

// SetOuterAgreementNo sets the "outerAgreementNo" field.
func (alc *AgreementLogCreate) SetOuterAgreementNo(s string) *AgreementLogCreate {
	alc.mutation.SetOuterAgreementNo(s)
	return alc
}

// SetNillableOuterAgreementNo sets the "outerAgreementNo" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableOuterAgreementNo(s *string) *AgreementLogCreate {
	if s != nil {
		alc.SetOuterAgreementNo(*s)
	}
	return alc
}

// SetOrderId sets the "orderId" field.
func (alc *AgreementLogCreate) SetOrderId(s string) *AgreementLogCreate {
	alc.mutation.SetOrderId(s)
	return alc
}

// SetNillableOrderId sets the "orderId" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableOrderId(s *string) *AgreementLogCreate {
	if s != nil {
		alc.SetOrderId(*s)
	}
	return alc
}

// SetUserId sets the "userId" field.
func (alc *AgreementLogCreate) SetUserId(i int64) *AgreementLogCreate {
	alc.mutation.SetUserId(i)
	return alc
}

// SetNillableUserId sets the "userId" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableUserId(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetUserId(*i)
	}
	return alc
}

// SetChId sets the "chId" field.
func (alc *AgreementLogCreate) SetChId(i int64) *AgreementLogCreate {
	alc.mutation.SetChId(i)
	return alc
}

// SetNillableChId sets the "chId" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableChId(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetChId(*i)
	}
	return alc
}

// SetUserName sets the "userName" field.
func (alc *AgreementLogCreate) SetUserName(s string) *AgreementLogCreate {
	alc.mutation.SetUserName(s)
	return alc
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableUserName(s *string) *AgreementLogCreate {
	if s != nil {
		alc.SetUserName(*s)
	}
	return alc
}

// SetPaymentName sets the "paymentName" field.
func (alc *AgreementLogCreate) SetPaymentName(s string) *AgreementLogCreate {
	alc.mutation.SetPaymentName(s)
	return alc
}

// SetNillablePaymentName sets the "paymentName" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillablePaymentName(s *string) *AgreementLogCreate {
	if s != nil {
		alc.SetPaymentName(*s)
	}
	return alc
}

// SetPaymentId sets the "paymentId" field.
func (alc *AgreementLogCreate) SetPaymentId(i int64) *AgreementLogCreate {
	alc.mutation.SetPaymentId(i)
	return alc
}

// SetNillablePaymentId sets the "paymentId" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillablePaymentId(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetPaymentId(*i)
	}
	return alc
}

// SetState sets the "state" field.
func (alc *AgreementLogCreate) SetState(i int32) *AgreementLogCreate {
	alc.mutation.SetState(i)
	return alc
}

// SetNillableState sets the "state" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableState(i *int32) *AgreementLogCreate {
	if i != nil {
		alc.SetState(*i)
	}
	return alc
}

// SetPayment sets the "payment" field.
func (alc *AgreementLogCreate) SetPayment(i int64) *AgreementLogCreate {
	alc.mutation.SetPayment(i)
	return alc
}

// SetNillablePayment sets the "payment" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillablePayment(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetPayment(*i)
	}
	return alc
}

// SetAgreementType sets the "agreementType" field.
func (alc *AgreementLogCreate) SetAgreementType(at agreementlog.AgreementType) *AgreementLogCreate {
	alc.mutation.SetAgreementType(at)
	return alc
}

// SetNillableAgreementType sets the "agreementType" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableAgreementType(at *agreementlog.AgreementType) *AgreementLogCreate {
	if at != nil {
		alc.SetAgreementType(*at)
	}
	return alc
}

// SetVipType sets the "vipType" field.
func (alc *AgreementLogCreate) SetVipType(i int64) *AgreementLogCreate {
	alc.mutation.SetVipType(i)
	return alc
}

// SetNillableVipType sets the "vipType" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableVipType(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetVipType(*i)
	}
	return alc
}

// SetTimes sets the "times" field.
func (alc *AgreementLogCreate) SetTimes(i int64) *AgreementLogCreate {
	alc.mutation.SetTimes(i)
	return alc
}

// SetNillableTimes sets the "times" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableTimes(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetTimes(*i)
	}
	return alc
}

// SetCycleDays sets the "cycleDays" field.
func (alc *AgreementLogCreate) SetCycleDays(i int32) *AgreementLogCreate {
	alc.mutation.SetCycleDays(i)
	return alc
}

// SetNillableCycleDays sets the "cycleDays" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableCycleDays(i *int32) *AgreementLogCreate {
	if i != nil {
		alc.SetCycleDays(*i)
	}
	return alc
}

// SetNextExecTime sets the "nextExecTime" field.
func (alc *AgreementLogCreate) SetNextExecTime(t time.Time) *AgreementLogCreate {
	alc.mutation.SetNextExecTime(t)
	return alc
}

// SetNillableNextExecTime sets the "nextExecTime" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableNextExecTime(t *time.Time) *AgreementLogCreate {
	if t != nil {
		alc.SetNextExecTime(*t)
	}
	return alc
}

// SetRemark sets the "remark" field.
func (alc *AgreementLogCreate) SetRemark(s string) *AgreementLogCreate {
	alc.mutation.SetRemark(s)
	return alc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableRemark(s *string) *AgreementLogCreate {
	if s != nil {
		alc.SetRemark(*s)
	}
	return alc
}

// SetCreatedAt sets the "createdAt" field.
func (alc *AgreementLogCreate) SetCreatedAt(t time.Time) *AgreementLogCreate {
	alc.mutation.SetCreatedAt(t)
	return alc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableCreatedAt(t *time.Time) *AgreementLogCreate {
	if t != nil {
		alc.SetCreatedAt(*t)
	}
	return alc
}

// SetUpdatedAt sets the "updatedAt" field.
func (alc *AgreementLogCreate) SetUpdatedAt(t time.Time) *AgreementLogCreate {
	alc.mutation.SetUpdatedAt(t)
	return alc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableUpdatedAt(t *time.Time) *AgreementLogCreate {
	if t != nil {
		alc.SetUpdatedAt(*t)
	}
	return alc
}

// SetCreateBy sets the "createBy" field.
func (alc *AgreementLogCreate) SetCreateBy(i int64) *AgreementLogCreate {
	alc.mutation.SetCreateBy(i)
	return alc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableCreateBy(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetCreateBy(*i)
	}
	return alc
}

// SetUpdateBy sets the "updateBy" field.
func (alc *AgreementLogCreate) SetUpdateBy(i int64) *AgreementLogCreate {
	alc.mutation.SetUpdateBy(i)
	return alc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableUpdateBy(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetUpdateBy(*i)
	}
	return alc
}

// SetTenantId sets the "tenantId" field.
func (alc *AgreementLogCreate) SetTenantId(i int64) *AgreementLogCreate {
	alc.mutation.SetTenantId(i)
	return alc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (alc *AgreementLogCreate) SetNillableTenantId(i *int64) *AgreementLogCreate {
	if i != nil {
		alc.SetTenantId(*i)
	}
	return alc
}

// AddOrderIDs adds the "orders" edge to the PayOrder entity by IDs.
func (alc *AgreementLogCreate) AddOrderIDs(ids ...int64) *AgreementLogCreate {
	alc.mutation.AddOrderIDs(ids...)
	return alc
}

// AddOrders adds the "orders" edges to the PayOrder entity.
func (alc *AgreementLogCreate) AddOrders(p ...*PayOrder) *AgreementLogCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return alc.AddOrderIDs(ids...)
}

// Mutation returns the AgreementLogMutation object of the builder.
func (alc *AgreementLogCreate) Mutation() *AgreementLogMutation {
	return alc.mutation
}

// Save creates the AgreementLog in the database.
func (alc *AgreementLogCreate) Save(ctx context.Context) (*AgreementLog, error) {
	var (
		err  error
		node *AgreementLog
	)
	alc.defaults()
	if len(alc.hooks) == 0 {
		if err = alc.check(); err != nil {
			return nil, err
		}
		node, err = alc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgreementLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = alc.check(); err != nil {
				return nil, err
			}
			alc.mutation = mutation
			if node, err = alc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(alc.hooks) - 1; i >= 0; i-- {
			if alc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = alc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, alc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (alc *AgreementLogCreate) SaveX(ctx context.Context) *AgreementLog {
	v, err := alc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alc *AgreementLogCreate) Exec(ctx context.Context) error {
	_, err := alc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alc *AgreementLogCreate) ExecX(ctx context.Context) {
	if err := alc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alc *AgreementLogCreate) defaults() {
	if _, ok := alc.mutation.CreatedAt(); !ok {
		v := agreementlog.DefaultCreatedAt()
		alc.mutation.SetCreatedAt(v)
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		v := agreementlog.DefaultUpdatedAt()
		alc.mutation.SetUpdatedAt(v)
	}
	if _, ok := alc.mutation.CreateBy(); !ok {
		v := agreementlog.DefaultCreateBy
		alc.mutation.SetCreateBy(v)
	}
	if _, ok := alc.mutation.UpdateBy(); !ok {
		v := agreementlog.DefaultUpdateBy
		alc.mutation.SetUpdateBy(v)
	}
	if _, ok := alc.mutation.TenantId(); !ok {
		v := agreementlog.DefaultTenantId
		alc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (alc *AgreementLogCreate) check() error {
	if v, ok := alc.mutation.AgreementType(); ok {
		if err := agreementlog.AgreementTypeValidator(v); err != nil {
			return &ValidationError{Name: "agreementType", err: fmt.Errorf(`ent: validator failed for field "agreementType": %w`, err)}
		}
	}
	if _, ok := alc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := alc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := alc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := alc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := alc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (alc *AgreementLogCreate) sqlSave(ctx context.Context) (*AgreementLog, error) {
	_node, _spec := alc.createSpec()
	if err := sqlgraph.CreateNode(ctx, alc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (alc *AgreementLogCreate) createSpec() (*AgreementLog, *sqlgraph.CreateSpec) {
	var (
		_node = &AgreementLog{config: alc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: agreementlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: agreementlog.FieldID,
			},
		}
	)
	if value, ok := alc.mutation.OuterAgreementNo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agreementlog.FieldOuterAgreementNo,
		})
		_node.OuterAgreementNo = value
	}
	if value, ok := alc.mutation.OrderId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agreementlog.FieldOrderId,
		})
		_node.OrderId = value
	}
	if value, ok := alc.mutation.UserId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldUserId,
		})
		_node.UserId = value
	}
	if value, ok := alc.mutation.ChId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldChId,
		})
		_node.ChId = value
	}
	if value, ok := alc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agreementlog.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := alc.mutation.PaymentName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agreementlog.FieldPaymentName,
		})
		_node.PaymentName = value
	}
	if value, ok := alc.mutation.PaymentId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldPaymentId,
		})
		_node.PaymentId = value
	}
	if value, ok := alc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agreementlog.FieldState,
		})
		_node.State = value
	}
	if value, ok := alc.mutation.Payment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldPayment,
		})
		_node.Payment = value
	}
	if value, ok := alc.mutation.AgreementType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: agreementlog.FieldAgreementType,
		})
		_node.AgreementType = value
	}
	if value, ok := alc.mutation.VipType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldVipType,
		})
		_node.VipType = value
	}
	if value, ok := alc.mutation.Times(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldTimes,
		})
		_node.Times = value
	}
	if value, ok := alc.mutation.CycleDays(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agreementlog.FieldCycleDays,
		})
		_node.CycleDays = value
	}
	if value, ok := alc.mutation.NextExecTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agreementlog.FieldNextExecTime,
		})
		_node.NextExecTime = value
	}
	if value, ok := alc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agreementlog.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := alc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agreementlog.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := alc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: agreementlog.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := alc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := alc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := alc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: agreementlog.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := alc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   agreementlog.OrdersTable,
			Columns: []string{agreementlog.OrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: payorder.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AgreementLogCreateBulk is the builder for creating many AgreementLog entities in bulk.
type AgreementLogCreateBulk struct {
	config
	builders []*AgreementLogCreate
}

// Save creates the AgreementLog entities in the database.
func (alcb *AgreementLogCreateBulk) Save(ctx context.Context) ([]*AgreementLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(alcb.builders))
	nodes := make([]*AgreementLog, len(alcb.builders))
	mutators := make([]Mutator, len(alcb.builders))
	for i := range alcb.builders {
		func(i int, root context.Context) {
			builder := alcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AgreementLogMutation)
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
					_, err = mutators[i+1].Mutate(root, alcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, alcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, alcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (alcb *AgreementLogCreateBulk) SaveX(ctx context.Context) []*AgreementLog {
	v, err := alcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (alcb *AgreementLogCreateBulk) Exec(ctx context.Context) error {
	_, err := alcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alcb *AgreementLogCreateBulk) ExecX(ctx context.Context) {
	if err := alcb.Exec(ctx); err != nil {
		panic(err)
	}
}
