// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"hope/apps/novel/internal/data/ent/vipuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VipUserQuery is the builder for querying VipUser entities.
type VipUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.VipUser
	// eager-loading edges.
	withUser *SocialUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VipUserQuery builder.
func (vuq *VipUserQuery) Where(ps ...predicate.VipUser) *VipUserQuery {
	vuq.predicates = append(vuq.predicates, ps...)
	return vuq
}

// Limit adds a limit step to the query.
func (vuq *VipUserQuery) Limit(limit int) *VipUserQuery {
	vuq.limit = &limit
	return vuq
}

// Offset adds an offset step to the query.
func (vuq *VipUserQuery) Offset(offset int) *VipUserQuery {
	vuq.offset = &offset
	return vuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vuq *VipUserQuery) Unique(unique bool) *VipUserQuery {
	vuq.unique = &unique
	return vuq
}

// Order adds an order step to the query.
func (vuq *VipUserQuery) Order(o ...OrderFunc) *VipUserQuery {
	vuq.order = append(vuq.order, o...)
	return vuq
}

// QueryUser chains the current query on the "user" edge.
func (vuq *VipUserQuery) QueryUser() *SocialUserQuery {
	query := &SocialUserQuery{config: vuq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vipuser.Table, vipuser.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vipuser.UserTable, vipuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(vuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first VipUser entity from the query.
// Returns a *NotFoundError when no VipUser was found.
func (vuq *VipUserQuery) First(ctx context.Context) (*VipUser, error) {
	nodes, err := vuq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vipuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vuq *VipUserQuery) FirstX(ctx context.Context) *VipUser {
	node, err := vuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VipUser ID from the query.
// Returns a *NotFoundError when no VipUser ID was found.
func (vuq *VipUserQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vuq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vipuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vuq *VipUserQuery) FirstIDX(ctx context.Context) int64 {
	id, err := vuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VipUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one VipUser entity is not found.
// Returns a *NotFoundError when no VipUser entities are found.
func (vuq *VipUserQuery) Only(ctx context.Context) (*VipUser, error) {
	nodes, err := vuq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vipuser.Label}
	default:
		return nil, &NotSingularError{vipuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vuq *VipUserQuery) OnlyX(ctx context.Context) *VipUser {
	node, err := vuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VipUser ID in the query.
// Returns a *NotSingularError when exactly one VipUser ID is not found.
// Returns a *NotFoundError when no entities are found.
func (vuq *VipUserQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vuq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = &NotSingularError{vipuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vuq *VipUserQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := vuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VipUsers.
func (vuq *VipUserQuery) All(ctx context.Context) ([]*VipUser, error) {
	if err := vuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vuq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vuq *VipUserQuery) AllX(ctx context.Context) []*VipUser {
	nodes, err := vuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VipUser IDs.
func (vuq *VipUserQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := vuq.Select(vipuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vuq *VipUserQuery) IDsX(ctx context.Context) []int64 {
	ids, err := vuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vuq *VipUserQuery) Count(ctx context.Context) (int, error) {
	if err := vuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vuq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vuq *VipUserQuery) CountX(ctx context.Context) int {
	count, err := vuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vuq *VipUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := vuq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vuq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vuq *VipUserQuery) ExistX(ctx context.Context) bool {
	exist, err := vuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VipUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vuq *VipUserQuery) Clone() *VipUserQuery {
	if vuq == nil {
		return nil
	}
	return &VipUserQuery{
		config:     vuq.config,
		limit:      vuq.limit,
		offset:     vuq.offset,
		order:      append([]OrderFunc{}, vuq.order...),
		predicates: append([]predicate.VipUser{}, vuq.predicates...),
		withUser:   vuq.withUser.Clone(),
		// clone intermediate query.
		sql:  vuq.sql.Clone(),
		path: vuq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (vuq *VipUserQuery) WithUser(opts ...func(*SocialUserQuery)) *VipUserQuery {
	query := &SocialUserQuery{config: vuq.config}
	for _, opt := range opts {
		opt(query)
	}
	vuq.withUser = query
	return vuq
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
//	client.VipUser.Query().
//		GroupBy(vipuser.FieldUserId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vuq *VipUserQuery) GroupBy(field string, fields ...string) *VipUserGroupBy {
	group := &VipUserGroupBy{config: vuq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vuq.sqlQuery(ctx), nil
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
//	client.VipUser.Query().
//		Select(vipuser.FieldUserId).
//		Scan(ctx, &v)
//
func (vuq *VipUserQuery) Select(fields ...string) *VipUserSelect {
	vuq.fields = append(vuq.fields, fields...)
	return &VipUserSelect{VipUserQuery: vuq}
}

func (vuq *VipUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vuq.fields {
		if !vipuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vuq.path != nil {
		prev, err := vuq.path(ctx)
		if err != nil {
			return err
		}
		vuq.sql = prev
	}
	return nil
}

func (vuq *VipUserQuery) sqlAll(ctx context.Context) ([]*VipUser, error) {
	var (
		nodes       = []*VipUser{}
		_spec       = vuq.querySpec()
		loadedTypes = [1]bool{
			vuq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &VipUser{config: vuq.config}
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
	if err := sqlgraph.QueryNodes(ctx, vuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := vuq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*VipUser)
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

func (vuq *VipUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vuq.querySpec()
	_spec.Node.Columns = vuq.fields
	if len(vuq.fields) > 0 {
		_spec.Unique = vuq.unique != nil && *vuq.unique
	}
	return sqlgraph.CountNodes(ctx, vuq.driver, _spec)
}

func (vuq *VipUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vuq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vuq *VipUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vipuser.Table,
			Columns: vipuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: vipuser.FieldID,
			},
		},
		From:   vuq.sql,
		Unique: true,
	}
	if unique := vuq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vuq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vipuser.FieldID)
		for i := range fields {
			if fields[i] != vipuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vuq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vuq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vuq *VipUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vuq.driver.Dialect())
	t1 := builder.Table(vipuser.Table)
	columns := vuq.fields
	if len(columns) == 0 {
		columns = vipuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vuq.sql != nil {
		selector = vuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vuq.unique != nil && *vuq.unique {
		selector.Distinct()
	}
	for _, p := range vuq.predicates {
		p(selector)
	}
	for _, p := range vuq.order {
		p(selector)
	}
	if offset := vuq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vuq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VipUserGroupBy is the group-by builder for VipUser entities.
type VipUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vugb *VipUserGroupBy) Aggregate(fns ...AggregateFunc) *VipUserGroupBy {
	vugb.fns = append(vugb.fns, fns...)
	return vugb
}

// Scan applies the group-by query and scans the result into the given value.
func (vugb *VipUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vugb.path(ctx)
	if err != nil {
		return err
	}
	vugb.sql = query
	return vugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vugb *VipUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := vugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(vugb.fields) > 1 {
		return nil, errors.New("ent: VipUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := vugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vugb *VipUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := vugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vugb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vugb *VipUserGroupBy) StringX(ctx context.Context) string {
	v, err := vugb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(vugb.fields) > 1 {
		return nil, errors.New("ent: VipUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := vugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vugb *VipUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := vugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vugb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vugb *VipUserGroupBy) IntX(ctx context.Context) int {
	v, err := vugb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(vugb.fields) > 1 {
		return nil, errors.New("ent: VipUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := vugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vugb *VipUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := vugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vugb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vugb *VipUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := vugb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(vugb.fields) > 1 {
		return nil, errors.New("ent: VipUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := vugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vugb *VipUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := vugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (vugb *VipUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vugb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vugb *VipUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := vugb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vugb *VipUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vugb.fields {
		if !vipuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vugb *VipUserGroupBy) sqlQuery() *sql.Selector {
	selector := vugb.sql.Select()
	aggregation := make([]string, 0, len(vugb.fns))
	for _, fn := range vugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vugb.fields)+len(vugb.fns))
		for _, f := range vugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vugb.fields...)...)
}

// VipUserSelect is the builder for selecting fields of VipUser entities.
type VipUserSelect struct {
	*VipUserQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vus *VipUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vus.prepareQuery(ctx); err != nil {
		return err
	}
	vus.sql = vus.VipUserQuery.sqlQuery(ctx)
	return vus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (vus *VipUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := vus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(vus.fields) > 1 {
		return nil, errors.New("ent: VipUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := vus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (vus *VipUserSelect) StringsX(ctx context.Context) []string {
	v, err := vus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = vus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (vus *VipUserSelect) StringX(ctx context.Context) string {
	v, err := vus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(vus.fields) > 1 {
		return nil, errors.New("ent: VipUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := vus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (vus *VipUserSelect) IntsX(ctx context.Context) []int {
	v, err := vus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = vus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (vus *VipUserSelect) IntX(ctx context.Context) int {
	v, err := vus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(vus.fields) > 1 {
		return nil, errors.New("ent: VipUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := vus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (vus *VipUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := vus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = vus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (vus *VipUserSelect) Float64X(ctx context.Context) float64 {
	v, err := vus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(vus.fields) > 1 {
		return nil, errors.New("ent: VipUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := vus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (vus *VipUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := vus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (vus *VipUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = vus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{vipuser.Label}
	default:
		err = fmt.Errorf("ent: VipUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (vus *VipUserSelect) BoolX(ctx context.Context) bool {
	v, err := vus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (vus *VipUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vus.sql.Query()
	if err := vus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
