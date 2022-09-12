// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/basemodel"
	"gitlab.kylincloud.org/kylincloud/nucleus/src/ent/predicate"
)

// BaseModelQuery is the builder for querying BaseModel entities.
type BaseModelQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BaseModel
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BaseModelQuery builder.
func (bmq *BaseModelQuery) Where(ps ...predicate.BaseModel) *BaseModelQuery {
	bmq.predicates = append(bmq.predicates, ps...)
	return bmq
}

// Limit adds a limit step to the query.
func (bmq *BaseModelQuery) Limit(limit int) *BaseModelQuery {
	bmq.limit = &limit
	return bmq
}

// Offset adds an offset step to the query.
func (bmq *BaseModelQuery) Offset(offset int) *BaseModelQuery {
	bmq.offset = &offset
	return bmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bmq *BaseModelQuery) Unique(unique bool) *BaseModelQuery {
	bmq.unique = &unique
	return bmq
}

// Order adds an order step to the query.
func (bmq *BaseModelQuery) Order(o ...OrderFunc) *BaseModelQuery {
	bmq.order = append(bmq.order, o...)
	return bmq
}

// First returns the first BaseModel entity from the query.
// Returns a *NotFoundError when no BaseModel was found.
func (bmq *BaseModelQuery) First(ctx context.Context) (*BaseModel, error) {
	nodes, err := bmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{basemodel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bmq *BaseModelQuery) FirstX(ctx context.Context) *BaseModel {
	node, err := bmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BaseModel ID from the query.
// Returns a *NotFoundError when no BaseModel ID was found.
func (bmq *BaseModelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{basemodel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bmq *BaseModelQuery) FirstIDX(ctx context.Context) int {
	id, err := bmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BaseModel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BaseModel entity is found.
// Returns a *NotFoundError when no BaseModel entities are found.
func (bmq *BaseModelQuery) Only(ctx context.Context) (*BaseModel, error) {
	nodes, err := bmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{basemodel.Label}
	default:
		return nil, &NotSingularError{basemodel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bmq *BaseModelQuery) OnlyX(ctx context.Context) *BaseModel {
	node, err := bmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BaseModel ID in the query.
// Returns a *NotSingularError when more than one BaseModel ID is found.
// Returns a *NotFoundError when no entities are found.
func (bmq *BaseModelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{basemodel.Label}
	default:
		err = &NotSingularError{basemodel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bmq *BaseModelQuery) OnlyIDX(ctx context.Context) int {
	id, err := bmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BaseModels.
func (bmq *BaseModelQuery) All(ctx context.Context) ([]*BaseModel, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bmq *BaseModelQuery) AllX(ctx context.Context) []*BaseModel {
	nodes, err := bmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BaseModel IDs.
func (bmq *BaseModelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bmq.Select(basemodel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bmq *BaseModelQuery) IDsX(ctx context.Context) []int {
	ids, err := bmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bmq *BaseModelQuery) Count(ctx context.Context) (int, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bmq *BaseModelQuery) CountX(ctx context.Context) int {
	count, err := bmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bmq *BaseModelQuery) Exist(ctx context.Context) (bool, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bmq *BaseModelQuery) ExistX(ctx context.Context) bool {
	exist, err := bmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BaseModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bmq *BaseModelQuery) Clone() *BaseModelQuery {
	if bmq == nil {
		return nil
	}
	return &BaseModelQuery{
		config:     bmq.config,
		limit:      bmq.limit,
		offset:     bmq.offset,
		order:      append([]OrderFunc{}, bmq.order...),
		predicates: append([]predicate.BaseModel{}, bmq.predicates...),
		// clone intermediate query.
		sql:    bmq.sql.Clone(),
		path:   bmq.path,
		unique: bmq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BaseModel.Query().
//		GroupBy(basemodel.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bmq *BaseModelQuery) GroupBy(field string, fields ...string) *BaseModelGroupBy {
	grbuild := &BaseModelGroupBy{config: bmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bmq.sqlQuery(ctx), nil
	}
	grbuild.label = basemodel.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.BaseModel.Query().
//		Select(basemodel.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (bmq *BaseModelQuery) Select(fields ...string) *BaseModelSelect {
	bmq.fields = append(bmq.fields, fields...)
	selbuild := &BaseModelSelect{BaseModelQuery: bmq}
	selbuild.label = basemodel.Label
	selbuild.flds, selbuild.scan = &bmq.fields, selbuild.Scan
	return selbuild
}

func (bmq *BaseModelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bmq.fields {
		if !basemodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bmq.path != nil {
		prev, err := bmq.path(ctx)
		if err != nil {
			return err
		}
		bmq.sql = prev
	}
	return nil
}

func (bmq *BaseModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BaseModel, error) {
	var (
		nodes = []*BaseModel{}
		_spec = bmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BaseModel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BaseModel{config: bmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bmq *BaseModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bmq.querySpec()
	_spec.Node.Columns = bmq.fields
	if len(bmq.fields) > 0 {
		_spec.Unique = bmq.unique != nil && *bmq.unique
	}
	return sqlgraph.CountNodes(ctx, bmq.driver, _spec)
}

func (bmq *BaseModelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bmq *BaseModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   basemodel.Table,
			Columns: basemodel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: basemodel.FieldID,
			},
		},
		From:   bmq.sql,
		Unique: true,
	}
	if unique := bmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, basemodel.FieldID)
		for i := range fields {
			if fields[i] != basemodel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bmq *BaseModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bmq.driver.Dialect())
	t1 := builder.Table(basemodel.Table)
	columns := bmq.fields
	if len(columns) == 0 {
		columns = basemodel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bmq.sql != nil {
		selector = bmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bmq.unique != nil && *bmq.unique {
		selector.Distinct()
	}
	for _, p := range bmq.predicates {
		p(selector)
	}
	for _, p := range bmq.order {
		p(selector)
	}
	if offset := bmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BaseModelGroupBy is the group-by builder for BaseModel entities.
type BaseModelGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bmgb *BaseModelGroupBy) Aggregate(fns ...AggregateFunc) *BaseModelGroupBy {
	bmgb.fns = append(bmgb.fns, fns...)
	return bmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bmgb *BaseModelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bmgb.path(ctx)
	if err != nil {
		return err
	}
	bmgb.sql = query
	return bmgb.sqlScan(ctx, v)
}

func (bmgb *BaseModelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bmgb.fields {
		if !basemodel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bmgb *BaseModelGroupBy) sqlQuery() *sql.Selector {
	selector := bmgb.sql.Select()
	aggregation := make([]string, 0, len(bmgb.fns))
	for _, fn := range bmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bmgb.fields)+len(bmgb.fns))
		for _, f := range bmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bmgb.fields...)...)
}

// BaseModelSelect is the builder for selecting fields of BaseModel entities.
type BaseModelSelect struct {
	*BaseModelQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bms *BaseModelSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bms.prepareQuery(ctx); err != nil {
		return err
	}
	bms.sql = bms.BaseModelQuery.sqlQuery(ctx)
	return bms.sqlScan(ctx, v)
}

func (bms *BaseModelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bms.sql.Query()
	if err := bms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
