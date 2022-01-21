// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/agreementlog"
	"hope/apps/novel/internal/data/ent/payorder"
	"hope/apps/novel/internal/data/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgreementLogQuery is the builder for querying AgreementLog entities.
type AgreementLogQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AgreementLog
	// eager-loading edges.
	withOrders *PayOrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AgreementLogQuery builder.
func (alq *AgreementLogQuery) Where(ps ...predicate.AgreementLog) *AgreementLogQuery {
	alq.predicates = append(alq.predicates, ps...)
	return alq
}

// Limit adds a limit step to the query.
func (alq *AgreementLogQuery) Limit(limit int) *AgreementLogQuery {
	alq.limit = &limit
	return alq
}

// Offset adds an offset step to the query.
func (alq *AgreementLogQuery) Offset(offset int) *AgreementLogQuery {
	alq.offset = &offset
	return alq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (alq *AgreementLogQuery) Unique(unique bool) *AgreementLogQuery {
	alq.unique = &unique
	return alq
}

// Order adds an order step to the query.
func (alq *AgreementLogQuery) Order(o ...OrderFunc) *AgreementLogQuery {
	alq.order = append(alq.order, o...)
	return alq
}

// QueryOrders chains the current query on the "orders" edge.
func (alq *AgreementLogQuery) QueryOrders() *PayOrderQuery {
	query := &PayOrderQuery{config: alq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := alq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := alq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(agreementlog.Table, agreementlog.FieldID, selector),
			sqlgraph.To(payorder.Table, payorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, agreementlog.OrdersTable, agreementlog.OrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(alq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AgreementLog entity from the query.
// Returns a *NotFoundError when no AgreementLog was found.
func (alq *AgreementLogQuery) First(ctx context.Context) (*AgreementLog, error) {
	nodes, err := alq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{agreementlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (alq *AgreementLogQuery) FirstX(ctx context.Context) *AgreementLog {
	node, err := alq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AgreementLog ID from the query.
// Returns a *NotFoundError when no AgreementLog ID was found.
func (alq *AgreementLogQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = alq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{agreementlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (alq *AgreementLogQuery) FirstIDX(ctx context.Context) int64 {
	id, err := alq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AgreementLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AgreementLog entity is not found.
// Returns a *NotFoundError when no AgreementLog entities are found.
func (alq *AgreementLogQuery) Only(ctx context.Context) (*AgreementLog, error) {
	nodes, err := alq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{agreementlog.Label}
	default:
		return nil, &NotSingularError{agreementlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (alq *AgreementLogQuery) OnlyX(ctx context.Context) *AgreementLog {
	node, err := alq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AgreementLog ID in the query.
// Returns a *NotSingularError when exactly one AgreementLog ID is not found.
// Returns a *NotFoundError when no entities are found.
func (alq *AgreementLogQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = alq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = &NotSingularError{agreementlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (alq *AgreementLogQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := alq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AgreementLogs.
func (alq *AgreementLogQuery) All(ctx context.Context) ([]*AgreementLog, error) {
	if err := alq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return alq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (alq *AgreementLogQuery) AllX(ctx context.Context) []*AgreementLog {
	nodes, err := alq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AgreementLog IDs.
func (alq *AgreementLogQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := alq.Select(agreementlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (alq *AgreementLogQuery) IDsX(ctx context.Context) []int64 {
	ids, err := alq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (alq *AgreementLogQuery) Count(ctx context.Context) (int, error) {
	if err := alq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return alq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (alq *AgreementLogQuery) CountX(ctx context.Context) int {
	count, err := alq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (alq *AgreementLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := alq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return alq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (alq *AgreementLogQuery) ExistX(ctx context.Context) bool {
	exist, err := alq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AgreementLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (alq *AgreementLogQuery) Clone() *AgreementLogQuery {
	if alq == nil {
		return nil
	}
	return &AgreementLogQuery{
		config:     alq.config,
		limit:      alq.limit,
		offset:     alq.offset,
		order:      append([]OrderFunc{}, alq.order...),
		predicates: append([]predicate.AgreementLog{}, alq.predicates...),
		withOrders: alq.withOrders.Clone(),
		// clone intermediate query.
		sql:  alq.sql.Clone(),
		path: alq.path,
	}
}

// WithOrders tells the query-builder to eager-load the nodes that are connected to
// the "orders" edge. The optional arguments are used to configure the query builder of the edge.
func (alq *AgreementLogQuery) WithOrders(opts ...func(*PayOrderQuery)) *AgreementLogQuery {
	query := &PayOrderQuery{config: alq.config}
	for _, opt := range opts {
		opt(query)
	}
	alq.withOrders = query
	return alq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OuterAgreementNo string `json:"outerAgreementNo,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AgreementLog.Query().
//		GroupBy(agreementlog.FieldOuterAgreementNo).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (alq *AgreementLogQuery) GroupBy(field string, fields ...string) *AgreementLogGroupBy {
	group := &AgreementLogGroupBy{config: alq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := alq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return alq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OuterAgreementNo string `json:"outerAgreementNo,omitempty"`
//	}
//
//	client.AgreementLog.Query().
//		Select(agreementlog.FieldOuterAgreementNo).
//		Scan(ctx, &v)
//
func (alq *AgreementLogQuery) Select(fields ...string) *AgreementLogSelect {
	alq.fields = append(alq.fields, fields...)
	return &AgreementLogSelect{AgreementLogQuery: alq}
}

func (alq *AgreementLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range alq.fields {
		if !agreementlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if alq.path != nil {
		prev, err := alq.path(ctx)
		if err != nil {
			return err
		}
		alq.sql = prev
	}
	return nil
}

func (alq *AgreementLogQuery) sqlAll(ctx context.Context) ([]*AgreementLog, error) {
	var (
		nodes       = []*AgreementLog{}
		_spec       = alq.querySpec()
		loadedTypes = [1]bool{
			alq.withOrders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AgreementLog{config: alq.config}
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
	if err := sqlgraph.QueryNodes(ctx, alq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := alq.withOrders; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*AgreementLog)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Orders = []*PayOrder{}
		}
		query.withFKs = true
		query.Where(predicate.PayOrder(func(s *sql.Selector) {
			s.Where(sql.InValues(agreementlog.OrdersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.agreement_log_orders
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "agreement_log_orders" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "agreement_log_orders" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Orders = append(node.Edges.Orders, n)
		}
	}

	return nodes, nil
}

func (alq *AgreementLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := alq.querySpec()
	return sqlgraph.CountNodes(ctx, alq.driver, _spec)
}

func (alq *AgreementLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := alq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (alq *AgreementLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agreementlog.Table,
			Columns: agreementlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: agreementlog.FieldID,
			},
		},
		From:   alq.sql,
		Unique: true,
	}
	if unique := alq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := alq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, agreementlog.FieldID)
		for i := range fields {
			if fields[i] != agreementlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := alq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := alq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := alq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := alq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (alq *AgreementLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(alq.driver.Dialect())
	t1 := builder.Table(agreementlog.Table)
	columns := alq.fields
	if len(columns) == 0 {
		columns = agreementlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if alq.sql != nil {
		selector = alq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range alq.predicates {
		p(selector)
	}
	for _, p := range alq.order {
		p(selector)
	}
	if offset := alq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := alq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AgreementLogGroupBy is the group-by builder for AgreementLog entities.
type AgreementLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (algb *AgreementLogGroupBy) Aggregate(fns ...AggregateFunc) *AgreementLogGroupBy {
	algb.fns = append(algb.fns, fns...)
	return algb
}

// Scan applies the group-by query and scans the result into the given value.
func (algb *AgreementLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := algb.path(ctx)
	if err != nil {
		return err
	}
	algb.sql = query
	return algb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (algb *AgreementLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := algb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("ent: AgreementLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (algb *AgreementLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := algb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = algb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (algb *AgreementLogGroupBy) StringX(ctx context.Context) string {
	v, err := algb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("ent: AgreementLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (algb *AgreementLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := algb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = algb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (algb *AgreementLogGroupBy) IntX(ctx context.Context) int {
	v, err := algb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("ent: AgreementLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (algb *AgreementLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := algb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = algb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (algb *AgreementLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := algb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("ent: AgreementLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (algb *AgreementLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := algb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AgreementLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = algb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (algb *AgreementLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := algb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (algb *AgreementLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range algb.fields {
		if !agreementlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := algb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := algb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (algb *AgreementLogGroupBy) sqlQuery() *sql.Selector {
	selector := algb.sql.Select()
	aggregation := make([]string, 0, len(algb.fns))
	for _, fn := range algb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(algb.fields)+len(algb.fns))
		for _, f := range algb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(algb.fields...)...)
}

// AgreementLogSelect is the builder for selecting fields of AgreementLog entities.
type AgreementLogSelect struct {
	*AgreementLogQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (als *AgreementLogSelect) Scan(ctx context.Context, v interface{}) error {
	if err := als.prepareQuery(ctx); err != nil {
		return err
	}
	als.sql = als.AgreementLogQuery.sqlQuery(ctx)
	return als.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (als *AgreementLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := als.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("ent: AgreementLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (als *AgreementLogSelect) StringsX(ctx context.Context) []string {
	v, err := als.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = als.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (als *AgreementLogSelect) StringX(ctx context.Context) string {
	v, err := als.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("ent: AgreementLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (als *AgreementLogSelect) IntsX(ctx context.Context) []int {
	v, err := als.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = als.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (als *AgreementLogSelect) IntX(ctx context.Context) int {
	v, err := als.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("ent: AgreementLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (als *AgreementLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := als.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = als.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (als *AgreementLogSelect) Float64X(ctx context.Context) float64 {
	v, err := als.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("ent: AgreementLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (als *AgreementLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := als.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (als *AgreementLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = als.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{agreementlog.Label}
	default:
		err = fmt.Errorf("ent: AgreementLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (als *AgreementLogSelect) BoolX(ctx context.Context) bool {
	v, err := als.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (als *AgreementLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := als.sql.Query()
	if err := als.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}