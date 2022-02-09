// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/resourcestorage"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResourceStorageQuery is the builder for querying ResourceStorage entities.
type ResourceStorageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ResourceStorage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResourceStorageQuery builder.
func (rsq *ResourceStorageQuery) Where(ps ...predicate.ResourceStorage) *ResourceStorageQuery {
	rsq.predicates = append(rsq.predicates, ps...)
	return rsq
}

// Limit adds a limit step to the query.
func (rsq *ResourceStorageQuery) Limit(limit int) *ResourceStorageQuery {
	rsq.limit = &limit
	return rsq
}

// Offset adds an offset step to the query.
func (rsq *ResourceStorageQuery) Offset(offset int) *ResourceStorageQuery {
	rsq.offset = &offset
	return rsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rsq *ResourceStorageQuery) Unique(unique bool) *ResourceStorageQuery {
	rsq.unique = &unique
	return rsq
}

// Order adds an order step to the query.
func (rsq *ResourceStorageQuery) Order(o ...OrderFunc) *ResourceStorageQuery {
	rsq.order = append(rsq.order, o...)
	return rsq
}

// First returns the first ResourceStorage entity from the query.
// Returns a *NotFoundError when no ResourceStorage was found.
func (rsq *ResourceStorageQuery) First(ctx context.Context) (*ResourceStorage, error) {
	nodes, err := rsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resourcestorage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rsq *ResourceStorageQuery) FirstX(ctx context.Context) *ResourceStorage {
	node, err := rsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ResourceStorage ID from the query.
// Returns a *NotFoundError when no ResourceStorage ID was found.
func (rsq *ResourceStorageQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resourcestorage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rsq *ResourceStorageQuery) FirstIDX(ctx context.Context) int64 {
	id, err := rsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ResourceStorage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one ResourceStorage entity is not found.
// Returns a *NotFoundError when no ResourceStorage entities are found.
func (rsq *ResourceStorageQuery) Only(ctx context.Context) (*ResourceStorage, error) {
	nodes, err := rsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resourcestorage.Label}
	default:
		return nil, &NotSingularError{resourcestorage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rsq *ResourceStorageQuery) OnlyX(ctx context.Context) *ResourceStorage {
	node, err := rsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ResourceStorage ID in the query.
// Returns a *NotSingularError when exactly one ResourceStorage ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rsq *ResourceStorageQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = rsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = &NotSingularError{resourcestorage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rsq *ResourceStorageQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := rsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ResourceStorages.
func (rsq *ResourceStorageQuery) All(ctx context.Context) ([]*ResourceStorage, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rsq *ResourceStorageQuery) AllX(ctx context.Context) []*ResourceStorage {
	nodes, err := rsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ResourceStorage IDs.
func (rsq *ResourceStorageQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := rsq.Select(resourcestorage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rsq *ResourceStorageQuery) IDsX(ctx context.Context) []int64 {
	ids, err := rsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rsq *ResourceStorageQuery) Count(ctx context.Context) (int, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rsq *ResourceStorageQuery) CountX(ctx context.Context) int {
	count, err := rsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rsq *ResourceStorageQuery) Exist(ctx context.Context) (bool, error) {
	if err := rsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rsq *ResourceStorageQuery) ExistX(ctx context.Context) bool {
	exist, err := rsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResourceStorageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rsq *ResourceStorageQuery) Clone() *ResourceStorageQuery {
	if rsq == nil {
		return nil
	}
	return &ResourceStorageQuery{
		config:     rsq.config,
		limit:      rsq.limit,
		offset:     rsq.offset,
		order:      append([]OrderFunc{}, rsq.order...),
		predicates: append([]predicate.ResourceStorage{}, rsq.predicates...),
		// clone intermediate query.
		sql:  rsq.sql.Clone(),
		path: rsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GroupId int32 `json:"groupId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ResourceStorage.Query().
//		GroupBy(resourcestorage.FieldGroupId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rsq *ResourceStorageQuery) GroupBy(field string, fields ...string) *ResourceStorageGroupBy {
	group := &ResourceStorageGroupBy{config: rsq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rsq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GroupId int32 `json:"groupId,omitempty"`
//	}
//
//	client.ResourceStorage.Query().
//		Select(resourcestorage.FieldGroupId).
//		Scan(ctx, &v)
//
func (rsq *ResourceStorageQuery) Select(fields ...string) *ResourceStorageSelect {
	rsq.fields = append(rsq.fields, fields...)
	return &ResourceStorageSelect{ResourceStorageQuery: rsq}
}

func (rsq *ResourceStorageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rsq.fields {
		if !resourcestorage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rsq.path != nil {
		prev, err := rsq.path(ctx)
		if err != nil {
			return err
		}
		rsq.sql = prev
	}
	return nil
}

func (rsq *ResourceStorageQuery) sqlAll(ctx context.Context) ([]*ResourceStorage, error) {
	var (
		nodes = []*ResourceStorage{}
		_spec = rsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &ResourceStorage{config: rsq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rsq *ResourceStorageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rsq.querySpec()
	_spec.Node.Columns = rsq.fields
	if len(rsq.fields) > 0 {
		_spec.Unique = rsq.unique != nil && *rsq.unique
	}
	return sqlgraph.CountNodes(ctx, rsq.driver, _spec)
}

func (rsq *ResourceStorageQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rsq *ResourceStorageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resourcestorage.Table,
			Columns: resourcestorage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: resourcestorage.FieldID,
			},
		},
		From:   rsq.sql,
		Unique: true,
	}
	if unique := rsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resourcestorage.FieldID)
		for i := range fields {
			if fields[i] != resourcestorage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rsq *ResourceStorageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rsq.driver.Dialect())
	t1 := builder.Table(resourcestorage.Table)
	columns := rsq.fields
	if len(columns) == 0 {
		columns = resourcestorage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rsq.sql != nil {
		selector = rsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rsq.unique != nil && *rsq.unique {
		selector.Distinct()
	}
	for _, p := range rsq.predicates {
		p(selector)
	}
	for _, p := range rsq.order {
		p(selector)
	}
	if offset := rsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ResourceStorageGroupBy is the group-by builder for ResourceStorage entities.
type ResourceStorageGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rsgb *ResourceStorageGroupBy) Aggregate(fns ...AggregateFunc) *ResourceStorageGroupBy {
	rsgb.fns = append(rsgb.fns, fns...)
	return rsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rsgb *ResourceStorageGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rsgb.path(ctx)
	if err != nil {
		return err
	}
	rsgb.sql = query
	return rsgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rsgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) StringsX(ctx context.Context) []string {
	v, err := rsgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rsgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) StringX(ctx context.Context) string {
	v, err := rsgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) IntsX(ctx context.Context) []int {
	v, err := rsgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rsgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) IntX(ctx context.Context) int {
	v, err := rsgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rsgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rsgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rsgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rsgb.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rsgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rsgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rsgb *ResourceStorageGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rsgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rsgb *ResourceStorageGroupBy) BoolX(ctx context.Context) bool {
	v, err := rsgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rsgb *ResourceStorageGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rsgb.fields {
		if !resourcestorage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rsgb *ResourceStorageGroupBy) sqlQuery() *sql.Selector {
	selector := rsgb.sql.Select()
	aggregation := make([]string, 0, len(rsgb.fns))
	for _, fn := range rsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rsgb.fields)+len(rsgb.fns))
		for _, f := range rsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rsgb.fields...)...)
}

// ResourceStorageSelect is the builder for selecting fields of ResourceStorage entities.
type ResourceStorageSelect struct {
	*ResourceStorageQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rss *ResourceStorageSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rss.prepareQuery(ctx); err != nil {
		return err
	}
	rss.sql = rss.ResourceStorageQuery.sqlQuery(ctx)
	return rss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rss *ResourceStorageSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rss *ResourceStorageSelect) StringsX(ctx context.Context) []string {
	v, err := rss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rss *ResourceStorageSelect) StringX(ctx context.Context) string {
	v, err := rss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rss *ResourceStorageSelect) IntsX(ctx context.Context) []int {
	v, err := rss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rss *ResourceStorageSelect) IntX(ctx context.Context) int {
	v, err := rss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rss *ResourceStorageSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rss *ResourceStorageSelect) Float64X(ctx context.Context) float64 {
	v, err := rss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rss.fields) > 1 {
		return nil, errors.New("ent: ResourceStorageSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rss *ResourceStorageSelect) BoolsX(ctx context.Context) []bool {
	v, err := rss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rss *ResourceStorageSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{resourcestorage.Label}
	default:
		err = fmt.Errorf("ent: ResourceStorageSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rss *ResourceStorageSelect) BoolX(ctx context.Context) bool {
	v, err := rss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rss *ResourceStorageSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rss.sql.Query()
	if err := rss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
