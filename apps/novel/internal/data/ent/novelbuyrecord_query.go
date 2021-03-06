// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelbuyrecord"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelBuyRecordQuery is the builder for querying NovelBuyRecord entities.
type NovelBuyRecordQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NovelBuyRecord
	// eager-loading edges.
	withUser *SocialUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NovelBuyRecordQuery builder.
func (nbrq *NovelBuyRecordQuery) Where(ps ...predicate.NovelBuyRecord) *NovelBuyRecordQuery {
	nbrq.predicates = append(nbrq.predicates, ps...)
	return nbrq
}

// Limit adds a limit step to the query.
func (nbrq *NovelBuyRecordQuery) Limit(limit int) *NovelBuyRecordQuery {
	nbrq.limit = &limit
	return nbrq
}

// Offset adds an offset step to the query.
func (nbrq *NovelBuyRecordQuery) Offset(offset int) *NovelBuyRecordQuery {
	nbrq.offset = &offset
	return nbrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nbrq *NovelBuyRecordQuery) Unique(unique bool) *NovelBuyRecordQuery {
	nbrq.unique = &unique
	return nbrq
}

// Order adds an order step to the query.
func (nbrq *NovelBuyRecordQuery) Order(o ...OrderFunc) *NovelBuyRecordQuery {
	nbrq.order = append(nbrq.order, o...)
	return nbrq
}

// QueryUser chains the current query on the "user" edge.
func (nbrq *NovelBuyRecordQuery) QueryUser() *SocialUserQuery {
	query := &SocialUserQuery{config: nbrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nbrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nbrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novelbuyrecord.Table, novelbuyrecord.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, novelbuyrecord.UserTable, novelbuyrecord.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(nbrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NovelBuyRecord entity from the query.
// Returns a *NotFoundError when no NovelBuyRecord was found.
func (nbrq *NovelBuyRecordQuery) First(ctx context.Context) (*NovelBuyRecord, error) {
	nodes, err := nbrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{novelbuyrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) FirstX(ctx context.Context) *NovelBuyRecord {
	node, err := nbrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NovelBuyRecord ID from the query.
// Returns a *NotFoundError when no NovelBuyRecord ID was found.
func (nbrq *NovelBuyRecordQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nbrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{novelbuyrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) FirstIDX(ctx context.Context) int64 {
	id, err := nbrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NovelBuyRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NovelBuyRecord entity is not found.
// Returns a *NotFoundError when no NovelBuyRecord entities are found.
func (nbrq *NovelBuyRecordQuery) Only(ctx context.Context) (*NovelBuyRecord, error) {
	nodes, err := nbrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{novelbuyrecord.Label}
	default:
		return nil, &NotSingularError{novelbuyrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) OnlyX(ctx context.Context) *NovelBuyRecord {
	node, err := nbrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NovelBuyRecord ID in the query.
// Returns a *NotSingularError when exactly one NovelBuyRecord ID is not found.
// Returns a *NotFoundError when no entities are found.
func (nbrq *NovelBuyRecordQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nbrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = &NotSingularError{novelbuyrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := nbrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NovelBuyRecords.
func (nbrq *NovelBuyRecordQuery) All(ctx context.Context) ([]*NovelBuyRecord, error) {
	if err := nbrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nbrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) AllX(ctx context.Context) []*NovelBuyRecord {
	nodes, err := nbrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NovelBuyRecord IDs.
func (nbrq *NovelBuyRecordQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := nbrq.Select(novelbuyrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) IDsX(ctx context.Context) []int64 {
	ids, err := nbrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nbrq *NovelBuyRecordQuery) Count(ctx context.Context) (int, error) {
	if err := nbrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nbrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) CountX(ctx context.Context) int {
	count, err := nbrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nbrq *NovelBuyRecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := nbrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nbrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nbrq *NovelBuyRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := nbrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NovelBuyRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nbrq *NovelBuyRecordQuery) Clone() *NovelBuyRecordQuery {
	if nbrq == nil {
		return nil
	}
	return &NovelBuyRecordQuery{
		config:     nbrq.config,
		limit:      nbrq.limit,
		offset:     nbrq.offset,
		order:      append([]OrderFunc{}, nbrq.order...),
		predicates: append([]predicate.NovelBuyRecord{}, nbrq.predicates...),
		withUser:   nbrq.withUser.Clone(),
		// clone intermediate query.
		sql:  nbrq.sql.Clone(),
		path: nbrq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (nbrq *NovelBuyRecordQuery) WithUser(opts ...func(*SocialUserQuery)) *NovelBuyRecordQuery {
	query := &SocialUserQuery{config: nbrq.config}
	for _, opt := range opts {
		opt(query)
	}
	nbrq.withUser = query
	return nbrq
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
//	client.NovelBuyRecord.Query().
//		GroupBy(novelbuyrecord.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nbrq *NovelBuyRecordQuery) GroupBy(field string, fields ...string) *NovelBuyRecordGroupBy {
	group := &NovelBuyRecordGroupBy{config: nbrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nbrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nbrq.sqlQuery(ctx), nil
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
//	client.NovelBuyRecord.Query().
//		Select(novelbuyrecord.FieldUserId).
//		Scan(ctx, &v)
//
func (nbrq *NovelBuyRecordQuery) Select(fields ...string) *NovelBuyRecordSelect {
	nbrq.fields = append(nbrq.fields, fields...)
	return &NovelBuyRecordSelect{NovelBuyRecordQuery: nbrq}
}

func (nbrq *NovelBuyRecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nbrq.fields {
		if !novelbuyrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nbrq.path != nil {
		prev, err := nbrq.path(ctx)
		if err != nil {
			return err
		}
		nbrq.sql = prev
	}
	return nil
}

func (nbrq *NovelBuyRecordQuery) sqlAll(ctx context.Context) ([]*NovelBuyRecord, error) {
	var (
		nodes       = []*NovelBuyRecord{}
		_spec       = nbrq.querySpec()
		loadedTypes = [1]bool{
			nbrq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NovelBuyRecord{config: nbrq.config}
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
	if err := sqlgraph.QueryNodes(ctx, nbrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nbrq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*NovelBuyRecord)
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

func (nbrq *NovelBuyRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nbrq.querySpec()
	_spec.Node.Columns = nbrq.fields
	if len(nbrq.fields) > 0 {
		_spec.Unique = nbrq.unique != nil && *nbrq.unique
	}
	return sqlgraph.CountNodes(ctx, nbrq.driver, _spec)
}

func (nbrq *NovelBuyRecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nbrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nbrq *NovelBuyRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelbuyrecord.Table,
			Columns: novelbuyrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbuyrecord.FieldID,
			},
		},
		From:   nbrq.sql,
		Unique: true,
	}
	if unique := nbrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nbrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelbuyrecord.FieldID)
		for i := range fields {
			if fields[i] != novelbuyrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nbrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nbrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nbrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nbrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nbrq *NovelBuyRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nbrq.driver.Dialect())
	t1 := builder.Table(novelbuyrecord.Table)
	columns := nbrq.fields
	if len(columns) == 0 {
		columns = novelbuyrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nbrq.sql != nil {
		selector = nbrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nbrq.unique != nil && *nbrq.unique {
		selector.Distinct()
	}
	for _, p := range nbrq.predicates {
		p(selector)
	}
	for _, p := range nbrq.order {
		p(selector)
	}
	if offset := nbrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nbrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NovelBuyRecordGroupBy is the group-by builder for NovelBuyRecord entities.
type NovelBuyRecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nbrgb *NovelBuyRecordGroupBy) Aggregate(fns ...AggregateFunc) *NovelBuyRecordGroupBy {
	nbrgb.fns = append(nbrgb.fns, fns...)
	return nbrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (nbrgb *NovelBuyRecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nbrgb.path(ctx)
	if err != nil {
		return err
	}
	nbrgb.sql = query
	return nbrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := nbrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(nbrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := nbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := nbrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nbrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) StringX(ctx context.Context) string {
	v, err := nbrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(nbrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := nbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := nbrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nbrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) IntX(ctx context.Context) int {
	v, err := nbrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(nbrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := nbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := nbrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nbrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := nbrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(nbrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := nbrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := nbrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbrgb *NovelBuyRecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nbrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nbrgb *NovelBuyRecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := nbrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nbrgb *NovelBuyRecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nbrgb.fields {
		if !novelbuyrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nbrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nbrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nbrgb *NovelBuyRecordGroupBy) sqlQuery() *sql.Selector {
	selector := nbrgb.sql.Select()
	aggregation := make([]string, 0, len(nbrgb.fns))
	for _, fn := range nbrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(nbrgb.fields)+len(nbrgb.fns))
		for _, f := range nbrgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(nbrgb.fields...)...)
}

// NovelBuyRecordSelect is the builder for selecting fields of NovelBuyRecord entities.
type NovelBuyRecordSelect struct {
	*NovelBuyRecordQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nbrs *NovelBuyRecordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nbrs.prepareQuery(ctx); err != nil {
		return err
	}
	nbrs.sql = nbrs.NovelBuyRecordQuery.sqlQuery(ctx)
	return nbrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := nbrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(nbrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) StringsX(ctx context.Context) []string {
	v, err := nbrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nbrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) StringX(ctx context.Context) string {
	v, err := nbrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(nbrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) IntsX(ctx context.Context) []int {
	v, err := nbrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nbrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) IntX(ctx context.Context) int {
	v, err := nbrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nbrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := nbrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nbrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) Float64X(ctx context.Context) float64 {
	v, err := nbrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nbrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyRecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nbrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := nbrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nbrs *NovelBuyRecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nbrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuyrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyRecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nbrs *NovelBuyRecordSelect) BoolX(ctx context.Context) bool {
	v, err := nbrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nbrs *NovelBuyRecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nbrs.sql.Query()
	if err := nbrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
