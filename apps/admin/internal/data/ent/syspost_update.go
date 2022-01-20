// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/syspost"
	"hope/apps/admin/internal/data/ent/sysuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysPostUpdate is the builder for updating SysPost entities.
type SysPostUpdate struct {
	config
	hooks    []Hook
	mutation *SysPostMutation
}

// Where appends a list predicates to the SysPostUpdate builder.
func (spu *SysPostUpdate) Where(ps ...predicate.SysPost) *SysPostUpdate {
	spu.mutation.Where(ps...)
	return spu
}

// SetPostName sets the "postName" field.
func (spu *SysPostUpdate) SetPostName(s string) *SysPostUpdate {
	spu.mutation.SetPostName(s)
	return spu
}

// SetNillablePostName sets the "postName" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillablePostName(s *string) *SysPostUpdate {
	if s != nil {
		spu.SetPostName(*s)
	}
	return spu
}

// ClearPostName clears the value of the "postName" field.
func (spu *SysPostUpdate) ClearPostName() *SysPostUpdate {
	spu.mutation.ClearPostName()
	return spu
}

// SetPostCode sets the "postCode" field.
func (spu *SysPostUpdate) SetPostCode(s string) *SysPostUpdate {
	spu.mutation.SetPostCode(s)
	return spu
}

// SetNillablePostCode sets the "postCode" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillablePostCode(s *string) *SysPostUpdate {
	if s != nil {
		spu.SetPostCode(*s)
	}
	return spu
}

// ClearPostCode clears the value of the "postCode" field.
func (spu *SysPostUpdate) ClearPostCode() *SysPostUpdate {
	spu.mutation.ClearPostCode()
	return spu
}

// SetSort sets the "sort" field.
func (spu *SysPostUpdate) SetSort(i int32) *SysPostUpdate {
	spu.mutation.ResetSort()
	spu.mutation.SetSort(i)
	return spu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillableSort(i *int32) *SysPostUpdate {
	if i != nil {
		spu.SetSort(*i)
	}
	return spu
}

// AddSort adds i to the "sort" field.
func (spu *SysPostUpdate) AddSort(i int32) *SysPostUpdate {
	spu.mutation.AddSort(i)
	return spu
}

// ClearSort clears the value of the "sort" field.
func (spu *SysPostUpdate) ClearSort() *SysPostUpdate {
	spu.mutation.ClearSort()
	return spu
}

// SetStatus sets the "status" field.
func (spu *SysPostUpdate) SetStatus(i int32) *SysPostUpdate {
	spu.mutation.ResetStatus()
	spu.mutation.SetStatus(i)
	return spu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillableStatus(i *int32) *SysPostUpdate {
	if i != nil {
		spu.SetStatus(*i)
	}
	return spu
}

// AddStatus adds i to the "status" field.
func (spu *SysPostUpdate) AddStatus(i int32) *SysPostUpdate {
	spu.mutation.AddStatus(i)
	return spu
}

// ClearStatus clears the value of the "status" field.
func (spu *SysPostUpdate) ClearStatus() *SysPostUpdate {
	spu.mutation.ClearStatus()
	return spu
}

// SetRemark sets the "remark" field.
func (spu *SysPostUpdate) SetRemark(s string) *SysPostUpdate {
	spu.mutation.SetRemark(s)
	return spu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillableRemark(s *string) *SysPostUpdate {
	if s != nil {
		spu.SetRemark(*s)
	}
	return spu
}

// ClearRemark clears the value of the "remark" field.
func (spu *SysPostUpdate) ClearRemark() *SysPostUpdate {
	spu.mutation.ClearRemark()
	return spu
}

// SetUpdatedAt sets the "updatedAt" field.
func (spu *SysPostUpdate) SetUpdatedAt(t time.Time) *SysPostUpdate {
	spu.mutation.SetUpdatedAt(t)
	return spu
}

// SetCreateBy sets the "createBy" field.
func (spu *SysPostUpdate) SetCreateBy(i int64) *SysPostUpdate {
	spu.mutation.ResetCreateBy()
	spu.mutation.SetCreateBy(i)
	return spu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillableCreateBy(i *int64) *SysPostUpdate {
	if i != nil {
		spu.SetCreateBy(*i)
	}
	return spu
}

// AddCreateBy adds i to the "createBy" field.
func (spu *SysPostUpdate) AddCreateBy(i int64) *SysPostUpdate {
	spu.mutation.AddCreateBy(i)
	return spu
}

// SetUpdateBy sets the "updateBy" field.
func (spu *SysPostUpdate) SetUpdateBy(i int64) *SysPostUpdate {
	spu.mutation.ResetUpdateBy()
	spu.mutation.SetUpdateBy(i)
	return spu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillableUpdateBy(i *int64) *SysPostUpdate {
	if i != nil {
		spu.SetUpdateBy(*i)
	}
	return spu
}

// AddUpdateBy adds i to the "updateBy" field.
func (spu *SysPostUpdate) AddUpdateBy(i int64) *SysPostUpdate {
	spu.mutation.AddUpdateBy(i)
	return spu
}

// SetTenantId sets the "tenantId" field.
func (spu *SysPostUpdate) SetTenantId(i int64) *SysPostUpdate {
	spu.mutation.ResetTenantId()
	spu.mutation.SetTenantId(i)
	return spu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (spu *SysPostUpdate) SetNillableTenantId(i *int64) *SysPostUpdate {
	if i != nil {
		spu.SetTenantId(*i)
	}
	return spu
}

// AddTenantId adds i to the "tenantId" field.
func (spu *SysPostUpdate) AddTenantId(i int64) *SysPostUpdate {
	spu.mutation.AddTenantId(i)
	return spu
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (spu *SysPostUpdate) AddUserIDs(ids ...int64) *SysPostUpdate {
	spu.mutation.AddUserIDs(ids...)
	return spu
}

// AddUsers adds the "users" edges to the SysUser entity.
func (spu *SysPostUpdate) AddUsers(s ...*SysUser) *SysPostUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return spu.AddUserIDs(ids...)
}

// Mutation returns the SysPostMutation object of the builder.
func (spu *SysPostUpdate) Mutation() *SysPostMutation {
	return spu.mutation
}

// ClearUsers clears all "users" edges to the SysUser entity.
func (spu *SysPostUpdate) ClearUsers() *SysPostUpdate {
	spu.mutation.ClearUsers()
	return spu
}

// RemoveUserIDs removes the "users" edge to SysUser entities by IDs.
func (spu *SysPostUpdate) RemoveUserIDs(ids ...int64) *SysPostUpdate {
	spu.mutation.RemoveUserIDs(ids...)
	return spu
}

// RemoveUsers removes "users" edges to SysUser entities.
func (spu *SysPostUpdate) RemoveUsers(s ...*SysUser) *SysPostUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return spu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spu *SysPostUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	spu.defaults()
	if len(spu.hooks) == 0 {
		affected, err = spu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysPostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spu.mutation = mutation
			affected, err = spu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spu.hooks) - 1; i >= 0; i-- {
			if spu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (spu *SysPostUpdate) SaveX(ctx context.Context) int {
	affected, err := spu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spu *SysPostUpdate) Exec(ctx context.Context) error {
	_, err := spu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spu *SysPostUpdate) ExecX(ctx context.Context) {
	if err := spu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spu *SysPostUpdate) defaults() {
	if _, ok := spu.mutation.UpdatedAt(); !ok {
		v := syspost.UpdateDefaultUpdatedAt()
		spu.mutation.SetUpdatedAt(v)
	}
}

func (spu *SysPostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   syspost.Table,
			Columns: syspost.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: syspost.FieldID,
			},
		},
	}
	if ps := spu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spu.mutation.PostName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldPostName,
		})
	}
	if spu.mutation.PostNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: syspost.FieldPostName,
		})
	}
	if value, ok := spu.mutation.PostCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldPostCode,
		})
	}
	if spu.mutation.PostCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: syspost.FieldPostCode,
		})
	}
	if value, ok := spu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldSort,
		})
	}
	if value, ok := spu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldSort,
		})
	}
	if spu.mutation.SortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: syspost.FieldSort,
		})
	}
	if value, ok := spu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldStatus,
		})
	}
	if value, ok := spu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldStatus,
		})
	}
	if spu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: syspost.FieldStatus,
		})
	}
	if value, ok := spu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldRemark,
		})
	}
	if spu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: syspost.FieldRemark,
		})
	}
	if value, ok := spu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: syspost.FieldUpdatedAt,
		})
	}
	if value, ok := spu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldCreateBy,
		})
	}
	if value, ok := spu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldCreateBy,
		})
	}
	if value, ok := spu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldUpdateBy,
		})
	}
	if value, ok := spu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldUpdateBy,
		})
	}
	if value, ok := spu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldTenantId,
		})
	}
	if value, ok := spu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldTenantId,
		})
	}
	if spu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   syspost.UsersTable,
			Columns: []string{syspost.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !spu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   syspost.UsersTable,
			Columns: []string{syspost.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   syspost.UsersTable,
			Columns: []string{syspost.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syspost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SysPostUpdateOne is the builder for updating a single SysPost entity.
type SysPostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysPostMutation
}

// SetPostName sets the "postName" field.
func (spuo *SysPostUpdateOne) SetPostName(s string) *SysPostUpdateOne {
	spuo.mutation.SetPostName(s)
	return spuo
}

// SetNillablePostName sets the "postName" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillablePostName(s *string) *SysPostUpdateOne {
	if s != nil {
		spuo.SetPostName(*s)
	}
	return spuo
}

// ClearPostName clears the value of the "postName" field.
func (spuo *SysPostUpdateOne) ClearPostName() *SysPostUpdateOne {
	spuo.mutation.ClearPostName()
	return spuo
}

// SetPostCode sets the "postCode" field.
func (spuo *SysPostUpdateOne) SetPostCode(s string) *SysPostUpdateOne {
	spuo.mutation.SetPostCode(s)
	return spuo
}

// SetNillablePostCode sets the "postCode" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillablePostCode(s *string) *SysPostUpdateOne {
	if s != nil {
		spuo.SetPostCode(*s)
	}
	return spuo
}

// ClearPostCode clears the value of the "postCode" field.
func (spuo *SysPostUpdateOne) ClearPostCode() *SysPostUpdateOne {
	spuo.mutation.ClearPostCode()
	return spuo
}

// SetSort sets the "sort" field.
func (spuo *SysPostUpdateOne) SetSort(i int32) *SysPostUpdateOne {
	spuo.mutation.ResetSort()
	spuo.mutation.SetSort(i)
	return spuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillableSort(i *int32) *SysPostUpdateOne {
	if i != nil {
		spuo.SetSort(*i)
	}
	return spuo
}

// AddSort adds i to the "sort" field.
func (spuo *SysPostUpdateOne) AddSort(i int32) *SysPostUpdateOne {
	spuo.mutation.AddSort(i)
	return spuo
}

// ClearSort clears the value of the "sort" field.
func (spuo *SysPostUpdateOne) ClearSort() *SysPostUpdateOne {
	spuo.mutation.ClearSort()
	return spuo
}

// SetStatus sets the "status" field.
func (spuo *SysPostUpdateOne) SetStatus(i int32) *SysPostUpdateOne {
	spuo.mutation.ResetStatus()
	spuo.mutation.SetStatus(i)
	return spuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillableStatus(i *int32) *SysPostUpdateOne {
	if i != nil {
		spuo.SetStatus(*i)
	}
	return spuo
}

// AddStatus adds i to the "status" field.
func (spuo *SysPostUpdateOne) AddStatus(i int32) *SysPostUpdateOne {
	spuo.mutation.AddStatus(i)
	return spuo
}

// ClearStatus clears the value of the "status" field.
func (spuo *SysPostUpdateOne) ClearStatus() *SysPostUpdateOne {
	spuo.mutation.ClearStatus()
	return spuo
}

// SetRemark sets the "remark" field.
func (spuo *SysPostUpdateOne) SetRemark(s string) *SysPostUpdateOne {
	spuo.mutation.SetRemark(s)
	return spuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillableRemark(s *string) *SysPostUpdateOne {
	if s != nil {
		spuo.SetRemark(*s)
	}
	return spuo
}

// ClearRemark clears the value of the "remark" field.
func (spuo *SysPostUpdateOne) ClearRemark() *SysPostUpdateOne {
	spuo.mutation.ClearRemark()
	return spuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (spuo *SysPostUpdateOne) SetUpdatedAt(t time.Time) *SysPostUpdateOne {
	spuo.mutation.SetUpdatedAt(t)
	return spuo
}

// SetCreateBy sets the "createBy" field.
func (spuo *SysPostUpdateOne) SetCreateBy(i int64) *SysPostUpdateOne {
	spuo.mutation.ResetCreateBy()
	spuo.mutation.SetCreateBy(i)
	return spuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillableCreateBy(i *int64) *SysPostUpdateOne {
	if i != nil {
		spuo.SetCreateBy(*i)
	}
	return spuo
}

// AddCreateBy adds i to the "createBy" field.
func (spuo *SysPostUpdateOne) AddCreateBy(i int64) *SysPostUpdateOne {
	spuo.mutation.AddCreateBy(i)
	return spuo
}

// SetUpdateBy sets the "updateBy" field.
func (spuo *SysPostUpdateOne) SetUpdateBy(i int64) *SysPostUpdateOne {
	spuo.mutation.ResetUpdateBy()
	spuo.mutation.SetUpdateBy(i)
	return spuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillableUpdateBy(i *int64) *SysPostUpdateOne {
	if i != nil {
		spuo.SetUpdateBy(*i)
	}
	return spuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (spuo *SysPostUpdateOne) AddUpdateBy(i int64) *SysPostUpdateOne {
	spuo.mutation.AddUpdateBy(i)
	return spuo
}

// SetTenantId sets the "tenantId" field.
func (spuo *SysPostUpdateOne) SetTenantId(i int64) *SysPostUpdateOne {
	spuo.mutation.ResetTenantId()
	spuo.mutation.SetTenantId(i)
	return spuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (spuo *SysPostUpdateOne) SetNillableTenantId(i *int64) *SysPostUpdateOne {
	if i != nil {
		spuo.SetTenantId(*i)
	}
	return spuo
}

// AddTenantId adds i to the "tenantId" field.
func (spuo *SysPostUpdateOne) AddTenantId(i int64) *SysPostUpdateOne {
	spuo.mutation.AddTenantId(i)
	return spuo
}

// AddUserIDs adds the "users" edge to the SysUser entity by IDs.
func (spuo *SysPostUpdateOne) AddUserIDs(ids ...int64) *SysPostUpdateOne {
	spuo.mutation.AddUserIDs(ids...)
	return spuo
}

// AddUsers adds the "users" edges to the SysUser entity.
func (spuo *SysPostUpdateOne) AddUsers(s ...*SysUser) *SysPostUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return spuo.AddUserIDs(ids...)
}

// Mutation returns the SysPostMutation object of the builder.
func (spuo *SysPostUpdateOne) Mutation() *SysPostMutation {
	return spuo.mutation
}

// ClearUsers clears all "users" edges to the SysUser entity.
func (spuo *SysPostUpdateOne) ClearUsers() *SysPostUpdateOne {
	spuo.mutation.ClearUsers()
	return spuo
}

// RemoveUserIDs removes the "users" edge to SysUser entities by IDs.
func (spuo *SysPostUpdateOne) RemoveUserIDs(ids ...int64) *SysPostUpdateOne {
	spuo.mutation.RemoveUserIDs(ids...)
	return spuo
}

// RemoveUsers removes "users" edges to SysUser entities.
func (spuo *SysPostUpdateOne) RemoveUsers(s ...*SysUser) *SysPostUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return spuo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spuo *SysPostUpdateOne) Select(field string, fields ...string) *SysPostUpdateOne {
	spuo.fields = append([]string{field}, fields...)
	return spuo
}

// Save executes the query and returns the updated SysPost entity.
func (spuo *SysPostUpdateOne) Save(ctx context.Context) (*SysPost, error) {
	var (
		err  error
		node *SysPost
	)
	spuo.defaults()
	if len(spuo.hooks) == 0 {
		node, err = spuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysPostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spuo.mutation = mutation
			node, err = spuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(spuo.hooks) - 1; i >= 0; i-- {
			if spuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (spuo *SysPostUpdateOne) SaveX(ctx context.Context) *SysPost {
	node, err := spuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spuo *SysPostUpdateOne) Exec(ctx context.Context) error {
	_, err := spuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spuo *SysPostUpdateOne) ExecX(ctx context.Context) {
	if err := spuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (spuo *SysPostUpdateOne) defaults() {
	if _, ok := spuo.mutation.UpdatedAt(); !ok {
		v := syspost.UpdateDefaultUpdatedAt()
		spuo.mutation.SetUpdatedAt(v)
	}
}

func (spuo *SysPostUpdateOne) sqlSave(ctx context.Context) (_node *SysPost, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   syspost.Table,
			Columns: syspost.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: syspost.FieldID,
			},
		},
	}
	id, ok := spuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysPost.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := spuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, syspost.FieldID)
		for _, f := range fields {
			if !syspost.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != syspost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spuo.mutation.PostName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldPostName,
		})
	}
	if spuo.mutation.PostNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: syspost.FieldPostName,
		})
	}
	if value, ok := spuo.mutation.PostCode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldPostCode,
		})
	}
	if spuo.mutation.PostCodeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: syspost.FieldPostCode,
		})
	}
	if value, ok := spuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldSort,
		})
	}
	if value, ok := spuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldSort,
		})
	}
	if spuo.mutation.SortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: syspost.FieldSort,
		})
	}
	if value, ok := spuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldStatus,
		})
	}
	if value, ok := spuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: syspost.FieldStatus,
		})
	}
	if spuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: syspost.FieldStatus,
		})
	}
	if value, ok := spuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: syspost.FieldRemark,
		})
	}
	if spuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: syspost.FieldRemark,
		})
	}
	if value, ok := spuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: syspost.FieldUpdatedAt,
		})
	}
	if value, ok := spuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldCreateBy,
		})
	}
	if value, ok := spuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldCreateBy,
		})
	}
	if value, ok := spuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldUpdateBy,
		})
	}
	if value, ok := spuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldUpdateBy,
		})
	}
	if value, ok := spuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldTenantId,
		})
	}
	if value, ok := spuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: syspost.FieldTenantId,
		})
	}
	if spuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   syspost.UsersTable,
			Columns: []string{syspost.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !spuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   syspost.UsersTable,
			Columns: []string{syspost.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   syspost.UsersTable,
			Columns: []string{syspost.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: sysuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SysPost{config: spuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{syspost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
