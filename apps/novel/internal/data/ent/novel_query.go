// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/bookpackage"
	"hope/apps/novel/internal/data/ent/novel"
	"hope/apps/novel/internal/data/ent/novelchapter"
	"hope/apps/novel/internal/data/ent/novelclassify"
	"hope/apps/novel/internal/data/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelQuery is the builder for querying Novel entities.
type NovelQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Novel
	// eager-loading edges.
	withChapters *NovelChapterQuery
	withPkgs     *BookPackageQuery
	withClassify *NovelClassifyQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NovelQuery builder.
func (nq *NovelQuery) Where(ps ...predicate.Novel) *NovelQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit adds a limit step to the query.
func (nq *NovelQuery) Limit(limit int) *NovelQuery {
	nq.limit = &limit
	return nq
}

// Offset adds an offset step to the query.
func (nq *NovelQuery) Offset(offset int) *NovelQuery {
	nq.offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NovelQuery) Unique(unique bool) *NovelQuery {
	nq.unique = &unique
	return nq
}

// Order adds an order step to the query.
func (nq *NovelQuery) Order(o ...OrderFunc) *NovelQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryChapters chains the current query on the "chapters" edge.
func (nq *NovelQuery) QueryChapters() *NovelChapterQuery {
	query := &NovelChapterQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novel.Table, novel.FieldID, selector),
			sqlgraph.To(novelchapter.Table, novelchapter.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, novel.ChaptersTable, novel.ChaptersColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPkgs chains the current query on the "pkgs" edge.
func (nq *NovelQuery) QueryPkgs() *BookPackageQuery {
	query := &BookPackageQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novel.Table, novel.FieldID, selector),
			sqlgraph.To(bookpackage.Table, bookpackage.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, novel.PkgsTable, novel.PkgsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClassify chains the current query on the "classify" edge.
func (nq *NovelQuery) QueryClassify() *NovelClassifyQuery {
	query := &NovelClassifyQuery{config: nq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novel.Table, novel.FieldID, selector),
			sqlgraph.To(novelclassify.Table, novelclassify.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, novel.ClassifyTable, novel.ClassifyColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Novel entity from the query.
// Returns a *NotFoundError when no Novel was found.
func (nq *NovelQuery) First(ctx context.Context) (*Novel, error) {
	nodes, err := nq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{novel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NovelQuery) FirstX(ctx context.Context) *Novel {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Novel ID from the query.
// Returns a *NotFoundError when no Novel ID was found.
func (nq *NovelQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{novel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NovelQuery) FirstIDX(ctx context.Context) int64 {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Novel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Novel entity is not found.
// Returns a *NotFoundError when no Novel entities are found.
func (nq *NovelQuery) Only(ctx context.Context) (*Novel, error) {
	nodes, err := nq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{novel.Label}
	default:
		return nil, &NotSingularError{novel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NovelQuery) OnlyX(ctx context.Context) *Novel {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Novel ID in the query.
// Returns a *NotSingularError when exactly one Novel ID is not found.
// Returns a *NotFoundError when no entities are found.
func (nq *NovelQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = &NotSingularError{novel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NovelQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Novels.
func (nq *NovelQuery) All(ctx context.Context) ([]*Novel, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nq *NovelQuery) AllX(ctx context.Context) []*Novel {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Novel IDs.
func (nq *NovelQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := nq.Select(novel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NovelQuery) IDsX(ctx context.Context) []int64 {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NovelQuery) Count(ctx context.Context) (int, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NovelQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NovelQuery) Exist(ctx context.Context) (bool, error) {
	if err := nq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NovelQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NovelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NovelQuery) Clone() *NovelQuery {
	if nq == nil {
		return nil
	}
	return &NovelQuery{
		config:       nq.config,
		limit:        nq.limit,
		offset:       nq.offset,
		order:        append([]OrderFunc{}, nq.order...),
		predicates:   append([]predicate.Novel{}, nq.predicates...),
		withChapters: nq.withChapters.Clone(),
		withPkgs:     nq.withPkgs.Clone(),
		withClassify: nq.withClassify.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithChapters tells the query-builder to eager-load the nodes that are connected to
// the "chapters" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NovelQuery) WithChapters(opts ...func(*NovelChapterQuery)) *NovelQuery {
	query := &NovelChapterQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withChapters = query
	return nq
}

// WithPkgs tells the query-builder to eager-load the nodes that are connected to
// the "pkgs" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NovelQuery) WithPkgs(opts ...func(*BookPackageQuery)) *NovelQuery {
	query := &BookPackageQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withPkgs = query
	return nq
}

// WithClassify tells the query-builder to eager-load the nodes that are connected to
// the "classify" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NovelQuery) WithClassify(opts ...func(*NovelClassifyQuery)) *NovelQuery {
	query := &NovelClassifyQuery{config: nq.config}
	for _, opt := range opts {
		opt(query)
	}
	nq.withClassify = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ClassifyId int64 `json:"classifyId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Novel.Query().
//		GroupBy(novel.FieldClassifyId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nq *NovelQuery) GroupBy(field string, fields ...string) *NovelGroupBy {
	group := &NovelGroupBy{config: nq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ClassifyId int64 `json:"classifyId,omitempty"`
//	}
//
//	client.Novel.Query().
//		Select(novel.FieldClassifyId).
//		Scan(ctx, &v)
//
func (nq *NovelQuery) Select(fields ...string) *NovelSelect {
	nq.fields = append(nq.fields, fields...)
	return &NovelSelect{NovelQuery: nq}
}

func (nq *NovelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nq.fields {
		if !novel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NovelQuery) sqlAll(ctx context.Context) ([]*Novel, error) {
	var (
		nodes       = []*Novel{}
		withFKs     = nq.withFKs
		_spec       = nq.querySpec()
		loadedTypes = [3]bool{
			nq.withChapters != nil,
			nq.withPkgs != nil,
			nq.withClassify != nil,
		}
	)
	if nq.withClassify != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, novel.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Novel{config: nq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nq.withChapters; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*Novel)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Chapters = []*NovelChapter{}
		}
		query.withFKs = true
		query.Where(predicate.NovelChapter(func(s *sql.Selector) {
			s.Where(sql.InValues(novel.ChaptersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.novel_chapters
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "novel_chapters" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "novel_chapters" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Chapters = append(node.Edges.Chapters, n)
		}
	}

	if query := nq.withPkgs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*Novel, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Pkgs = []*BookPackage{}
		}
		var (
			edgeids []int64
			edges   = make(map[int64][]*Novel)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   novel.PkgsTable,
				Columns: novel.PkgsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(novel.PkgsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.Int64
				inValue := ein.Int64
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, nq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "pkgs": %w`, err)
		}
		query.Where(bookpackage.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "pkgs" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Pkgs = append(nodes[i].Edges.Pkgs, n)
			}
		}
	}

	if query := nq.withClassify; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*Novel)
		for i := range nodes {
			if nodes[i].novel_classify_novels == nil {
				continue
			}
			fk := *nodes[i].novel_classify_novels
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(novelclassify.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "novel_classify_novels" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Classify = n
			}
		}
	}

	return nodes, nil
}

func (nq *NovelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NovelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nq *NovelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novel.Table,
			Columns: novel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novel.FieldID,
			},
		},
		From:   nq.sql,
		Unique: true,
	}
	if unique := nq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novel.FieldID)
		for i := range fields {
			if fields[i] != novel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NovelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(novel.Table)
	columns := nq.fields
	if len(columns) == 0 {
		columns = novel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NovelGroupBy is the group-by builder for Novel entities.
type NovelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NovelGroupBy) Aggregate(fns ...AggregateFunc) *NovelGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the group-by query and scans the result into the given value.
func (ngb *NovelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ngb.path(ctx)
	if err != nil {
		return err
	}
	ngb.sql = query
	return ngb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ngb *NovelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ngb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NovelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ngb *NovelGroupBy) StringsX(ctx context.Context) []string {
	v, err := ngb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ngb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ngb *NovelGroupBy) StringX(ctx context.Context) string {
	v, err := ngb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NovelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ngb *NovelGroupBy) IntsX(ctx context.Context) []int {
	v, err := ngb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ngb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ngb *NovelGroupBy) IntX(ctx context.Context) int {
	v, err := ngb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NovelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ngb *NovelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ngb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ngb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ngb *NovelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ngb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ngb.fields) > 1 {
		return nil, errors.New("ent: NovelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ngb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ngb *NovelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ngb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ngb *NovelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ngb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ngb *NovelGroupBy) BoolX(ctx context.Context) bool {
	v, err := ngb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ngb *NovelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ngb.fields {
		if !novel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ngb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ngb *NovelGroupBy) sqlQuery() *sql.Selector {
	selector := ngb.sql.Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ngb.fields)+len(ngb.fns))
		for _, f := range ngb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ngb.fields...)...)
}

// NovelSelect is the builder for selecting fields of Novel entities.
type NovelSelect struct {
	*NovelQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NovelSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	ns.sql = ns.NovelQuery.sqlQuery(ctx)
	return ns.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ns *NovelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ns.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NovelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ns *NovelSelect) StringsX(ctx context.Context) []string {
	v, err := ns.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ns.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ns *NovelSelect) StringX(ctx context.Context) string {
	v, err := ns.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NovelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ns *NovelSelect) IntsX(ctx context.Context) []int {
	v, err := ns.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ns.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ns *NovelSelect) IntX(ctx context.Context) int {
	v, err := ns.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NovelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ns *NovelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ns.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ns.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ns *NovelSelect) Float64X(ctx context.Context) float64 {
	v, err := ns.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ns.fields) > 1 {
		return nil, errors.New("ent: NovelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ns.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ns *NovelSelect) BoolsX(ctx context.Context) []bool {
	v, err := ns.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ns *NovelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ns.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novel.Label}
	default:
		err = fmt.Errorf("ent: NovelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ns *NovelSelect) BoolX(ctx context.Context) bool {
	v, err := ns.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ns *NovelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ns.sql.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
