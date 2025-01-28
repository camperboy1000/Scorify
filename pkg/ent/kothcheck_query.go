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
	"github.com/scorify/scorify/pkg/ent/kothcheck"
	"github.com/scorify/scorify/pkg/ent/kothstatus"
	"github.com/scorify/scorify/pkg/ent/predicate"
)

// KothCheckQuery is the builder for querying KothCheck entities.
type KothCheckQuery struct {
	config
	ctx          *QueryContext
	order        []kothcheck.OrderOption
	inters       []Interceptor
	predicates   []predicate.KothCheck
	withStatuses *KothStatusQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KothCheckQuery builder.
func (kcq *KothCheckQuery) Where(ps ...predicate.KothCheck) *KothCheckQuery {
	kcq.predicates = append(kcq.predicates, ps...)
	return kcq
}

// Limit the number of records to be returned by this query.
func (kcq *KothCheckQuery) Limit(limit int) *KothCheckQuery {
	kcq.ctx.Limit = &limit
	return kcq
}

// Offset to start from.
func (kcq *KothCheckQuery) Offset(offset int) *KothCheckQuery {
	kcq.ctx.Offset = &offset
	return kcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kcq *KothCheckQuery) Unique(unique bool) *KothCheckQuery {
	kcq.ctx.Unique = &unique
	return kcq
}

// Order specifies how the records should be ordered.
func (kcq *KothCheckQuery) Order(o ...kothcheck.OrderOption) *KothCheckQuery {
	kcq.order = append(kcq.order, o...)
	return kcq
}

// QueryStatuses chains the current query on the "statuses" edge.
func (kcq *KothCheckQuery) QueryStatuses() *KothStatusQuery {
	query := (&KothStatusClient{config: kcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := kcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := kcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(kothcheck.Table, kothcheck.FieldID, selector),
			sqlgraph.To(kothstatus.Table, kothstatus.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, kothcheck.StatusesTable, kothcheck.StatusesColumn),
		)
		fromU = sqlgraph.SetNeighbors(kcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first KothCheck entity from the query.
// Returns a *NotFoundError when no KothCheck was found.
func (kcq *KothCheckQuery) First(ctx context.Context) (*KothCheck, error) {
	nodes, err := kcq.Limit(1).All(setContextOp(ctx, kcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{kothcheck.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kcq *KothCheckQuery) FirstX(ctx context.Context) *KothCheck {
	node, err := kcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first KothCheck ID from the query.
// Returns a *NotFoundError when no KothCheck ID was found.
func (kcq *KothCheckQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = kcq.Limit(1).IDs(setContextOp(ctx, kcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{kothcheck.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kcq *KothCheckQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := kcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single KothCheck entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one KothCheck entity is found.
// Returns a *NotFoundError when no KothCheck entities are found.
func (kcq *KothCheckQuery) Only(ctx context.Context) (*KothCheck, error) {
	nodes, err := kcq.Limit(2).All(setContextOp(ctx, kcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{kothcheck.Label}
	default:
		return nil, &NotSingularError{kothcheck.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kcq *KothCheckQuery) OnlyX(ctx context.Context) *KothCheck {
	node, err := kcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only KothCheck ID in the query.
// Returns a *NotSingularError when more than one KothCheck ID is found.
// Returns a *NotFoundError when no entities are found.
func (kcq *KothCheckQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = kcq.Limit(2).IDs(setContextOp(ctx, kcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{kothcheck.Label}
	default:
		err = &NotSingularError{kothcheck.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kcq *KothCheckQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := kcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KothChecks.
func (kcq *KothCheckQuery) All(ctx context.Context) ([]*KothCheck, error) {
	ctx = setContextOp(ctx, kcq.ctx, "All")
	if err := kcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*KothCheck, *KothCheckQuery]()
	return withInterceptors[[]*KothCheck](ctx, kcq, qr, kcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (kcq *KothCheckQuery) AllX(ctx context.Context) []*KothCheck {
	nodes, err := kcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of KothCheck IDs.
func (kcq *KothCheckQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if kcq.ctx.Unique == nil && kcq.path != nil {
		kcq.Unique(true)
	}
	ctx = setContextOp(ctx, kcq.ctx, "IDs")
	if err = kcq.Select(kothcheck.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kcq *KothCheckQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := kcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kcq *KothCheckQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, kcq.ctx, "Count")
	if err := kcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, kcq, querierCount[*KothCheckQuery](), kcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (kcq *KothCheckQuery) CountX(ctx context.Context) int {
	count, err := kcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kcq *KothCheckQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, kcq.ctx, "Exist")
	switch _, err := kcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (kcq *KothCheckQuery) ExistX(ctx context.Context) bool {
	exist, err := kcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KothCheckQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kcq *KothCheckQuery) Clone() *KothCheckQuery {
	if kcq == nil {
		return nil
	}
	return &KothCheckQuery{
		config:       kcq.config,
		ctx:          kcq.ctx.Clone(),
		order:        append([]kothcheck.OrderOption{}, kcq.order...),
		inters:       append([]Interceptor{}, kcq.inters...),
		predicates:   append([]predicate.KothCheck{}, kcq.predicates...),
		withStatuses: kcq.withStatuses.Clone(),
		// clone intermediate query.
		sql:  kcq.sql.Clone(),
		path: kcq.path,
	}
}

// WithStatuses tells the query-builder to eager-load the nodes that are connected to
// the "statuses" edge. The optional arguments are used to configure the query builder of the edge.
func (kcq *KothCheckQuery) WithStatuses(opts ...func(*KothStatusQuery)) *KothCheckQuery {
	query := (&KothStatusClient{config: kcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	kcq.withStatuses = query
	return kcq
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
//	client.KothCheck.Query().
//		GroupBy(kothcheck.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (kcq *KothCheckQuery) GroupBy(field string, fields ...string) *KothCheckGroupBy {
	kcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &KothCheckGroupBy{build: kcq}
	grbuild.flds = &kcq.ctx.Fields
	grbuild.label = kothcheck.Label
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
//	client.KothCheck.Query().
//		Select(kothcheck.FieldCreateTime).
//		Scan(ctx, &v)
func (kcq *KothCheckQuery) Select(fields ...string) *KothCheckSelect {
	kcq.ctx.Fields = append(kcq.ctx.Fields, fields...)
	sbuild := &KothCheckSelect{KothCheckQuery: kcq}
	sbuild.label = kothcheck.Label
	sbuild.flds, sbuild.scan = &kcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a KothCheckSelect configured with the given aggregations.
func (kcq *KothCheckQuery) Aggregate(fns ...AggregateFunc) *KothCheckSelect {
	return kcq.Select().Aggregate(fns...)
}

func (kcq *KothCheckQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range kcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, kcq); err != nil {
				return err
			}
		}
	}
	for _, f := range kcq.ctx.Fields {
		if !kothcheck.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kcq.path != nil {
		prev, err := kcq.path(ctx)
		if err != nil {
			return err
		}
		kcq.sql = prev
	}
	return nil
}

func (kcq *KothCheckQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*KothCheck, error) {
	var (
		nodes       = []*KothCheck{}
		_spec       = kcq.querySpec()
		loadedTypes = [1]bool{
			kcq.withStatuses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*KothCheck).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &KothCheck{config: kcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := kcq.withStatuses; query != nil {
		if err := kcq.loadStatuses(ctx, query, nodes,
			func(n *KothCheck) { n.Edges.Statuses = []*KothStatus{} },
			func(n *KothCheck, e *KothStatus) { n.Edges.Statuses = append(n.Edges.Statuses, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (kcq *KothCheckQuery) loadStatuses(ctx context.Context, query *KothStatusQuery, nodes []*KothCheck, init func(*KothCheck), assign func(*KothCheck, *KothStatus)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*KothCheck)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(kothstatus.FieldCheckID)
	}
	query.Where(predicate.KothStatus(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(kothcheck.StatusesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CheckID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "check_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (kcq *KothCheckQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kcq.querySpec()
	_spec.Node.Columns = kcq.ctx.Fields
	if len(kcq.ctx.Fields) > 0 {
		_spec.Unique = kcq.ctx.Unique != nil && *kcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, kcq.driver, _spec)
}

func (kcq *KothCheckQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(kothcheck.Table, kothcheck.Columns, sqlgraph.NewFieldSpec(kothcheck.FieldID, field.TypeUUID))
	_spec.From = kcq.sql
	if unique := kcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if kcq.path != nil {
		_spec.Unique = true
	}
	if fields := kcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kothcheck.FieldID)
		for i := range fields {
			if fields[i] != kothcheck.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kcq *KothCheckQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kcq.driver.Dialect())
	t1 := builder.Table(kothcheck.Table)
	columns := kcq.ctx.Fields
	if len(columns) == 0 {
		columns = kothcheck.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kcq.sql != nil {
		selector = kcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kcq.ctx.Unique != nil && *kcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range kcq.predicates {
		p(selector)
	}
	for _, p := range kcq.order {
		p(selector)
	}
	if offset := kcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KothCheckGroupBy is the group-by builder for KothCheck entities.
type KothCheckGroupBy struct {
	selector
	build *KothCheckQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kcgb *KothCheckGroupBy) Aggregate(fns ...AggregateFunc) *KothCheckGroupBy {
	kcgb.fns = append(kcgb.fns, fns...)
	return kcgb
}

// Scan applies the selector query and scans the result into the given value.
func (kcgb *KothCheckGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, kcgb.build.ctx, "GroupBy")
	if err := kcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KothCheckQuery, *KothCheckGroupBy](ctx, kcgb.build, kcgb, kcgb.build.inters, v)
}

func (kcgb *KothCheckGroupBy) sqlScan(ctx context.Context, root *KothCheckQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(kcgb.fns))
	for _, fn := range kcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*kcgb.flds)+len(kcgb.fns))
		for _, f := range *kcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*kcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// KothCheckSelect is the builder for selecting fields of KothCheck entities.
type KothCheckSelect struct {
	*KothCheckQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (kcs *KothCheckSelect) Aggregate(fns ...AggregateFunc) *KothCheckSelect {
	kcs.fns = append(kcs.fns, fns...)
	return kcs
}

// Scan applies the selector query and scans the result into the given value.
func (kcs *KothCheckSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, kcs.ctx, "Select")
	if err := kcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KothCheckQuery, *KothCheckSelect](ctx, kcs.KothCheckQuery, kcs, kcs.inters, v)
}

func (kcs *KothCheckSelect) sqlScan(ctx context.Context, root *KothCheckQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(kcs.fns))
	for _, fn := range kcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*kcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
