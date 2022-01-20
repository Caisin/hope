// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"hope/apps/novel/internal/data/ent/novelcomment"
	"hope/apps/novel/internal/data/ent/predicate"
	"hope/apps/novel/internal/data/ent/socialuser"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NovelCommentQuery is the builder for querying NovelComment entities.
type NovelCommentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NovelComment
	// eager-loading edges.
	withParent  *NovelCommentQuery
	withChildes *NovelCommentQuery
	withUser    *SocialUserQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NovelCommentQuery builder.
func (ncq *NovelCommentQuery) Where(ps ...predicate.NovelComment) *NovelCommentQuery {
	ncq.predicates = append(ncq.predicates, ps...)
	return ncq
}

// Limit adds a limit step to the query.
func (ncq *NovelCommentQuery) Limit(limit int) *NovelCommentQuery {
	ncq.limit = &limit
	return ncq
}

// Offset adds an offset step to the query.
func (ncq *NovelCommentQuery) Offset(offset int) *NovelCommentQuery {
	ncq.offset = &offset
	return ncq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ncq *NovelCommentQuery) Unique(unique bool) *NovelCommentQuery {
	ncq.unique = &unique
	return ncq
}

// Order adds an order step to the query.
func (ncq *NovelCommentQuery) Order(o ...OrderFunc) *NovelCommentQuery {
	ncq.order = append(ncq.order, o...)
	return ncq
}

// QueryParent chains the current query on the "parent" edge.
func (ncq *NovelCommentQuery) QueryParent() *NovelCommentQuery {
	query := &NovelCommentQuery{config: ncq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ncq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novelcomment.Table, novelcomment.FieldID, selector),
			sqlgraph.To(novelcomment.Table, novelcomment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, novelcomment.ParentTable, novelcomment.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(ncq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildes chains the current query on the "childes" edge.
func (ncq *NovelCommentQuery) QueryChildes() *NovelCommentQuery {
	query := &NovelCommentQuery{config: ncq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ncq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novelcomment.Table, novelcomment.FieldID, selector),
			sqlgraph.To(novelcomment.Table, novelcomment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, novelcomment.ChildesTable, novelcomment.ChildesColumn),
		)
		fromU = sqlgraph.SetNeighbors(ncq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ncq *NovelCommentQuery) QueryUser() *SocialUserQuery {
	query := &SocialUserQuery{config: ncq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ncq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novelcomment.Table, novelcomment.FieldID, selector),
			sqlgraph.To(socialuser.Table, socialuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, novelcomment.UserTable, novelcomment.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ncq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NovelComment entity from the query.
// Returns a *NotFoundError when no NovelComment was found.
func (ncq *NovelCommentQuery) First(ctx context.Context) (*NovelComment, error) {
	nodes, err := ncq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{novelcomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ncq *NovelCommentQuery) FirstX(ctx context.Context) *NovelComment {
	node, err := ncq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NovelComment ID from the query.
// Returns a *NotFoundError when no NovelComment ID was found.
func (ncq *NovelCommentQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ncq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{novelcomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ncq *NovelCommentQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ncq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NovelComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NovelComment entity is not found.
// Returns a *NotFoundError when no NovelComment entities are found.
func (ncq *NovelCommentQuery) Only(ctx context.Context) (*NovelComment, error) {
	nodes, err := ncq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{novelcomment.Label}
	default:
		return nil, &NotSingularError{novelcomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ncq *NovelCommentQuery) OnlyX(ctx context.Context) *NovelComment {
	node, err := ncq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NovelComment ID in the query.
// Returns a *NotSingularError when exactly one NovelComment ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ncq *NovelCommentQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ncq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = &NotSingularError{novelcomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ncq *NovelCommentQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ncq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NovelComments.
func (ncq *NovelCommentQuery) All(ctx context.Context) ([]*NovelComment, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ncq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ncq *NovelCommentQuery) AllX(ctx context.Context) []*NovelComment {
	nodes, err := ncq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NovelComment IDs.
func (ncq *NovelCommentQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := ncq.Select(novelcomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ncq *NovelCommentQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ncq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ncq *NovelCommentQuery) Count(ctx context.Context) (int, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ncq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ncq *NovelCommentQuery) CountX(ctx context.Context) int {
	count, err := ncq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ncq *NovelCommentQuery) Exist(ctx context.Context) (bool, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ncq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ncq *NovelCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := ncq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NovelCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ncq *NovelCommentQuery) Clone() *NovelCommentQuery {
	if ncq == nil {
		return nil
	}
	return &NovelCommentQuery{
		config:      ncq.config,
		limit:       ncq.limit,
		offset:      ncq.offset,
		order:       append([]OrderFunc{}, ncq.order...),
		predicates:  append([]predicate.NovelComment{}, ncq.predicates...),
		withParent:  ncq.withParent.Clone(),
		withChildes: ncq.withChildes.Clone(),
		withUser:    ncq.withUser.Clone(),
		// clone intermediate query.
		sql:  ncq.sql.Clone(),
		path: ncq.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (ncq *NovelCommentQuery) WithParent(opts ...func(*NovelCommentQuery)) *NovelCommentQuery {
	query := &NovelCommentQuery{config: ncq.config}
	for _, opt := range opts {
		opt(query)
	}
	ncq.withParent = query
	return ncq
}

// WithChildes tells the query-builder to eager-load the nodes that are connected to
// the "childes" edge. The optional arguments are used to configure the query builder of the edge.
func (ncq *NovelCommentQuery) WithChildes(opts ...func(*NovelCommentQuery)) *NovelCommentQuery {
	query := &NovelCommentQuery{config: ncq.config}
	for _, opt := range opts {
		opt(query)
	}
	ncq.withChildes = query
	return ncq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ncq *NovelCommentQuery) WithUser(opts ...func(*SocialUserQuery)) *NovelCommentQuery {
	query := &SocialUserQuery{config: ncq.config}
	for _, opt := range opts {
		opt(query)
	}
	ncq.withUser = query
	return ncq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NovelId int64 `json:"novelId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NovelComment.Query().
//		GroupBy(novelcomment.FieldNovelId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ncq *NovelCommentQuery) GroupBy(field string, fields ...string) *NovelCommentGroupBy {
	group := &NovelCommentGroupBy{config: ncq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ncq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NovelId int64 `json:"novelId,omitempty"`
//	}
//
//	client.NovelComment.Query().
//		Select(novelcomment.FieldNovelId).
//		Scan(ctx, &v)
//
func (ncq *NovelCommentQuery) Select(fields ...string) *NovelCommentSelect {
	ncq.fields = append(ncq.fields, fields...)
	return &NovelCommentSelect{NovelCommentQuery: ncq}
}

func (ncq *NovelCommentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ncq.fields {
		if !novelcomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ncq.path != nil {
		prev, err := ncq.path(ctx)
		if err != nil {
			return err
		}
		ncq.sql = prev
	}
	return nil
}

func (ncq *NovelCommentQuery) sqlAll(ctx context.Context) ([]*NovelComment, error) {
	var (
		nodes       = []*NovelComment{}
		withFKs     = ncq.withFKs
		_spec       = ncq.querySpec()
		loadedTypes = [3]bool{
			ncq.withParent != nil,
			ncq.withChildes != nil,
			ncq.withUser != nil,
		}
	)
	if ncq.withParent != nil || ncq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, novelcomment.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NovelComment{config: ncq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ncq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ncq.withParent; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*NovelComment)
		for i := range nodes {
			if nodes[i].novel_comment_childes == nil {
				continue
			}
			fk := *nodes[i].novel_comment_childes
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(novelcomment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "novel_comment_childes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := ncq.withChildes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*NovelComment)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Childes = []*NovelComment{}
		}
		query.withFKs = true
		query.Where(predicate.NovelComment(func(s *sql.Selector) {
			s.Where(sql.InValues(novelcomment.ChildesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.novel_comment_childes
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "novel_comment_childes" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "novel_comment_childes" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Childes = append(node.Edges.Childes, n)
		}
	}

	if query := ncq.withUser; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*NovelComment)
		for i := range nodes {
			if nodes[i].social_user_comments == nil {
				continue
			}
			fk := *nodes[i].social_user_comments
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
				return nil, fmt.Errorf(`unexpected foreign-key "social_user_comments" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (ncq *NovelCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ncq.querySpec()
	return sqlgraph.CountNodes(ctx, ncq.driver, _spec)
}

func (ncq *NovelCommentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ncq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ncq *NovelCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelcomment.Table,
			Columns: novelcomment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: novelcomment.FieldID,
			},
		},
		From:   ncq.sql,
		Unique: true,
	}
	if unique := ncq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ncq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelcomment.FieldID)
		for i := range fields {
			if fields[i] != novelcomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ncq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ncq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ncq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ncq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ncq *NovelCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ncq.driver.Dialect())
	t1 := builder.Table(novelcomment.Table)
	columns := ncq.fields
	if len(columns) == 0 {
		columns = novelcomment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ncq.sql != nil {
		selector = ncq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ncq.predicates {
		p(selector)
	}
	for _, p := range ncq.order {
		p(selector)
	}
	if offset := ncq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ncq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NovelCommentGroupBy is the group-by builder for NovelComment entities.
type NovelCommentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ncgb *NovelCommentGroupBy) Aggregate(fns ...AggregateFunc) *NovelCommentGroupBy {
	ncgb.fns = append(ncgb.fns, fns...)
	return ncgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ncgb *NovelCommentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ncgb.path(ctx)
	if err != nil {
		return err
	}
	ncgb.sql = query
	return ncgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ncgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelCommentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) StringsX(ctx context.Context) []string {
	v, err := ncgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ncgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) StringX(ctx context.Context) string {
	v, err := ncgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelCommentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) IntsX(ctx context.Context) []int {
	v, err := ncgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ncgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) IntX(ctx context.Context) int {
	v, err := ncgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelCommentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ncgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ncgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ncgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelCommentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ncgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelCommentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ncgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ncgb *NovelCommentGroupBy) BoolX(ctx context.Context) bool {
	v, err := ncgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ncgb *NovelCommentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ncgb.fields {
		if !novelcomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ncgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ncgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ncgb *NovelCommentGroupBy) sqlQuery() *sql.Selector {
	selector := ncgb.sql.Select()
	aggregation := make([]string, 0, len(ncgb.fns))
	for _, fn := range ncgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ncgb.fields)+len(ncgb.fns))
		for _, f := range ncgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ncgb.fields...)...)
}

// NovelCommentSelect is the builder for selecting fields of NovelComment entities.
type NovelCommentSelect struct {
	*NovelCommentQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ncs *NovelCommentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ncs.prepareQuery(ctx); err != nil {
		return err
	}
	ncs.sql = ncs.NovelCommentQuery.sqlQuery(ctx)
	return ncs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ncs *NovelCommentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ncs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelCommentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ncs *NovelCommentSelect) StringsX(ctx context.Context) []string {
	v, err := ncs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ncs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ncs *NovelCommentSelect) StringX(ctx context.Context) string {
	v, err := ncs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelCommentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ncs *NovelCommentSelect) IntsX(ctx context.Context) []int {
	v, err := ncs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ncs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ncs *NovelCommentSelect) IntX(ctx context.Context) int {
	v, err := ncs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelCommentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ncs *NovelCommentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ncs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ncs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ncs *NovelCommentSelect) Float64X(ctx context.Context) float64 {
	v, err := ncs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelCommentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ncs *NovelCommentSelect) BoolsX(ctx context.Context) []bool {
	v, err := ncs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ncs *NovelCommentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ncs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelcomment.Label}
	default:
		err = fmt.Errorf("ent: NovelCommentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ncs *NovelCommentSelect) BoolX(ctx context.Context) bool {
	v, err := ncs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ncs *NovelCommentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ncs.sql.Query()
	if err := ncs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
