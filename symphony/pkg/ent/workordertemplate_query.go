// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/symphony/pkg/ent/checklistcategorydefinition"
	"github.com/facebookincubator/symphony/pkg/ent/predicate"
	"github.com/facebookincubator/symphony/pkg/ent/propertytype"
	"github.com/facebookincubator/symphony/pkg/ent/workordertemplate"
	"github.com/facebookincubator/symphony/pkg/ent/workordertype"
)

// WorkOrderTemplateQuery is the builder for querying WorkOrderTemplate entities.
type WorkOrderTemplateQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.WorkOrderTemplate
	// eager-loading edges.
	withPropertyTypes                *PropertyTypeQuery
	withCheckListCategoryDefinitions *CheckListCategoryDefinitionQuery
	withType                         *WorkOrderTypeQuery
	withFKs                          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (wotq *WorkOrderTemplateQuery) Where(ps ...predicate.WorkOrderTemplate) *WorkOrderTemplateQuery {
	wotq.predicates = append(wotq.predicates, ps...)
	return wotq
}

// Limit adds a limit step to the query.
func (wotq *WorkOrderTemplateQuery) Limit(limit int) *WorkOrderTemplateQuery {
	wotq.limit = &limit
	return wotq
}

// Offset adds an offset step to the query.
func (wotq *WorkOrderTemplateQuery) Offset(offset int) *WorkOrderTemplateQuery {
	wotq.offset = &offset
	return wotq
}

// Order adds an order step to the query.
func (wotq *WorkOrderTemplateQuery) Order(o ...OrderFunc) *WorkOrderTemplateQuery {
	wotq.order = append(wotq.order, o...)
	return wotq
}

// QueryPropertyTypes chains the current query on the property_types edge.
func (wotq *WorkOrderTemplateQuery) QueryPropertyTypes() *PropertyTypeQuery {
	query := &PropertyTypeQuery{config: wotq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wotq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workordertemplate.Table, workordertemplate.FieldID, wotq.sqlQuery()),
			sqlgraph.To(propertytype.Table, propertytype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workordertemplate.PropertyTypesTable, workordertemplate.PropertyTypesColumn),
		)
		fromU = sqlgraph.SetNeighbors(wotq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCheckListCategoryDefinitions chains the current query on the check_list_category_definitions edge.
func (wotq *WorkOrderTemplateQuery) QueryCheckListCategoryDefinitions() *CheckListCategoryDefinitionQuery {
	query := &CheckListCategoryDefinitionQuery{config: wotq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wotq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workordertemplate.Table, workordertemplate.FieldID, wotq.sqlQuery()),
			sqlgraph.To(checklistcategorydefinition.Table, checklistcategorydefinition.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, workordertemplate.CheckListCategoryDefinitionsTable, workordertemplate.CheckListCategoryDefinitionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wotq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryType chains the current query on the type edge.
func (wotq *WorkOrderTemplateQuery) QueryType() *WorkOrderTypeQuery {
	query := &WorkOrderTypeQuery{config: wotq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wotq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(workordertemplate.Table, workordertemplate.FieldID, wotq.sqlQuery()),
			sqlgraph.To(workordertype.Table, workordertype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, workordertemplate.TypeTable, workordertemplate.TypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(wotq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WorkOrderTemplate entity in the query. Returns *NotFoundError when no workordertemplate was found.
func (wotq *WorkOrderTemplateQuery) First(ctx context.Context) (*WorkOrderTemplate, error) {
	wots, err := wotq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(wots) == 0 {
		return nil, &NotFoundError{workordertemplate.Label}
	}
	return wots[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) FirstX(ctx context.Context) *WorkOrderTemplate {
	wot, err := wotq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return wot
}

// FirstID returns the first WorkOrderTemplate id in the query. Returns *NotFoundError when no id was found.
func (wotq *WorkOrderTemplateQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wotq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{workordertemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) FirstXID(ctx context.Context) int {
	id, err := wotq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only WorkOrderTemplate entity in the query, returns an error if not exactly one entity was returned.
func (wotq *WorkOrderTemplateQuery) Only(ctx context.Context) (*WorkOrderTemplate, error) {
	wots, err := wotq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(wots) {
	case 1:
		return wots[0], nil
	case 0:
		return nil, &NotFoundError{workordertemplate.Label}
	default:
		return nil, &NotSingularError{workordertemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) OnlyX(ctx context.Context) *WorkOrderTemplate {
	wot, err := wotq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return wot
}

// OnlyID returns the only WorkOrderTemplate id in the query, returns an error if not exactly one id was returned.
func (wotq *WorkOrderTemplateQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wotq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = &NotSingularError{workordertemplate.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) OnlyXID(ctx context.Context) int {
	id, err := wotq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WorkOrderTemplates.
func (wotq *WorkOrderTemplateQuery) All(ctx context.Context) ([]*WorkOrderTemplate, error) {
	if err := wotq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return wotq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) AllX(ctx context.Context) []*WorkOrderTemplate {
	wots, err := wotq.All(ctx)
	if err != nil {
		panic(err)
	}
	return wots
}

// IDs executes the query and returns a list of WorkOrderTemplate ids.
func (wotq *WorkOrderTemplateQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wotq.Select(workordertemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) IDsX(ctx context.Context) []int {
	ids, err := wotq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wotq *WorkOrderTemplateQuery) Count(ctx context.Context) (int, error) {
	if err := wotq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return wotq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) CountX(ctx context.Context) int {
	count, err := wotq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wotq *WorkOrderTemplateQuery) Exist(ctx context.Context) (bool, error) {
	if err := wotq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return wotq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wotq *WorkOrderTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := wotq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wotq *WorkOrderTemplateQuery) Clone() *WorkOrderTemplateQuery {
	return &WorkOrderTemplateQuery{
		config:     wotq.config,
		limit:      wotq.limit,
		offset:     wotq.offset,
		order:      append([]OrderFunc{}, wotq.order...),
		unique:     append([]string{}, wotq.unique...),
		predicates: append([]predicate.WorkOrderTemplate{}, wotq.predicates...),
		// clone intermediate query.
		sql:  wotq.sql.Clone(),
		path: wotq.path,
	}
}

//  WithPropertyTypes tells the query-builder to eager-loads the nodes that are connected to
// the "property_types" edge. The optional arguments used to configure the query builder of the edge.
func (wotq *WorkOrderTemplateQuery) WithPropertyTypes(opts ...func(*PropertyTypeQuery)) *WorkOrderTemplateQuery {
	query := &PropertyTypeQuery{config: wotq.config}
	for _, opt := range opts {
		opt(query)
	}
	wotq.withPropertyTypes = query
	return wotq
}

//  WithCheckListCategoryDefinitions tells the query-builder to eager-loads the nodes that are connected to
// the "check_list_category_definitions" edge. The optional arguments used to configure the query builder of the edge.
func (wotq *WorkOrderTemplateQuery) WithCheckListCategoryDefinitions(opts ...func(*CheckListCategoryDefinitionQuery)) *WorkOrderTemplateQuery {
	query := &CheckListCategoryDefinitionQuery{config: wotq.config}
	for _, opt := range opts {
		opt(query)
	}
	wotq.withCheckListCategoryDefinitions = query
	return wotq
}

//  WithType tells the query-builder to eager-loads the nodes that are connected to
// the "type" edge. The optional arguments used to configure the query builder of the edge.
func (wotq *WorkOrderTemplateQuery) WithType(opts ...func(*WorkOrderTypeQuery)) *WorkOrderTemplateQuery {
	query := &WorkOrderTypeQuery{config: wotq.config}
	for _, opt := range opts {
		opt(query)
	}
	wotq.withType = query
	return wotq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WorkOrderTemplate.Query().
//		GroupBy(workordertemplate.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wotq *WorkOrderTemplateQuery) GroupBy(field string, fields ...string) *WorkOrderTemplateGroupBy {
	group := &WorkOrderTemplateGroupBy{config: wotq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wotq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wotq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.WorkOrderTemplate.Query().
//		Select(workordertemplate.FieldName).
//		Scan(ctx, &v)
//
func (wotq *WorkOrderTemplateQuery) Select(field string, fields ...string) *WorkOrderTemplateSelect {
	selector := &WorkOrderTemplateSelect{config: wotq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := wotq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return wotq.sqlQuery(), nil
	}
	return selector
}

func (wotq *WorkOrderTemplateQuery) prepareQuery(ctx context.Context) error {
	if wotq.path != nil {
		prev, err := wotq.path(ctx)
		if err != nil {
			return err
		}
		wotq.sql = prev
	}
	if err := workordertemplate.Policy.EvalQuery(ctx, wotq); err != nil {
		return err
	}
	return nil
}

func (wotq *WorkOrderTemplateQuery) sqlAll(ctx context.Context) ([]*WorkOrderTemplate, error) {
	var (
		nodes       = []*WorkOrderTemplate{}
		withFKs     = wotq.withFKs
		_spec       = wotq.querySpec()
		loadedTypes = [3]bool{
			wotq.withPropertyTypes != nil,
			wotq.withCheckListCategoryDefinitions != nil,
			wotq.withType != nil,
		}
	)
	if wotq.withType != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, workordertemplate.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &WorkOrderTemplate{config: wotq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, wotq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := wotq.withPropertyTypes; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkOrderTemplate)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.PropertyType(func(s *sql.Selector) {
			s.Where(sql.InValues(workordertemplate.PropertyTypesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.work_order_template_property_types
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "work_order_template_property_types" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "work_order_template_property_types" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.PropertyTypes = append(node.Edges.PropertyTypes, n)
		}
	}

	if query := wotq.withCheckListCategoryDefinitions; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*WorkOrderTemplate)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.CheckListCategoryDefinition(func(s *sql.Selector) {
			s.Where(sql.InValues(workordertemplate.CheckListCategoryDefinitionsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.work_order_template_check_list_category_definitions
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "work_order_template_check_list_category_definitions" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "work_order_template_check_list_category_definitions" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.CheckListCategoryDefinitions = append(node.Edges.CheckListCategoryDefinitions, n)
		}
	}

	if query := wotq.withType; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*WorkOrderTemplate)
		for i := range nodes {
			if fk := nodes[i].work_order_template_type; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(workordertype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "work_order_template_type" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Type = n
			}
		}
	}

	return nodes, nil
}

func (wotq *WorkOrderTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wotq.querySpec()
	return sqlgraph.CountNodes(ctx, wotq.driver, _spec)
}

func (wotq *WorkOrderTemplateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wotq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (wotq *WorkOrderTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   workordertemplate.Table,
			Columns: workordertemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: workordertemplate.FieldID,
			},
		},
		From:   wotq.sql,
		Unique: true,
	}
	if ps := wotq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wotq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wotq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wotq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wotq *WorkOrderTemplateQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wotq.driver.Dialect())
	t1 := builder.Table(workordertemplate.Table)
	selector := builder.Select(t1.Columns(workordertemplate.Columns...)...).From(t1)
	if wotq.sql != nil {
		selector = wotq.sql
		selector.Select(selector.Columns(workordertemplate.Columns...)...)
	}
	for _, p := range wotq.predicates {
		p(selector)
	}
	for _, p := range wotq.order {
		p(selector)
	}
	if offset := wotq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wotq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WorkOrderTemplateGroupBy is the builder for group-by WorkOrderTemplate entities.
type WorkOrderTemplateGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wotgb *WorkOrderTemplateGroupBy) Aggregate(fns ...AggregateFunc) *WorkOrderTemplateGroupBy {
	wotgb.fns = append(wotgb.fns, fns...)
	return wotgb
}

// Scan applies the group-by query and scan the result into the given value.
func (wotgb *WorkOrderTemplateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := wotgb.path(ctx)
	if err != nil {
		return err
	}
	wotgb.sql = query
	return wotgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wotgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wotgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wotgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) StringsX(ctx context.Context) []string {
	v, err := wotgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wotgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) StringX(ctx context.Context) string {
	v, err := wotgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wotgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wotgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) IntsX(ctx context.Context) []int {
	v, err := wotgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wotgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) IntX(ctx context.Context) int {
	v, err := wotgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wotgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wotgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wotgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wotgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) Float64X(ctx context.Context) float64 {
	v, err := wotgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wotgb.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wotgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wotgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (wotgb *WorkOrderTemplateGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wotgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wotgb *WorkOrderTemplateGroupBy) BoolX(ctx context.Context) bool {
	v, err := wotgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wotgb *WorkOrderTemplateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wotgb.sqlQuery().Query()
	if err := wotgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wotgb *WorkOrderTemplateGroupBy) sqlQuery() *sql.Selector {
	selector := wotgb.sql
	columns := make([]string, 0, len(wotgb.fields)+len(wotgb.fns))
	columns = append(columns, wotgb.fields...)
	for _, fn := range wotgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wotgb.fields...)
}

// WorkOrderTemplateSelect is the builder for select fields of WorkOrderTemplate entities.
type WorkOrderTemplateSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (wots *WorkOrderTemplateSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := wots.path(ctx)
	if err != nil {
		return err
	}
	wots.sql = query
	return wots.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wots.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wots.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) StringsX(ctx context.Context) []string {
	v, err := wots.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = wots.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) StringX(ctx context.Context) string {
	v, err := wots.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wots.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) IntsX(ctx context.Context) []int {
	v, err := wots.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = wots.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) IntX(ctx context.Context) int {
	v, err := wots.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wots.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wots.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = wots.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) Float64X(ctx context.Context) float64 {
	v, err := wots.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wots.fields) > 1 {
		return nil, errors.New("ent: WorkOrderTemplateSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wots.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) BoolsX(ctx context.Context) []bool {
	v, err := wots.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (wots *WorkOrderTemplateSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = wots.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{workordertemplate.Label}
	default:
		err = fmt.Errorf("ent: WorkOrderTemplateSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (wots *WorkOrderTemplateSelect) BoolX(ctx context.Context) bool {
	v, err := wots.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wots *WorkOrderTemplateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wots.sqlQuery().Query()
	if err := wots.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wots *WorkOrderTemplateSelect) sqlQuery() sql.Querier {
	selector := wots.sql
	selector.Select(selector.Columns(wots.fields...)...)
	return selector
}