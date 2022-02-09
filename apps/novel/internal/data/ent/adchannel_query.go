// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/adchannel"
	"hope/apps/novel/internal/data/ent/payorder"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdChannelQuery is the builder for querying AdChannel entities.
type AdChannelQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AdChannel
	// eager-loading edges.
	withUsers  *SocialUserQuery
	withOrders *PayOrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdChannelQuery builder.
func (acq *AdChannelQuery) Where(ps ...predicate.AdChannel) *AdChannelQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit adds a limit step to the query.
func (acq *AdChannelQuery) Limit(limit int) *AdChannelQuery {
	acq.limit = &limit
	return acq
}

// Offset adds an offset step to the query.
func (acq *AdChannelQuery) Offset(offset int) *AdChannelQuery {
	acq.offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AdChannelQuery) Unique(unique bool) *AdChannelQuery {
	acq.unique = &unique
	return acq
}

// Order adds an order step to the query.
func (acq *AdChannelQuery) Order(o ...OrderFunc) *AdChannelQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// QueryUsers chains the current query on the "users" edge.
func (acq *AdChannelQuery) QueryUsers() *SocialUserQuery {
	query := &SocialUserQuery{config: acq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adchannel.Table, adchannel.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adchannel.UsersTable, adchannel.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrders chains the current query on the "orders" edge.
func (acq *AdChannelQuery) QueryOrders() *PayOrderQuery {
	query := &PayOrderQuery{config: acq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := acq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(adchannel.Table, adchannel.FieldID, selector),
			sqlgraph.To(payorder.Table, payorder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, adchannel.OrdersTable, adchannel.OrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(acq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AdChannel entity from the query.
// Returns a *NotFoundError when no AdChannel was found.
func (acq *AdChannelQuery) First(ctx context.Context) (*AdChannel, error) {
	nodes, err := acq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adchannel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AdChannelQuery) FirstX(ctx context.Context) *AdChannel {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdChannel ID from the query.
// Returns a *NotFoundError when no AdChannel ID was found.
func (acq *AdChannelQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = acq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adchannel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AdChannelQuery) FirstIDX(ctx context.Context) int64 {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdChannel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AdChannel entity is not found.
// Returns a *NotFoundError when no AdChannel entities are found.
func (acq *AdChannelQuery) Only(ctx context.Context) (*AdChannel, error) {
	nodes, err := acq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adchannel.Label}
	default:
		return nil, &NotSingularError{adchannel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AdChannelQuery) OnlyX(ctx context.Context) *AdChannel {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdChannel ID in the query.
// Returns a *NotSingularError when exactly one AdChannel ID is not found.
// Returns a *NotFoundError when no entities are found.
func (acq *AdChannelQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = acq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = &NotSingularError{adchannel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AdChannelQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdChannels.
func (acq *AdChannelQuery) All(ctx context.Context) ([]*AdChannel, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return acq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (acq *AdChannelQuery) AllX(ctx context.Context) []*AdChannel {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdChannel IDs.
func (acq *AdChannelQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := acq.Select(adchannel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AdChannelQuery) IDsX(ctx context.Context) []int64 {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AdChannelQuery) Count(ctx context.Context) (int, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return acq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AdChannelQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AdChannelQuery) Exist(ctx context.Context) (bool, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return acq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AdChannelQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdChannelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AdChannelQuery) Clone() *AdChannelQuery {
	if acq == nil {
		return nil
	}
	return &AdChannelQuery{
		config:     acq.config,
		limit:      acq.limit,
		offset:     acq.offset,
		order:      append([]OrderFunc{}, acq.order...),
		predicates: append([]predicate.AdChannel{}, acq.predicates...),
		withUsers:  acq.withUsers.Clone(),
		withOrders: acq.withOrders.Clone(),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AdChannelQuery) WithUsers(opts ...func(*SocialUserQuery)) *AdChannelQuery {
	query := &SocialUserQuery{config: acq.config}
	for _, opt := range opts {
		opt(query)
	}
	acq.withUsers = query
	return acq
}

// WithOrders tells the query-builder to eager-load the nodes that are connected to
// the "orders" edge. The optional arguments are used to configure the query builder of the edge.
func (acq *AdChannelQuery) WithOrders(opts ...func(*PayOrderQuery)) *AdChannelQuery {
	query := &PayOrderQuery{config: acq.config}
	for _, opt := range opts {
		opt(query)
	}
	acq.withOrders = query
	return acq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ChannelName string `json:"channelName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdChannel.Query().
//		GroupBy(adchannel.FieldChannelName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (acq *AdChannelQuery) GroupBy(field string, fields ...string) *AdChannelGroupBy {
	group := &AdChannelGroupBy{config: acq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return acq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ChannelName string `json:"channelName,omitempty"`
//	}
//
//	client.AdChannel.Query().
//		Select(adchannel.FieldChannelName).
//		Scan(ctx, &v)
//
func (acq *AdChannelQuery) Select(fields ...string) *AdChannelSelect {
	acq.fields = append(acq.fields, fields...)
	return &AdChannelSelect{AdChannelQuery: acq}
}

func (acq *AdChannelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range acq.fields {
		if !adchannel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AdChannelQuery) sqlAll(ctx context.Context) ([]*AdChannel, error) {
	var (
		nodes       = []*AdChannel{}
		_spec       = acq.querySpec()
		loadedTypes = [2]bool{
			acq.withUsers != nil,
			acq.withOrders != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AdChannel{config: acq.config}
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
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := acq.withUsers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*AdChannel)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Users = []*SocialUser{}
		}
		query.Where(predicate.SocialUser(func(s *sql.Selector) {
			s.Where(sql.InValues(adchannel.UsersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ChId
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "chId" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Users = append(node.Edges.Users, n)
		}
	}

	if query := acq.withOrders; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*AdChannel)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Orders = []*PayOrder{}
		}
		query.Where(predicate.PayOrder(func(s *sql.Selector) {
			s.Where(sql.InValues(adchannel.OrdersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ChId
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "chId" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Orders = append(node.Edges.Orders, n)
		}
	}

	return nodes, nil
}

func (acq *AdChannelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	_spec.Node.Columns = acq.fields
	if len(acq.fields) > 0 {
		_spec.Unique = acq.unique != nil && *acq.unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AdChannelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := acq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (acq *AdChannelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adchannel.Table,
			Columns: adchannel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: adchannel.FieldID,
			},
		},
		From:   acq.sql,
		Unique: true,
	}
	if unique := acq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := acq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adchannel.FieldID)
		for i := range fields {
			if fields[i] != adchannel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AdChannelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(adchannel.Table)
	columns := acq.fields
	if len(columns) == 0 {
		columns = adchannel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.unique != nil && *acq.unique {
		selector.Distinct()
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdChannelGroupBy is the group-by builder for AdChannel entities.
type AdChannelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AdChannelGroupBy) Aggregate(fns ...AggregateFunc) *AdChannelGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the group-by query and scans the result into the given value.
func (acgb *AdChannelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := acgb.path(ctx)
	if err != nil {
		return err
	}
	acgb.sql = query
	return acgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acgb *AdChannelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := acgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AdChannelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acgb *AdChannelGroupBy) StringsX(ctx context.Context) []string {
	v, err := acgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acgb *AdChannelGroupBy) StringX(ctx context.Context) string {
	v, err := acgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AdChannelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acgb *AdChannelGroupBy) IntsX(ctx context.Context) []int {
	v, err := acgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acgb *AdChannelGroupBy) IntX(ctx context.Context) int {
	v, err := acgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AdChannelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acgb *AdChannelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := acgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acgb *AdChannelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := acgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AdChannelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acgb *AdChannelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := acgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AdChannelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acgb *AdChannelGroupBy) BoolX(ctx context.Context) bool {
	v, err := acgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acgb *AdChannelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range acgb.fields {
		if !adchannel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := acgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (acgb *AdChannelGroupBy) sqlQuery() *sql.Selector {
	selector := acgb.sql.Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(acgb.fields)+len(acgb.fns))
		for _, f := range acgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(acgb.fields...)...)
}

// AdChannelSelect is the builder for selecting fields of AdChannel entities.
type AdChannelSelect struct {
	*AdChannelQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AdChannelSelect) Scan(ctx context.Context, v interface{}) error {
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	acs.sql = acs.AdChannelQuery.sqlQuery(ctx)
	return acs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acs *AdChannelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := acs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AdChannelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acs *AdChannelSelect) StringsX(ctx context.Context) []string {
	v, err := acs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acs *AdChannelSelect) StringX(ctx context.Context) string {
	v, err := acs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AdChannelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acs *AdChannelSelect) IntsX(ctx context.Context) []int {
	v, err := acs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acs *AdChannelSelect) IntX(ctx context.Context) int {
	v, err := acs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AdChannelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acs *AdChannelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := acs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acs *AdChannelSelect) Float64X(ctx context.Context) float64 {
	v, err := acs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AdChannelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acs *AdChannelSelect) BoolsX(ctx context.Context) []bool {
	v, err := acs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (acs *AdChannelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{adchannel.Label}
	default:
		err = fmt.Errorf("ent: AdChannelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acs *AdChannelSelect) BoolX(ctx context.Context) bool {
	v, err := acs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acs *AdChannelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := acs.sql.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
