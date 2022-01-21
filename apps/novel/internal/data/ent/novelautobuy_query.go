// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelautobuy"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelAutoBuyQuery is the builder for querying NovelAutoBuy entities.
type NovelAutoBuyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NovelAutoBuy
	// eager-loading edges.
	withUser *SocialUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NovelAutoBuyQuery builder.
func (nabq *NovelAutoBuyQuery) Where(ps ...predicate.NovelAutoBuy) *NovelAutoBuyQuery {
	nabq.predicates = append(nabq.predicates, ps...)
	return nabq
}

// Limit adds a limit step to the query.
func (nabq *NovelAutoBuyQuery) Limit(limit int) *NovelAutoBuyQuery {
	nabq.limit = &limit
	return nabq
}

// Offset adds an offset step to the query.
func (nabq *NovelAutoBuyQuery) Offset(offset int) *NovelAutoBuyQuery {
	nabq.offset = &offset
	return nabq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nabq *NovelAutoBuyQuery) Unique(unique bool) *NovelAutoBuyQuery {
	nabq.unique = &unique
	return nabq
}

// Order adds an order step to the query.
func (nabq *NovelAutoBuyQuery) Order(o ...OrderFunc) *NovelAutoBuyQuery {
	nabq.order = append(nabq.order, o...)
	return nabq
}

// QueryUser chains the current query on the "user" edge.
func (nabq *NovelAutoBuyQuery) QueryUser() *SocialUserQuery {
	query := &SocialUserQuery{config: nabq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nabq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nabq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novelautobuy.Table, novelautobuy.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, novelautobuy.UserTable, novelautobuy.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(nabq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NovelAutoBuy entity from the query.
// Returns a *NotFoundError when no NovelAutoBuy was found.
func (nabq *NovelAutoBuyQuery) First(ctx context.Context) (*NovelAutoBuy, error) {
	nodes, err := nabq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{novelautobuy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) FirstX(ctx context.Context) *NovelAutoBuy {
	node, err := nabq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NovelAutoBuy ID from the query.
// Returns a *NotFoundError when no NovelAutoBuy ID was found.
func (nabq *NovelAutoBuyQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nabq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{novelautobuy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) FirstIDX(ctx context.Context) int64 {
	id, err := nabq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NovelAutoBuy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NovelAutoBuy entity is not found.
// Returns a *NotFoundError when no NovelAutoBuy entities are found.
func (nabq *NovelAutoBuyQuery) Only(ctx context.Context) (*NovelAutoBuy, error) {
	nodes, err := nabq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{novelautobuy.Label}
	default:
		return nil, &NotSingularError{novelautobuy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) OnlyX(ctx context.Context) *NovelAutoBuy {
	node, err := nabq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NovelAutoBuy ID in the query.
// Returns a *NotSingularError when exactly one NovelAutoBuy ID is not found.
// Returns a *NotFoundError when no entities are found.
func (nabq *NovelAutoBuyQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nabq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = &NotSingularError{novelautobuy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := nabq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NovelAutoBuys.
func (nabq *NovelAutoBuyQuery) All(ctx context.Context) ([]*NovelAutoBuy, error) {
	if err := nabq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nabq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) AllX(ctx context.Context) []*NovelAutoBuy {
	nodes, err := nabq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NovelAutoBuy IDs.
func (nabq *NovelAutoBuyQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := nabq.Select(novelautobuy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) IDsX(ctx context.Context) []int64 {
	ids, err := nabq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nabq *NovelAutoBuyQuery) Count(ctx context.Context) (int, error) {
	if err := nabq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nabq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) CountX(ctx context.Context) int {
	count, err := nabq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nabq *NovelAutoBuyQuery) Exist(ctx context.Context) (bool, error) {
	if err := nabq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nabq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nabq *NovelAutoBuyQuery) ExistX(ctx context.Context) bool {
	exist, err := nabq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NovelAutoBuyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nabq *NovelAutoBuyQuery) Clone() *NovelAutoBuyQuery {
	if nabq == nil {
		return nil
	}
	return &NovelAutoBuyQuery{
		config:     nabq.config,
		limit:      nabq.limit,
		offset:     nabq.offset,
		order:      append([]OrderFunc{}, nabq.order...),
		predicates: append([]predicate.NovelAutoBuy{}, nabq.predicates...),
		withUser:   nabq.withUser.Clone(),
		// clone intermediate query.
		sql:  nabq.sql.Clone(),
		path: nabq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (nabq *NovelAutoBuyQuery) WithUser(opts ...func(*SocialUserQuery)) *NovelAutoBuyQuery {
	query := &SocialUserQuery{config: nabq.config}
	for _, opt := range opts {
		opt(query)
	}
	nabq.withUser = query
	return nabq
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
//	client.NovelAutoBuy.Query().
//		GroupBy(novelautobuy.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nabq *NovelAutoBuyQuery) GroupBy(field string, fields ...string) *NovelAutoBuyGroupBy {
	group := &NovelAutoBuyGroupBy{config: nabq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nabq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nabq.sqlQuery(ctx), nil
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
//	client.NovelAutoBuy.Query().
//		Select(novelautobuy.FieldUserId).
//		Scan(ctx, &v)
//
func (nabq *NovelAutoBuyQuery) Select(fields ...string) *NovelAutoBuySelect {
	nabq.fields = append(nabq.fields, fields...)
	return &NovelAutoBuySelect{NovelAutoBuyQuery: nabq}
}

func (nabq *NovelAutoBuyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nabq.fields {
		if !novelautobuy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nabq.path != nil {
		prev, err := nabq.path(ctx)
		if err != nil {
			return err
		}
		nabq.sql = prev
	}
	return nil
}

func (nabq *NovelAutoBuyQuery) sqlAll(ctx context.Context) ([]*NovelAutoBuy, error) {
	var (
		nodes       = []*NovelAutoBuy{}
		_spec       = nabq.querySpec()
		loadedTypes = [1]bool{
			nabq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NovelAutoBuy{config: nabq.config}
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
	if err := sqlgraph.QueryNodes(ctx, nabq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nabq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*NovelAutoBuy)
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

func (nabq *NovelAutoBuyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nabq.querySpec()
	return sqlgraph.CountNodes(ctx, nabq.driver, _spec)
}

func (nabq *NovelAutoBuyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nabq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nabq *NovelAutoBuyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelautobuy.Table,
			Columns: novelautobuy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelautobuy.FieldID,
			},
		},
		From:   nabq.sql,
		Unique: true,
	}
	if unique := nabq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nabq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelautobuy.FieldID)
		for i := range fields {
			if fields[i] != novelautobuy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nabq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nabq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nabq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nabq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nabq *NovelAutoBuyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nabq.driver.Dialect())
	t1 := builder.Table(novelautobuy.Table)
	columns := nabq.fields
	if len(columns) == 0 {
		columns = novelautobuy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nabq.sql != nil {
		selector = nabq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range nabq.predicates {
		p(selector)
	}
	for _, p := range nabq.order {
		p(selector)
	}
	if offset := nabq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nabq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NovelAutoBuyGroupBy is the group-by builder for NovelAutoBuy entities.
type NovelAutoBuyGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nabgb *NovelAutoBuyGroupBy) Aggregate(fns ...AggregateFunc) *NovelAutoBuyGroupBy {
	nabgb.fns = append(nabgb.fns, fns...)
	return nabgb
}

// Scan applies the group-by query and scans the result into the given value.
func (nabgb *NovelAutoBuyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nabgb.path(ctx)
	if err != nil {
		return err
	}
	nabgb.sql = query
	return nabgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := nabgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(nabgb.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuyGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := nabgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) StringsX(ctx context.Context) []string {
	v, err := nabgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nabgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuyGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) StringX(ctx context.Context) string {
	v, err := nabgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(nabgb.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuyGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := nabgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) IntsX(ctx context.Context) []int {
	v, err := nabgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nabgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuyGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) IntX(ctx context.Context) int {
	v, err := nabgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(nabgb.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuyGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := nabgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := nabgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nabgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuyGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) Float64X(ctx context.Context) float64 {
	v, err := nabgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(nabgb.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuyGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := nabgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := nabgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nabgb *NovelAutoBuyGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nabgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuyGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nabgb *NovelAutoBuyGroupBy) BoolX(ctx context.Context) bool {
	v, err := nabgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nabgb *NovelAutoBuyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nabgb.fields {
		if !novelautobuy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nabgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nabgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nabgb *NovelAutoBuyGroupBy) sqlQuery() *sql.Selector {
	selector := nabgb.sql.Select()
	aggregation := make([]string, 0, len(nabgb.fns))
	for _, fn := range nabgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(nabgb.fields)+len(nabgb.fns))
		for _, f := range nabgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(nabgb.fields...)...)
}

// NovelAutoBuySelect is the builder for selecting fields of NovelAutoBuy entities.
type NovelAutoBuySelect struct {
	*NovelAutoBuyQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nabs *NovelAutoBuySelect) Scan(ctx context.Context, v interface{}) error {
	if err := nabs.prepareQuery(ctx); err != nil {
		return err
	}
	nabs.sql = nabs.NovelAutoBuyQuery.sqlQuery(ctx)
	return nabs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) ScanX(ctx context.Context, v interface{}) {
	if err := nabs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) Strings(ctx context.Context) ([]string, error) {
	if len(nabs.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nabs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) StringsX(ctx context.Context) []string {
	v, err := nabs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nabs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) StringX(ctx context.Context) string {
	v, err := nabs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) Ints(ctx context.Context) ([]int, error) {
	if len(nabs.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nabs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) IntsX(ctx context.Context) []int {
	v, err := nabs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nabs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) IntX(ctx context.Context) int {
	v, err := nabs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nabs.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nabs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) Float64sX(ctx context.Context) []float64 {
	v, err := nabs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nabs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) Float64X(ctx context.Context) float64 {
	v, err := nabs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nabs.fields) > 1 {
		return nil, errors.New("ent: NovelAutoBuySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nabs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) BoolsX(ctx context.Context) []bool {
	v, err := nabs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nabs *NovelAutoBuySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nabs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelautobuy.Label}
	default:
		err = fmt.Errorf("ent: NovelAutoBuySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nabs *NovelAutoBuySelect) BoolX(ctx context.Context) bool {
	v, err := nabs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nabs *NovelAutoBuySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nabs.sql.Query()
	if err := nabs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
