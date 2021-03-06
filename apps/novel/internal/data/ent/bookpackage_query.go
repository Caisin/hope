// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/bookpackage"
	"hope/apps/novel/internal/data/ent/novel"
	"hope/apps/novel/internal/data/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookPackageQuery is the builder for querying BookPackage entities.
type BookPackageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BookPackage
	// eager-loading edges.
	withBooks *NovelQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookPackageQuery builder.
func (bpq *BookPackageQuery) Where(ps ...predicate.BookPackage) *BookPackageQuery {
	bpq.predicates = append(bpq.predicates, ps...)
	return bpq
}

// Limit adds a limit step to the query.
func (bpq *BookPackageQuery) Limit(limit int) *BookPackageQuery {
	bpq.limit = &limit
	return bpq
}

// Offset adds an offset step to the query.
func (bpq *BookPackageQuery) Offset(offset int) *BookPackageQuery {
	bpq.offset = &offset
	return bpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bpq *BookPackageQuery) Unique(unique bool) *BookPackageQuery {
	bpq.unique = &unique
	return bpq
}

// Order adds an order step to the query.
func (bpq *BookPackageQuery) Order(o ...OrderFunc) *BookPackageQuery {
	bpq.order = append(bpq.order, o...)
	return bpq
}

// QueryBooks chains the current query on the "books" edge.
func (bpq *BookPackageQuery) QueryBooks() *NovelQuery {
	query := &NovelQuery{config: bpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookpackage.Table, bookpackage.FieldID, selector),
			sqlgraph.To(novel.Table, novel.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, bookpackage.BooksTable, bookpackage.BooksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BookPackage entity from the query.
// Returns a *NotFoundError when no BookPackage was found.
func (bpq *BookPackageQuery) First(ctx context.Context) (*BookPackage, error) {
	nodes, err := bpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bookpackage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bpq *BookPackageQuery) FirstX(ctx context.Context) *BookPackage {
	node, err := bpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BookPackage ID from the query.
// Returns a *NotFoundError when no BookPackage ID was found.
func (bpq *BookPackageQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bookpackage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bpq *BookPackageQuery) FirstIDX(ctx context.Context) int64 {
	id, err := bpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BookPackage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BookPackage entity is not found.
// Returns a *NotFoundError when no BookPackage entities are found.
func (bpq *BookPackageQuery) Only(ctx context.Context) (*BookPackage, error) {
	nodes, err := bpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bookpackage.Label}
	default:
		return nil, &NotSingularError{bookpackage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bpq *BookPackageQuery) OnlyX(ctx context.Context) *BookPackage {
	node, err := bpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BookPackage ID in the query.
// Returns a *NotSingularError when exactly one BookPackage ID is not found.
// Returns a *NotFoundError when no entities are found.
func (bpq *BookPackageQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = &NotSingularError{bookpackage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bpq *BookPackageQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := bpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BookPackages.
func (bpq *BookPackageQuery) All(ctx context.Context) ([]*BookPackage, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bpq *BookPackageQuery) AllX(ctx context.Context) []*BookPackage {
	nodes, err := bpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BookPackage IDs.
func (bpq *BookPackageQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := bpq.Select(bookpackage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bpq *BookPackageQuery) IDsX(ctx context.Context) []int64 {
	ids, err := bpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bpq *BookPackageQuery) Count(ctx context.Context) (int, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bpq *BookPackageQuery) CountX(ctx context.Context) int {
	count, err := bpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bpq *BookPackageQuery) Exist(ctx context.Context) (bool, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bpq *BookPackageQuery) ExistX(ctx context.Context) bool {
	exist, err := bpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookPackageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bpq *BookPackageQuery) Clone() *BookPackageQuery {
	if bpq == nil {
		return nil
	}
	return &BookPackageQuery{
		config:     bpq.config,
		limit:      bpq.limit,
		offset:     bpq.offset,
		order:      append([]OrderFunc{}, bpq.order...),
		predicates: append([]predicate.BookPackage{}, bpq.predicates...),
		withBooks:  bpq.withBooks.Clone(),
		// clone intermediate query.
		sql:  bpq.sql.Clone(),
		path: bpq.path,
	}
}

// WithBooks tells the query-builder to eager-load the nodes that are connected to
// the "books" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BookPackageQuery) WithBooks(opts ...func(*NovelQuery)) *BookPackageQuery {
	query := &NovelQuery{config: bpq.config}
	for _, opt := range opts {
		opt(query)
	}
	bpq.withBooks = query
	return bpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ActivityCode string `json:"activityCode,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BookPackage.Query().
//		GroupBy(bookpackage.FieldActivityCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bpq *BookPackageQuery) GroupBy(field string, fields ...string) *BookPackageGroupBy {
	group := &BookPackageGroupBy{config: bpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bpq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ActivityCode string `json:"activityCode,omitempty"`
//	}
//
//	client.BookPackage.Query().
//		Select(bookpackage.FieldActivityCode).
//		Scan(ctx, &v)
//
func (bpq *BookPackageQuery) Select(fields ...string) *BookPackageSelect {
	bpq.fields = append(bpq.fields, fields...)
	return &BookPackageSelect{BookPackageQuery: bpq}
}

func (bpq *BookPackageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bpq.fields {
		if !bookpackage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bpq.path != nil {
		prev, err := bpq.path(ctx)
		if err != nil {
			return err
		}
		bpq.sql = prev
	}
	return nil
}

func (bpq *BookPackageQuery) sqlAll(ctx context.Context) ([]*BookPackage, error) {
	var (
		nodes       = []*BookPackage{}
		_spec       = bpq.querySpec()
		loadedTypes = [1]bool{
			bpq.withBooks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BookPackage{config: bpq.config}
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
	if err := sqlgraph.QueryNodes(ctx, bpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bpq.withBooks; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*BookPackage, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Books = []*Novel{}
		}
		var (
			edgeids []int64
			edges   = make(map[int64][]*BookPackage)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   bookpackage.BooksTable,
				Columns: bookpackage.BooksPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(bookpackage.BooksPrimaryKey[1], fks...))
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
		if err := sqlgraph.QueryEdges(ctx, bpq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "books": %w`, err)
		}
		query.Where(novel.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "books" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Books = append(nodes[i].Edges.Books, n)
			}
		}
	}

	return nodes, nil
}

func (bpq *BookPackageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bpq.querySpec()
	_spec.Node.Columns = bpq.fields
	if len(bpq.fields) > 0 {
		_spec.Unique = bpq.unique != nil && *bpq.unique
	}
	return sqlgraph.CountNodes(ctx, bpq.driver, _spec)
}

func (bpq *BookPackageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bpq *BookPackageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bookpackage.Table,
			Columns: bookpackage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: bookpackage.FieldID,
			},
		},
		From:   bpq.sql,
		Unique: true,
	}
	if unique := bpq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bpq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookpackage.FieldID)
		for i := range fields {
			if fields[i] != bookpackage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bpq *BookPackageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bpq.driver.Dialect())
	t1 := builder.Table(bookpackage.Table)
	columns := bpq.fields
	if len(columns) == 0 {
		columns = bookpackage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bpq.sql != nil {
		selector = bpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bpq.unique != nil && *bpq.unique {
		selector.Distinct()
	}
	for _, p := range bpq.predicates {
		p(selector)
	}
	for _, p := range bpq.order {
		p(selector)
	}
	if offset := bpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookPackageGroupBy is the group-by builder for BookPackage entities.
type BookPackageGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bpgb *BookPackageGroupBy) Aggregate(fns ...AggregateFunc) *BookPackageGroupBy {
	bpgb.fns = append(bpgb.fns, fns...)
	return bpgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bpgb *BookPackageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bpgb.path(ctx)
	if err != nil {
		return err
	}
	bpgb.sql = query
	return bpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BookPackageGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) StringsX(ctx context.Context) []string {
	v, err := bpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) StringX(ctx context.Context) string {
	v, err := bpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BookPackageGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) IntsX(ctx context.Context) []int {
	v, err := bpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) IntX(ctx context.Context) int {
	v, err := bpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BookPackageGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BookPackageGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bpgb *BookPackageGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bpgb *BookPackageGroupBy) BoolX(ctx context.Context) bool {
	v, err := bpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bpgb *BookPackageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bpgb.fields {
		if !bookpackage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bpgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bpgb *BookPackageGroupBy) sqlQuery() *sql.Selector {
	selector := bpgb.sql.Select()
	aggregation := make([]string, 0, len(bpgb.fns))
	for _, fn := range bpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bpgb.fields)+len(bpgb.fns))
		for _, f := range bpgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bpgb.fields...)...)
}

// BookPackageSelect is the builder for selecting fields of BookPackage entities.
type BookPackageSelect struct {
	*BookPackageQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bps *BookPackageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bps.prepareQuery(ctx); err != nil {
		return err
	}
	bps.sql = bps.BookPackageQuery.sqlQuery(ctx)
	return bps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bps *BookPackageSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BookPackageSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bps *BookPackageSelect) StringsX(ctx context.Context) []string {
	v, err := bps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bps *BookPackageSelect) StringX(ctx context.Context) string {
	v, err := bps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BookPackageSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bps *BookPackageSelect) IntsX(ctx context.Context) []int {
	v, err := bps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bps *BookPackageSelect) IntX(ctx context.Context) int {
	v, err := bps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BookPackageSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bps *BookPackageSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bps *BookPackageSelect) Float64X(ctx context.Context) float64 {
	v, err := bps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BookPackageSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bps *BookPackageSelect) BoolsX(ctx context.Context) []bool {
	v, err := bps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bps *BookPackageSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookpackage.Label}
	default:
		err = fmt.Errorf("ent: BookPackageSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bps *BookPackageSelect) BoolX(ctx context.Context) bool {
	v, err := bps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bps *BookPackageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bps.sql.Query()
	if err := bps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
