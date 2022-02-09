// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/datasource"
	"hope/apps/novel/internal/data/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DataSourceQuery is the builder for querying DataSource entities.
type DataSourceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DataSource
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DataSourceQuery builder.
func (dsq *DataSourceQuery) Where(ps ...predicate.DataSource) *DataSourceQuery {
	dsq.predicates = append(dsq.predicates, ps...)
	return dsq
}

// Limit adds a limit step to the query.
func (dsq *DataSourceQuery) Limit(limit int) *DataSourceQuery {
	dsq.limit = &limit
	return dsq
}

// Offset adds an offset step to the query.
func (dsq *DataSourceQuery) Offset(offset int) *DataSourceQuery {
	dsq.offset = &offset
	return dsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dsq *DataSourceQuery) Unique(unique bool) *DataSourceQuery {
	dsq.unique = &unique
	return dsq
}

// Order adds an order step to the query.
func (dsq *DataSourceQuery) Order(o ...OrderFunc) *DataSourceQuery {
	dsq.order = append(dsq.order, o...)
	return dsq
}

// First returns the first DataSource entity from the query.
// Returns a *NotFoundError when no DataSource was found.
func (dsq *DataSourceQuery) First(ctx context.Context) (*DataSource, error) {
	nodes, err := dsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{datasource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dsq *DataSourceQuery) FirstX(ctx context.Context) *DataSource {
	node, err := dsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DataSource ID from the query.
// Returns a *NotFoundError when no DataSource ID was found.
func (dsq *DataSourceQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{datasource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dsq *DataSourceQuery) FirstIDX(ctx context.Context) int64 {
	id, err := dsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DataSource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one DataSource entity is not found.
// Returns a *NotFoundError when no DataSource entities are found.
func (dsq *DataSourceQuery) Only(ctx context.Context) (*DataSource, error) {
	nodes, err := dsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{datasource.Label}
	default:
		return nil, &NotSingularError{datasource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dsq *DataSourceQuery) OnlyX(ctx context.Context) *DataSource {
	node, err := dsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DataSource ID in the query.
// Returns a *NotSingularError when exactly one DataSource ID is not found.
// Returns a *NotFoundError when no entities are found.
func (dsq *DataSourceQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = dsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = &NotSingularError{datasource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dsq *DataSourceQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := dsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DataSources.
func (dsq *DataSourceQuery) All(ctx context.Context) ([]*DataSource, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dsq *DataSourceQuery) AllX(ctx context.Context) []*DataSource {
	nodes, err := dsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DataSource IDs.
func (dsq *DataSourceQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := dsq.Select(datasource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dsq *DataSourceQuery) IDsX(ctx context.Context) []int64 {
	ids, err := dsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dsq *DataSourceQuery) Count(ctx context.Context) (int, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dsq *DataSourceQuery) CountX(ctx context.Context) int {
	count, err := dsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dsq *DataSourceQuery) Exist(ctx context.Context) (bool, error) {
	if err := dsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dsq *DataSourceQuery) ExistX(ctx context.Context) bool {
	exist, err := dsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DataSourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dsq *DataSourceQuery) Clone() *DataSourceQuery {
	if dsq == nil {
		return nil
	}
	return &DataSourceQuery{
		config:     dsq.config,
		limit:      dsq.limit,
		offset:     dsq.offset,
		order:      append([]OrderFunc{}, dsq.order...),
		predicates: append([]predicate.DataSource{}, dsq.predicates...),
		// clone intermediate query.
		sql:  dsq.sql.Clone(),
		path: dsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DbName string `json:"dbName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DataSource.Query().
//		GroupBy(datasource.FieldDbName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dsq *DataSourceQuery) GroupBy(field string, fields ...string) *DataSourceGroupBy {
	group := &DataSourceGroupBy{config: dsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DbName string `json:"dbName,omitempty"`
//	}
//
//	client.DataSource.Query().
//		Select(datasource.FieldDbName).
//		Scan(ctx, &v)
//
func (dsq *DataSourceQuery) Select(fields ...string) *DataSourceSelect {
	dsq.fields = append(dsq.fields, fields...)
	return &DataSourceSelect{DataSourceQuery: dsq}
}

func (dsq *DataSourceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dsq.fields {
		if !datasource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dsq.path != nil {
		prev, err := dsq.path(ctx)
		if err != nil {
			return err
		}
		dsq.sql = prev
	}
	return nil
}

func (dsq *DataSourceQuery) sqlAll(ctx context.Context) ([]*DataSource, error) {
	var (
		nodes = []*DataSource{}
		_spec = dsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &DataSource{config: dsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, dsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dsq *DataSourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dsq.querySpec()
	_spec.Node.Columns = dsq.fields
	if len(dsq.fields) > 0 {
		_spec.Unique = dsq.unique != nil && *dsq.unique
	}
	return sqlgraph.CountNodes(ctx, dsq.driver, _spec)
}

func (dsq *DataSourceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dsq *DataSourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   datasource.Table,
			Columns: datasource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: datasource.FieldID,
			},
		},
		From:   dsq.sql,
		Unique: true,
	}
	if unique := dsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, datasource.FieldID)
		for i := range fields {
			if fields[i] != datasource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dsq *DataSourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dsq.driver.Dialect())
	t1 := builder.Table(datasource.Table)
	columns := dsq.fields
	if len(columns) == 0 {
		columns = datasource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dsq.sql != nil {
		selector = dsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dsq.unique != nil && *dsq.unique {
		selector.Distinct()
	}
	for _, p := range dsq.predicates {
		p(selector)
	}
	for _, p := range dsq.order {
		p(selector)
	}
	if offset := dsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DataSourceGroupBy is the group-by builder for DataSource entities.
type DataSourceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dsgb *DataSourceGroupBy) Aggregate(fns ...AggregateFunc) *DataSourceGroupBy {
	dsgb.fns = append(dsgb.fns, fns...)
	return dsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dsgb *DataSourceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dsgb.path(ctx)
	if err != nil {
		return err
	}
	dsgb.sql = query
	return dsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DataSourceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) StringsX(ctx context.Context) []string {
	v, err := dsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) StringX(ctx context.Context) string {
	v, err := dsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DataSourceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) IntsX(ctx context.Context) []int {
	v, err := dsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) IntX(ctx context.Context) int {
	v, err := dsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DataSourceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dsgb.fields) > 1 {
		return nil, errors.New("ent: DataSourceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dsgb *DataSourceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dsgb *DataSourceGroupBy) BoolX(ctx context.Context) bool {
	v, err := dsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dsgb *DataSourceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dsgb.fields {
		if !datasource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dsgb *DataSourceGroupBy) sqlQuery() *sql.Selector {
	selector := dsgb.sql.Select()
	aggregation := make([]string, 0, len(dsgb.fns))
	for _, fn := range dsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dsgb.fields)+len(dsgb.fns))
		for _, f := range dsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dsgb.fields...)...)
}

// DataSourceSelect is the builder for selecting fields of DataSource entities.
type DataSourceSelect struct {
	*DataSourceQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dss *DataSourceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dss.prepareQuery(ctx); err != nil {
		return err
	}
	dss.sql = dss.DataSourceQuery.sqlQuery(ctx)
	return dss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dss *DataSourceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := dss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DataSourceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dss *DataSourceSelect) StringsX(ctx context.Context) []string {
	v, err := dss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dss *DataSourceSelect) StringX(ctx context.Context) string {
	v, err := dss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DataSourceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dss *DataSourceSelect) IntsX(ctx context.Context) []int {
	v, err := dss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dss *DataSourceSelect) IntX(ctx context.Context) int {
	v, err := dss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DataSourceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dss *DataSourceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := dss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dss *DataSourceSelect) Float64X(ctx context.Context) float64 {
	v, err := dss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(dss.fields) > 1 {
		return nil, errors.New("ent: DataSourceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := dss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dss *DataSourceSelect) BoolsX(ctx context.Context) []bool {
	v, err := dss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (dss *DataSourceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{datasource.Label}
	default:
		err = fmt.Errorf("ent: DataSourceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dss *DataSourceSelect) BoolX(ctx context.Context) bool {
	v, err := dss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dss *DataSourceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dss.sql.Query()
	if err := dss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
