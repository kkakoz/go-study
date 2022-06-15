// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"fmt"
	"learn-go/db/entdemo/ent/education"
	"learn-go/db/entdemo/ent/institution"
	"learn-go/db/entdemo/ent/predicate"
)

// EducationUpdate is the builder for updating Education entities.
type EducationUpdate struct {
	config
	hooks    []Hook
	mutation *EducationMutation
}

// Where appends a list predicates to the EducationUpdate builder.
func (eu *EducationUpdate) Where(ps ...predicate.Education) *EducationUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EducationUpdate) SetName(s string) *EducationUpdate {
	eu.mutation.SetName(s)
	return eu
}

// AddInstitutionIDs adds the "institution" edge to the Institution entity by IDs.
func (eu *EducationUpdate) AddInstitutionIDs(ids ...int) *EducationUpdate {
	eu.mutation.AddInstitutionIDs(ids...)
	return eu
}

// AddInstitution adds the "institution" edges to the Institution entity.
func (eu *EducationUpdate) AddInstitution(i ...*Institution) *EducationUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return eu.AddInstitutionIDs(ids...)
}

// Mutation returns the EducationMutation object of the builder.
func (eu *EducationUpdate) Mutation() *EducationMutation {
	return eu.mutation
}

// ClearInstitution clears all "institution" edges to the Institution entity.
func (eu *EducationUpdate) ClearInstitution() *EducationUpdate {
	eu.mutation.ClearInstitution()
	return eu
}

// RemoveInstitutionIDs removes the "institution" edge to Institution entities by IDs.
func (eu *EducationUpdate) RemoveInstitutionIDs(ids ...int) *EducationUpdate {
	eu.mutation.RemoveInstitutionIDs(ids...)
	return eu
}

// RemoveInstitution removes "institution" edges to Institution entities.
func (eu *EducationUpdate) RemoveInstitution(i ...*Institution) *EducationUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return eu.RemoveInstitutionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EducationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EducationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EducationUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EducationUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EducationUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EducationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   education.Table,
			Columns: education.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: education.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldName,
		})
	}
	if eu.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   education.InstitutionTable,
			Columns: []string{education.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institution.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedInstitutionIDs(); len(nodes) > 0 && !eu.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   education.InstitutionTable,
			Columns: []string{education.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   education.InstitutionTable,
			Columns: []string{education.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{education.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EducationUpdateOne is the builder for updating a single Education entity.
type EducationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EducationMutation
}

// SetName sets the "name" field.
func (euo *EducationUpdateOne) SetName(s string) *EducationUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// AddInstitutionIDs adds the "institution" edge to the Institution entity by IDs.
func (euo *EducationUpdateOne) AddInstitutionIDs(ids ...int) *EducationUpdateOne {
	euo.mutation.AddInstitutionIDs(ids...)
	return euo
}

// AddInstitution adds the "institution" edges to the Institution entity.
func (euo *EducationUpdateOne) AddInstitution(i ...*Institution) *EducationUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return euo.AddInstitutionIDs(ids...)
}

// Mutation returns the EducationMutation object of the builder.
func (euo *EducationUpdateOne) Mutation() *EducationMutation {
	return euo.mutation
}

// ClearInstitution clears all "institution" edges to the Institution entity.
func (euo *EducationUpdateOne) ClearInstitution() *EducationUpdateOne {
	euo.mutation.ClearInstitution()
	return euo
}

// RemoveInstitutionIDs removes the "institution" edge to Institution entities by IDs.
func (euo *EducationUpdateOne) RemoveInstitutionIDs(ids ...int) *EducationUpdateOne {
	euo.mutation.RemoveInstitutionIDs(ids...)
	return euo
}

// RemoveInstitution removes "institution" edges to Institution entities.
func (euo *EducationUpdateOne) RemoveInstitution(i ...*Institution) *EducationUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return euo.RemoveInstitutionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EducationUpdateOne) Select(field string, fields ...string) *EducationUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Education entity.
func (euo *EducationUpdateOne) Save(ctx context.Context) (*Education, error) {
	var (
		err  error
		node *Education
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EducationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EducationUpdateOne) SaveX(ctx context.Context) *Education {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EducationUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EducationUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EducationUpdateOne) sqlSave(ctx context.Context) (_node *Education, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   education.Table,
			Columns: education.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: education.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Education.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, education.FieldID)
		for _, f := range fields {
			if !education.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != education.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: education.FieldName,
		})
	}
	if euo.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   education.InstitutionTable,
			Columns: []string{education.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institution.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedInstitutionIDs(); len(nodes) > 0 && !euo.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   education.InstitutionTable,
			Columns: []string{education.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   education.InstitutionTable,
			Columns: []string{education.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: institution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Education{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{education.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}