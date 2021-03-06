// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelmsg"
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelMsgUpdate is the builder for updating NovelMsg entities.
type NovelMsgUpdate struct {
	config
	hooks    []Hook
	mutation *NovelMsgMutation
}

// Where appends a list predicates to the NovelMsgUpdate builder.
func (nmu *NovelMsgUpdate) Where(ps ...predicate.NovelMsg) *NovelMsgUpdate {
	nmu.mutation.Where(ps...)
	return nmu
}

// SetTitle sets the "title" field.
func (nmu *NovelMsgUpdate) SetTitle(s string) *NovelMsgUpdate {
	nmu.mutation.SetTitle(s)
	return nmu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nmu *NovelMsgUpdate) SetNillableTitle(s *string) *NovelMsgUpdate {
	if s != nil {
		nmu.SetTitle(*s)
	}
	return nmu
}

// ClearTitle clears the value of the "title" field.
func (nmu *NovelMsgUpdate) ClearTitle() *NovelMsgUpdate {
	nmu.mutation.ClearTitle()
	return nmu
}

// SetMsg sets the "msg" field.
func (nmu *NovelMsgUpdate) SetMsg(s string) *NovelMsgUpdate {
	nmu.mutation.SetMsg(s)
	return nmu
}

// SetNillableMsg sets the "msg" field if the given value is not nil.
func (nmu *NovelMsgUpdate) SetNillableMsg(s *string) *NovelMsgUpdate {
	if s != nil {
		nmu.SetMsg(*s)
	}
	return nmu
}

// ClearMsg clears the value of the "msg" field.
func (nmu *NovelMsgUpdate) ClearMsg() *NovelMsgUpdate {
	nmu.mutation.ClearMsg()
	return nmu
}

// SetMsgType sets the "msgType" field.
func (nmu *NovelMsgUpdate) SetMsgType(s string) *NovelMsgUpdate {
	nmu.mutation.SetMsgType(s)
	return nmu
}

// SetNillableMsgType sets the "msgType" field if the given value is not nil.
func (nmu *NovelMsgUpdate) SetNillableMsgType(s *string) *NovelMsgUpdate {
	if s != nil {
		nmu.SetMsgType(*s)
	}
	return nmu
}

// ClearMsgType clears the value of the "msgType" field.
func (nmu *NovelMsgUpdate) ClearMsgType() *NovelMsgUpdate {
	nmu.mutation.ClearMsgType()
	return nmu
}

// SetStatus sets the "status" field.
func (nmu *NovelMsgUpdate) SetStatus(b bool) *NovelMsgUpdate {
	nmu.mutation.SetStatus(b)
	return nmu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nmu *NovelMsgUpdate) SetNillableStatus(b *bool) *NovelMsgUpdate {
	if b != nil {
		nmu.SetStatus(*b)
	}
	return nmu
}

// ClearStatus clears the value of the "status" field.
func (nmu *NovelMsgUpdate) ClearStatus() *NovelMsgUpdate {
	nmu.mutation.ClearStatus()
	return nmu
}

// SetUpdatedAt sets the "updatedAt" field.
func (nmu *NovelMsgUpdate) SetUpdatedAt(t time.Time) *NovelMsgUpdate {
	nmu.mutation.SetUpdatedAt(t)
	return nmu
}

// SetCreateBy sets the "createBy" field.
func (nmu *NovelMsgUpdate) SetCreateBy(i int64) *NovelMsgUpdate {
	nmu.mutation.ResetCreateBy()
	nmu.mutation.SetCreateBy(i)
	return nmu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nmu *NovelMsgUpdate) SetNillableCreateBy(i *int64) *NovelMsgUpdate {
	if i != nil {
		nmu.SetCreateBy(*i)
	}
	return nmu
}

// AddCreateBy adds i to the "createBy" field.
func (nmu *NovelMsgUpdate) AddCreateBy(i int64) *NovelMsgUpdate {
	nmu.mutation.AddCreateBy(i)
	return nmu
}

// SetUpdateBy sets the "updateBy" field.
func (nmu *NovelMsgUpdate) SetUpdateBy(i int64) *NovelMsgUpdate {
	nmu.mutation.ResetUpdateBy()
	nmu.mutation.SetUpdateBy(i)
	return nmu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nmu *NovelMsgUpdate) SetNillableUpdateBy(i *int64) *NovelMsgUpdate {
	if i != nil {
		nmu.SetUpdateBy(*i)
	}
	return nmu
}

// AddUpdateBy adds i to the "updateBy" field.
func (nmu *NovelMsgUpdate) AddUpdateBy(i int64) *NovelMsgUpdate {
	nmu.mutation.AddUpdateBy(i)
	return nmu
}

// SetTenantId sets the "tenantId" field.
func (nmu *NovelMsgUpdate) SetTenantId(i int64) *NovelMsgUpdate {
	nmu.mutation.ResetTenantId()
	nmu.mutation.SetTenantId(i)
	return nmu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nmu *NovelMsgUpdate) SetNillableTenantId(i *int64) *NovelMsgUpdate {
	if i != nil {
		nmu.SetTenantId(*i)
	}
	return nmu
}

// AddTenantId adds i to the "tenantId" field.
func (nmu *NovelMsgUpdate) AddTenantId(i int64) *NovelMsgUpdate {
	nmu.mutation.AddTenantId(i)
	return nmu
}

// Mutation returns the NovelMsgMutation object of the builder.
func (nmu *NovelMsgUpdate) Mutation() *NovelMsgMutation {
	return nmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nmu *NovelMsgUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	nmu.defaults()
	if len(nmu.hooks) == 0 {
		affected, err = nmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelMsgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nmu.mutation = mutation
			affected, err = nmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nmu.hooks) - 1; i >= 0; i-- {
			if nmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nmu *NovelMsgUpdate) SaveX(ctx context.Context) int {
	affected, err := nmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nmu *NovelMsgUpdate) Exec(ctx context.Context) error {
	_, err := nmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmu *NovelMsgUpdate) ExecX(ctx context.Context) {
	if err := nmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nmu *NovelMsgUpdate) defaults() {
	if _, ok := nmu.mutation.UpdatedAt(); !ok {
		v := novelmsg.UpdateDefaultUpdatedAt()
		nmu.mutation.SetUpdatedAt(v)
	}
}

func (nmu *NovelMsgUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelmsg.Table,
			Columns: novelmsg.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelmsg.FieldID,
			},
		},
	}
	if ps := nmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nmu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelmsg.FieldTitle,
		})
	}
	if nmu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelmsg.FieldTitle,
		})
	}
	if value, ok := nmu.mutation.Msg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelmsg.FieldMsg,
		})
	}
	if nmu.mutation.MsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelmsg.FieldMsg,
		})
	}
	if value, ok := nmu.mutation.MsgType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelmsg.FieldMsgType,
		})
	}
	if nmu.mutation.MsgTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelmsg.FieldMsgType,
		})
	}
	if value, ok := nmu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: novelmsg.FieldStatus,
		})
	}
	if nmu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: novelmsg.FieldStatus,
		})
	}
	if value, ok := nmu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelmsg.FieldUpdatedAt,
		})
	}
	if value, ok := nmu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldCreateBy,
		})
	}
	if value, ok := nmu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldCreateBy,
		})
	}
	if value, ok := nmu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldUpdateBy,
		})
	}
	if value, ok := nmu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldUpdateBy,
		})
	}
	if value, ok := nmu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldTenantId,
		})
	}
	if value, ok := nmu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldTenantId,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NovelMsgUpdateOne is the builder for updating a single NovelMsg entity.
type NovelMsgUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NovelMsgMutation
}

// SetTitle sets the "title" field.
func (nmuo *NovelMsgUpdateOne) SetTitle(s string) *NovelMsgUpdateOne {
	nmuo.mutation.SetTitle(s)
	return nmuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nmuo *NovelMsgUpdateOne) SetNillableTitle(s *string) *NovelMsgUpdateOne {
	if s != nil {
		nmuo.SetTitle(*s)
	}
	return nmuo
}

// ClearTitle clears the value of the "title" field.
func (nmuo *NovelMsgUpdateOne) ClearTitle() *NovelMsgUpdateOne {
	nmuo.mutation.ClearTitle()
	return nmuo
}

// SetMsg sets the "msg" field.
func (nmuo *NovelMsgUpdateOne) SetMsg(s string) *NovelMsgUpdateOne {
	nmuo.mutation.SetMsg(s)
	return nmuo
}

// SetNillableMsg sets the "msg" field if the given value is not nil.
func (nmuo *NovelMsgUpdateOne) SetNillableMsg(s *string) *NovelMsgUpdateOne {
	if s != nil {
		nmuo.SetMsg(*s)
	}
	return nmuo
}

// ClearMsg clears the value of the "msg" field.
func (nmuo *NovelMsgUpdateOne) ClearMsg() *NovelMsgUpdateOne {
	nmuo.mutation.ClearMsg()
	return nmuo
}

// SetMsgType sets the "msgType" field.
func (nmuo *NovelMsgUpdateOne) SetMsgType(s string) *NovelMsgUpdateOne {
	nmuo.mutation.SetMsgType(s)
	return nmuo
}

// SetNillableMsgType sets the "msgType" field if the given value is not nil.
func (nmuo *NovelMsgUpdateOne) SetNillableMsgType(s *string) *NovelMsgUpdateOne {
	if s != nil {
		nmuo.SetMsgType(*s)
	}
	return nmuo
}

// ClearMsgType clears the value of the "msgType" field.
func (nmuo *NovelMsgUpdateOne) ClearMsgType() *NovelMsgUpdateOne {
	nmuo.mutation.ClearMsgType()
	return nmuo
}

// SetStatus sets the "status" field.
func (nmuo *NovelMsgUpdateOne) SetStatus(b bool) *NovelMsgUpdateOne {
	nmuo.mutation.SetStatus(b)
	return nmuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (nmuo *NovelMsgUpdateOne) SetNillableStatus(b *bool) *NovelMsgUpdateOne {
	if b != nil {
		nmuo.SetStatus(*b)
	}
	return nmuo
}

// ClearStatus clears the value of the "status" field.
func (nmuo *NovelMsgUpdateOne) ClearStatus() *NovelMsgUpdateOne {
	nmuo.mutation.ClearStatus()
	return nmuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (nmuo *NovelMsgUpdateOne) SetUpdatedAt(t time.Time) *NovelMsgUpdateOne {
	nmuo.mutation.SetUpdatedAt(t)
	return nmuo
}

// SetCreateBy sets the "createBy" field.
func (nmuo *NovelMsgUpdateOne) SetCreateBy(i int64) *NovelMsgUpdateOne {
	nmuo.mutation.ResetCreateBy()
	nmuo.mutation.SetCreateBy(i)
	return nmuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nmuo *NovelMsgUpdateOne) SetNillableCreateBy(i *int64) *NovelMsgUpdateOne {
	if i != nil {
		nmuo.SetCreateBy(*i)
	}
	return nmuo
}

// AddCreateBy adds i to the "createBy" field.
func (nmuo *NovelMsgUpdateOne) AddCreateBy(i int64) *NovelMsgUpdateOne {
	nmuo.mutation.AddCreateBy(i)
	return nmuo
}

// SetUpdateBy sets the "updateBy" field.
func (nmuo *NovelMsgUpdateOne) SetUpdateBy(i int64) *NovelMsgUpdateOne {
	nmuo.mutation.ResetUpdateBy()
	nmuo.mutation.SetUpdateBy(i)
	return nmuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nmuo *NovelMsgUpdateOne) SetNillableUpdateBy(i *int64) *NovelMsgUpdateOne {
	if i != nil {
		nmuo.SetUpdateBy(*i)
	}
	return nmuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (nmuo *NovelMsgUpdateOne) AddUpdateBy(i int64) *NovelMsgUpdateOne {
	nmuo.mutation.AddUpdateBy(i)
	return nmuo
}

// SetTenantId sets the "tenantId" field.
func (nmuo *NovelMsgUpdateOne) SetTenantId(i int64) *NovelMsgUpdateOne {
	nmuo.mutation.ResetTenantId()
	nmuo.mutation.SetTenantId(i)
	return nmuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nmuo *NovelMsgUpdateOne) SetNillableTenantId(i *int64) *NovelMsgUpdateOne {
	if i != nil {
		nmuo.SetTenantId(*i)
	}
	return nmuo
}

// AddTenantId adds i to the "tenantId" field.
func (nmuo *NovelMsgUpdateOne) AddTenantId(i int64) *NovelMsgUpdateOne {
	nmuo.mutation.AddTenantId(i)
	return nmuo
}

// Mutation returns the NovelMsgMutation object of the builder.
func (nmuo *NovelMsgUpdateOne) Mutation() *NovelMsgMutation {
	return nmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nmuo *NovelMsgUpdateOne) Select(field string, fields ...string) *NovelMsgUpdateOne {
	nmuo.fields = append([]string{field}, fields...)
	return nmuo
}

// Save executes the query and returns the updated NovelMsg entity.
func (nmuo *NovelMsgUpdateOne) Save(ctx context.Context) (*NovelMsg, error) {
	var (
		err  error
		node *NovelMsg
	)
	nmuo.defaults()
	if len(nmuo.hooks) == 0 {
		node, err = nmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelMsgMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nmuo.mutation = mutation
			node, err = nmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nmuo.hooks) - 1; i >= 0; i-- {
			if nmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nmuo *NovelMsgUpdateOne) SaveX(ctx context.Context) *NovelMsg {
	node, err := nmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nmuo *NovelMsgUpdateOne) Exec(ctx context.Context) error {
	_, err := nmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nmuo *NovelMsgUpdateOne) ExecX(ctx context.Context) {
	if err := nmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nmuo *NovelMsgUpdateOne) defaults() {
	if _, ok := nmuo.mutation.UpdatedAt(); !ok {
		v := novelmsg.UpdateDefaultUpdatedAt()
		nmuo.mutation.SetUpdatedAt(v)
	}
}

func (nmuo *NovelMsgUpdateOne) sqlSave(ctx context.Context) (_node *NovelMsg, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelmsg.Table,
			Columns: novelmsg.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelmsg.FieldID,
			},
		},
	}
	id, ok := nmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NovelMsg.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelmsg.FieldID)
		for _, f := range fields {
			if !novelmsg.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != novelmsg.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nmuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelmsg.FieldTitle,
		})
	}
	if nmuo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelmsg.FieldTitle,
		})
	}
	if value, ok := nmuo.mutation.Msg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelmsg.FieldMsg,
		})
	}
	if nmuo.mutation.MsgCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelmsg.FieldMsg,
		})
	}
	if value, ok := nmuo.mutation.MsgType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelmsg.FieldMsgType,
		})
	}
	if nmuo.mutation.MsgTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelmsg.FieldMsgType,
		})
	}
	if value, ok := nmuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: novelmsg.FieldStatus,
		})
	}
	if nmuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: novelmsg.FieldStatus,
		})
	}
	if value, ok := nmuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelmsg.FieldUpdatedAt,
		})
	}
	if value, ok := nmuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldCreateBy,
		})
	}
	if value, ok := nmuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldCreateBy,
		})
	}
	if value, ok := nmuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldUpdateBy,
		})
	}
	if value, ok := nmuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldUpdateBy,
		})
	}
	if value, ok := nmuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldTenantId,
		})
	}
	if value, ok := nmuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelmsg.FieldTenantId,
		})
	}
	_node = &NovelMsg{config: nmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelmsg.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
