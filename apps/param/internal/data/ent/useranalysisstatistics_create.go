// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/useranalysisstatistics"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAnalysisStatisticsCreate is the builder for creating a UserAnalysisStatistics entity.
type UserAnalysisStatisticsCreate struct {
	config
	mutation *UserAnalysisStatisticsMutation
	hooks    []Hook
}

// SetStatisticsDate sets the "statisticsDate" field.
func (uasc *UserAnalysisStatisticsCreate) SetStatisticsDate(s string) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetStatisticsDate(s)
	return uasc
}

// SetAllUserNum sets the "allUserNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetAllUserNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetAllUserNum(i)
	return uasc
}

// SetNillableAllUserNum sets the "allUserNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableAllUserNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetAllUserNum(*i)
	}
	return uasc
}

// SetAllPayment sets the "allPayment" field.
func (uasc *UserAnalysisStatisticsCreate) SetAllPayment(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetAllPayment(i)
	return uasc
}

// SetNillableAllPayment sets the "allPayment" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableAllPayment(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetAllPayment(*i)
	}
	return uasc
}

// SetAllPayUser sets the "allPayUser" field.
func (uasc *UserAnalysisStatisticsCreate) SetAllPayUser(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetAllPayUser(i)
	return uasc
}

// SetNillableAllPayUser sets the "allPayUser" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableAllPayUser(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetAllPayUser(*i)
	}
	return uasc
}

// SetAllOrderNum sets the "allOrderNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetAllOrderNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetAllOrderNum(i)
	return uasc
}

// SetNillableAllOrderNum sets the "allOrderNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableAllOrderNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetAllOrderNum(*i)
	}
	return uasc
}

// SetDayUserNum sets the "dayUserNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayUserNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayUserNum(i)
	return uasc
}

// SetNillableDayUserNum sets the "dayUserNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayUserNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayUserNum(*i)
	}
	return uasc
}

// SetDayPayment sets the "dayPayment" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayPayment(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayPayment(i)
	return uasc
}

// SetNillableDayPayment sets the "dayPayment" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayPayment(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayPayment(*i)
	}
	return uasc
}

// SetDayOrderNum sets the "dayOrderNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayOrderNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayOrderNum(i)
	return uasc
}

// SetNillableDayOrderNum sets the "dayOrderNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayOrderNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayOrderNum(*i)
	}
	return uasc
}

// SetDayPayUser sets the "dayPayUser" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayPayUser(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayPayUser(i)
	return uasc
}

// SetNillableDayPayUser sets the "dayPayUser" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayPayUser(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayPayUser(*i)
	}
	return uasc
}

// SetDayRegPayment sets the "dayRegPayment" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayRegPayment(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayRegPayment(i)
	return uasc
}

// SetNillableDayRegPayment sets the "dayRegPayment" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayRegPayment(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayRegPayment(*i)
	}
	return uasc
}

// SetDayRegUserNum sets the "dayRegUserNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayRegUserNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayRegUserNum(i)
	return uasc
}

// SetNillableDayRegUserNum sets the "dayRegUserNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayRegUserNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayRegUserNum(*i)
	}
	return uasc
}

// SetDayRegOrderNum sets the "dayRegOrderNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayRegOrderNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayRegOrderNum(i)
	return uasc
}

// SetNillableDayRegOrderNum sets the "dayRegOrderNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayRegOrderNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayRegOrderNum(*i)
	}
	return uasc
}

// SetOldRegPayment sets the "oldRegPayment" field.
func (uasc *UserAnalysisStatisticsCreate) SetOldRegPayment(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetOldRegPayment(i)
	return uasc
}

// SetNillableOldRegPayment sets the "oldRegPayment" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableOldRegPayment(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetOldRegPayment(*i)
	}
	return uasc
}

// SetOldRegUserNum sets the "oldRegUserNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetOldRegUserNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetOldRegUserNum(i)
	return uasc
}

// SetNillableOldRegUserNum sets the "oldRegUserNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableOldRegUserNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetOldRegUserNum(*i)
	}
	return uasc
}

// SetOldRegOrderNum sets the "oldRegOrderNum" field.
func (uasc *UserAnalysisStatisticsCreate) SetOldRegOrderNum(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetOldRegOrderNum(i)
	return uasc
}

// SetNillableOldRegOrderNum sets the "oldRegOrderNum" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableOldRegOrderNum(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetOldRegOrderNum(*i)
	}
	return uasc
}

// SetPayRate sets the "payRate" field.
func (uasc *UserAnalysisStatisticsCreate) SetPayRate(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetPayRate(i)
	return uasc
}

// SetNillablePayRate sets the "payRate" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillablePayRate(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetPayRate(*i)
	}
	return uasc
}

// SetArpu sets the "arpu" field.
func (uasc *UserAnalysisStatisticsCreate) SetArpu(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetArpu(i)
	return uasc
}

// SetNillableArpu sets the "arpu" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableArpu(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetArpu(*i)
	}
	return uasc
}

// SetDayRegArpu sets the "dayRegArpu" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayRegArpu(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayRegArpu(i)
	return uasc
}

// SetNillableDayRegArpu sets the "dayRegArpu" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayRegArpu(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayRegArpu(*i)
	}
	return uasc
}

// SetDayArpu sets the "dayArpu" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayArpu(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayArpu(i)
	return uasc
}

// SetNillableDayArpu sets the "dayArpu" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayArpu(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayArpu(*i)
	}
	return uasc
}

// SetDayOldArpu sets the "dayOldArpu" field.
func (uasc *UserAnalysisStatisticsCreate) SetDayOldArpu(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetDayOldArpu(i)
	return uasc
}

// SetNillableDayOldArpu sets the "dayOldArpu" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableDayOldArpu(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetDayOldArpu(*i)
	}
	return uasc
}

// SetCreatedAt sets the "createdAt" field.
func (uasc *UserAnalysisStatisticsCreate) SetCreatedAt(t time.Time) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetCreatedAt(t)
	return uasc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableCreatedAt(t *time.Time) *UserAnalysisStatisticsCreate {
	if t != nil {
		uasc.SetCreatedAt(*t)
	}
	return uasc
}

// SetUpdatedAt sets the "updatedAt" field.
func (uasc *UserAnalysisStatisticsCreate) SetUpdatedAt(t time.Time) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetUpdatedAt(t)
	return uasc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableUpdatedAt(t *time.Time) *UserAnalysisStatisticsCreate {
	if t != nil {
		uasc.SetUpdatedAt(*t)
	}
	return uasc
}

// SetCreateBy sets the "createBy" field.
func (uasc *UserAnalysisStatisticsCreate) SetCreateBy(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetCreateBy(i)
	return uasc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableCreateBy(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetCreateBy(*i)
	}
	return uasc
}

// SetUpdateBy sets the "updateBy" field.
func (uasc *UserAnalysisStatisticsCreate) SetUpdateBy(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetUpdateBy(i)
	return uasc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableUpdateBy(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetUpdateBy(*i)
	}
	return uasc
}

// SetTenantId sets the "tenantId" field.
func (uasc *UserAnalysisStatisticsCreate) SetTenantId(i int64) *UserAnalysisStatisticsCreate {
	uasc.mutation.SetTenantId(i)
	return uasc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (uasc *UserAnalysisStatisticsCreate) SetNillableTenantId(i *int64) *UserAnalysisStatisticsCreate {
	if i != nil {
		uasc.SetTenantId(*i)
	}
	return uasc
}

// Mutation returns the UserAnalysisStatisticsMutation object of the builder.
func (uasc *UserAnalysisStatisticsCreate) Mutation() *UserAnalysisStatisticsMutation {
	return uasc.mutation
}

// Save creates the UserAnalysisStatistics in the database.
func (uasc *UserAnalysisStatisticsCreate) Save(ctx context.Context) (*UserAnalysisStatistics, error) {
	var (
		err  error
		node *UserAnalysisStatistics
	)
	uasc.defaults()
	if len(uasc.hooks) == 0 {
		if err = uasc.check(); err != nil {
			return nil, err
		}
		node, err = uasc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserAnalysisStatisticsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uasc.check(); err != nil {
				return nil, err
			}
			uasc.mutation = mutation
			if node, err = uasc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uasc.hooks) - 1; i >= 0; i-- {
			if uasc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uasc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uasc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uasc *UserAnalysisStatisticsCreate) SaveX(ctx context.Context) *UserAnalysisStatistics {
	v, err := uasc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uasc *UserAnalysisStatisticsCreate) Exec(ctx context.Context) error {
	_, err := uasc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uasc *UserAnalysisStatisticsCreate) ExecX(ctx context.Context) {
	if err := uasc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uasc *UserAnalysisStatisticsCreate) defaults() {
	if _, ok := uasc.mutation.CreatedAt(); !ok {
		v := useranalysisstatistics.DefaultCreatedAt()
		uasc.mutation.SetCreatedAt(v)
	}
	if _, ok := uasc.mutation.UpdatedAt(); !ok {
		v := useranalysisstatistics.DefaultUpdatedAt()
		uasc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uasc.mutation.CreateBy(); !ok {
		v := useranalysisstatistics.DefaultCreateBy
		uasc.mutation.SetCreateBy(v)
	}
	if _, ok := uasc.mutation.UpdateBy(); !ok {
		v := useranalysisstatistics.DefaultUpdateBy
		uasc.mutation.SetUpdateBy(v)
	}
	if _, ok := uasc.mutation.TenantId(); !ok {
		v := useranalysisstatistics.DefaultTenantId
		uasc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uasc *UserAnalysisStatisticsCreate) check() error {
	if _, ok := uasc.mutation.StatisticsDate(); !ok {
		return &ValidationError{Name: "statisticsDate", err: errors.New(`ent: missing required field "statisticsDate"`)}
	}
	if _, ok := uasc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := uasc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := uasc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := uasc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := uasc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	return nil
}

func (uasc *UserAnalysisStatisticsCreate) sqlSave(ctx context.Context) (*UserAnalysisStatistics, error) {
	_node, _spec := uasc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uasc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (uasc *UserAnalysisStatisticsCreate) createSpec() (*UserAnalysisStatistics, *sqlgraph.CreateSpec) {
	var (
		_node = &UserAnalysisStatistics{config: uasc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: useranalysisstatistics.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: useranalysisstatistics.FieldID,
			},
		}
	)
	if value, ok := uasc.mutation.StatisticsDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useranalysisstatistics.FieldStatisticsDate,
		})
		_node.StatisticsDate = value
	}
	if value, ok := uasc.mutation.AllUserNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldAllUserNum,
		})
		_node.AllUserNum = value
	}
	if value, ok := uasc.mutation.AllPayment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldAllPayment,
		})
		_node.AllPayment = value
	}
	if value, ok := uasc.mutation.AllPayUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldAllPayUser,
		})
		_node.AllPayUser = value
	}
	if value, ok := uasc.mutation.AllOrderNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldAllOrderNum,
		})
		_node.AllOrderNum = value
	}
	if value, ok := uasc.mutation.DayUserNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayUserNum,
		})
		_node.DayUserNum = value
	}
	if value, ok := uasc.mutation.DayPayment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayPayment,
		})
		_node.DayPayment = value
	}
	if value, ok := uasc.mutation.DayOrderNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayOrderNum,
		})
		_node.DayOrderNum = value
	}
	if value, ok := uasc.mutation.DayPayUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayPayUser,
		})
		_node.DayPayUser = value
	}
	if value, ok := uasc.mutation.DayRegPayment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayRegPayment,
		})
		_node.DayRegPayment = value
	}
	if value, ok := uasc.mutation.DayRegUserNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayRegUserNum,
		})
		_node.DayRegUserNum = value
	}
	if value, ok := uasc.mutation.DayRegOrderNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayRegOrderNum,
		})
		_node.DayRegOrderNum = value
	}
	if value, ok := uasc.mutation.OldRegPayment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldOldRegPayment,
		})
		_node.OldRegPayment = value
	}
	if value, ok := uasc.mutation.OldRegUserNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldOldRegUserNum,
		})
		_node.OldRegUserNum = value
	}
	if value, ok := uasc.mutation.OldRegOrderNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldOldRegOrderNum,
		})
		_node.OldRegOrderNum = value
	}
	if value, ok := uasc.mutation.PayRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldPayRate,
		})
		_node.PayRate = value
	}
	if value, ok := uasc.mutation.Arpu(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldArpu,
		})
		_node.Arpu = value
	}
	if value, ok := uasc.mutation.DayRegArpu(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayRegArpu,
		})
		_node.DayRegArpu = value
	}
	if value, ok := uasc.mutation.DayArpu(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayArpu,
		})
		_node.DayArpu = value
	}
	if value, ok := uasc.mutation.DayOldArpu(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldDayOldArpu,
		})
		_node.DayOldArpu = value
	}
	if value, ok := uasc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useranalysisstatistics.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uasc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useranalysisstatistics.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uasc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := uasc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := uasc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: useranalysisstatistics.FieldTenantId,
		})
		_node.TenantId = value
	}
	return _node, _spec
}

// UserAnalysisStatisticsCreateBulk is the builder for creating many UserAnalysisStatistics entities in bulk.
type UserAnalysisStatisticsCreateBulk struct {
	config
	builders []*UserAnalysisStatisticsCreate
}

// Save creates the UserAnalysisStatistics entities in the database.
func (uascb *UserAnalysisStatisticsCreateBulk) Save(ctx context.Context) ([]*UserAnalysisStatistics, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uascb.builders))
	nodes := make([]*UserAnalysisStatistics, len(uascb.builders))
	mutators := make([]Mutator, len(uascb.builders))
	for i := range uascb.builders {
		func(i int, root context.Context) {
			builder := uascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserAnalysisStatisticsMutation)
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
					_, err = mutators[i+1].Mutate(root, uascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uascb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uascb *UserAnalysisStatisticsCreateBulk) SaveX(ctx context.Context) []*UserAnalysisStatistics {
	v, err := uascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uascb *UserAnalysisStatisticsCreateBulk) Exec(ctx context.Context) error {
	_, err := uascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uascb *UserAnalysisStatisticsCreateBulk) ExecX(ctx context.Context) {
	if err := uascb.Exec(ctx); err != nil {
		panic(err)
	}
}