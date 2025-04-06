// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"gojeksrepo/ent/driverprofile"
	"gojeksrepo/ent/payment"
	"gojeksrepo/ent/predicate"
	"gojeksrepo/ent/trip"
	"gojeksrepo/ent/triprating"
	"gojeksrepo/ent/user"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TripQuery is the builder for querying Trip entities.
type TripQuery struct {
	config
	ctx         *QueryContext
	order       []trip.OrderOption
	inters      []Interceptor
	predicates  []predicate.Trip
	withUser    *UserQuery
	withDriver  *DriverProfileQuery
	withPayment *PaymentQuery
	withRatings *TripRatingQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TripQuery builder.
func (tq *TripQuery) Where(ps ...predicate.Trip) *TripQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TripQuery) Limit(limit int) *TripQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TripQuery) Offset(offset int) *TripQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TripQuery) Unique(unique bool) *TripQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TripQuery) Order(o ...trip.OrderOption) *TripQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryUser chains the current query on the "user" edge.
func (tq *TripQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trip.Table, trip.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, trip.UserTable, trip.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDriver chains the current query on the "driver" edge.
func (tq *TripQuery) QueryDriver() *DriverProfileQuery {
	query := (&DriverProfileClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trip.Table, trip.FieldID, selector),
			sqlgraph.To(driverprofile.Table, driverprofile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, trip.DriverTable, trip.DriverColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPayment chains the current query on the "payment" edge.
func (tq *TripQuery) QueryPayment() *PaymentQuery {
	query := (&PaymentClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trip.Table, trip.FieldID, selector),
			sqlgraph.To(payment.Table, payment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, trip.PaymentTable, trip.PaymentColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRatings chains the current query on the "ratings" edge.
func (tq *TripQuery) QueryRatings() *TripRatingQuery {
	query := (&TripRatingClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(trip.Table, trip.FieldID, selector),
			sqlgraph.To(triprating.Table, triprating.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, trip.RatingsTable, trip.RatingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Trip entity from the query.
// Returns a *NotFoundError when no Trip was found.
func (tq *TripQuery) First(ctx context.Context) (*Trip, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{trip.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TripQuery) FirstX(ctx context.Context) *Trip {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Trip ID from the query.
// Returns a *NotFoundError when no Trip ID was found.
func (tq *TripQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{trip.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TripQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Trip entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Trip entity is found.
// Returns a *NotFoundError when no Trip entities are found.
func (tq *TripQuery) Only(ctx context.Context) (*Trip, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{trip.Label}
	default:
		return nil, &NotSingularError{trip.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TripQuery) OnlyX(ctx context.Context) *Trip {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Trip ID in the query.
// Returns a *NotSingularError when more than one Trip ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TripQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{trip.Label}
	default:
		err = &NotSingularError{trip.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TripQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Trips.
func (tq *TripQuery) All(ctx context.Context) ([]*Trip, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Trip, *TripQuery]()
	return withInterceptors[[]*Trip](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TripQuery) AllX(ctx context.Context) []*Trip {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Trip IDs.
func (tq *TripQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryIDs)
	if err = tq.Select(trip.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TripQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TripQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TripQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TripQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TripQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryExist)
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TripQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TripQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TripQuery) Clone() *TripQuery {
	if tq == nil {
		return nil
	}
	return &TripQuery{
		config:      tq.config,
		ctx:         tq.ctx.Clone(),
		order:       append([]trip.OrderOption{}, tq.order...),
		inters:      append([]Interceptor{}, tq.inters...),
		predicates:  append([]predicate.Trip{}, tq.predicates...),
		withUser:    tq.withUser.Clone(),
		withDriver:  tq.withDriver.Clone(),
		withPayment: tq.withPayment.Clone(),
		withRatings: tq.withRatings.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TripQuery) WithUser(opts ...func(*UserQuery)) *TripQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUser = query
	return tq
}

// WithDriver tells the query-builder to eager-load the nodes that are connected to
// the "driver" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TripQuery) WithDriver(opts ...func(*DriverProfileQuery)) *TripQuery {
	query := (&DriverProfileClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withDriver = query
	return tq
}

// WithPayment tells the query-builder to eager-load the nodes that are connected to
// the "payment" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TripQuery) WithPayment(opts ...func(*PaymentQuery)) *TripQuery {
	query := (&PaymentClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withPayment = query
	return tq
}

// WithRatings tells the query-builder to eager-load the nodes that are connected to
// the "ratings" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TripQuery) WithRatings(opts ...func(*TripRatingQuery)) *TripQuery {
	query := (&TripRatingClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withRatings = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Trip.Query().
//		GroupBy(trip.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TripQuery) GroupBy(field string, fields ...string) *TripGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TripGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = trip.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.Trip.Query().
//		Select(trip.FieldUserID).
//		Scan(ctx, &v)
func (tq *TripQuery) Select(fields ...string) *TripSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TripSelect{TripQuery: tq}
	sbuild.label = trip.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TripSelect configured with the given aggregations.
func (tq *TripQuery) Aggregate(fns ...AggregateFunc) *TripSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TripQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !trip.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TripQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Trip, error) {
	var (
		nodes       = []*Trip{}
		_spec       = tq.querySpec()
		loadedTypes = [4]bool{
			tq.withUser != nil,
			tq.withDriver != nil,
			tq.withPayment != nil,
			tq.withRatings != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Trip).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Trip{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withUser; query != nil {
		if err := tq.loadUser(ctx, query, nodes, nil,
			func(n *Trip, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withDriver; query != nil {
		if err := tq.loadDriver(ctx, query, nodes, nil,
			func(n *Trip, e *DriverProfile) { n.Edges.Driver = e }); err != nil {
			return nil, err
		}
	}
	if query := tq.withPayment; query != nil {
		if err := tq.loadPayment(ctx, query, nodes,
			func(n *Trip) { n.Edges.Payment = []*Payment{} },
			func(n *Trip, e *Payment) { n.Edges.Payment = append(n.Edges.Payment, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withRatings; query != nil {
		if err := tq.loadRatings(ctx, query, nodes,
			func(n *Trip) { n.Edges.Ratings = []*TripRating{} },
			func(n *Trip, e *TripRating) { n.Edges.Ratings = append(n.Edges.Ratings, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TripQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Trip, init func(*Trip), assign func(*Trip, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trip)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TripQuery) loadDriver(ctx context.Context, query *DriverProfileQuery, nodes []*Trip, init func(*Trip), assign func(*Trip, *DriverProfile)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Trip)
	for i := range nodes {
		if nodes[i].DriverID == nil {
			continue
		}
		fk := *nodes[i].DriverID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(driverprofile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "driver_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tq *TripQuery) loadPayment(ctx context.Context, query *PaymentQuery, nodes []*Trip, init func(*Trip), assign func(*Trip, *Payment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Trip)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(payment.FieldTripID)
	}
	query.Where(predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(trip.PaymentColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TripID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "trip_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tq *TripQuery) loadRatings(ctx context.Context, query *TripRatingQuery, nodes []*Trip, init func(*Trip), assign func(*Trip, *TripRating)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Trip)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(triprating.FieldTripID)
	}
	query.Where(predicate.TripRating(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(trip.RatingsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TripID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "trip_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TripQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TripQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(trip.Table, trip.Columns, sqlgraph.NewFieldSpec(trip.FieldID, field.TypeUUID))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trip.FieldID)
		for i := range fields {
			if fields[i] != trip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tq.withUser != nil {
			_spec.Node.AddColumnOnce(trip.FieldUserID)
		}
		if tq.withDriver != nil {
			_spec.Node.AddColumnOnce(trip.FieldDriverID)
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TripQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(trip.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = trip.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TripGroupBy is the group-by builder for Trip entities.
type TripGroupBy struct {
	selector
	build *TripQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TripGroupBy) Aggregate(fns ...AggregateFunc) *TripGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TripGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TripQuery, *TripGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TripGroupBy) sqlScan(ctx context.Context, root *TripQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TripSelect is the builder for selecting fields of Trip entities.
type TripSelect struct {
	*TripQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TripSelect) Aggregate(fns ...AggregateFunc) *TripSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TripSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TripQuery, *TripSelect](ctx, ts.TripQuery, ts, ts.inters, v)
}

func (ts *TripSelect) sqlScan(ctx context.Context, root *TripQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
