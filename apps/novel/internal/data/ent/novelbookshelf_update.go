// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelbookshelf"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelBookshelfUpdate is the builder for updating NovelBookshelf entities.
type NovelBookshelfUpdate struct {
	config
	hooks    []Hook
	mutation *NovelBookshelfMutation
}

// Where appends a list predicates to the NovelBookshelfUpdate builder.
func (nbu *NovelBookshelfUpdate) Where(ps ...predicate.NovelBookshelf) *NovelBookshelfUpdate {
	nbu.mutation.Where(ps...)
	return nbu
}

// SetUserId sets the "userId" field.
func (nbu *NovelBookshelfUpdate) SetUserId(i int64) *NovelBookshelfUpdate {
	nbu.mutation.SetUserId(i)
	return nbu
}

// SetUserName sets the "userName" field.
func (nbu *NovelBookshelfUpdate) SetUserName(s string) *NovelBookshelfUpdate {
	nbu.mutation.SetUserName(s)
	return nbu
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableUserName(s *string) *NovelBookshelfUpdate {
	if s != nil {
		nbu.SetUserName(*s)
	}
	return nbu
}

// ClearUserName clears the value of the "userName" field.
func (nbu *NovelBookshelfUpdate) ClearUserName() *NovelBookshelfUpdate {
	nbu.mutation.ClearUserName()
	return nbu
}

// SetNovelId sets the "novelId" field.
func (nbu *NovelBookshelfUpdate) SetNovelId(i int64) *NovelBookshelfUpdate {
	nbu.mutation.ResetNovelId()
	nbu.mutation.SetNovelId(i)
	return nbu
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableNovelId(i *int64) *NovelBookshelfUpdate {
	if i != nil {
		nbu.SetNovelId(*i)
	}
	return nbu
}

// AddNovelId adds i to the "novelId" field.
func (nbu *NovelBookshelfUpdate) AddNovelId(i int64) *NovelBookshelfUpdate {
	nbu.mutation.AddNovelId(i)
	return nbu
}

// ClearNovelId clears the value of the "novelId" field.
func (nbu *NovelBookshelfUpdate) ClearNovelId() *NovelBookshelfUpdate {
	nbu.mutation.ClearNovelId()
	return nbu
}

// SetLastReadTime sets the "lastReadTime" field.
func (nbu *NovelBookshelfUpdate) SetLastReadTime(t time.Time) *NovelBookshelfUpdate {
	nbu.mutation.SetLastReadTime(t)
	return nbu
}

// SetNillableLastReadTime sets the "lastReadTime" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableLastReadTime(t *time.Time) *NovelBookshelfUpdate {
	if t != nil {
		nbu.SetLastReadTime(*t)
	}
	return nbu
}

// ClearLastReadTime clears the value of the "lastReadTime" field.
func (nbu *NovelBookshelfUpdate) ClearLastReadTime() *NovelBookshelfUpdate {
	nbu.mutation.ClearLastReadTime()
	return nbu
}

// SetLastChapterOrder sets the "lastChapterOrder" field.
func (nbu *NovelBookshelfUpdate) SetLastChapterOrder(i int32) *NovelBookshelfUpdate {
	nbu.mutation.ResetLastChapterOrder()
	nbu.mutation.SetLastChapterOrder(i)
	return nbu
}

// SetNillableLastChapterOrder sets the "lastChapterOrder" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableLastChapterOrder(i *int32) *NovelBookshelfUpdate {
	if i != nil {
		nbu.SetLastChapterOrder(*i)
	}
	return nbu
}

// AddLastChapterOrder adds i to the "lastChapterOrder" field.
func (nbu *NovelBookshelfUpdate) AddLastChapterOrder(i int32) *NovelBookshelfUpdate {
	nbu.mutation.AddLastChapterOrder(i)
	return nbu
}

// ClearLastChapterOrder clears the value of the "lastChapterOrder" field.
func (nbu *NovelBookshelfUpdate) ClearLastChapterOrder() *NovelBookshelfUpdate {
	nbu.mutation.ClearLastChapterOrder()
	return nbu
}

// SetLastChapterId sets the "lastChapterId" field.
func (nbu *NovelBookshelfUpdate) SetLastChapterId(i int64) *NovelBookshelfUpdate {
	nbu.mutation.ResetLastChapterId()
	nbu.mutation.SetLastChapterId(i)
	return nbu
}

// SetNillableLastChapterId sets the "lastChapterId" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableLastChapterId(i *int64) *NovelBookshelfUpdate {
	if i != nil {
		nbu.SetLastChapterId(*i)
	}
	return nbu
}

// AddLastChapterId adds i to the "lastChapterId" field.
func (nbu *NovelBookshelfUpdate) AddLastChapterId(i int64) *NovelBookshelfUpdate {
	nbu.mutation.AddLastChapterId(i)
	return nbu
}

// ClearLastChapterId clears the value of the "lastChapterId" field.
func (nbu *NovelBookshelfUpdate) ClearLastChapterId() *NovelBookshelfUpdate {
	nbu.mutation.ClearLastChapterId()
	return nbu
}

// SetLastChapterName sets the "lastChapterName" field.
func (nbu *NovelBookshelfUpdate) SetLastChapterName(s string) *NovelBookshelfUpdate {
	nbu.mutation.SetLastChapterName(s)
	return nbu
}

// SetNillableLastChapterName sets the "lastChapterName" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableLastChapterName(s *string) *NovelBookshelfUpdate {
	if s != nil {
		nbu.SetLastChapterName(*s)
	}
	return nbu
}

// ClearLastChapterName clears the value of the "lastChapterName" field.
func (nbu *NovelBookshelfUpdate) ClearLastChapterName() *NovelBookshelfUpdate {
	nbu.mutation.ClearLastChapterName()
	return nbu
}

// SetRemark sets the "remark" field.
func (nbu *NovelBookshelfUpdate) SetRemark(s string) *NovelBookshelfUpdate {
	nbu.mutation.SetRemark(s)
	return nbu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableRemark(s *string) *NovelBookshelfUpdate {
	if s != nil {
		nbu.SetRemark(*s)
	}
	return nbu
}

// ClearRemark clears the value of the "remark" field.
func (nbu *NovelBookshelfUpdate) ClearRemark() *NovelBookshelfUpdate {
	nbu.mutation.ClearRemark()
	return nbu
}

// SetUpdatedAt sets the "updatedAt" field.
func (nbu *NovelBookshelfUpdate) SetUpdatedAt(t time.Time) *NovelBookshelfUpdate {
	nbu.mutation.SetUpdatedAt(t)
	return nbu
}

// SetCreateBy sets the "createBy" field.
func (nbu *NovelBookshelfUpdate) SetCreateBy(i int64) *NovelBookshelfUpdate {
	nbu.mutation.ResetCreateBy()
	nbu.mutation.SetCreateBy(i)
	return nbu
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableCreateBy(i *int64) *NovelBookshelfUpdate {
	if i != nil {
		nbu.SetCreateBy(*i)
	}
	return nbu
}

// AddCreateBy adds i to the "createBy" field.
func (nbu *NovelBookshelfUpdate) AddCreateBy(i int64) *NovelBookshelfUpdate {
	nbu.mutation.AddCreateBy(i)
	return nbu
}

// SetUpdateBy sets the "updateBy" field.
func (nbu *NovelBookshelfUpdate) SetUpdateBy(i int64) *NovelBookshelfUpdate {
	nbu.mutation.ResetUpdateBy()
	nbu.mutation.SetUpdateBy(i)
	return nbu
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableUpdateBy(i *int64) *NovelBookshelfUpdate {
	if i != nil {
		nbu.SetUpdateBy(*i)
	}
	return nbu
}

// AddUpdateBy adds i to the "updateBy" field.
func (nbu *NovelBookshelfUpdate) AddUpdateBy(i int64) *NovelBookshelfUpdate {
	nbu.mutation.AddUpdateBy(i)
	return nbu
}

// SetTenantId sets the "tenantId" field.
func (nbu *NovelBookshelfUpdate) SetTenantId(i int64) *NovelBookshelfUpdate {
	nbu.mutation.ResetTenantId()
	nbu.mutation.SetTenantId(i)
	return nbu
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nbu *NovelBookshelfUpdate) SetNillableTenantId(i *int64) *NovelBookshelfUpdate {
	if i != nil {
		nbu.SetTenantId(*i)
	}
	return nbu
}

// AddTenantId adds i to the "tenantId" field.
func (nbu *NovelBookshelfUpdate) AddTenantId(i int64) *NovelBookshelfUpdate {
	nbu.mutation.AddTenantId(i)
	return nbu
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (nbu *NovelBookshelfUpdate) SetUserID(id int64) *NovelBookshelfUpdate {
	nbu.mutation.SetUserID(id)
	return nbu
}

// SetUser sets the "user" edge to the SocialUser entity.
func (nbu *NovelBookshelfUpdate) SetUser(s *SocialUser) *NovelBookshelfUpdate {
	return nbu.SetUserID(s.ID)
}

// Mutation returns the NovelBookshelfMutation object of the builder.
func (nbu *NovelBookshelfUpdate) Mutation() *NovelBookshelfMutation {
	return nbu.mutation
}

// ClearUser clears the "user" edge to the SocialUser entity.
func (nbu *NovelBookshelfUpdate) ClearUser() *NovelBookshelfUpdate {
	nbu.mutation.ClearUser()
	return nbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nbu *NovelBookshelfUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	nbu.defaults()
	if len(nbu.hooks) == 0 {
		if err = nbu.check(); err != nil {
			return 0, err
		}
		affected, err = nbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelBookshelfMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nbu.check(); err != nil {
				return 0, err
			}
			nbu.mutation = mutation
			affected, err = nbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nbu.hooks) - 1; i >= 0; i-- {
			if nbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nbu *NovelBookshelfUpdate) SaveX(ctx context.Context) int {
	affected, err := nbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nbu *NovelBookshelfUpdate) Exec(ctx context.Context) error {
	_, err := nbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbu *NovelBookshelfUpdate) ExecX(ctx context.Context) {
	if err := nbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nbu *NovelBookshelfUpdate) defaults() {
	if _, ok := nbu.mutation.UpdatedAt(); !ok {
		v := novelbookshelf.UpdateDefaultUpdatedAt()
		nbu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nbu *NovelBookshelfUpdate) check() error {
	if _, ok := nbu.mutation.UserID(); nbu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NovelBookshelf.user"`)
	}
	return nil
}

func (nbu *NovelBookshelfUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelbookshelf.Table,
			Columns: novelbookshelf.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbookshelf.FieldID,
			},
		},
	}
	if ps := nbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nbu.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldUserName,
		})
	}
	if nbu.mutation.UserNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbookshelf.FieldUserName,
		})
	}
	if value, ok := nbu.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldNovelId,
		})
	}
	if value, ok := nbu.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldNovelId,
		})
	}
	if nbu.mutation.NovelIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbookshelf.FieldNovelId,
		})
	}
	if value, ok := nbu.mutation.LastReadTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbookshelf.FieldLastReadTime,
		})
	}
	if nbu.mutation.LastReadTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: novelbookshelf.FieldLastReadTime,
		})
	}
	if value, ok := nbu.mutation.LastChapterOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterOrder,
		})
	}
	if value, ok := nbu.mutation.AddedLastChapterOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterOrder,
		})
	}
	if nbu.mutation.LastChapterOrderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: novelbookshelf.FieldLastChapterOrder,
		})
	}
	if value, ok := nbu.mutation.LastChapterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterId,
		})
	}
	if value, ok := nbu.mutation.AddedLastChapterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterId,
		})
	}
	if nbu.mutation.LastChapterIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbookshelf.FieldLastChapterId,
		})
	}
	if value, ok := nbu.mutation.LastChapterName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterName,
		})
	}
	if nbu.mutation.LastChapterNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbookshelf.FieldLastChapterName,
		})
	}
	if value, ok := nbu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldRemark,
		})
	}
	if nbu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbookshelf.FieldRemark,
		})
	}
	if value, ok := nbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbookshelf.FieldUpdatedAt,
		})
	}
	if value, ok := nbu.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldCreateBy,
		})
	}
	if value, ok := nbu.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldCreateBy,
		})
	}
	if value, ok := nbu.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldUpdateBy,
		})
	}
	if value, ok := nbu.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldUpdateBy,
		})
	}
	if value, ok := nbu.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldTenantId,
		})
	}
	if value, ok := nbu.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldTenantId,
		})
	}
	if nbu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbookshelf.UserTable,
			Columns: []string{novelbookshelf.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: socialuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nbu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbookshelf.UserTable,
			Columns: []string{novelbookshelf.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelbookshelf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NovelBookshelfUpdateOne is the builder for updating a single NovelBookshelf entity.
type NovelBookshelfUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NovelBookshelfMutation
}

// SetUserId sets the "userId" field.
func (nbuo *NovelBookshelfUpdateOne) SetUserId(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.SetUserId(i)
	return nbuo
}

// SetUserName sets the "userName" field.
func (nbuo *NovelBookshelfUpdateOne) SetUserName(s string) *NovelBookshelfUpdateOne {
	nbuo.mutation.SetUserName(s)
	return nbuo
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableUserName(s *string) *NovelBookshelfUpdateOne {
	if s != nil {
		nbuo.SetUserName(*s)
	}
	return nbuo
}

// ClearUserName clears the value of the "userName" field.
func (nbuo *NovelBookshelfUpdateOne) ClearUserName() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearUserName()
	return nbuo
}

// SetNovelId sets the "novelId" field.
func (nbuo *NovelBookshelfUpdateOne) SetNovelId(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.ResetNovelId()
	nbuo.mutation.SetNovelId(i)
	return nbuo
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableNovelId(i *int64) *NovelBookshelfUpdateOne {
	if i != nil {
		nbuo.SetNovelId(*i)
	}
	return nbuo
}

// AddNovelId adds i to the "novelId" field.
func (nbuo *NovelBookshelfUpdateOne) AddNovelId(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.AddNovelId(i)
	return nbuo
}

// ClearNovelId clears the value of the "novelId" field.
func (nbuo *NovelBookshelfUpdateOne) ClearNovelId() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearNovelId()
	return nbuo
}

// SetLastReadTime sets the "lastReadTime" field.
func (nbuo *NovelBookshelfUpdateOne) SetLastReadTime(t time.Time) *NovelBookshelfUpdateOne {
	nbuo.mutation.SetLastReadTime(t)
	return nbuo
}

// SetNillableLastReadTime sets the "lastReadTime" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableLastReadTime(t *time.Time) *NovelBookshelfUpdateOne {
	if t != nil {
		nbuo.SetLastReadTime(*t)
	}
	return nbuo
}

// ClearLastReadTime clears the value of the "lastReadTime" field.
func (nbuo *NovelBookshelfUpdateOne) ClearLastReadTime() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearLastReadTime()
	return nbuo
}

// SetLastChapterOrder sets the "lastChapterOrder" field.
func (nbuo *NovelBookshelfUpdateOne) SetLastChapterOrder(i int32) *NovelBookshelfUpdateOne {
	nbuo.mutation.ResetLastChapterOrder()
	nbuo.mutation.SetLastChapterOrder(i)
	return nbuo
}

// SetNillableLastChapterOrder sets the "lastChapterOrder" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableLastChapterOrder(i *int32) *NovelBookshelfUpdateOne {
	if i != nil {
		nbuo.SetLastChapterOrder(*i)
	}
	return nbuo
}

// AddLastChapterOrder adds i to the "lastChapterOrder" field.
func (nbuo *NovelBookshelfUpdateOne) AddLastChapterOrder(i int32) *NovelBookshelfUpdateOne {
	nbuo.mutation.AddLastChapterOrder(i)
	return nbuo
}

// ClearLastChapterOrder clears the value of the "lastChapterOrder" field.
func (nbuo *NovelBookshelfUpdateOne) ClearLastChapterOrder() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearLastChapterOrder()
	return nbuo
}

// SetLastChapterId sets the "lastChapterId" field.
func (nbuo *NovelBookshelfUpdateOne) SetLastChapterId(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.ResetLastChapterId()
	nbuo.mutation.SetLastChapterId(i)
	return nbuo
}

// SetNillableLastChapterId sets the "lastChapterId" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableLastChapterId(i *int64) *NovelBookshelfUpdateOne {
	if i != nil {
		nbuo.SetLastChapterId(*i)
	}
	return nbuo
}

// AddLastChapterId adds i to the "lastChapterId" field.
func (nbuo *NovelBookshelfUpdateOne) AddLastChapterId(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.AddLastChapterId(i)
	return nbuo
}

// ClearLastChapterId clears the value of the "lastChapterId" field.
func (nbuo *NovelBookshelfUpdateOne) ClearLastChapterId() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearLastChapterId()
	return nbuo
}

// SetLastChapterName sets the "lastChapterName" field.
func (nbuo *NovelBookshelfUpdateOne) SetLastChapterName(s string) *NovelBookshelfUpdateOne {
	nbuo.mutation.SetLastChapterName(s)
	return nbuo
}

// SetNillableLastChapterName sets the "lastChapterName" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableLastChapterName(s *string) *NovelBookshelfUpdateOne {
	if s != nil {
		nbuo.SetLastChapterName(*s)
	}
	return nbuo
}

// ClearLastChapterName clears the value of the "lastChapterName" field.
func (nbuo *NovelBookshelfUpdateOne) ClearLastChapterName() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearLastChapterName()
	return nbuo
}

// SetRemark sets the "remark" field.
func (nbuo *NovelBookshelfUpdateOne) SetRemark(s string) *NovelBookshelfUpdateOne {
	nbuo.mutation.SetRemark(s)
	return nbuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableRemark(s *string) *NovelBookshelfUpdateOne {
	if s != nil {
		nbuo.SetRemark(*s)
	}
	return nbuo
}

// ClearRemark clears the value of the "remark" field.
func (nbuo *NovelBookshelfUpdateOne) ClearRemark() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearRemark()
	return nbuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (nbuo *NovelBookshelfUpdateOne) SetUpdatedAt(t time.Time) *NovelBookshelfUpdateOne {
	nbuo.mutation.SetUpdatedAt(t)
	return nbuo
}

// SetCreateBy sets the "createBy" field.
func (nbuo *NovelBookshelfUpdateOne) SetCreateBy(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.ResetCreateBy()
	nbuo.mutation.SetCreateBy(i)
	return nbuo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableCreateBy(i *int64) *NovelBookshelfUpdateOne {
	if i != nil {
		nbuo.SetCreateBy(*i)
	}
	return nbuo
}

// AddCreateBy adds i to the "createBy" field.
func (nbuo *NovelBookshelfUpdateOne) AddCreateBy(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.AddCreateBy(i)
	return nbuo
}

// SetUpdateBy sets the "updateBy" field.
func (nbuo *NovelBookshelfUpdateOne) SetUpdateBy(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.ResetUpdateBy()
	nbuo.mutation.SetUpdateBy(i)
	return nbuo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableUpdateBy(i *int64) *NovelBookshelfUpdateOne {
	if i != nil {
		nbuo.SetUpdateBy(*i)
	}
	return nbuo
}

// AddUpdateBy adds i to the "updateBy" field.
func (nbuo *NovelBookshelfUpdateOne) AddUpdateBy(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.AddUpdateBy(i)
	return nbuo
}

// SetTenantId sets the "tenantId" field.
func (nbuo *NovelBookshelfUpdateOne) SetTenantId(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.ResetTenantId()
	nbuo.mutation.SetTenantId(i)
	return nbuo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nbuo *NovelBookshelfUpdateOne) SetNillableTenantId(i *int64) *NovelBookshelfUpdateOne {
	if i != nil {
		nbuo.SetTenantId(*i)
	}
	return nbuo
}

// AddTenantId adds i to the "tenantId" field.
func (nbuo *NovelBookshelfUpdateOne) AddTenantId(i int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.AddTenantId(i)
	return nbuo
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (nbuo *NovelBookshelfUpdateOne) SetUserID(id int64) *NovelBookshelfUpdateOne {
	nbuo.mutation.SetUserID(id)
	return nbuo
}

// SetUser sets the "user" edge to the SocialUser entity.
func (nbuo *NovelBookshelfUpdateOne) SetUser(s *SocialUser) *NovelBookshelfUpdateOne {
	return nbuo.SetUserID(s.ID)
}

// Mutation returns the NovelBookshelfMutation object of the builder.
func (nbuo *NovelBookshelfUpdateOne) Mutation() *NovelBookshelfMutation {
	return nbuo.mutation
}

// ClearUser clears the "user" edge to the SocialUser entity.
func (nbuo *NovelBookshelfUpdateOne) ClearUser() *NovelBookshelfUpdateOne {
	nbuo.mutation.ClearUser()
	return nbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nbuo *NovelBookshelfUpdateOne) Select(field string, fields ...string) *NovelBookshelfUpdateOne {
	nbuo.fields = append([]string{field}, fields...)
	return nbuo
}

// Save executes the query and returns the updated NovelBookshelf entity.
func (nbuo *NovelBookshelfUpdateOne) Save(ctx context.Context) (*NovelBookshelf, error) {
	var (
		err  error
		node *NovelBookshelf
	)
	nbuo.defaults()
	if len(nbuo.hooks) == 0 {
		if err = nbuo.check(); err != nil {
			return nil, err
		}
		node, err = nbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelBookshelfMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nbuo.check(); err != nil {
				return nil, err
			}
			nbuo.mutation = mutation
			node, err = nbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nbuo.hooks) - 1; i >= 0; i-- {
			if nbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nbuo *NovelBookshelfUpdateOne) SaveX(ctx context.Context) *NovelBookshelf {
	node, err := nbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nbuo *NovelBookshelfUpdateOne) Exec(ctx context.Context) error {
	_, err := nbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbuo *NovelBookshelfUpdateOne) ExecX(ctx context.Context) {
	if err := nbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nbuo *NovelBookshelfUpdateOne) defaults() {
	if _, ok := nbuo.mutation.UpdatedAt(); !ok {
		v := novelbookshelf.UpdateDefaultUpdatedAt()
		nbuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nbuo *NovelBookshelfUpdateOne) check() error {
	if _, ok := nbuo.mutation.UserID(); nbuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NovelBookshelf.user"`)
	}
	return nil
}

func (nbuo *NovelBookshelfUpdateOne) sqlSave(ctx context.Context) (_node *NovelBookshelf, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelbookshelf.Table,
			Columns: novelbookshelf.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbookshelf.FieldID,
			},
		},
	}
	id, ok := nbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NovelBookshelf.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelbookshelf.FieldID)
		for _, f := range fields {
			if !novelbookshelf.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != novelbookshelf.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nbuo.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldUserName,
		})
	}
	if nbuo.mutation.UserNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbookshelf.FieldUserName,
		})
	}
	if value, ok := nbuo.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldNovelId,
		})
	}
	if value, ok := nbuo.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldNovelId,
		})
	}
	if nbuo.mutation.NovelIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbookshelf.FieldNovelId,
		})
	}
	if value, ok := nbuo.mutation.LastReadTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbookshelf.FieldLastReadTime,
		})
	}
	if nbuo.mutation.LastReadTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: novelbookshelf.FieldLastReadTime,
		})
	}
	if value, ok := nbuo.mutation.LastChapterOrder(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterOrder,
		})
	}
	if value, ok := nbuo.mutation.AddedLastChapterOrder(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterOrder,
		})
	}
	if nbuo.mutation.LastChapterOrderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: novelbookshelf.FieldLastChapterOrder,
		})
	}
	if value, ok := nbuo.mutation.LastChapterId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterId,
		})
	}
	if value, ok := nbuo.mutation.AddedLastChapterId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterId,
		})
	}
	if nbuo.mutation.LastChapterIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbookshelf.FieldLastChapterId,
		})
	}
	if value, ok := nbuo.mutation.LastChapterName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldLastChapterName,
		})
	}
	if nbuo.mutation.LastChapterNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbookshelf.FieldLastChapterName,
		})
	}
	if value, ok := nbuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbookshelf.FieldRemark,
		})
	}
	if nbuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbookshelf.FieldRemark,
		})
	}
	if value, ok := nbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbookshelf.FieldUpdatedAt,
		})
	}
	if value, ok := nbuo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldCreateBy,
		})
	}
	if value, ok := nbuo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldCreateBy,
		})
	}
	if value, ok := nbuo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldUpdateBy,
		})
	}
	if value, ok := nbuo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldUpdateBy,
		})
	}
	if value, ok := nbuo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldTenantId,
		})
	}
	if value, ok := nbuo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbookshelf.FieldTenantId,
		})
	}
	if nbuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbookshelf.UserTable,
			Columns: []string{novelbookshelf.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: socialuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nbuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbookshelf.UserTable,
			Columns: []string{novelbookshelf.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NovelBookshelf{config: nbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelbookshelf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
