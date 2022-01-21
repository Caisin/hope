// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelbuychapterrecord"
	"hope/apps/novel/internal/data/ent/socialuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelBuyChapterRecordCreate is the builder for creating a NovelBuyChapterRecord entity.
type NovelBuyChapterRecordCreate struct {
	config
	mutation *NovelBuyChapterRecordMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetUserId(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetUserId(i)
	return nbcrc
}

// SetUserName sets the "userName" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetUserName(s string) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetUserName(s)
	return nbcrc
}

// SetNillableUserName sets the "userName" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableUserName(s *string) *NovelBuyChapterRecordCreate {
	if s != nil {
		nbcrc.SetUserName(*s)
	}
	return nbcrc
}

// SetChapterId sets the "chapterId" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetChapterId(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetChapterId(i)
	return nbcrc
}

// SetNillableChapterId sets the "chapterId" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableChapterId(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetChapterId(*i)
	}
	return nbcrc
}

// SetChapterOrderNum sets the "chapterOrderNum" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetChapterOrderNum(i int32) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetChapterOrderNum(i)
	return nbcrc
}

// SetNillableChapterOrderNum sets the "chapterOrderNum" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableChapterOrderNum(i *int32) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetChapterOrderNum(*i)
	}
	return nbcrc
}

// SetNovelId sets the "novelId" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetNovelId(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetNovelId(i)
	return nbcrc
}

// SetNillableNovelId sets the "novelId" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableNovelId(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetNovelId(*i)
	}
	return nbcrc
}

// SetNovelName sets the "novelName" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetNovelName(s string) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetNovelName(s)
	return nbcrc
}

// SetNillableNovelName sets the "novelName" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableNovelName(s *string) *NovelBuyChapterRecordCreate {
	if s != nil {
		nbcrc.SetNovelName(*s)
	}
	return nbcrc
}

// SetChapterName sets the "chapterName" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetChapterName(s string) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetChapterName(s)
	return nbcrc
}

// SetNillableChapterName sets the "chapterName" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableChapterName(s *string) *NovelBuyChapterRecordCreate {
	if s != nil {
		nbcrc.SetChapterName(*s)
	}
	return nbcrc
}

// SetIsSvip sets the "isSvip" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetIsSvip(b bool) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetIsSvip(b)
	return nbcrc
}

// SetNillableIsSvip sets the "isSvip" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableIsSvip(b *bool) *NovelBuyChapterRecordCreate {
	if b != nil {
		nbcrc.SetIsSvip(*b)
	}
	return nbcrc
}

// SetCoin sets the "coin" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetCoin(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetCoin(i)
	return nbcrc
}

// SetNillableCoin sets the "coin" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableCoin(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetCoin(*i)
	}
	return nbcrc
}

// SetCoupon sets the "coupon" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetCoupon(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetCoupon(i)
	return nbcrc
}

// SetNillableCoupon sets the "coupon" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableCoupon(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetCoupon(*i)
	}
	return nbcrc
}

// SetDiscount sets the "discount" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetDiscount(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetDiscount(i)
	return nbcrc
}

// SetNillableDiscount sets the "discount" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableDiscount(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetDiscount(*i)
	}
	return nbcrc
}

// SetRemark sets the "remark" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetRemark(s string) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetRemark(s)
	return nbcrc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableRemark(s *string) *NovelBuyChapterRecordCreate {
	if s != nil {
		nbcrc.SetRemark(*s)
	}
	return nbcrc
}

// SetCreatedAt sets the "createdAt" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetCreatedAt(t time.Time) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetCreatedAt(t)
	return nbcrc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableCreatedAt(t *time.Time) *NovelBuyChapterRecordCreate {
	if t != nil {
		nbcrc.SetCreatedAt(*t)
	}
	return nbcrc
}

// SetUpdatedAt sets the "updatedAt" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetUpdatedAt(t time.Time) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetUpdatedAt(t)
	return nbcrc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableUpdatedAt(t *time.Time) *NovelBuyChapterRecordCreate {
	if t != nil {
		nbcrc.SetUpdatedAt(*t)
	}
	return nbcrc
}

// SetCreateBy sets the "createBy" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetCreateBy(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetCreateBy(i)
	return nbcrc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableCreateBy(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetCreateBy(*i)
	}
	return nbcrc
}

// SetUpdateBy sets the "updateBy" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetUpdateBy(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetUpdateBy(i)
	return nbcrc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableUpdateBy(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetUpdateBy(*i)
	}
	return nbcrc
}

// SetTenantId sets the "tenantId" field.
func (nbcrc *NovelBuyChapterRecordCreate) SetTenantId(i int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetTenantId(i)
	return nbcrc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nbcrc *NovelBuyChapterRecordCreate) SetNillableTenantId(i *int64) *NovelBuyChapterRecordCreate {
	if i != nil {
		nbcrc.SetTenantId(*i)
	}
	return nbcrc
}

// SetUserID sets the "user" edge to the SocialUser entity by ID.
func (nbcrc *NovelBuyChapterRecordCreate) SetUserID(id int64) *NovelBuyChapterRecordCreate {
	nbcrc.mutation.SetUserID(id)
	return nbcrc
}

// SetUser sets the "user" edge to the SocialUser entity.
func (nbcrc *NovelBuyChapterRecordCreate) SetUser(s *SocialUser) *NovelBuyChapterRecordCreate {
	return nbcrc.SetUserID(s.ID)
}

// Mutation returns the NovelBuyChapterRecordMutation object of the builder.
func (nbcrc *NovelBuyChapterRecordCreate) Mutation() *NovelBuyChapterRecordMutation {
	return nbcrc.mutation
}

// Save creates the NovelBuyChapterRecord in the database.
func (nbcrc *NovelBuyChapterRecordCreate) Save(ctx context.Context) (*NovelBuyChapterRecord, error) {
	var (
		err  error
		node *NovelBuyChapterRecord
	)
	nbcrc.defaults()
	if len(nbcrc.hooks) == 0 {
		if err = nbcrc.check(); err != nil {
			return nil, err
		}
		node, err = nbcrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelBuyChapterRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nbcrc.check(); err != nil {
				return nil, err
			}
			nbcrc.mutation = mutation
			if node, err = nbcrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nbcrc.hooks) - 1; i >= 0; i-- {
			if nbcrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nbcrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nbcrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nbcrc *NovelBuyChapterRecordCreate) SaveX(ctx context.Context) *NovelBuyChapterRecord {
	v, err := nbcrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nbcrc *NovelBuyChapterRecordCreate) Exec(ctx context.Context) error {
	_, err := nbcrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbcrc *NovelBuyChapterRecordCreate) ExecX(ctx context.Context) {
	if err := nbcrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nbcrc *NovelBuyChapterRecordCreate) defaults() {
	if _, ok := nbcrc.mutation.CreatedAt(); !ok {
		v := novelbuychapterrecord.DefaultCreatedAt()
		nbcrc.mutation.SetCreatedAt(v)
	}
	if _, ok := nbcrc.mutation.UpdatedAt(); !ok {
		v := novelbuychapterrecord.DefaultUpdatedAt()
		nbcrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nbcrc.mutation.CreateBy(); !ok {
		v := novelbuychapterrecord.DefaultCreateBy
		nbcrc.mutation.SetCreateBy(v)
	}
	if _, ok := nbcrc.mutation.UpdateBy(); !ok {
		v := novelbuychapterrecord.DefaultUpdateBy
		nbcrc.mutation.SetUpdateBy(v)
	}
	if _, ok := nbcrc.mutation.TenantId(); !ok {
		v := novelbuychapterrecord.DefaultTenantId
		nbcrc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nbcrc *NovelBuyChapterRecordCreate) check() error {
	if _, ok := nbcrc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "userId"`)}
	}
	if _, ok := nbcrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := nbcrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := nbcrc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "createBy"`)}
	}
	if _, ok := nbcrc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "updateBy"`)}
	}
	if _, ok := nbcrc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "tenantId"`)}
	}
	if _, ok := nbcrc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (nbcrc *NovelBuyChapterRecordCreate) sqlSave(ctx context.Context) (*NovelBuyChapterRecord, error) {
	_node, _spec := nbcrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nbcrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (nbcrc *NovelBuyChapterRecordCreate) createSpec() (*NovelBuyChapterRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &NovelBuyChapterRecord{config: nbcrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: novelbuychapterrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbuychapterrecord.FieldID,
			},
		}
	)
	if value, ok := nbcrc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuychapterrecord.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := nbcrc.mutation.ChapterId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldChapterId,
		})
		_node.ChapterId = value
	}
	if value, ok := nbcrc.mutation.ChapterOrderNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novelbuychapterrecord.FieldChapterOrderNum,
		})
		_node.ChapterOrderNum = value
	}
	if value, ok := nbcrc.mutation.NovelId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldNovelId,
		})
		_node.NovelId = value
	}
	if value, ok := nbcrc.mutation.NovelName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuychapterrecord.FieldNovelName,
		})
		_node.NovelName = value
	}
	if value, ok := nbcrc.mutation.ChapterName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuychapterrecord.FieldChapterName,
		})
		_node.ChapterName = value
	}
	if value, ok := nbcrc.mutation.IsSvip(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: novelbuychapterrecord.FieldIsSvip,
		})
		_node.IsSvip = value
	}
	if value, ok := nbcrc.mutation.Coin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldCoin,
		})
		_node.Coin = value
	}
	if value, ok := nbcrc.mutation.Coupon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldCoupon,
		})
		_node.Coupon = value
	}
	if value, ok := nbcrc.mutation.Discount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldDiscount,
		})
		_node.Discount = value
	}
	if value, ok := nbcrc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelbuychapterrecord.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := nbcrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbuychapterrecord.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := nbcrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelbuychapterrecord.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := nbcrc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := nbcrc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := nbcrc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novelbuychapterrecord.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := nbcrc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novelbuychapterrecord.UserTable,
			Columns: []string{novelbuychapterrecord.UserColumn},
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

// NovelBuyChapterRecordCreateBulk is the builder for creating many NovelBuyChapterRecord entities in bulk.
type NovelBuyChapterRecordCreateBulk struct {
	config
	builders []*NovelBuyChapterRecordCreate
}

// Save creates the NovelBuyChapterRecord entities in the database.
func (nbcrcb *NovelBuyChapterRecordCreateBulk) Save(ctx context.Context) ([]*NovelBuyChapterRecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nbcrcb.builders))
	nodes := make([]*NovelBuyChapterRecord, len(nbcrcb.builders))
	mutators := make([]Mutator, len(nbcrcb.builders))
	for i := range nbcrcb.builders {
		func(i int, root context.Context) {
			builder := nbcrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NovelBuyChapterRecordMutation)
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
					_, err = mutators[i+1].Mutate(root, nbcrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nbcrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nbcrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nbcrcb *NovelBuyChapterRecordCreateBulk) SaveX(ctx context.Context) []*NovelBuyChapterRecord {
	v, err := nbcrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nbcrcb *NovelBuyChapterRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := nbcrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nbcrcb *NovelBuyChapterRecordCreateBulk) ExecX(ctx context.Context) {
	if err := nbcrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
