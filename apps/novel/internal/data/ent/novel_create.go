// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/bookpackage"
	"hope/apps/novel/internal/data/ent/novel"
	"hope/apps/novel/internal/data/ent/novelchapter"
	"hope/apps/novel/internal/data/ent/novelclassify"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelCreate is the builder for creating a Novel entity.
type NovelCreate struct {
	config
	mutation *NovelMutation
	hooks    []Hook
}

// SetClassifyId sets the "classifyId" field.
func (nc *NovelCreate) SetClassifyId(i int64) *NovelCreate {
	nc.mutation.SetClassifyId(i)
	return nc
}

// SetNillableClassifyId sets the "classifyId" field if the given value is not nil.
func (nc *NovelCreate) SetNillableClassifyId(i *int64) *NovelCreate {
	if i != nil {
		nc.SetClassifyId(*i)
	}
	return nc
}

// SetClassifyName sets the "classifyName" field.
func (nc *NovelCreate) SetClassifyName(s string) *NovelCreate {
	nc.mutation.SetClassifyName(s)
	return nc
}

// SetNillableClassifyName sets the "classifyName" field if the given value is not nil.
func (nc *NovelCreate) SetNillableClassifyName(s *string) *NovelCreate {
	if s != nil {
		nc.SetClassifyName(*s)
	}
	return nc
}

// SetAuthorId sets the "authorId" field.
func (nc *NovelCreate) SetAuthorId(s string) *NovelCreate {
	nc.mutation.SetAuthorId(s)
	return nc
}

// SetNillableAuthorId sets the "authorId" field if the given value is not nil.
func (nc *NovelCreate) SetNillableAuthorId(s *string) *NovelCreate {
	if s != nil {
		nc.SetAuthorId(*s)
	}
	return nc
}

// SetTitle sets the "title" field.
func (nc *NovelCreate) SetTitle(s string) *NovelCreate {
	nc.mutation.SetTitle(s)
	return nc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (nc *NovelCreate) SetNillableTitle(s *string) *NovelCreate {
	if s != nil {
		nc.SetTitle(*s)
	}
	return nc
}

// SetSummary sets the "summary" field.
func (nc *NovelCreate) SetSummary(s string) *NovelCreate {
	nc.mutation.SetSummary(s)
	return nc
}

// SetNillableSummary sets the "summary" field if the given value is not nil.
func (nc *NovelCreate) SetNillableSummary(s *string) *NovelCreate {
	if s != nil {
		nc.SetSummary(*s)
	}
	return nc
}

// SetAuthor sets the "author" field.
func (nc *NovelCreate) SetAuthor(s string) *NovelCreate {
	nc.mutation.SetAuthor(s)
	return nc
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (nc *NovelCreate) SetNillableAuthor(s *string) *NovelCreate {
	if s != nil {
		nc.SetAuthor(*s)
	}
	return nc
}

// SetAnchor sets the "anchor" field.
func (nc *NovelCreate) SetAnchor(s string) *NovelCreate {
	nc.mutation.SetAnchor(s)
	return nc
}

// SetNillableAnchor sets the "anchor" field if the given value is not nil.
func (nc *NovelCreate) SetNillableAnchor(s *string) *NovelCreate {
	if s != nil {
		nc.SetAnchor(*s)
	}
	return nc
}

// SetHits sets the "hits" field.
func (nc *NovelCreate) SetHits(i int64) *NovelCreate {
	nc.mutation.SetHits(i)
	return nc
}

// SetNillableHits sets the "hits" field if the given value is not nil.
func (nc *NovelCreate) SetNillableHits(i *int64) *NovelCreate {
	if i != nil {
		nc.SetHits(*i)
	}
	return nc
}

// SetKeywords sets the "keywords" field.
func (nc *NovelCreate) SetKeywords(s string) *NovelCreate {
	nc.mutation.SetKeywords(s)
	return nc
}

// SetNillableKeywords sets the "keywords" field if the given value is not nil.
func (nc *NovelCreate) SetNillableKeywords(s *string) *NovelCreate {
	if s != nil {
		nc.SetKeywords(*s)
	}
	return nc
}

// SetSource sets the "source" field.
func (nc *NovelCreate) SetSource(s string) *NovelCreate {
	nc.mutation.SetSource(s)
	return nc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (nc *NovelCreate) SetNillableSource(s *string) *NovelCreate {
	if s != nil {
		nc.SetSource(*s)
	}
	return nc
}

// SetScore sets the "score" field.
func (nc *NovelCreate) SetScore(i int32) *NovelCreate {
	nc.mutation.SetScore(i)
	return nc
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (nc *NovelCreate) SetNillableScore(i *int32) *NovelCreate {
	if i != nil {
		nc.SetScore(*i)
	}
	return nc
}

// SetCover sets the "cover" field.
func (nc *NovelCreate) SetCover(s string) *NovelCreate {
	nc.mutation.SetCover(s)
	return nc
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (nc *NovelCreate) SetNillableCover(s *string) *NovelCreate {
	if s != nil {
		nc.SetCover(*s)
	}
	return nc
}

// SetTagIds sets the "tagIds" field.
func (nc *NovelCreate) SetTagIds(s string) *NovelCreate {
	nc.mutation.SetTagIds(s)
	return nc
}

// SetNillableTagIds sets the "tagIds" field if the given value is not nil.
func (nc *NovelCreate) SetNillableTagIds(s *string) *NovelCreate {
	if s != nil {
		nc.SetTagIds(*s)
	}
	return nc
}

// SetWordNum sets the "wordNum" field.
func (nc *NovelCreate) SetWordNum(i int32) *NovelCreate {
	nc.mutation.SetWordNum(i)
	return nc
}

// SetNillableWordNum sets the "wordNum" field if the given value is not nil.
func (nc *NovelCreate) SetNillableWordNum(i *int32) *NovelCreate {
	if i != nil {
		nc.SetWordNum(*i)
	}
	return nc
}

// SetFreeNum sets the "freeNum" field.
func (nc *NovelCreate) SetFreeNum(i int32) *NovelCreate {
	nc.mutation.SetFreeNum(i)
	return nc
}

// SetNillableFreeNum sets the "freeNum" field if the given value is not nil.
func (nc *NovelCreate) SetNillableFreeNum(i *int32) *NovelCreate {
	if i != nil {
		nc.SetFreeNum(*i)
	}
	return nc
}

// SetOnlineState sets the "onlineState" field.
func (nc *NovelCreate) SetOnlineState(i int32) *NovelCreate {
	nc.mutation.SetOnlineState(i)
	return nc
}

// SetNillableOnlineState sets the "onlineState" field if the given value is not nil.
func (nc *NovelCreate) SetNillableOnlineState(i *int32) *NovelCreate {
	if i != nil {
		nc.SetOnlineState(*i)
	}
	return nc
}

// SetPrice sets the "price" field.
func (nc *NovelCreate) SetPrice(i int64) *NovelCreate {
	nc.mutation.SetPrice(i)
	return nc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (nc *NovelCreate) SetNillablePrice(i *int64) *NovelCreate {
	if i != nil {
		nc.SetPrice(*i)
	}
	return nc
}

// SetPublish sets the "publish" field.
func (nc *NovelCreate) SetPublish(i int32) *NovelCreate {
	nc.mutation.SetPublish(i)
	return nc
}

// SetNillablePublish sets the "publish" field if the given value is not nil.
func (nc *NovelCreate) SetNillablePublish(i *int32) *NovelCreate {
	if i != nil {
		nc.SetPublish(*i)
	}
	return nc
}

// SetOriginalPrice sets the "originalPrice" field.
func (nc *NovelCreate) SetOriginalPrice(i int64) *NovelCreate {
	nc.mutation.SetOriginalPrice(i)
	return nc
}

// SetNillableOriginalPrice sets the "originalPrice" field if the given value is not nil.
func (nc *NovelCreate) SetNillableOriginalPrice(i *int64) *NovelCreate {
	if i != nil {
		nc.SetOriginalPrice(*i)
	}
	return nc
}

// SetChapterPrice sets the "chapterPrice" field.
func (nc *NovelCreate) SetChapterPrice(i int32) *NovelCreate {
	nc.mutation.SetChapterPrice(i)
	return nc
}

// SetNillableChapterPrice sets the "chapterPrice" field if the given value is not nil.
func (nc *NovelCreate) SetNillableChapterPrice(i *int32) *NovelCreate {
	if i != nil {
		nc.SetChapterPrice(*i)
	}
	return nc
}

// SetChapterCount sets the "chapterCount" field.
func (nc *NovelCreate) SetChapterCount(i int32) *NovelCreate {
	nc.mutation.SetChapterCount(i)
	return nc
}

// SetNillableChapterCount sets the "chapterCount" field if the given value is not nil.
func (nc *NovelCreate) SetNillableChapterCount(i *int32) *NovelCreate {
	if i != nil {
		nc.SetChapterCount(*i)
	}
	return nc
}

// SetSignType sets the "signType" field.
func (nc *NovelCreate) SetSignType(i int32) *NovelCreate {
	nc.mutation.SetSignType(i)
	return nc
}

// SetNillableSignType sets the "signType" field if the given value is not nil.
func (nc *NovelCreate) SetNillableSignType(i *int32) *NovelCreate {
	if i != nil {
		nc.SetSignType(*i)
	}
	return nc
}

// SetSignDate sets the "signDate" field.
func (nc *NovelCreate) SetSignDate(t time.Time) *NovelCreate {
	nc.mutation.SetSignDate(t)
	return nc
}

// SetNillableSignDate sets the "signDate" field if the given value is not nil.
func (nc *NovelCreate) SetNillableSignDate(t *time.Time) *NovelCreate {
	if t != nil {
		nc.SetSignDate(*t)
	}
	return nc
}

// SetLeadingMan sets the "leadingMan" field.
func (nc *NovelCreate) SetLeadingMan(s string) *NovelCreate {
	nc.mutation.SetLeadingMan(s)
	return nc
}

// SetNillableLeadingMan sets the "leadingMan" field if the given value is not nil.
func (nc *NovelCreate) SetNillableLeadingMan(s *string) *NovelCreate {
	if s != nil {
		nc.SetLeadingMan(*s)
	}
	return nc
}

// SetLeadingLady sets the "leadingLady" field.
func (nc *NovelCreate) SetLeadingLady(s string) *NovelCreate {
	nc.mutation.SetLeadingLady(s)
	return nc
}

// SetNillableLeadingLady sets the "leadingLady" field if the given value is not nil.
func (nc *NovelCreate) SetNillableLeadingLady(s *string) *NovelCreate {
	if s != nil {
		nc.SetLeadingLady(*s)
	}
	return nc
}

// SetRemark sets the "remark" field.
func (nc *NovelCreate) SetRemark(s string) *NovelCreate {
	nc.mutation.SetRemark(s)
	return nc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (nc *NovelCreate) SetNillableRemark(s *string) *NovelCreate {
	if s != nil {
		nc.SetRemark(*s)
	}
	return nc
}

// SetMediaKey sets the "mediaKey" field.
func (nc *NovelCreate) SetMediaKey(s string) *NovelCreate {
	nc.mutation.SetMediaKey(s)
	return nc
}

// SetNillableMediaKey sets the "mediaKey" field if the given value is not nil.
func (nc *NovelCreate) SetNillableMediaKey(s *string) *NovelCreate {
	if s != nil {
		nc.SetMediaKey(*s)
	}
	return nc
}

// SetCreatedAt sets the "createdAt" field.
func (nc *NovelCreate) SetCreatedAt(t time.Time) *NovelCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (nc *NovelCreate) SetNillableCreatedAt(t *time.Time) *NovelCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updatedAt" field.
func (nc *NovelCreate) SetUpdatedAt(t time.Time) *NovelCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (nc *NovelCreate) SetNillableUpdatedAt(t *time.Time) *NovelCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetCreateBy sets the "createBy" field.
func (nc *NovelCreate) SetCreateBy(i int64) *NovelCreate {
	nc.mutation.SetCreateBy(i)
	return nc
}

// SetNillableCreateBy sets the "createBy" field if the given value is not nil.
func (nc *NovelCreate) SetNillableCreateBy(i *int64) *NovelCreate {
	if i != nil {
		nc.SetCreateBy(*i)
	}
	return nc
}

// SetUpdateBy sets the "updateBy" field.
func (nc *NovelCreate) SetUpdateBy(i int64) *NovelCreate {
	nc.mutation.SetUpdateBy(i)
	return nc
}

// SetNillableUpdateBy sets the "updateBy" field if the given value is not nil.
func (nc *NovelCreate) SetNillableUpdateBy(i *int64) *NovelCreate {
	if i != nil {
		nc.SetUpdateBy(*i)
	}
	return nc
}

// SetTenantId sets the "tenantId" field.
func (nc *NovelCreate) SetTenantId(i int64) *NovelCreate {
	nc.mutation.SetTenantId(i)
	return nc
}

// SetNillableTenantId sets the "tenantId" field if the given value is not nil.
func (nc *NovelCreate) SetNillableTenantId(i *int64) *NovelCreate {
	if i != nil {
		nc.SetTenantId(*i)
	}
	return nc
}

// AddChapterIDs adds the "chapters" edge to the NovelChapter entity by IDs.
func (nc *NovelCreate) AddChapterIDs(ids ...int64) *NovelCreate {
	nc.mutation.AddChapterIDs(ids...)
	return nc
}

// AddChapters adds the "chapters" edges to the NovelChapter entity.
func (nc *NovelCreate) AddChapters(n ...*NovelChapter) *NovelCreate {
	ids := make([]int64, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nc.AddChapterIDs(ids...)
}

// AddPkgIDs adds the "pkgs" edge to the BookPackage entity by IDs.
func (nc *NovelCreate) AddPkgIDs(ids ...int64) *NovelCreate {
	nc.mutation.AddPkgIDs(ids...)
	return nc
}

// AddPkgs adds the "pkgs" edges to the BookPackage entity.
func (nc *NovelCreate) AddPkgs(b ...*BookPackage) *NovelCreate {
	ids := make([]int64, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return nc.AddPkgIDs(ids...)
}

// SetClassifyID sets the "classify" edge to the NovelClassify entity by ID.
func (nc *NovelCreate) SetClassifyID(id int64) *NovelCreate {
	nc.mutation.SetClassifyID(id)
	return nc
}

// SetNillableClassifyID sets the "classify" edge to the NovelClassify entity by ID if the given value is not nil.
func (nc *NovelCreate) SetNillableClassifyID(id *int64) *NovelCreate {
	if id != nil {
		nc = nc.SetClassifyID(*id)
	}
	return nc
}

// SetClassify sets the "classify" edge to the NovelClassify entity.
func (nc *NovelCreate) SetClassify(n *NovelClassify) *NovelCreate {
	return nc.SetClassifyID(n.ID)
}

// Mutation returns the NovelMutation object of the builder.
func (nc *NovelCreate) Mutation() *NovelMutation {
	return nc.mutation
}

// Save creates the Novel in the database.
func (nc *NovelCreate) Save(ctx context.Context) (*Novel, error) {
	var (
		err  error
		node *Novel
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NovelCreate) SaveX(ctx context.Context) *Novel {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NovelCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NovelCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NovelCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := novel.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := novel.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.CreateBy(); !ok {
		v := novel.DefaultCreateBy
		nc.mutation.SetCreateBy(v)
	}
	if _, ok := nc.mutation.UpdateBy(); !ok {
		v := novel.DefaultUpdateBy
		nc.mutation.SetUpdateBy(v)
	}
	if _, ok := nc.mutation.TenantId(); !ok {
		v := novel.DefaultTenantId
		nc.mutation.SetTenantId(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NovelCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Novel.createdAt"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Novel.updatedAt"`)}
	}
	if _, ok := nc.mutation.CreateBy(); !ok {
		return &ValidationError{Name: "createBy", err: errors.New(`ent: missing required field "Novel.createBy"`)}
	}
	if _, ok := nc.mutation.UpdateBy(); !ok {
		return &ValidationError{Name: "updateBy", err: errors.New(`ent: missing required field "Novel.updateBy"`)}
	}
	if _, ok := nc.mutation.TenantId(); !ok {
		return &ValidationError{Name: "tenantId", err: errors.New(`ent: missing required field "Novel.tenantId"`)}
	}
	return nil
}

func (nc *NovelCreate) sqlSave(ctx context.Context) (*Novel, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (nc *NovelCreate) createSpec() (*Novel, *sqlgraph.CreateSpec) {
	var (
		_node = &Novel{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: novel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novel.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.ClassifyName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldClassifyName,
		})
		_node.ClassifyName = value
	}
	if value, ok := nc.mutation.AuthorId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldAuthorId,
		})
		_node.AuthorId = value
	}
	if value, ok := nc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := nc.mutation.Summary(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldSummary,
		})
		_node.Summary = value
	}
	if value, ok := nc.mutation.Author(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldAuthor,
		})
		_node.Author = value
	}
	if value, ok := nc.mutation.Anchor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldAnchor,
		})
		_node.Anchor = value
	}
	if value, ok := nc.mutation.Hits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novel.FieldHits,
		})
		_node.Hits = value
	}
	if value, ok := nc.mutation.Keywords(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldKeywords,
		})
		_node.Keywords = value
	}
	if value, ok := nc.mutation.Source(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldSource,
		})
		_node.Source = value
	}
	if value, ok := nc.mutation.Score(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldScore,
		})
		_node.Score = value
	}
	if value, ok := nc.mutation.Cover(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldCover,
		})
		_node.Cover = value
	}
	if value, ok := nc.mutation.TagIds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldTagIds,
		})
		_node.TagIds = value
	}
	if value, ok := nc.mutation.WordNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldWordNum,
		})
		_node.WordNum = value
	}
	if value, ok := nc.mutation.FreeNum(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldFreeNum,
		})
		_node.FreeNum = value
	}
	if value, ok := nc.mutation.OnlineState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldOnlineState,
		})
		_node.OnlineState = value
	}
	if value, ok := nc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novel.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := nc.mutation.Publish(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldPublish,
		})
		_node.Publish = value
	}
	if value, ok := nc.mutation.OriginalPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novel.FieldOriginalPrice,
		})
		_node.OriginalPrice = value
	}
	if value, ok := nc.mutation.ChapterPrice(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldChapterPrice,
		})
		_node.ChapterPrice = value
	}
	if value, ok := nc.mutation.ChapterCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldChapterCount,
		})
		_node.ChapterCount = value
	}
	if value, ok := nc.mutation.SignType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: novel.FieldSignType,
		})
		_node.SignType = value
	}
	if value, ok := nc.mutation.SignDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novel.FieldSignDate,
		})
		_node.SignDate = value
	}
	if value, ok := nc.mutation.LeadingMan(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldLeadingMan,
		})
		_node.LeadingMan = value
	}
	if value, ok := nc.mutation.LeadingLady(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldLeadingLady,
		})
		_node.LeadingLady = value
	}
	if value, ok := nc.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := nc.mutation.MediaKey(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novel.FieldMediaKey,
		})
		_node.MediaKey = value
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novel.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novel.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.CreateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novel.FieldCreateBy,
		})
		_node.CreateBy = value
	}
	if value, ok := nc.mutation.UpdateBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novel.FieldUpdateBy,
		})
		_node.UpdateBy = value
	}
	if value, ok := nc.mutation.TenantId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: novel.FieldTenantId,
		})
		_node.TenantId = value
	}
	if nodes := nc.mutation.ChaptersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   novel.ChaptersTable,
			Columns: []string{novel.ChaptersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: novelchapter.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.PkgsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   novel.PkgsTable,
			Columns: novel.PkgsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: bookpackage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ClassifyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   novel.ClassifyTable,
			Columns: []string{novel.ClassifyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: novelclassify.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ClassifyId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NovelCreateBulk is the builder for creating many Novel entities in bulk.
type NovelCreateBulk struct {
	config
	builders []*NovelCreate
}

// Save creates the Novel entities in the database.
func (ncb *NovelCreateBulk) Save(ctx context.Context) ([]*Novel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Novel, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NovelMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NovelCreateBulk) SaveX(ctx context.Context) []*Novel {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NovelCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NovelCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
