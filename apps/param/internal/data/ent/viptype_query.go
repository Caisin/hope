// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/param/internal/data/ent/predicate"
	"hope/apps/param/internal/data/ent/viptype"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VipTypeQuery is the builder for querying VipType entities.
type VipTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.VipType
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VipTypeQuery builder.
func (vtq *VipTypeQuery) Where(ps ...predicate.VipType) *VipTypeQuery {
	vtq.predicates = append(vtq.predicates, ps...)
	return vtq
}

// Limit adds a limit step to the query.
func (vtq *VipTypeQuery) Limit(limit int) *VipTypeQuery {
	vtq.limit = &limit
	return vtq
}

// Offset adds an offset step to the query.
func (vtq *VipTypeQuery) Offset(offset int) *VipTypeQuery {
	vtq.offset = &offset
	return vtq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vtq *VipTypeQuery) Unique(unique bool) *VipTypeQuery {
	vtq.unique = &unique
	return vtq
}

// Order adds an order step to the query.
func (vtq *VipTypeQuery) Order(o ...OrderFunc) *VipTypeQuery {
	vtq.order = append(vtq.order, o...)
	return vtq
}

// First returns the first VipType entity from the query.
// Returns a *NotFoundError when no VipType was found.
func (vtq *VipTypeQuery) First(ctx context.Context) (*VipType, error) {
	nodes, err := vtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{viptype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vtq *VipTypeQuery) FirstX(ctx context.Context) *VipType {
	node, err := vtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VipType ID from the query.
// Returns a *NotFoundError when no VipType ID was found.
func (vtq *VipTypeQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{viptype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vtq *VipTypeQuery) FirstIDX(ctx context.Context) int64 {
	id, err := vtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VipType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one VipType entity is not found.
// Returns a *NotFoundError when no VipType entities are found.
func (vtq *VipTypeQuery) Only(ctx context.Context) (*VipType, error) {
	nodes, err := vtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{viptype.Label}
	default:
		return nil, &NotSingularError{viptype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vtq *VipTypeQuery) OnlyX(ctx context.Context) *VipType {
	node, err := vtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VipType ID in the query.
// Returns a *NotSingularError when exactly one VipType ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vtq *VipTypeQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = &NotSingularError{viptype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vtq *VipTypeQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := vtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VipTypes.
func (vtq *VipTypeQuery) All(ctx context.Context) ([]*VipType, error) {
	if err := vtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vtq *VipTypeQuery) AllX(ctx context.Context) []*VipType {
	nodes, err := vtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VipType IDs.
func (vtq *VipTypeQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := vtq.Select(viptype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vtq *VipTypeQuery) IDsX(ctx context.Context) []int64 {
	ids, err := vtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vtq *VipTypeQuery) Count(ctx context.Context) (int, error) {
	if err := vtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vtq *VipTypeQuery) CountX(ctx context.Context) int {
	count, err := vtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vtq *VipTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := vtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vtq *VipTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := vtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VipTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vtq *VipTypeQuery) Clone() *VipTypeQuery {
	if vtq == nil {
		return nil
	}
	return &VipTypeQuery{
		config:     vtq.config,
		limit:      vtq.limit,
		offset:     vtq.offset,
		order:      append([]OrderFunc{}, vtq.order...),
		predicates: append([]predicate.VipType{}, vtq.predicates...),
		// clone intermediate query.
		sql:  vtq.sql.Clone(),
		path: vtq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VipName string `json:"vipName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.VipType.Query().
//		GroupBy(viptype.FieldVipName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vtq *VipTypeQuery) GroupBy(field string, fields ...string) *VipTypeGroupBy {
	group := &VipTypeGroupBy{config: vtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vtq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VipName string `json:"vipName,omitempty"`
//	}
//
//	client.VipType.Query().
//		Select(viptype.FieldVipName).
//		Scan(ctx, &v)
//
func (vtq *VipTypeQuery) Select(fields ...string) *VipTypeSelect {
	vtq.fields = append(vtq.fields, fields...)
	return &VipTypeSelect{VipTypeQuery: vtq}
}

func (vtq *VipTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vtq.fields {
		if !viptype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vtq.path != nil {
		prev, err := vtq.path(ctx)
		if err != nil {
			return err
		}
		vtq.sql = prev
	}
	return nil
}

func (vtq *VipTypeQuery) sqlAll(ctx context.Context) ([]*VipType, error) {
	var (
		nodes = []*VipType{}
		_spec = vtq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &VipType{config: vtq.config}
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
	if err := sqlgraph.QueryNodes(ctx, vtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vtq *VipTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vtq.querySpec()
	return sqlgraph.CountNodes(ctx, vtq.driver, _spec)
}

func (vtq *VipTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vtq *VipTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   viptype.Table,
			Columns: viptype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: viptype.FieldID,
			},
		},
		From:   vtq.sql,
		Unique: true,
	}
	if unique := vtq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vtq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, viptype.FieldID)
		for i := range fields {
			if fields[i] != viptype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vtq *VipTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vtq.driver.Dialect())
	t1 := builder.Table(viptype.Table)
	columns := vtq.fields
	if len(columns) == 0 {
		columns = viptype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vtq.sql != nil {
		selector = vtq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range vtq.predicates {
		p(selector)
	}
	for _, p := range vtq.order {
		p(selector)
	}
	if offset := vtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VipTypeGroupBy is the group-by builder for VipType entities.
type VipTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vtgb *VipTypeGroupBy) Aggregate(fns ...AggregateFunc) *VipTypeGroupBy {
	vtgb.fns = append(vtgb.fns, fns...)
	return vtgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vtgb *VipTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vtgb.path(ctx)
	if err != nil {
		return err
	}
	vtgb.sql = query
	return vtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vtgb.fields) > 1 {
		return nil, errors.New("ent: VipTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := vtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) StringX(ctx context.Context) string {
	v, err := vtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vtgb.fields) > 1 {
		return nil, errors.New("ent: VipTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := vtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) IntX(ctx context.Context) int {
	v, err := vtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vtgb.fields) > 1 {
		return nil, errors.New("ent: VipTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vtgb.fields) > 1 {
		return nil, errors.New("ent: VipTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vtgb *VipTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vtgb *VipTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := vtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vtgb *VipTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vtgb.fields {
		if !viptype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vtgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vtgb *VipTypeGroupBy) sqlQuery() *sql.Selector {
	selector := vtgb.sql.Select()
	aggregation := make([]string, 0, len(vtgb.fns))
	for _, fn := range vtgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vtgb.fields)+len(vtgb.fns))
		for _, f := range vtgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vtgb.fields...)...)
}

// VipTypeSelect is the builder for selecting fields of VipType entities.
type VipTypeSelect struct {
	*VipTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vts *VipTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vts.prepareQuery(ctx); err != nil {
		return err
	}
	vts.sql = vts.VipTypeQuery.sqlQuery(ctx)
	return vts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vts *VipTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vts.fields) > 1 {
		return nil, errors.New("ent: VipTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vts *VipTypeSelect) StringsX(ctx context.Context) []string {
	v, err := vts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vts *VipTypeSelect) StringX(ctx context.Context) string {
	v, err := vts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vts.fields) > 1 {
		return nil, errors.New("ent: VipTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vts *VipTypeSelect) IntsX(ctx context.Context) []int {
	v, err := vts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vts *VipTypeSelect) IntX(ctx context.Context) int {
	v, err := vts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vts.fields) > 1 {
		return nil, errors.New("ent: VipTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vts *VipTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vts *VipTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := vts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vts.fields) > 1 {
		return nil, errors.New("ent: VipTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vts *VipTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := vts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vts *VipTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{viptype.Label}
	default:
		err = fmt.Errorf("ent: VipTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vts *VipTypeSelect) BoolX(ctx context.Context) bool {
	v, err := vts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vts *VipTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vts.sql.Query()
	if err := vts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}