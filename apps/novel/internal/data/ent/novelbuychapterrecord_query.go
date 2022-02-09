// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelbuychapterrecord"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelBuyChapterRecordQuery is the builder for querying NovelBuyChapterRecord entities.
type NovelBuyChapterRecordQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NovelBuyChapterRecord
	// eager-loading edges.
	withUser *SocialUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NovelBuyChapterRecordQuery builder.
func (nbcrq *NovelBuyChapterRecordQuery) Where(ps ...predicate.NovelBuyChapterRecord) *NovelBuyChapterRecordQuery {
	nbcrq.predicates = append(nbcrq.predicates, ps...)
	return nbcrq
}

// Limit adds a limit step to the query.
func (nbcrq *NovelBuyChapterRecordQuery) Limit(limit int) *NovelBuyChapterRecordQuery {
	nbcrq.limit = &limit
	return nbcrq
}

// Offset adds an offset step to the query.
func (nbcrq *NovelBuyChapterRecordQuery) Offset(offset int) *NovelBuyChapterRecordQuery {
	nbcrq.offset = &offset
	return nbcrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nbcrq *NovelBuyChapterRecordQuery) Unique(unique bool) *NovelBuyChapterRecordQuery {
	nbcrq.unique = &unique
	return nbcrq
}

// Order adds an order step to the query.
func (nbcrq *NovelBuyChapterRecordQuery) Order(o ...OrderFunc) *NovelBuyChapterRecordQuery {
	nbcrq.order = append(nbcrq.order, o...)
	return nbcrq
}

// QueryUser chains the current query on the "user" edge.
func (nbcrq *NovelBuyChapterRecordQuery) QueryUser() *SocialUserQuery {
	query := &SocialUserQuery{config: nbcrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nbcrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nbcrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novelbuychapterrecord.Table, novelbuychapterrecord.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, novelbuychapterrecord.UserTable, novelbuychapterrecord.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(nbcrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NovelBuyChapterRecord entity from the query.
// Returns a *NotFoundError when no NovelBuyChapterRecord was found.
func (nbcrq *NovelBuyChapterRecordQuery) First(ctx context.Context) (*NovelBuyChapterRecord, error) {
	nodes, err := nbcrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{novelbuychapterrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) FirstX(ctx context.Context) *NovelBuyChapterRecord {
	node, err := nbcrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NovelBuyChapterRecord ID from the query.
// Returns a *NotFoundError when no NovelBuyChapterRecord ID was found.
func (nbcrq *NovelBuyChapterRecordQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nbcrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{novelbuychapterrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) FirstIDX(ctx context.Context) int64 {
	id, err := nbcrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NovelBuyChapterRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NovelBuyChapterRecord entity is not found.
// Returns a *NotFoundError when no NovelBuyChapterRecord entities are found.
func (nbcrq *NovelBuyChapterRecordQuery) Only(ctx context.Context) (*NovelBuyChapterRecord, error) {
	nodes, err := nbcrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{novelbuychapterrecord.Label}
	default:
		return nil, &NotSingularError{novelbuychapterrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) OnlyX(ctx context.Context) *NovelBuyChapterRecord {
	node, err := nbcrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NovelBuyChapterRecord ID in the query.
// Returns a *NotSingularError when exactly one NovelBuyChapterRecord ID is not found.
// Returns a *NotFoundError when no entities are found.
func (nbcrq *NovelBuyChapterRecordQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = nbcrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = &NotSingularError{novelbuychapterrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := nbcrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NovelBuyChapterRecords.
func (nbcrq *NovelBuyChapterRecordQuery) All(ctx context.Context) ([]*NovelBuyChapterRecord, error) {
	if err := nbcrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return nbcrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) AllX(ctx context.Context) []*NovelBuyChapterRecord {
	nodes, err := nbcrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NovelBuyChapterRecord IDs.
func (nbcrq *NovelBuyChapterRecordQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := nbcrq.Select(novelbuychapterrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) IDsX(ctx context.Context) []int64 {
	ids, err := nbcrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nbcrq *NovelBuyChapterRecordQuery) Count(ctx context.Context) (int, error) {
	if err := nbcrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return nbcrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) CountX(ctx context.Context) int {
	count, err := nbcrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nbcrq *NovelBuyChapterRecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := nbcrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return nbcrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (nbcrq *NovelBuyChapterRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := nbcrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NovelBuyChapterRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nbcrq *NovelBuyChapterRecordQuery) Clone() *NovelBuyChapterRecordQuery {
	if nbcrq == nil {
		return nil
	}
	return &NovelBuyChapterRecordQuery{
		config:     nbcrq.config,
		limit:      nbcrq.limit,
		offset:     nbcrq.offset,
		order:      append([]OrderFunc{}, nbcrq.order...),
		predicates: append([]predicate.NovelBuyChapterRecord{}, nbcrq.predicates...),
		withUser:   nbcrq.withUser.Clone(),
		// clone intermediate query.
		sql:  nbcrq.sql.Clone(),
		path: nbcrq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (nbcrq *NovelBuyChapterRecordQuery) WithUser(opts ...func(*SocialUserQuery)) *NovelBuyChapterRecordQuery {
	query := &SocialUserQuery{config: nbcrq.config}
	for _, opt := range opts {
		opt(query)
	}
	nbcrq.withUser = query
	return nbcrq
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
//	client.NovelBuyChapterRecord.Query().
//		GroupBy(novelbuychapterrecord.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (nbcrq *NovelBuyChapterRecordQuery) GroupBy(field string, fields ...string) *NovelBuyChapterRecordGroupBy {
	group := &NovelBuyChapterRecordGroupBy{config: nbcrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := nbcrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return nbcrq.sqlQuery(ctx), nil
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
//	client.NovelBuyChapterRecord.Query().
//		Select(novelbuychapterrecord.FieldUserId).
//		Scan(ctx, &v)
//
func (nbcrq *NovelBuyChapterRecordQuery) Select(fields ...string) *NovelBuyChapterRecordSelect {
	nbcrq.fields = append(nbcrq.fields, fields...)
	return &NovelBuyChapterRecordSelect{NovelBuyChapterRecordQuery: nbcrq}
}

func (nbcrq *NovelBuyChapterRecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range nbcrq.fields {
		if !novelbuychapterrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nbcrq.path != nil {
		prev, err := nbcrq.path(ctx)
		if err != nil {
			return err
		}
		nbcrq.sql = prev
	}
	return nil
}

func (nbcrq *NovelBuyChapterRecordQuery) sqlAll(ctx context.Context) ([]*NovelBuyChapterRecord, error) {
	var (
		nodes       = []*NovelBuyChapterRecord{}
		_spec       = nbcrq.querySpec()
		loadedTypes = [1]bool{
			nbcrq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NovelBuyChapterRecord{config: nbcrq.config}
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
	if err := sqlgraph.QueryNodes(ctx, nbcrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := nbcrq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*NovelBuyChapterRecord)
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

func (nbcrq *NovelBuyChapterRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nbcrq.querySpec()
	_spec.Node.Columns = nbcrq.fields
	if len(nbcrq.fields) > 0 {
		_spec.Unique = nbcrq.unique != nil && *nbcrq.unique
	}
	return sqlgraph.CountNodes(ctx, nbcrq.driver, _spec)
}

func (nbcrq *NovelBuyChapterRecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := nbcrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (nbcrq *NovelBuyChapterRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelbuychapterrecord.Table,
			Columns: novelbuychapterrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelbuychapterrecord.FieldID,
			},
		},
		From:   nbcrq.sql,
		Unique: true,
	}
	if unique := nbcrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := nbcrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelbuychapterrecord.FieldID)
		for i := range fields {
			if fields[i] != novelbuychapterrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nbcrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nbcrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nbcrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nbcrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nbcrq *NovelBuyChapterRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nbcrq.driver.Dialect())
	t1 := builder.Table(novelbuychapterrecord.Table)
	columns := nbcrq.fields
	if len(columns) == 0 {
		columns = novelbuychapterrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nbcrq.sql != nil {
		selector = nbcrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nbcrq.unique != nil && *nbcrq.unique {
		selector.Distinct()
	}
	for _, p := range nbcrq.predicates {
		p(selector)
	}
	for _, p := range nbcrq.order {
		p(selector)
	}
	if offset := nbcrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nbcrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NovelBuyChapterRecordGroupBy is the group-by builder for NovelBuyChapterRecord entities.
type NovelBuyChapterRecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Aggregate(fns ...AggregateFunc) *NovelBuyChapterRecordGroupBy {
	nbcrgb.fns = append(nbcrgb.fns, fns...)
	return nbcrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := nbcrgb.path(ctx)
	if err != nil {
		return err
	}
	nbcrgb.sql = query
	return nbcrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := nbcrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(nbcrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := nbcrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := nbcrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nbcrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) StringX(ctx context.Context) string {
	v, err := nbcrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(nbcrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := nbcrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := nbcrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nbcrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) IntX(ctx context.Context) int {
	v, err := nbcrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(nbcrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := nbcrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := nbcrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nbcrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := nbcrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(nbcrgb.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := nbcrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := nbcrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (nbcrgb *NovelBuyChapterRecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nbcrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nbcrgb *NovelBuyChapterRecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := nbcrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nbcrgb *NovelBuyChapterRecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range nbcrgb.fields {
		if !novelbuychapterrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := nbcrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nbcrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (nbcrgb *NovelBuyChapterRecordGroupBy) sqlQuery() *sql.Selector {
	selector := nbcrgb.sql.Select()
	aggregation := make([]string, 0, len(nbcrgb.fns))
	for _, fn := range nbcrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(nbcrgb.fields)+len(nbcrgb.fns))
		for _, f := range nbcrgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(nbcrgb.fields...)...)
}

// NovelBuyChapterRecordSelect is the builder for selecting fields of NovelBuyChapterRecord entities.
type NovelBuyChapterRecordSelect struct {
	*NovelBuyChapterRecordQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (nbcrs *NovelBuyChapterRecordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := nbcrs.prepareQuery(ctx); err != nil {
		return err
	}
	nbcrs.sql = nbcrs.NovelBuyChapterRecordQuery.sqlQuery(ctx)
	return nbcrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := nbcrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(nbcrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := nbcrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) StringsX(ctx context.Context) []string {
	v, err := nbcrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = nbcrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) StringX(ctx context.Context) string {
	v, err := nbcrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(nbcrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := nbcrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) IntsX(ctx context.Context) []int {
	v, err := nbcrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = nbcrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) IntX(ctx context.Context) int {
	v, err := nbcrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(nbcrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := nbcrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := nbcrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = nbcrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) Float64X(ctx context.Context) float64 {
	v, err := nbcrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(nbcrs.fields) > 1 {
		return nil, errors.New("ent: NovelBuyChapterRecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := nbcrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := nbcrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (nbcrs *NovelBuyChapterRecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = nbcrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelbuychapterrecord.Label}
	default:
		err = fmt.Errorf("ent: NovelBuyChapterRecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (nbcrs *NovelBuyChapterRecordSelect) BoolX(ctx context.Context) bool {
	v, err := nbcrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nbcrs *NovelBuyChapterRecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := nbcrs.sql.Query()
	if err := nbcrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
