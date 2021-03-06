// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelbuyrecord"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelBuyRecordUpdate is the builder for updating NovelBuyRecord entities.
type NovelBuyRecordUpdate struct {
	config
	hooks    []Hook
	mutation *NovelBuyRecordMutation
}

// Where appends a list predicates to the NovelBuyRecordUpdate builder.
func (nbru *NovelBuyRecordUpdate) Where(ps ...predicate.NovelBuyRecord) *NovelBuyRecordUpdate {
	nbru.mutation.Where(ps...)
	return nbru
}

// SetUserId sets the "userId" field.
func (nbru *NovelBuyRecordUpdate) SetUserId(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.SetUserId(i)
	return nbru
}

// SetUserName sets the "userName" field.
func (nbru *NovelBuyRecordUpdate) SetUserName(s string) *NovelBuyRecordUpdate {
	nbru.mutation.SetUserName(s)
	return nbru
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableUserName(s *string) *NovelBuyRecordUpdate {
	if s != nil {
		nbru.SetUserName(*s)
	}
	return nbru
}

// ClearUserName clears the value of the "userName" field.
func (nbru *NovelBuyRecordUpdate) ClearUserName() *NovelBuyRecordUpdate {
	nbru.mutation.ClearUserName()
	return nbru
}

// SetNovelId sets the "novelId" field.
func (nbru *NovelBuyRecordUpdate) SetNovelId(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.ResetNovelId()
	nbru.mutation.SetNovelId(i)
	return nbru
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableNovelId(i *int64) *NovelBuyRecordUpdate {
	if i != nil {
		nbru.SetNovelId(*i)
	}
	return nbru
}

// AddNovelId adds i to the "novelId" field.
func (nbru *NovelBuyRecordUpdate) AddNovelId(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.AddNovelId(i)
	return nbru
}

// ClearNovelId clears the value of the "novelId" field.
func (nbru *NovelBuyRecordUpdate) ClearNovelId() *NovelBuyRecordUpdate {
	nbru.mutation.ClearNovelId()
	return nbru
}

// SetNovelName sets the "novelName" field.
func (nbru *NovelBuyRecordUpdate) SetNovelName(s string) *NovelBuyRecordUpdate {
	nbru.mutation.SetNovelName(s)
	return nbru
}

// SetNillableNovelName sets the "novelName" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableNovelName(s *string) *NovelBuyRecordUpdate {
	if s != nil {
		nbru.SetNovelName(*s)
	}
	return nbru
}

// ClearNovelName clears the value of the "novelName" field.
func (nbru *NovelBuyRecordUpdate) ClearNovelName() *NovelBuyRecordUpdate {
	nbru.mutation.ClearNovelName()
	return nbru
}

// SetPackageId sets the "packageId" field.
func (nbru *NovelBuyRecordUpdate) SetPackageId(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.ResetPackageId()
	nbru.mutation.SetPackageId(i)
	return nbru
}

// SetNillablePackageId sets the "packageId" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillablePackageId(i *int64) *NovelBuyRecordUpdate {
	if i != nil {
		nbru.SetPackageId(*i)
	}
	return nbru
}

// AddPackageId adds i to the "packageId" field.
func (nbru *NovelBuyRecordUpdate) AddPackageId(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.AddPackageId(i)
	return nbru
}

// ClearPackageId clears the value of the "packageId" field.
func (nbru *NovelBuyRecordUpdate) ClearPackageId() *NovelBuyRecordUpdate {
	nbru.mutation.ClearPackageId()
	return nbru
}

// SetCover sets the "cover" field.
func (nbru *NovelBuyRecordUpdate) SetCover(s string) *NovelBuyRecordUpdate {
	nbru.mutation.SetCover(s)
	return nbru
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableCover(s *string) *NovelBuyRecordUpdate {
	if s != nil {
		nbru.SetCover(*s)
	}
	return nbru
}

// ClearCover clears the value of the "cover" field.
func (nbru *NovelBuyRecordUpdate) ClearCover() *NovelBuyRecordUpdate {
	nbru.mutation.ClearCover()
	return nbru
}

// SetCoin sets the "coin" field.
func (nbru *NovelBuyRecordUpdate) SetCoin(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.ResetCoin()
	nbru.mutation.SetCoin(i)
	return nbru
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableCoin(i *int64) *NovelBuyRecordUpdate {
	if i != nil {
		nbru.SetCoin(*i)
	}
	return nbru
}

// AddCoin adds i to the "coin" field.
func (nbru *NovelBuyRecordUpdate) AddCoin(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.AddCoin(i)
	return nbru
}

// ClearCoin clears the value of the "coin" field.
func (nbru *NovelBuyRecordUpdate) ClearCoin() *NovelBuyRecordUpdate {
	nbru.mutation.ClearCoin()
	return nbru
}

// SetCoupon sets the "coupon" field.
func (nbru *NovelBuyRecordUpdate) SetCoupon(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.ResetCoupon()
	nbru.mutation.SetCoupon(i)
	return nbru
}

// SetNillableCoupon sets the "coupon" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableCoupon(i *int64) *NovelBuyRecordUpdate {
	if i != nil {
		nbru.SetCoupon(*i)
	}
	return nbru
}

// AddCoupon adds i to the "coupon" field.
func (nbru *NovelBuyRecordUpdate) AddCoupon(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.AddCoupon(i)
	return nbru
}

// ClearCoupon clears the value of the "coupon" field.
func (nbru *NovelBuyRecordUpdate) ClearCoupon() *NovelBuyRecordUpdate {
	nbru.mutation.ClearCoupon()
	return nbru
}

// SetRemark sets the "remark" field.
func (nbru *NovelBuyRecordUpdate) SetRemark(s string) *NovelBuyRecordUpdate {
	nbru.mutation.SetRemark(s)
	return nbru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableRemark(s *string) *NovelBuyRecordUpdate {
	if s != nil {
		nbru.SetRemark(*s)
	}
	return nbru
}

// ClearRemark clears the value of the "remark" field.
func (nbru *NovelBuyRecordUpdate) ClearRemark() *NovelBuyRecordUpdate {
	nbru.mutation.ClearRemark()
	return nbru
}

// SetUpdatedAt sets the "updatedAt" field.
func (nbru *NovelBuyRecordUpdate) SetUpdatedAt(t time.Time) *NovelBuyRecordUpdate {
	nbru.mutation.SetUpdatedAt(t)
	return nbru
}

// SetCreateBy sets the "createBy" field.
func (nbru *NovelBuyRecordUpdate) SetCreateBy(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.ResetCreateBy()
	nbru.mutation.SetCreateBy(i)
	return nbru
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableCreateBy(i *int64) *NovelBuyRecordUpdate {
	if i != nil {
		nbru.SetCreateBy(*i)
	}
	return nbru
}

// AddCreateBy adds i to the "createBy" field.
func (nbru *NovelBuyRecordUpdate) AddCreateBy(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.AddCreateBy(i)
	return nbru
}

// SetUpdateBy sets the "updateBy" field.
func (nbru *NovelBuyRecordUpdate) SetUpdateBy(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.ResetUpdateBy()
	nbru.mutation.SetUpdateBy(i)
	return nbru
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableUpdateBy(i *int64) *NovelBuyRecordUpdate {
	if i != nil {
		nbru.SetUpdateBy(*i)
	}
	return nbru
}

// AddUpdateBy adds i to the "updateBy" field.
func (nbru *NovelBuyRecordUpdate) AddUpdateBy(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.AddUpdateBy(i)
	return nbru
}

// SetTenantId sets the "tenantId" field.
func (nbru *NovelBuyRecordUpdate) SetTenantId(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.ResetTenantId()
	nbru.mutation.SetTenantId(i)
	return nbru
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nbru *NovelBuyRecordUpdate) SetNillableTenantId(i *int64) *NovelBuyRecordUpdate {
	if i != nil {
		nbru.SetTenantId(*i)
	}
	return nbru
}

// AddTenantId adds i to the "tenantId" field.
func (nbru *NovelBuyRecordUpdate) AddTenantId(i int64) *NovelBuyRecordUpdate {
	nbru.mutation.AddTenantId(i)
	return nbru
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (nbru *NovelBuyRecordUpdate) SetUserID(id int64) *NovelBuyRecordUpdate {
	nbru.mutation.SetUserID(id)
	return nbru
}

// SetUser sets the "user" edge to the SocialUser entity.
func (nbru *NovelBuyRecordUpdate) SetUser(s *SocialUser) *NovelBuyRecordUpdate {
	return nbru.SetUserID(s.ID)
}

// Mutation returns the NovelBuyRecordMutation object of the builder.
func (nbru *NovelBuyRecordUpdate) Mutation() *NovelBuyRecordMutation {
	return nbru.mutation
}

// ClearUser clears the "user" edge to the SocialUser entity.
func (nbru *NovelBuyRecordUpdate) ClearUser() *NovelBuyRecordUpdate {
	nbru.mutation.ClearUser()
	return nbru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nbru *NovelBuyRecordUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	nbru.defaults()
	if len(nbru.hooks) == 0 {
		if err = nbru.check(); err != nil {
			return 0, err
		}
		affected, err = nbru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelBuyRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nbru.check(); err != nil {
				return 0, err
			}
			nbru.mutation = mutation
			affected, err = nbru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nbru.hooks) - 1; i >= 0; i-- {
			if nbru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nbru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nbru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nbru *NovelBuyRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := nbru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nbru *NovelBuyRecordUpdate) Exec(ctx context.Context) error {
	_, err := nbru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbru *NovelBuyRecordUpdate) ExecX(ctx context.Context) {
	if err := nbru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nbru *NovelBuyRecordUpdate) defaults() {
	if _, ok := nbru.mutation.UpdatedAt(); !ok {
		v := novelbuyrecord.UpdateDefaultUpdatedAt()
		nbru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nbru *NovelBuyRecordUpdate) check() error {
	if _, ok := nbru.mutation.UserID(); nbru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NovelBuyRecord.user"`)
	}
	return nil
}

func (nbru *NovelBuyRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelbuyrecord.Table,
			Columns: novelbuyrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbuyrecord.FieldID,
			},
		},
	}
	if ps := nbru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nbru.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldUserName,
		})
	}
	if nbru.mutation.UserNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldUserName,
		})
	}
	if value, ok := nbru.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldNovelId,
		})
	}
	if value, ok := nbru.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldNovelId,
		})
	}
	if nbru.mutation.NovelIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldNovelId,
		})
	}
	if value, ok := nbru.mutation.NovelName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldNovelName,
		})
	}
	if nbru.mutation.NovelNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldNovelName,
		})
	}
	if value, ok := nbru.mutation.PackageId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldPackageId,
		})
	}
	if value, ok := nbru.mutation.AddedPackageId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldPackageId,
		})
	}
	if nbru.mutation.PackageIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldPackageId,
		})
	}
	if value, ok := nbru.mutation.Cover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldCover,
		})
	}
	if nbru.mutation.CoverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldCover,
		})
	}
	if value, ok := nbru.mutation.Coin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoin,
		})
	}
	if value, ok := nbru.mutation.AddedCoin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoin,
		})
	}
	if nbru.mutation.CoinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldCoin,
		})
	}
	if value, ok := nbru.mutation.Coupon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoupon,
		})
	}
	if value, ok := nbru.mutation.AddedCoupon(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoupon,
		})
	}
	if nbru.mutation.CouponCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldCoupon,
		})
	}
	if value, ok := nbru.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldRemark,
		})
	}
	if nbru.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldRemark,
		})
	}
	if value, ok := nbru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbuyrecord.FieldUpdatedAt,
		})
	}
	if value, ok := nbru.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCreateBy,
		})
	}
	if value, ok := nbru.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCreateBy,
		})
	}
	if value, ok := nbru.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldUpdateBy,
		})
	}
	if value, ok := nbru.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldUpdateBy,
		})
	}
	if value, ok := nbru.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldTenantId,
		})
	}
	if value, ok := nbru.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldTenantId,
		})
	}
	if nbru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbuyrecord.UserTable,
			Columns: []string{novelbuyrecord.UserColumn},
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
	if nodes := nbru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbuyrecord.UserTable,
			Columns: []string{novelbuyrecord.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, nbru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelbuyrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NovelBuyRecordUpdateOne is the builder for updating a single NovelBuyRecord entity.
type NovelBuyRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NovelBuyRecordMutation
}

// SetUserId sets the "userId" field.
func (nbruo *NovelBuyRecordUpdateOne) SetUserId(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.SetUserId(i)
	return nbruo
}

// SetUserName sets the "userName" field.
func (nbruo *NovelBuyRecordUpdateOne) SetUserName(s string) *NovelBuyRecordUpdateOne {
	nbruo.mutation.SetUserName(s)
	return nbruo
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableUserName(s *string) *NovelBuyRecordUpdateOne {
	if s != nil {
		nbruo.SetUserName(*s)
	}
	return nbruo
}

// ClearUserName clears the value of the "userName" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearUserName() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearUserName()
	return nbruo
}

// SetNovelId sets the "novelId" field.
func (nbruo *NovelBuyRecordUpdateOne) SetNovelId(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.ResetNovelId()
	nbruo.mutation.SetNovelId(i)
	return nbruo
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableNovelId(i *int64) *NovelBuyRecordUpdateOne {
	if i != nil {
		nbruo.SetNovelId(*i)
	}
	return nbruo
}

// AddNovelId adds i to the "novelId" field.
func (nbruo *NovelBuyRecordUpdateOne) AddNovelId(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.AddNovelId(i)
	return nbruo
}

// ClearNovelId clears the value of the "novelId" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearNovelId() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearNovelId()
	return nbruo
}

// SetNovelName sets the "novelName" field.
func (nbruo *NovelBuyRecordUpdateOne) SetNovelName(s string) *NovelBuyRecordUpdateOne {
	nbruo.mutation.SetNovelName(s)
	return nbruo
}

// SetNillableNovelName sets the "novelName" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableNovelName(s *string) *NovelBuyRecordUpdateOne {
	if s != nil {
		nbruo.SetNovelName(*s)
	}
	return nbruo
}

// ClearNovelName clears the value of the "novelName" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearNovelName() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearNovelName()
	return nbruo
}

// SetPackageId sets the "packageId" field.
func (nbruo *NovelBuyRecordUpdateOne) SetPackageId(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.ResetPackageId()
	nbruo.mutation.SetPackageId(i)
	return nbruo
}

// SetNillablePackageId sets the "packageId" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillablePackageId(i *int64) *NovelBuyRecordUpdateOne {
	if i != nil {
		nbruo.SetPackageId(*i)
	}
	return nbruo
}

// AddPackageId adds i to the "packageId" field.
func (nbruo *NovelBuyRecordUpdateOne) AddPackageId(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.AddPackageId(i)
	return nbruo
}

// ClearPackageId clears the value of the "packageId" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearPackageId() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearPackageId()
	return nbruo
}

// SetCover sets the "cover" field.
func (nbruo *NovelBuyRecordUpdateOne) SetCover(s string) *NovelBuyRecordUpdateOne {
	nbruo.mutation.SetCover(s)
	return nbruo
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableCover(s *string) *NovelBuyRecordUpdateOne {
	if s != nil {
		nbruo.SetCover(*s)
	}
	return nbruo
}

// ClearCover clears the value of the "cover" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearCover() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearCover()
	return nbruo
}

// SetCoin sets the "coin" field.
func (nbruo *NovelBuyRecordUpdateOne) SetCoin(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.ResetCoin()
	nbruo.mutation.SetCoin(i)
	return nbruo
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableCoin(i *int64) *NovelBuyRecordUpdateOne {
	if i != nil {
		nbruo.SetCoin(*i)
	}
	return nbruo
}

// AddCoin adds i to the "coin" field.
func (nbruo *NovelBuyRecordUpdateOne) AddCoin(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.AddCoin(i)
	return nbruo
}

// ClearCoin clears the value of the "coin" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearCoin() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearCoin()
	return nbruo
}

// SetCoupon sets the "coupon" field.
func (nbruo *NovelBuyRecordUpdateOne) SetCoupon(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.ResetCoupon()
	nbruo.mutation.SetCoupon(i)
	return nbruo
}

// SetNillableCoupon sets the "coupon" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableCoupon(i *int64) *NovelBuyRecordUpdateOne {
	if i != nil {
		nbruo.SetCoupon(*i)
	}
	return nbruo
}

// AddCoupon adds i to the "coupon" field.
func (nbruo *NovelBuyRecordUpdateOne) AddCoupon(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.AddCoupon(i)
	return nbruo
}

// ClearCoupon clears the value of the "coupon" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearCoupon() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearCoupon()
	return nbruo
}

// SetRemark sets the "remark" field.
func (nbruo *NovelBuyRecordUpdateOne) SetRemark(s string) *NovelBuyRecordUpdateOne {
	nbruo.mutation.SetRemark(s)
	return nbruo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableRemark(s *string) *NovelBuyRecordUpdateOne {
	if s != nil {
		nbruo.SetRemark(*s)
	}
	return nbruo
}

// ClearRemark clears the value of the "remark" field.
func (nbruo *NovelBuyRecordUpdateOne) ClearRemark() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearRemark()
	return nbruo
}

// SetUpdatedAt sets the "updatedAt" field.
func (nbruo *NovelBuyRecordUpdateOne) SetUpdatedAt(t time.Time) *NovelBuyRecordUpdateOne {
	nbruo.mutation.SetUpdatedAt(t)
	return nbruo
}

// SetCreateBy sets the "createBy" field.
func (nbruo *NovelBuyRecordUpdateOne) SetCreateBy(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.ResetCreateBy()
	nbruo.mutation.SetCreateBy(i)
	return nbruo
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableCreateBy(i *int64) *NovelBuyRecordUpdateOne {
	if i != nil {
		nbruo.SetCreateBy(*i)
	}
	return nbruo
}

// AddCreateBy adds i to the "createBy" field.
func (nbruo *NovelBuyRecordUpdateOne) AddCreateBy(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.AddCreateBy(i)
	return nbruo
}

// SetUpdateBy sets the "updateBy" field.
func (nbruo *NovelBuyRecordUpdateOne) SetUpdateBy(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.ResetUpdateBy()
	nbruo.mutation.SetUpdateBy(i)
	return nbruo
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableUpdateBy(i *int64) *NovelBuyRecordUpdateOne {
	if i != nil {
		nbruo.SetUpdateBy(*i)
	}
	return nbruo
}

// AddUpdateBy adds i to the "updateBy" field.
func (nbruo *NovelBuyRecordUpdateOne) AddUpdateBy(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.AddUpdateBy(i)
	return nbruo
}

// SetTenantId sets the "tenantId" field.
func (nbruo *NovelBuyRecordUpdateOne) SetTenantId(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.ResetTenantId()
	nbruo.mutation.SetTenantId(i)
	return nbruo
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nbruo *NovelBuyRecordUpdateOne) SetNillableTenantId(i *int64) *NovelBuyRecordUpdateOne {
	if i != nil {
		nbruo.SetTenantId(*i)
	}
	return nbruo
}

// AddTenantId adds i to the "tenantId" field.
func (nbruo *NovelBuyRecordUpdateOne) AddTenantId(i int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.AddTenantId(i)
	return nbruo
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (nbruo *NovelBuyRecordUpdateOne) SetUserID(id int64) *NovelBuyRecordUpdateOne {
	nbruo.mutation.SetUserID(id)
	return nbruo
}

// SetUser sets the "user" edge to the SocialUser entity.
func (nbruo *NovelBuyRecordUpdateOne) SetUser(s *SocialUser) *NovelBuyRecordUpdateOne {
	return nbruo.SetUserID(s.ID)
}

// Mutation returns the NovelBuyRecordMutation object of the builder.
func (nbruo *NovelBuyRecordUpdateOne) Mutation() *NovelBuyRecordMutation {
	return nbruo.mutation
}

// ClearUser clears the "user" edge to the SocialUser entity.
func (nbruo *NovelBuyRecordUpdateOne) ClearUser() *NovelBuyRecordUpdateOne {
	nbruo.mutation.ClearUser()
	return nbruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nbruo *NovelBuyRecordUpdateOne) Select(field string, fields ...string) *NovelBuyRecordUpdateOne {
	nbruo.fields = append([]string{field}, fields...)
	return nbruo
}

// Save executes the query and returns the updated NovelBuyRecord entity.
func (nbruo *NovelBuyRecordUpdateOne) Save(ctx context.Context) (*NovelBuyRecord, error) {
	var (
		err  error
		node *NovelBuyRecord
	)
	nbruo.defaults()
	if len(nbruo.hooks) == 0 {
		if err = nbruo.check(); err != nil {
			return nil, err
		}
		node, err = nbruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelBuyRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nbruo.check(); err != nil {
				return nil, err
			}
			nbruo.mutation = mutation
			node, err = nbruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nbruo.hooks) - 1; i >= 0; i-- {
			if nbruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nbruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nbruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nbruo *NovelBuyRecordUpdateOne) SaveX(ctx context.Context) *NovelBuyRecord {
	node, err := nbruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nbruo *NovelBuyRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := nbruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbruo *NovelBuyRecordUpdateOne) ExecX(ctx context.Context) {
	if err := nbruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nbruo *NovelBuyRecordUpdateOne) defaults() {
	if _, ok := nbruo.mutation.UpdatedAt(); !ok {
		v := novelbuyrecord.UpdateDefaultUpdatedAt()
		nbruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nbruo *NovelBuyRecordUpdateOne) check() error {
	if _, ok := nbruo.mutation.UserID(); nbruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NovelBuyRecord.user"`)
	}
	return nil
}

func (nbruo *NovelBuyRecordUpdateOne) sqlSave(ctx context.Context) (_node *NovelBuyRecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelbuyrecord.Table,
			Columns: novelbuyrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbuyrecord.FieldID,
			},
		},
	}
	id, ok := nbruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NovelBuyRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nbruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelbuyrecord.FieldID)
		for _, f := range fields {
			if !novelbuyrecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != novelbuyrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nbruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nbruo.mutation.UserName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldUserName,
		})
	}
	if nbruo.mutation.UserNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldUserName,
		})
	}
	if value, ok := nbruo.mutation.NovelId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldNovelId,
		})
	}
	if value, ok := nbruo.mutation.AddedNovelId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldNovelId,
		})
	}
	if nbruo.mutation.NovelIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldNovelId,
		})
	}
	if value, ok := nbruo.mutation.NovelName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldNovelName,
		})
	}
	if nbruo.mutation.NovelNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldNovelName,
		})
	}
	if value, ok := nbruo.mutation.PackageId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldPackageId,
		})
	}
	if value, ok := nbruo.mutation.AddedPackageId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldPackageId,
		})
	}
	if nbruo.mutation.PackageIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldPackageId,
		})
	}
	if value, ok := nbruo.mutation.Cover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldCover,
		})
	}
	if nbruo.mutation.CoverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldCover,
		})
	}
	if value, ok := nbruo.mutation.Coin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoin,
		})
	}
	if value, ok := nbruo.mutation.AddedCoin(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoin,
		})
	}
	if nbruo.mutation.CoinCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldCoin,
		})
	}
	if value, ok := nbruo.mutation.Coupon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoupon,
		})
	}
	if value, ok := nbruo.mutation.AddedCoupon(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCoupon,
		})
	}
	if nbruo.mutation.CouponCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: novelbuyrecord.FieldCoupon,
		})
	}
	if value, ok := nbruo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuyrecord.FieldRemark,
		})
	}
	if nbruo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: novelbuyrecord.FieldRemark,
		})
	}
	if value, ok := nbruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbuyrecord.FieldUpdatedAt,
		})
	}
	if value, ok := nbruo.mutation.CreateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCreateBy,
		})
	}
	if value, ok := nbruo.mutation.AddedCreateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldCreateBy,
		})
	}
	if value, ok := nbruo.mutation.UpdateBy(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldUpdateBy,
		})
	}
	if value, ok := nbruo.mutation.AddedUpdateBy(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldUpdateBy,
		})
	}
	if value, ok := nbruo.mutation.TenantId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldTenantId,
		})
	}
	if value, ok := nbruo.mutation.AddedTenantId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuyrecord.FieldTenantId,
		})
	}
	if nbruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbuyrecord.UserTable,
			Columns: []string{novelbuyrecord.UserColumn},
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
	if nodes := nbruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbuyrecord.UserTable,
			Columns: []string{novelbuyrecord.UserColumn},
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
	_node = &NovelBuyRecord{config: nbruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nbruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelbuyrecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
