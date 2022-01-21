// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/ambalance"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AmBalanceQuery is the builder for querying AmBalance entities.
type AmBalanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AmBalance
	// eager-loading edges.
	withUser *SocialUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AmBalanceQuery builder.
func (abq *AmBalanceQuery) Where(ps ...predicate.AmBalance) *AmBalanceQuery {
	abq.predicates = append(abq.predicates, ps...)
	return abq
}

// Limit adds a limit step to the query.
func (abq *AmBalanceQuery) Limit(limit int) *AmBalanceQuery {
	abq.limit = &limit
	return abq
}

// Offset adds an offset step to the query.
func (abq *AmBalanceQuery) Offset(offset int) *AmBalanceQuery {
	abq.offset = &offset
	return abq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (abq *AmBalanceQuery) Unique(unique bool) *AmBalanceQuery {
	abq.unique = &unique
	return abq
}

// Order adds an order step to the query.
func (abq *AmBalanceQuery) Order(o ...OrderFunc) *AmBalanceQuery {
	abq.order = append(abq.order, o...)
	return abq
}

// QueryUser chains the current query on the "user" edge.
func (abq *AmBalanceQuery) QueryUser() *SocialUserQuery {
	query := &SocialUserQuery{config: abq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := abq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := abq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ambalance.Table, ambalance.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ambalance.UserTable, ambalance.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(abq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AmBalance entity from the query.
// Returns a *NotFoundError when no AmBalance was found.
func (abq *AmBalanceQuery) First(ctx context.Context) (*AmBalance, error) {
	nodes, err := abq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ambalance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (abq *AmBalanceQuery) FirstX(ctx context.Context) *AmBalance {
	node, err := abq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AmBalance ID from the query.
// Returns a *NotFoundError when no AmBalance ID was found.
func (abq *AmBalanceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = abq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ambalance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (abq *AmBalanceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := abq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AmBalance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AmBalance entity is not found.
// Returns a *NotFoundError when no AmBalance entities are found.
func (abq *AmBalanceQuery) Only(ctx context.Context) (*AmBalance, error) {
	nodes, err := abq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ambalance.Label}
	default:
		return nil, &NotSingularError{ambalance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (abq *AmBalanceQuery) OnlyX(ctx context.Context) *AmBalance {
	node, err := abq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AmBalance ID in the query.
// Returns a *NotSingularError when exactly one AmBalance ID is not found.
// Returns a *NotFoundError when no entities are found.
func (abq *AmBalanceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = abq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = &NotSingularError{ambalance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (abq *AmBalanceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := abq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AmBalances.
func (abq *AmBalanceQuery) All(ctx context.Context) ([]*AmBalance, error) {
	if err := abq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return abq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (abq *AmBalanceQuery) AllX(ctx context.Context) []*AmBalance {
	nodes, err := abq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AmBalance IDs.
func (abq *AmBalanceQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := abq.Select(ambalance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (abq *AmBalanceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := abq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (abq *AmBalanceQuery) Count(ctx context.Context) (int, error) {
	if err := abq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return abq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (abq *AmBalanceQuery) CountX(ctx context.Context) int {
	count, err := abq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (abq *AmBalanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := abq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return abq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (abq *AmBalanceQuery) ExistX(ctx context.Context) bool {
	exist, err := abq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AmBalanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (abq *AmBalanceQuery) Clone() *AmBalanceQuery {
	if abq == nil {
		return nil
	}
	return &AmBalanceQuery{
		config:     abq.config,
		limit:      abq.limit,
		offset:     abq.offset,
		order:      append([]OrderFunc{}, abq.order...),
		predicates: append([]predicate.AmBalance{}, abq.predicates...),
		withUser:   abq.withUser.Clone(),
		// clone intermediate query.
		sql:  abq.sql.Clone(),
		path: abq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (abq *AmBalanceQuery) WithUser(opts ...func(*SocialUserQuery)) *AmBalanceQuery {
	query := &SocialUserQuery{config: abq.config}
	for _, opt := range opts {
		opt(query)
	}
	abq.withUser = query
	return abq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserId int64 `json:"userId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AmBalance.Query().
//		GroupBy(ambalance.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (abq *AmBalanceQuery) GroupBy(field string, fields ...string) *AmBalanceGroupBy {
	group := &AmBalanceGroupBy{config: abq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := abq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return abq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserId int64 `json:"userId,omitempty"`
//	}
//
//	client.AmBalance.Query().
//		Select(ambalance.FieldUserId).
//		Scan(ctx, &v)
//
func (abq *AmBalanceQuery) Select(fields ...string) *AmBalanceSelect {
	abq.fields = append(abq.fields, fields...)
	return &AmBalanceSelect{AmBalanceQuery: abq}
}

func (abq *AmBalanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range abq.fields {
		if !ambalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if abq.path != nil {
		prev, err := abq.path(ctx)
		if err != nil {
			return err
		}
		abq.sql = prev
	}
	return nil
}

func (abq *AmBalanceQuery) sqlAll(ctx context.Context) ([]*AmBalance, error) {
	var (
		nodes       = []*AmBalance{}
		_spec       = abq.querySpec()
		loadedTypes = [1]bool{
			abq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AmBalance{config: abq.config}
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
	if err := sqlgraph.QueryNodes(ctx, abq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := abq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AmBalance)
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

func (abq *AmBalanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := abq.querySpec()
	return sqlgraph.CountNodes(ctx, abq.driver, _spec)
}

func (abq *AmBalanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := abq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (abq *AmBalanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ambalance.Table,
			Columns: ambalance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: ambalance.FieldID,
			},
		},
		From:   abq.sql,
		Unique: true,
	}
	if unique := abq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := abq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ambalance.FieldID)
		for i := range fields {
			if fields[i] != ambalance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := abq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := abq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := abq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := abq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (abq *AmBalanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(abq.driver.Dialect())
	t1 := builder.Table(ambalance.Table)
	columns := abq.fields
	if len(columns) == 0 {
		columns = ambalance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if abq.sql != nil {
		selector = abq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range abq.predicates {
		p(selector)
	}
	for _, p := range abq.order {
		p(selector)
	}
	if offset := abq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := abq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AmBalanceGroupBy is the group-by builder for AmBalance entities.
type AmBalanceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (abgb *AmBalanceGroupBy) Aggregate(fns ...AggregateFunc) *AmBalanceGroupBy {
	abgb.fns = append(abgb.fns, fns...)
	return abgb
}

// Scan applies the group-by query and scans the result into the given value.
func (abgb *AmBalanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := abgb.path(ctx)
	if err != nil {
		return err
	}
	abgb.sql = query
	return abgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := abgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(abgb.fields) > 1 {
		return nil, errors.New("ent: AmBalanceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := abgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) StringsX(ctx context.Context) []string {
	v, err := abgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = abgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) StringX(ctx context.Context) string {
	v, err := abgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(abgb.fields) > 1 {
		return nil, errors.New("ent: AmBalanceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := abgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) IntsX(ctx context.Context) []int {
	v, err := abgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = abgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) IntX(ctx context.Context) int {
	v, err := abgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(abgb.fields) > 1 {
		return nil, errors.New("ent: AmBalanceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := abgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := abgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = abgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := abgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(abgb.fields) > 1 {
		return nil, errors.New("ent: AmBalanceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := abgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := abgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (abgb *AmBalanceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = abgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (abgb *AmBalanceGroupBy) BoolX(ctx context.Context) bool {
	v, err := abgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (abgb *AmBalanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range abgb.fields {
		if !ambalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := abgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := abgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (abgb *AmBalanceGroupBy) sqlQuery() *sql.Selector {
	selector := abgb.sql.Select()
	aggregation := make([]string, 0, len(abgb.fns))
	for _, fn := range abgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(abgb.fields)+len(abgb.fns))
		for _, f := range abgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(abgb.fields...)...)
}

// AmBalanceSelect is the builder for selecting fields of AmBalance entities.
type AmBalanceSelect struct {
	*AmBalanceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (abs *AmBalanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := abs.prepareQuery(ctx); err != nil {
		return err
	}
	abs.sql = abs.AmBalanceQuery.sqlQuery(ctx)
	return abs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (abs *AmBalanceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := abs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(abs.fields) > 1 {
		return nil, errors.New("ent: AmBalanceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := abs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (abs *AmBalanceSelect) StringsX(ctx context.Context) []string {
	v, err := abs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = abs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (abs *AmBalanceSelect) StringX(ctx context.Context) string {
	v, err := abs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(abs.fields) > 1 {
		return nil, errors.New("ent: AmBalanceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := abs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (abs *AmBalanceSelect) IntsX(ctx context.Context) []int {
	v, err := abs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = abs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (abs *AmBalanceSelect) IntX(ctx context.Context) int {
	v, err := abs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(abs.fields) > 1 {
		return nil, errors.New("ent: AmBalanceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := abs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (abs *AmBalanceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := abs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = abs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (abs *AmBalanceSelect) Float64X(ctx context.Context) float64 {
	v, err := abs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(abs.fields) > 1 {
		return nil, errors.New("ent: AmBalanceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := abs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (abs *AmBalanceSelect) BoolsX(ctx context.Context) []bool {
	v, err := abs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (abs *AmBalanceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = abs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ambalance.Label}
	default:
		err = fmt.Errorf("ent: AmBalanceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (abs *AmBalanceSelect) BoolX(ctx context.Context) bool {
	v, err := abs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (abs *AmBalanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := abs.sql.Query()
	if err := abs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
