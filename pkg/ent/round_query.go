// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/kothstatus"
	"github.com/scorify/scorify/pkg/ent/predicate"
	"github.com/scorify/scorify/pkg/ent/round"
	"github.com/scorify/scorify/pkg/ent/scorecache"
	"github.com/scorify/scorify/pkg/ent/status"
)

// RoundQuery is the builder for querying Round entities.
type RoundQuery struct {
	config
	ctx              *QueryContext
	order            []round.OrderOption
	inters           []Interceptor
	predicates       []predicate.Round
	withStatuses     *StatusQuery
	withScoreCaches  *ScoreCacheQuery
	withKothStatuses *KothStatusQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoundQuery builder.
func (rq *RoundQuery) Where(ps ...predicate.Round) *RoundQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RoundQuery) Limit(limit int) *RoundQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RoundQuery) Offset(offset int) *RoundQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RoundQuery) Unique(unique bool) *RoundQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RoundQuery) Order(o ...round.OrderOption) *RoundQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryStatuses chains the current query on the "statuses" edge.
func (rq *RoundQuery) QueryStatuses() *StatusQuery {
	query := (&StatusClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(round.Table, round.FieldID, selector),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, round.StatusesTable, round.StatusesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScoreCaches chains the current query on the "scoreCaches" edge.
func (rq *RoundQuery) QueryScoreCaches() *ScoreCacheQuery {
	query := (&ScoreCacheClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(round.Table, round.FieldID, selector),
			sqlgraph.To(scorecache.Table, scorecache.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, round.ScoreCachesTable, round.ScoreCachesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryKothStatuses chains the current query on the "kothStatuses" edge.
func (rq *RoundQuery) QueryKothStatuses() *KothStatusQuery {
	query := (&KothStatusClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(round.Table, round.FieldID, selector),
			sqlgraph.To(kothstatus.Table, kothstatus.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, round.KothStatusesTable, round.KothStatusesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Round entity from the query.
// Returns a *NotFoundError when no Round was found.
func (rq *RoundQuery) First(ctx context.Context) (*Round, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{round.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RoundQuery) FirstX(ctx context.Context) *Round {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Round ID from the query.
// Returns a *NotFoundError when no Round ID was found.
func (rq *RoundQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{round.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RoundQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Round entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Round entity is found.
// Returns a *NotFoundError when no Round entities are found.
func (rq *RoundQuery) Only(ctx context.Context) (*Round, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{round.Label}
	default:
		return nil, &NotSingularError{round.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RoundQuery) OnlyX(ctx context.Context) *Round {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Round ID in the query.
// Returns a *NotSingularError when more than one Round ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RoundQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{round.Label}
	default:
		err = &NotSingularError{round.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RoundQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Rounds.
func (rq *RoundQuery) All(ctx context.Context) ([]*Round, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Round, *RoundQuery]()
	return withInterceptors[[]*Round](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RoundQuery) AllX(ctx context.Context) []*Round {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Round IDs.
func (rq *RoundQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(round.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RoundQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RoundQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RoundQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RoundQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RoundQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RoundQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoundQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RoundQuery) Clone() *RoundQuery {
	if rq == nil {
		return nil
	}
	return &RoundQuery{
		config:           rq.config,
		ctx:              rq.ctx.Clone(),
		order:            append([]round.OrderOption{}, rq.order...),
		inters:           append([]Interceptor{}, rq.inters...),
		predicates:       append([]predicate.Round{}, rq.predicates...),
		withStatuses:     rq.withStatuses.Clone(),
		withScoreCaches:  rq.withScoreCaches.Clone(),
		withKothStatuses: rq.withKothStatuses.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithStatuses tells the query-builder to eager-load the nodes that are connected to
// the "statuses" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoundQuery) WithStatuses(opts ...func(*StatusQuery)) *RoundQuery {
	query := (&StatusClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withStatuses = query
	return rq
}

// WithScoreCaches tells the query-builder to eager-load the nodes that are connected to
// the "scoreCaches" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoundQuery) WithScoreCaches(opts ...func(*ScoreCacheQuery)) *RoundQuery {
	query := (&ScoreCacheClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withScoreCaches = query
	return rq
}

// WithKothStatuses tells the query-builder to eager-load the nodes that are connected to
// the "kothStatuses" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RoundQuery) WithKothStatuses(opts ...func(*KothStatusQuery)) *RoundQuery {
	query := (&KothStatusClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withKothStatuses = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Round.Query().
//		GroupBy(round.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RoundQuery) GroupBy(field string, fields ...string) *RoundGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoundGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = round.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Round.Query().
//		Select(round.FieldCreateTime).
//		Scan(ctx, &v)
func (rq *RoundQuery) Select(fields ...string) *RoundSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RoundSelect{RoundQuery: rq}
	sbuild.label = round.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoundSelect configured with the given aggregations.
func (rq *RoundQuery) Aggregate(fns ...AggregateFunc) *RoundSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RoundQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !round.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RoundQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Round, error) {
	var (
		nodes       = []*Round{}
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withStatuses != nil,
			rq.withScoreCaches != nil,
			rq.withKothStatuses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Round).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Round{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withStatuses; query != nil {
		if err := rq.loadStatuses(ctx, query, nodes,
			func(n *Round) { n.Edges.Statuses = []*Status{} },
			func(n *Round, e *Status) { n.Edges.Statuses = append(n.Edges.Statuses, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withScoreCaches; query != nil {
		if err := rq.loadScoreCaches(ctx, query, nodes,
			func(n *Round) { n.Edges.ScoreCaches = []*ScoreCache{} },
			func(n *Round, e *ScoreCache) { n.Edges.ScoreCaches = append(n.Edges.ScoreCaches, e) }); err != nil {
			return nil, err
		}
	}
	if query := rq.withKothStatuses; query != nil {
		if err := rq.loadKothStatuses(ctx, query, nodes,
			func(n *Round) { n.Edges.KothStatuses = []*KothStatus{} },
			func(n *Round, e *KothStatus) { n.Edges.KothStatuses = append(n.Edges.KothStatuses, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RoundQuery) loadStatuses(ctx context.Context, query *StatusQuery, nodes []*Round, init func(*Round), assign func(*Round, *Status)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Round)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(status.FieldRoundID)
	}
	query.Where(predicate.Status(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(round.StatusesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoundID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "round_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *RoundQuery) loadScoreCaches(ctx context.Context, query *ScoreCacheQuery, nodes []*Round, init func(*Round), assign func(*Round, *ScoreCache)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Round)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(scorecache.FieldRoundID)
	}
	query.Where(predicate.ScoreCache(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(round.ScoreCachesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoundID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "round_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (rq *RoundQuery) loadKothStatuses(ctx context.Context, query *KothStatusQuery, nodes []*Round, init func(*Round), assign func(*Round, *KothStatus)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Round)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(kothstatus.FieldRoundID)
	}
	query.Where(predicate.KothStatus(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(round.KothStatusesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.RoundID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "round_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *RoundQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RoundQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(round.Table, round.Columns, sqlgraph.NewFieldSpec(round.FieldID, field.TypeUUID))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, round.FieldID)
		for i := range fields {
			if fields[i] != round.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RoundQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(round.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = round.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoundGroupBy is the group-by builder for Round entities.
type RoundGroupBy struct {
	selector
	build *RoundQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RoundGroupBy) Aggregate(fns ...AggregateFunc) *RoundGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RoundGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoundQuery, *RoundGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RoundGroupBy) sqlScan(ctx context.Context, root *RoundQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoundSelect is the builder for selecting fields of Round entities.
type RoundSelect struct {
	*RoundQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RoundSelect) Aggregate(fns ...AggregateFunc) *RoundSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RoundSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoundQuery, *RoundSelect](ctx, rs.RoundQuery, rs, rs.inters, v)
}

func (rs *RoundSelect) sqlScan(ctx context.Context, root *RoundQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
