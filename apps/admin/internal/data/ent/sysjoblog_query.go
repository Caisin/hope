// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysjob"
	"hope/apps/admin/internal/data/ent/sysjoblog"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysJobLogQuery is the builder for querying SysJobLog entities.
type SysJobLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysJobLog
	// eager-loading edges.
	withJob *SysJobQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysJobLogQuery builder.
func (sjlq *SysJobLogQuery) Where(ps ...predicate.SysJobLog) *SysJobLogQuery {
	sjlq.predicates = append(sjlq.predicates, ps...)
	return sjlq
}

// Limit adds a limit step to the query.
func (sjlq *SysJobLogQuery) Limit(limit int) *SysJobLogQuery {
	sjlq.limit = &limit
	return sjlq
}

// Offset adds an offset step to the query.
func (sjlq *SysJobLogQuery) Offset(offset int) *SysJobLogQuery {
	sjlq.offset = &offset
	return sjlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sjlq *SysJobLogQuery) Unique(unique bool) *SysJobLogQuery {
	sjlq.unique = &unique
	return sjlq
}

// Order adds an order step to the query.
func (sjlq *SysJobLogQuery) Order(o ...OrderFunc) *SysJobLogQuery {
	sjlq.order = append(sjlq.order, o...)
	return sjlq
}

// QueryJob chains the current query on the "job" edge.
func (sjlq *SysJobLogQuery) QueryJob() *SysJobQuery {
	query := &SysJobQuery{config: sjlq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sjlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sjlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysjoblog.Table, sysjoblog.FieldID, selector),
			sqlgraph.To(sysjob.Table, sysjob.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sysjoblog.JobTable, sysjoblog.JobColumn),
		)
		fromU = sqlgraph.SetNeighbors(sjlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysJobLog entity from the query.
// Returns a *NotFoundError when no SysJobLog was found.
func (sjlq *SysJobLogQuery) First(ctx context.Context) (*SysJobLog, error) {
	nodes, err := sjlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysjoblog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sjlq *SysJobLogQuery) FirstX(ctx context.Context) *SysJobLog {
	node, err := sjlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysJobLog ID from the query.
// Returns a *NotFoundError when no SysJobLog ID was found.
func (sjlq *SysJobLogQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sjlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysjoblog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sjlq *SysJobLogQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sjlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysJobLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SysJobLog entity is not found.
// Returns a *NotFoundError when no SysJobLog entities are found.
func (sjlq *SysJobLogQuery) Only(ctx context.Context) (*SysJobLog, error) {
	nodes, err := sjlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysjoblog.Label}
	default:
		return nil, &NotSingularError{sysjoblog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sjlq *SysJobLogQuery) OnlyX(ctx context.Context) *SysJobLog {
	node, err := sjlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysJobLog ID in the query.
// Returns a *NotSingularError when exactly one SysJobLog ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sjlq *SysJobLogQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sjlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = &NotSingularError{sysjoblog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sjlq *SysJobLogQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sjlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysJobLogs.
func (sjlq *SysJobLogQuery) All(ctx context.Context) ([]*SysJobLog, error) {
	if err := sjlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sjlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sjlq *SysJobLogQuery) AllX(ctx context.Context) []*SysJobLog {
	nodes, err := sjlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysJobLog IDs.
func (sjlq *SysJobLogQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := sjlq.Select(sysjoblog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sjlq *SysJobLogQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sjlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sjlq *SysJobLogQuery) Count(ctx context.Context) (int, error) {
	if err := sjlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sjlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sjlq *SysJobLogQuery) CountX(ctx context.Context) int {
	count, err := sjlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sjlq *SysJobLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := sjlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sjlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sjlq *SysJobLogQuery) ExistX(ctx context.Context) bool {
	exist, err := sjlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysJobLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sjlq *SysJobLogQuery) Clone() *SysJobLogQuery {
	if sjlq == nil {
		return nil
	}
	return &SysJobLogQuery{
		config:     sjlq.config,
		limit:      sjlq.limit,
		offset:     sjlq.offset,
		order:      append([]OrderFunc{}, sjlq.order...),
		predicates: append([]predicate.SysJobLog{}, sjlq.predicates...),
		withJob:    sjlq.withJob.Clone(),
		// clone intermediate query.
		sql:  sjlq.sql.Clone(),
		path: sjlq.path,
	}
}

// WithJob tells the query-builder to eager-load the nodes that are connected to
// the "job" edge. The optional arguments are used to configure the query builder of the edge.
func (sjlq *SysJobLogQuery) WithJob(opts ...func(*SysJobQuery)) *SysJobLogQuery {
	query := &SysJobQuery{config: sjlq.config}
	for _, opt := range opts {
		opt(query)
	}
	sjlq.withJob = query
	return sjlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JobId int64 `json:"jobId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysJobLog.Query().
//		GroupBy(sysjoblog.FieldJobId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sjlq *SysJobLogQuery) GroupBy(field string, fields ...string) *SysJobLogGroupBy {
	group := &SysJobLogGroupBy{config: sjlq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sjlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sjlq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JobId int64 `json:"jobId,omitempty"`
//	}
//
//	client.SysJobLog.Query().
//		Select(sysjoblog.FieldJobId).
//		Scan(ctx, &v)
//
func (sjlq *SysJobLogQuery) Select(fields ...string) *SysJobLogSelect {
	sjlq.fields = append(sjlq.fields, fields...)
	return &SysJobLogSelect{SysJobLogQuery: sjlq}
}

func (sjlq *SysJobLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sjlq.fields {
		if !sysjoblog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sjlq.path != nil {
		prev, err := sjlq.path(ctx)
		if err != nil {
			return err
		}
		sjlq.sql = prev
	}
	return nil
}

func (sjlq *SysJobLogQuery) sqlAll(ctx context.Context) ([]*SysJobLog, error) {
	var (
		nodes       = []*SysJobLog{}
		_spec       = sjlq.querySpec()
		loadedTypes = [1]bool{
			sjlq.withJob != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysJobLog{config: sjlq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sjlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sjlq.withJob; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*SysJobLog)
		for i := range nodes {
			fk := nodes[i].JobId
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(sysjob.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "jobId" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Job = n
			}
		}
	}

	return nodes, nil
}

func (sjlq *SysJobLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sjlq.querySpec()
	_spec.Node.Columns = sjlq.fields
	if len(sjlq.fields) > 0 {
		_spec.Unique = sjlq.unique != nil && *sjlq.unique
	}
	return sqlgraph.CountNodes(ctx, sjlq.driver, _spec)
}

func (sjlq *SysJobLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sjlq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sjlq *SysJobLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysjoblog.Table,
			Columns: sysjoblog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysjoblog.FieldID,
			},
		},
		From:   sjlq.sql,
		Unique: true,
	}
	if unique := sjlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sjlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysjoblog.FieldID)
		for i := range fields {
			if fields[i] != sysjoblog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sjlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sjlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sjlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sjlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sjlq *SysJobLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sjlq.driver.Dialect())
	t1 := builder.Table(sysjoblog.Table)
	columns := sjlq.fields
	if len(columns) == 0 {
		columns = sysjoblog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sjlq.sql != nil {
		selector = sjlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sjlq.unique != nil && *sjlq.unique {
		selector.Distinct()
	}
	for _, p := range sjlq.predicates {
		p(selector)
	}
	for _, p := range sjlq.order {
		p(selector)
	}
	if offset := sjlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sjlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysJobLogGroupBy is the group-by builder for SysJobLog entities.
type SysJobLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sjlgb *SysJobLogGroupBy) Aggregate(fns ...AggregateFunc) *SysJobLogGroupBy {
	sjlgb.fns = append(sjlgb.fns, fns...)
	return sjlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sjlgb *SysJobLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sjlgb.path(ctx)
	if err != nil {
		return err
	}
	sjlgb.sql = query
	return sjlgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sjlgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sjlgb.fields) > 1 {
		return nil, errors.New("ent: SysJobLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sjlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := sjlgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sjlgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) StringX(ctx context.Context) string {
	v, err := sjlgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sjlgb.fields) > 1 {
		return nil, errors.New("ent: SysJobLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sjlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := sjlgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sjlgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) IntX(ctx context.Context) int {
	v, err := sjlgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sjlgb.fields) > 1 {
		return nil, errors.New("ent: SysJobLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sjlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sjlgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sjlgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sjlgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sjlgb.fields) > 1 {
		return nil, errors.New("ent: SysJobLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sjlgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sjlgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sjlgb *SysJobLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sjlgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sjlgb *SysJobLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := sjlgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sjlgb *SysJobLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sjlgb.fields {
		if !sysjoblog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sjlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sjlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sjlgb *SysJobLogGroupBy) sqlQuery() *sql.Selector {
	selector := sjlgb.sql.Select()
	aggregation := make([]string, 0, len(sjlgb.fns))
	for _, fn := range sjlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sjlgb.fields)+len(sjlgb.fns))
		for _, f := range sjlgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sjlgb.fields...)...)
}

// SysJobLogSelect is the builder for selecting fields of SysJobLog entities.
type SysJobLogSelect struct {
	*SysJobLogQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sjls *SysJobLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sjls.prepareQuery(ctx); err != nil {
		return err
	}
	sjls.sql = sjls.SysJobLogQuery.sqlQuery(ctx)
	return sjls.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sjls *SysJobLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sjls.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sjls.fields) > 1 {
		return nil, errors.New("ent: SysJobLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sjls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sjls *SysJobLogSelect) StringsX(ctx context.Context) []string {
	v, err := sjls.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sjls.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sjls *SysJobLogSelect) StringX(ctx context.Context) string {
	v, err := sjls.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sjls.fields) > 1 {
		return nil, errors.New("ent: SysJobLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sjls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sjls *SysJobLogSelect) IntsX(ctx context.Context) []int {
	v, err := sjls.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sjls.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sjls *SysJobLogSelect) IntX(ctx context.Context) int {
	v, err := sjls.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sjls.fields) > 1 {
		return nil, errors.New("ent: SysJobLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sjls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sjls *SysJobLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sjls.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sjls.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sjls *SysJobLogSelect) Float64X(ctx context.Context) float64 {
	v, err := sjls.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sjls.fields) > 1 {
		return nil, errors.New("ent: SysJobLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sjls.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sjls *SysJobLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := sjls.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sjls *SysJobLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sjls.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysjoblog.Label}
	default:
		err = fmt.Errorf("ent: SysJobLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sjls *SysJobLogSelect) BoolX(ctx context.Context) bool {
	v, err := sjls.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sjls *SysJobLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sjls.sql.Query()
	if err := sjls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
