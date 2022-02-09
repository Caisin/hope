// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/assetchangelog"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AssetChangeLogQuery is the builder for querying AssetChangeLog entities.
type AssetChangeLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AssetChangeLog
	// eager-loading edges.
	withUser *SocialUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AssetChangeLogQuery builder.
func (aclq *AssetChangeLogQuery) Where(ps ...predicate.AssetChangeLog) *AssetChangeLogQuery {
	aclq.predicates = append(aclq.predicates, ps...)
	return aclq
}

// Limit adds a limit step to the query.
func (aclq *AssetChangeLogQuery) Limit(limit int) *AssetChangeLogQuery {
	aclq.limit = &limit
	return aclq
}

// Offset adds an offset step to the query.
func (aclq *AssetChangeLogQuery) Offset(offset int) *AssetChangeLogQuery {
	aclq.offset = &offset
	return aclq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aclq *AssetChangeLogQuery) Unique(unique bool) *AssetChangeLogQuery {
	aclq.unique = &unique
	return aclq
}

// Order adds an order step to the query.
func (aclq *AssetChangeLogQuery) Order(o ...OrderFunc) *AssetChangeLogQuery {
	aclq.order = append(aclq.order, o...)
	return aclq
}

// QueryUser chains the current query on the "user" edge.
func (aclq *AssetChangeLogQuery) QueryUser() *SocialUserQuery {
	query := &SocialUserQuery{config: aclq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aclq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aclq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(assetchangelog.Table, assetchangelog.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, assetchangelog.UserTable, assetchangelog.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aclq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AssetChangeLog entity from the query.
// Returns a *NotFoundError when no AssetChangeLog was found.
func (aclq *AssetChangeLogQuery) First(ctx context.Context) (*AssetChangeLog, error) {
	nodes, err := aclq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{assetchangelog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) FirstX(ctx context.Context) *AssetChangeLog {
	node, err := aclq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AssetChangeLog ID from the query.
// Returns a *NotFoundError when no AssetChangeLog ID was found.
func (aclq *AssetChangeLogQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aclq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{assetchangelog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) FirstIDX(ctx context.Context) int64 {
	id, err := aclq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AssetChangeLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AssetChangeLog entity is not found.
// Returns a *NotFoundError when no AssetChangeLog entities are found.
func (aclq *AssetChangeLogQuery) Only(ctx context.Context) (*AssetChangeLog, error) {
	nodes, err := aclq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{assetchangelog.Label}
	default:
		return nil, &NotSingularError{assetchangelog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) OnlyX(ctx context.Context) *AssetChangeLog {
	node, err := aclq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AssetChangeLog ID in the query.
// Returns a *NotSingularError when exactly one AssetChangeLog ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aclq *AssetChangeLogQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aclq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = &NotSingularError{assetchangelog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := aclq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AssetChangeLogs.
func (aclq *AssetChangeLogQuery) All(ctx context.Context) ([]*AssetChangeLog, error) {
	if err := aclq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aclq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) AllX(ctx context.Context) []*AssetChangeLog {
	nodes, err := aclq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AssetChangeLog IDs.
func (aclq *AssetChangeLogQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := aclq.Select(assetchangelog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) IDsX(ctx context.Context) []int64 {
	ids, err := aclq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aclq *AssetChangeLogQuery) Count(ctx context.Context) (int, error) {
	if err := aclq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aclq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) CountX(ctx context.Context) int {
	count, err := aclq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aclq *AssetChangeLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := aclq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aclq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aclq *AssetChangeLogQuery) ExistX(ctx context.Context) bool {
	exist, err := aclq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AssetChangeLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aclq *AssetChangeLogQuery) Clone() *AssetChangeLogQuery {
	if aclq == nil {
		return nil
	}
	return &AssetChangeLogQuery{
		config:     aclq.config,
		limit:      aclq.limit,
		offset:     aclq.offset,
		order:      append([]OrderFunc{}, aclq.order...),
		predicates: append([]predicate.AssetChangeLog{}, aclq.predicates...),
		withUser:   aclq.withUser.Clone(),
		// clone intermediate query.
		sql:  aclq.sql.Clone(),
		path: aclq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aclq *AssetChangeLogQuery) WithUser(opts ...func(*SocialUserQuery)) *AssetChangeLogQuery {
	query := &SocialUserQuery{config: aclq.config}
	for _, opt := range opts {
		opt(query)
	}
	aclq.withUser = query
	return aclq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderId string `json:"orderId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AssetChangeLog.Query().
//		GroupBy(assetchangelog.FieldOrderId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aclq *AssetChangeLogQuery) GroupBy(field string, fields ...string) *AssetChangeLogGroupBy {
	group := &AssetChangeLogGroupBy{config: aclq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aclq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aclq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderId string `json:"orderId,omitempty"`
//	}
//
//	client.AssetChangeLog.Query().
//		Select(assetchangelog.FieldOrderId).
//		Scan(ctx, &v)
//
func (aclq *AssetChangeLogQuery) Select(fields ...string) *AssetChangeLogSelect {
	aclq.fields = append(aclq.fields, fields...)
	return &AssetChangeLogSelect{AssetChangeLogQuery: aclq}
}

func (aclq *AssetChangeLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aclq.fields {
		if !assetchangelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aclq.path != nil {
		prev, err := aclq.path(ctx)
		if err != nil {
			return err
		}
		aclq.sql = prev
	}
	return nil
}

func (aclq *AssetChangeLogQuery) sqlAll(ctx context.Context) ([]*AssetChangeLog, error) {
	var (
		nodes       = []*AssetChangeLog{}
		_spec       = aclq.querySpec()
		loadedTypes = [1]bool{
			aclq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AssetChangeLog{config: aclq.config}
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
	if err := sqlgraph.QueryNodes(ctx, aclq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aclq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AssetChangeLog)
		for i := range nodes {
			fk := nodes[i].UserId
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(socialuser.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "userId" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (aclq *AssetChangeLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aclq.querySpec()
	_spec.Node.Columns = aclq.fields
	if len(aclq.fields) > 0 {
		_spec.Unique = aclq.unique != nil && *aclq.unique
	}
	return sqlgraph.CountNodes(ctx, aclq.driver, _spec)
}

func (aclq *AssetChangeLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aclq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aclq *AssetChangeLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   assetchangelog.Table,
			Columns: assetchangelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: assetchangelog.FieldID,
			},
		},
		From:   aclq.sql,
		Unique: true,
	}
	if unique := aclq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aclq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, assetchangelog.FieldID)
		for i := range fields {
			if fields[i] != assetchangelog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aclq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aclq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aclq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aclq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aclq *AssetChangeLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aclq.driver.Dialect())
	t1 := builder.Table(assetchangelog.Table)
	columns := aclq.fields
	if len(columns) == 0 {
		columns = assetchangelog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aclq.sql != nil {
		selector = aclq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aclq.unique != nil && *aclq.unique {
		selector.Distinct()
	}
	for _, p := range aclq.predicates {
		p(selector)
	}
	for _, p := range aclq.order {
		p(selector)
	}
	if offset := aclq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aclq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AssetChangeLogGroupBy is the group-by builder for AssetChangeLog entities.
type AssetChangeLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aclgb *AssetChangeLogGroupBy) Aggregate(fns ...AggregateFunc) *AssetChangeLogGroupBy {
	aclgb.fns = append(aclgb.fns, fns...)
	return aclgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aclgb *AssetChangeLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aclgb.path(ctx)
	if err != nil {
		return err
	}
	aclgb.sql = query
	return aclgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aclgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aclgb.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aclgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := aclgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aclgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) StringX(ctx context.Context) string {
	v, err := aclgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aclgb.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aclgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := aclgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aclgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) IntX(ctx context.Context) int {
	v, err := aclgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aclgb.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aclgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aclgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aclgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aclgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aclgb.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aclgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aclgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aclgb *AssetChangeLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aclgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aclgb *AssetChangeLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := aclgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aclgb *AssetChangeLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aclgb.fields {
		if !assetchangelog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aclgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aclgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aclgb *AssetChangeLogGroupBy) sqlQuery() *sql.Selector {
	selector := aclgb.sql.Select()
	aggregation := make([]string, 0, len(aclgb.fns))
	for _, fn := range aclgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aclgb.fields)+len(aclgb.fns))
		for _, f := range aclgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aclgb.fields...)...)
}

// AssetChangeLogSelect is the builder for selecting fields of AssetChangeLog entities.
type AssetChangeLogSelect struct {
	*AssetChangeLogQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (acls *AssetChangeLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := acls.prepareQuery(ctx); err != nil {
		return err
	}
	acls.sql = acls.AssetChangeLogQuery.sqlQuery(ctx)
	return acls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acls *AssetChangeLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := acls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(acls.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := acls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acls *AssetChangeLogSelect) StringsX(ctx context.Context) []string {
	v, err := acls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acls *AssetChangeLogSelect) StringX(ctx context.Context) string {
	v, err := acls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(acls.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := acls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acls *AssetChangeLogSelect) IntsX(ctx context.Context) []int {
	v, err := acls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acls *AssetChangeLogSelect) IntX(ctx context.Context) int {
	v, err := acls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(acls.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := acls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acls *AssetChangeLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := acls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acls *AssetChangeLogSelect) Float64X(ctx context.Context) float64 {
	v, err := acls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(acls.fields) > 1 {
		return nil, errors.New("ent: AssetChangeLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := acls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acls *AssetChangeLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := acls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (acls *AssetChangeLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{assetchangelog.Label}
	default:
		err = fmt.Errorf("ent: AssetChangeLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acls *AssetChangeLogSelect) BoolX(ctx context.Context) bool {
	v, err := acls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acls *AssetChangeLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := acls.sql.Query()
	if err := acls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
