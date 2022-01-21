// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/scoreproduct"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScoreProductQuery is the builder for querying ScoreProduct entities.
type ScoreProductQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ScoreProduct
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScoreProductQuery builder.
func (spq *ScoreProductQuery) Where(ps ...predicate.ScoreProduct) *ScoreProductQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit adds a limit step to the query.
func (spq *ScoreProductQuery) Limit(limit int) *ScoreProductQuery {
	spq.limit = &limit
	return spq
}

// Offset adds an offset step to the query.
func (spq *ScoreProductQuery) Offset(offset int) *ScoreProductQuery {
	spq.offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *ScoreProductQuery) Unique(unique bool) *ScoreProductQuery {
	spq.unique = &unique
	return spq
}

// Order adds an order step to the query.
func (spq *ScoreProductQuery) Order(o ...OrderFunc) *ScoreProductQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// First returns the first ScoreProduct entity from the query.
// Returns a *NotFoundError when no ScoreProduct was found.
func (spq *ScoreProductQuery) First(ctx context.Context) (*ScoreProduct, error) {
	nodes, err := spq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scoreproduct.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *ScoreProductQuery) FirstX(ctx context.Context) *ScoreProduct {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ScoreProduct ID from the query.
// Returns a *NotFoundError when no ScoreProduct ID was found.
func (spq *ScoreProductQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = spq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scoreproduct.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *ScoreProductQuery) FirstIDX(ctx context.Context) int64 {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ScoreProduct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ScoreProduct entity is not found.
// Returns a *NotFoundError when no ScoreProduct entities are found.
func (spq *ScoreProductQuery) Only(ctx context.Context) (*ScoreProduct, error) {
	nodes, err := spq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scoreproduct.Label}
	default:
		return nil, &NotSingularError{scoreproduct.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *ScoreProductQuery) OnlyX(ctx context.Context) *ScoreProduct {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ScoreProduct ID in the query.
// Returns a *NotSingularError when exactly one ScoreProduct ID is not found.
// Returns a *NotFoundError when no entities are found.
func (spq *ScoreProductQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = spq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = &NotSingularError{scoreproduct.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *ScoreProductQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScoreProducts.
func (spq *ScoreProductQuery) All(ctx context.Context) ([]*ScoreProduct, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return spq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (spq *ScoreProductQuery) AllX(ctx context.Context) []*ScoreProduct {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ScoreProduct IDs.
func (spq *ScoreProductQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := spq.Select(scoreproduct.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *ScoreProductQuery) IDsX(ctx context.Context) []int64 {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *ScoreProductQuery) Count(ctx context.Context) (int, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return spq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (spq *ScoreProductQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *ScoreProductQuery) Exist(ctx context.Context) (bool, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return spq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *ScoreProductQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScoreProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *ScoreProductQuery) Clone() *ScoreProductQuery {
	if spq == nil {
		return nil
	}
	return &ScoreProductQuery{
		config:     spq.config,
		limit:      spq.limit,
		offset:     spq.offset,
		order:      append([]OrderFunc{}, spq.order...),
		predicates: append([]predicate.ScoreProduct{}, spq.predicates...),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProductName string `json:"productName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ScoreProduct.Query().
//		GroupBy(scoreproduct.FieldProductName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (spq *ScoreProductQuery) GroupBy(field string, fields ...string) *ScoreProductGroupBy {
	group := &ScoreProductGroupBy{config: spq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return spq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProductName string `json:"productName,omitempty"`
//	}
//
//	client.ScoreProduct.Query().
//		Select(scoreproduct.FieldProductName).
//		Scan(ctx, &v)
//
func (spq *ScoreProductQuery) Select(fields ...string) *ScoreProductSelect {
	spq.fields = append(spq.fields, fields...)
	return &ScoreProductSelect{ScoreProductQuery: spq}
}

func (spq *ScoreProductQuery) prepareQuery(ctx context.Context) error {
	for _, f := range spq.fields {
		if !scoreproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *ScoreProductQuery) sqlAll(ctx context.Context) ([]*ScoreProduct, error) {
	var (
		nodes = []*ScoreProduct{}
		_spec = spq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ScoreProduct{config: spq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (spq *ScoreProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *ScoreProductQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := spq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (spq *ScoreProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scoreproduct.Table,
			Columns: scoreproduct.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: scoreproduct.FieldID,
			},
		},
		From:   spq.sql,
		Unique: true,
	}
	if unique := spq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := spq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scoreproduct.FieldID)
		for i := range fields {
			if fields[i] != scoreproduct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *ScoreProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(scoreproduct.Table)
	columns := spq.fields
	if len(columns) == 0 {
		columns = scoreproduct.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScoreProductGroupBy is the group-by builder for ScoreProduct entities.
type ScoreProductGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *ScoreProductGroupBy) Aggregate(fns ...AggregateFunc) *ScoreProductGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the group-by query and scans the result into the given value.
func (spgb *ScoreProductGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := spgb.path(ctx)
	if err != nil {
		return err
	}
	spgb.sql = query
	return spgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := spgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: ScoreProductGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) StringsX(ctx context.Context) []string {
	v, err := spgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = spgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) StringX(ctx context.Context) string {
	v, err := spgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: ScoreProductGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) IntsX(ctx context.Context) []int {
	v, err := spgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = spgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) IntX(ctx context.Context) int {
	v, err := spgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: ScoreProductGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := spgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = spgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) Float64X(ctx context.Context) float64 {
	v, err := spgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(spgb.fields) > 1 {
		return nil, errors.New("ent: ScoreProductGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := spgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := spgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spgb *ScoreProductGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = spgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (spgb *ScoreProductGroupBy) BoolX(ctx context.Context) bool {
	v, err := spgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (spgb *ScoreProductGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range spgb.fields {
		if !scoreproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := spgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (spgb *ScoreProductGroupBy) sqlQuery() *sql.Selector {
	selector := spgb.sql.Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(spgb.fields)+len(spgb.fns))
		for _, f := range spgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(spgb.fields...)...)
}

// ScoreProductSelect is the builder for selecting fields of ScoreProduct entities.
type ScoreProductSelect struct {
	*ScoreProductQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sps *ScoreProductSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	sps.sql = sps.ScoreProductQuery.sqlQuery(ctx)
	return sps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sps *ScoreProductSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: ScoreProductSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sps *ScoreProductSelect) StringsX(ctx context.Context) []string {
	v, err := sps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sps *ScoreProductSelect) StringX(ctx context.Context) string {
	v, err := sps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: ScoreProductSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sps *ScoreProductSelect) IntsX(ctx context.Context) []int {
	v, err := sps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sps *ScoreProductSelect) IntX(ctx context.Context) int {
	v, err := sps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: ScoreProductSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sps *ScoreProductSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sps *ScoreProductSelect) Float64X(ctx context.Context) float64 {
	v, err := sps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sps.fields) > 1 {
		return nil, errors.New("ent: ScoreProductSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sps *ScoreProductSelect) BoolsX(ctx context.Context) []bool {
	v, err := sps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sps *ScoreProductSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scoreproduct.Label}
	default:
		err = fmt.Errorf("ent: ScoreProductSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sps *ScoreProductSelect) BoolX(ctx context.Context) bool {
	v, err := sps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sps *ScoreProductSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sps.sql.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}