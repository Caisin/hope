// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"hope/apps/admin/internal/data/ent/predicate"
	"hope/apps/admin/internal/data/ent/sysmenu"
	"hope/apps/admin/internal/data/ent/sysrole"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SysMenuQuery is the builder for querying SysMenu entities.
type SysMenuQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysMenu
	// eager-loading edges.
	withRole    *SysRoleQuery
	withParent  *SysMenuQuery
	withChildes *SysMenuQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysMenuQuery builder.
func (smq *SysMenuQuery) Where(ps ...predicate.SysMenu) *SysMenuQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit adds a limit step to the query.
func (smq *SysMenuQuery) Limit(limit int) *SysMenuQuery {
	smq.limit = &limit
	return smq
}

// Offset adds an offset step to the query.
func (smq *SysMenuQuery) Offset(offset int) *SysMenuQuery {
	smq.offset = &offset
	return smq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smq *SysMenuQuery) Unique(unique bool) *SysMenuQuery {
	smq.unique = &unique
	return smq
}

// Order adds an order step to the query.
func (smq *SysMenuQuery) Order(o ...OrderFunc) *SysMenuQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// QueryRole chains the current query on the "role" edge.
func (smq *SysMenuQuery) QueryRole() *SysRoleQuery {
	query := &SysRoleQuery{config: smq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysrole.Table, sysrole.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, sysmenu.RoleTable, sysmenu.RolePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (smq *SysMenuQuery) QueryParent() *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysmenu.Table, sysmenu.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, sysmenu.ParentTable, sysmenu.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildes chains the current query on the "childes" edge.
func (smq *SysMenuQuery) QueryChildes() *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenu.Table, sysmenu.FieldID, selector),
			sqlgraph.To(sysmenu.Table, sysmenu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sysmenu.ChildesTable, sysmenu.ChildesColumn),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysMenu entity from the query.
// Returns a *NotFoundError when no SysMenu was found.
func (smq *SysMenuQuery) First(ctx context.Context) (*SysMenu, error) {
	nodes, err := smq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysmenu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *SysMenuQuery) FirstX(ctx context.Context) *SysMenu {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysMenu ID from the query.
// Returns a *NotFoundError when no SysMenu ID was found.
func (smq *SysMenuQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = smq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysmenu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *SysMenuQuery) FirstIDX(ctx context.Context) int64 {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysMenu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SysMenu entity is not found.
// Returns a *NotFoundError when no SysMenu entities are found.
func (smq *SysMenuQuery) Only(ctx context.Context) (*SysMenu, error) {
	nodes, err := smq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysmenu.Label}
	default:
		return nil, &NotSingularError{sysmenu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *SysMenuQuery) OnlyX(ctx context.Context) *SysMenu {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysMenu ID in the query.
// Returns a *NotSingularError when exactly one SysMenu ID is not found.
// Returns a *NotFoundError when no entities are found.
func (smq *SysMenuQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = smq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = &NotSingularError{sysmenu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *SysMenuQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysMenus.
func (smq *SysMenuQuery) All(ctx context.Context) ([]*SysMenu, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return smq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (smq *SysMenuQuery) AllX(ctx context.Context) []*SysMenu {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysMenu IDs.
func (smq *SysMenuQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := smq.Select(sysmenu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *SysMenuQuery) IDsX(ctx context.Context) []int64 {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *SysMenuQuery) Count(ctx context.Context) (int, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return smq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (smq *SysMenuQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *SysMenuQuery) Exist(ctx context.Context) (bool, error) {
	if err := smq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return smq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *SysMenuQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysMenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *SysMenuQuery) Clone() *SysMenuQuery {
	if smq == nil {
		return nil
	}
	return &SysMenuQuery{
		config:      smq.config,
		limit:       smq.limit,
		offset:      smq.offset,
		order:       append([]OrderFunc{}, smq.order...),
		predicates:  append([]predicate.SysMenu{}, smq.predicates...),
		withRole:    smq.withRole.Clone(),
		withParent:  smq.withParent.Clone(),
		withChildes: smq.withChildes.Clone(),
		// clone intermediate query.
		sql:  smq.sql.Clone(),
		path: smq.path,
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithRole(opts ...func(*SysRoleQuery)) *SysMenuQuery {
	query := &SysRoleQuery{config: smq.config}
	for _, opt := range opts {
		opt(query)
	}
	smq.withRole = query
	return smq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithParent(opts ...func(*SysMenuQuery)) *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	for _, opt := range opts {
		opt(query)
	}
	smq.withParent = query
	return smq
}

// WithChildes tells the query-builder to eager-load the nodes that are connected to
// the "childes" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SysMenuQuery) WithChildes(opts ...func(*SysMenuQuery)) *SysMenuQuery {
	query := &SysMenuQuery{config: smq.config}
	for _, opt := range opts {
		opt(query)
	}
	smq.withChildes = query
	return smq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MenuName string `json:"menuName,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysMenu.Query().
//		GroupBy(sysmenu.FieldMenuName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (smq *SysMenuQuery) GroupBy(field string, fields ...string) *SysMenuGroupBy {
	group := &SysMenuGroupBy{config: smq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return smq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MenuName string `json:"menuName,omitempty"`
//	}
//
//	client.SysMenu.Query().
//		Select(sysmenu.FieldMenuName).
//		Scan(ctx, &v)
//
func (smq *SysMenuQuery) Select(fields ...string) *SysMenuSelect {
	smq.fields = append(smq.fields, fields...)
	return &SysMenuSelect{SysMenuQuery: smq}
}

func (smq *SysMenuQuery) prepareQuery(ctx context.Context) error {
	for _, f := range smq.fields {
		if !sysmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if smq.path != nil {
		prev, err := smq.path(ctx)
		if err != nil {
			return err
		}
		smq.sql = prev
	}
	return nil
}

func (smq *SysMenuQuery) sqlAll(ctx context.Context) ([]*SysMenu, error) {
	var (
		nodes       = []*SysMenu{}
		withFKs     = smq.withFKs
		_spec       = smq.querySpec()
		loadedTypes = [3]bool{
			smq.withRole != nil,
			smq.withParent != nil,
			smq.withChildes != nil,
		}
	)
	if smq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenu.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysMenu{config: smq.config}
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
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := smq.withRole; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*SysMenu, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Role = []*SysRole{}
		}
		var (
			edgeids []int64
			edges   = make(map[int64][]*SysMenu)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   sysmenu.RoleTable,
				Columns: sysmenu.RolePrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(sysmenu.RolePrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.Int64
				inValue := ein.Int64
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, smq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "role": %w`, err)
		}
		query.Where(sysrole.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "role" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Role = append(nodes[i].Edges.Role, n)
			}
		}
	}

	if query := smq.withParent; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*SysMenu)
		for i := range nodes {
			if nodes[i].sys_menu_childes == nil {
				continue
			}
			fk := *nodes[i].sys_menu_childes
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(sysmenu.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "sys_menu_childes" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent = n
			}
		}
	}

	if query := smq.withChildes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*SysMenu)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Childes = []*SysMenu{}
		}
		query.withFKs = true
		query.Where(predicate.SysMenu(func(s *sql.Selector) {
			s.Where(sql.InValues(sysmenu.ChildesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.sys_menu_childes
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "sys_menu_childes" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "sys_menu_childes" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Childes = append(node.Edges.Childes, n)
		}
	}

	return nodes, nil
}

func (smq *SysMenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *SysMenuQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := smq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (smq *SysMenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenu.Table,
			Columns: sysmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: sysmenu.FieldID,
			},
		},
		From:   smq.sql,
		Unique: true,
	}
	if unique := smq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := smq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenu.FieldID)
		for i := range fields {
			if fields[i] != sysmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smq *SysMenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(sysmenu.Table)
	columns := smq.fields
	if len(columns) == 0 {
		columns = sysmenu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector)
	}
	if offset := smq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysMenuGroupBy is the group-by builder for SysMenu entities.
type SysMenuGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *SysMenuGroupBy) Aggregate(fns ...AggregateFunc) *SysMenuGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the group-by query and scans the result into the given value.
func (smgb *SysMenuGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := smgb.path(ctx)
	if err != nil {
		return err
	}
	smgb.sql = query
	return smgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (smgb *SysMenuGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := smgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: SysMenuGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (smgb *SysMenuGroupBy) StringsX(ctx context.Context) []string {
	v, err := smgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = smgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (smgb *SysMenuGroupBy) StringX(ctx context.Context) string {
	v, err := smgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: SysMenuGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (smgb *SysMenuGroupBy) IntsX(ctx context.Context) []int {
	v, err := smgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = smgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (smgb *SysMenuGroupBy) IntX(ctx context.Context) int {
	v, err := smgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: SysMenuGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (smgb *SysMenuGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := smgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = smgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (smgb *SysMenuGroupBy) Float64X(ctx context.Context) float64 {
	v, err := smgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(smgb.fields) > 1 {
		return nil, errors.New("ent: SysMenuGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := smgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (smgb *SysMenuGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := smgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smgb *SysMenuGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = smgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (smgb *SysMenuGroupBy) BoolX(ctx context.Context) bool {
	v, err := smgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (smgb *SysMenuGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range smgb.fields {
		if !sysmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := smgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (smgb *SysMenuGroupBy) sqlQuery() *sql.Selector {
	selector := smgb.sql.Select()
	aggregation := make([]string, 0, len(smgb.fns))
	for _, fn := range smgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(smgb.fields)+len(smgb.fns))
		for _, f := range smgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(smgb.fields...)...)
}

// SysMenuSelect is the builder for selecting fields of SysMenu entities.
type SysMenuSelect struct {
	*SysMenuQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sms *SysMenuSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	sms.sql = sms.SysMenuQuery.sqlQuery(ctx)
	return sms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sms *SysMenuSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: SysMenuSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sms *SysMenuSelect) StringsX(ctx context.Context) []string {
	v, err := sms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sms *SysMenuSelect) StringX(ctx context.Context) string {
	v, err := sms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: SysMenuSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sms *SysMenuSelect) IntsX(ctx context.Context) []int {
	v, err := sms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sms *SysMenuSelect) IntX(ctx context.Context) int {
	v, err := sms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: SysMenuSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sms *SysMenuSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sms *SysMenuSelect) Float64X(ctx context.Context) float64 {
	v, err := sms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sms.fields) > 1 {
		return nil, errors.New("ent: SysMenuSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sms *SysMenuSelect) BoolsX(ctx context.Context) []bool {
	v, err := sms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sms *SysMenuSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenu.Label}
	default:
		err = fmt.Errorf("ent: SysMenuSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sms *SysMenuSelect) BoolX(ctx context.Context) bool {
	v, err := sms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sms *SysMenuSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sms.sql.Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
