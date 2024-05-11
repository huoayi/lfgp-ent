// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/creation"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/predicate"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetCreatedBy sets the "created_by" field.
func (uu *UserUpdate) SetCreatedBy(i int64) *UserUpdate {
	uu.mutation.ResetCreatedBy()
	uu.mutation.SetCreatedBy(i)
	return uu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedBy(i *int64) *UserUpdate {
	if i != nil {
		uu.SetCreatedBy(*i)
	}
	return uu
}

// AddCreatedBy adds i to the "created_by" field.
func (uu *UserUpdate) AddCreatedBy(i int64) *UserUpdate {
	uu.mutation.AddCreatedBy(i)
	return uu
}

// SetUpdatedBy sets the "updated_by" field.
func (uu *UserUpdate) SetUpdatedBy(i int64) *UserUpdate {
	uu.mutation.ResetUpdatedBy()
	uu.mutation.SetUpdatedBy(i)
	return uu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedBy(i *int64) *UserUpdate {
	if i != nil {
		uu.SetUpdatedBy(*i)
	}
	return uu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (uu *UserUpdate) AddUpdatedBy(i int64) *UserUpdate {
	uu.mutation.AddUpdatedBy(i)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetNickName sets the "nick_name" field.
func (uu *UserUpdate) SetNickName(s string) *UserUpdate {
	uu.mutation.SetNickName(s)
	return uu
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickName(*s)
	}
	return uu
}

// SetJpgURL sets the "jpg_url" field.
func (uu *UserUpdate) SetJpgURL(s string) *UserUpdate {
	uu.mutation.SetJpgURL(s)
	return uu
}

// SetNillableJpgURL sets the "jpg_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillableJpgURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetJpgURL(*s)
	}
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// AddCreationIDs adds the "creations" edge to the Creation entity by IDs.
func (uu *UserUpdate) AddCreationIDs(ids ...int64) *UserUpdate {
	uu.mutation.AddCreationIDs(ids...)
	return uu
}

// AddCreations adds the "creations" edges to the Creation entity.
func (uu *UserUpdate) AddCreations(c ...*Creation) *UserUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCreationIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCreations clears all "creations" edges to the Creation entity.
func (uu *UserUpdate) ClearCreations() *UserUpdate {
	uu.mutation.ClearCreations()
	return uu
}

// RemoveCreationIDs removes the "creations" edge to Creation entities by IDs.
func (uu *UserUpdate) RemoveCreationIDs(ids ...int64) *UserUpdate {
	uu.mutation.RemoveCreationIDs(ids...)
	return uu
}

// RemoveCreations removes "creations" edges to Creation entities.
func (uu *UserUpdate) RemoveCreations(c ...*Creation) *UserUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCreationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(user.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(user.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if value, ok := uu.mutation.JpgURL(); ok {
		_spec.SetField(user.FieldJpgURL, field.TypeString, value)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.CreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreationsTable,
			Columns: []string{user.CreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creation.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedCreationsIDs(); len(nodes) > 0 && !uu.mutation.CreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreationsTable,
			Columns: []string{user.CreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creation.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CreationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreationsTable,
			Columns: []string{user.CreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creation.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(uu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (uuo *UserUpdateOne) SetCreatedBy(i int64) *UserUpdateOne {
	uuo.mutation.ResetCreatedBy()
	uuo.mutation.SetCreatedBy(i)
	return uuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedBy(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetCreatedBy(*i)
	}
	return uuo
}

// AddCreatedBy adds i to the "created_by" field.
func (uuo *UserUpdateOne) AddCreatedBy(i int64) *UserUpdateOne {
	uuo.mutation.AddCreatedBy(i)
	return uuo
}

// SetUpdatedBy sets the "updated_by" field.
func (uuo *UserUpdateOne) SetUpdatedBy(i int64) *UserUpdateOne {
	uuo.mutation.ResetUpdatedBy()
	uuo.mutation.SetUpdatedBy(i)
	return uuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedBy(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetUpdatedBy(*i)
	}
	return uuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (uuo *UserUpdateOne) AddUpdatedBy(i int64) *UserUpdateOne {
	uuo.mutation.AddUpdatedBy(i)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetNickName sets the "nick_name" field.
func (uuo *UserUpdateOne) SetNickName(s string) *UserUpdateOne {
	uuo.mutation.SetNickName(s)
	return uuo
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickName(*s)
	}
	return uuo
}

// SetJpgURL sets the "jpg_url" field.
func (uuo *UserUpdateOne) SetJpgURL(s string) *UserUpdateOne {
	uuo.mutation.SetJpgURL(s)
	return uuo
}

// SetNillableJpgURL sets the "jpg_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableJpgURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetJpgURL(*s)
	}
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// AddCreationIDs adds the "creations" edge to the Creation entity by IDs.
func (uuo *UserUpdateOne) AddCreationIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.AddCreationIDs(ids...)
	return uuo
}

// AddCreations adds the "creations" edges to the Creation entity.
func (uuo *UserUpdateOne) AddCreations(c ...*Creation) *UserUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCreationIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCreations clears all "creations" edges to the Creation entity.
func (uuo *UserUpdateOne) ClearCreations() *UserUpdateOne {
	uuo.mutation.ClearCreations()
	return uuo
}

// RemoveCreationIDs removes the "creations" edge to Creation entities by IDs.
func (uuo *UserUpdateOne) RemoveCreationIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.RemoveCreationIDs(ids...)
	return uuo
}

// RemoveCreations removes "creations" edges to Creation entities.
func (uuo *UserUpdateOne) RemoveCreations(c ...*Creation) *UserUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCreationIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent_work: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent_work: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(user.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(user.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.JpgURL(); ok {
		_spec.SetField(user.FieldJpgURL, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.CreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreationsTable,
			Columns: []string{user.CreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creation.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedCreationsIDs(); len(nodes) > 0 && !uuo.mutation.CreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreationsTable,
			Columns: []string{user.CreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creation.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CreationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.CreationsTable,
			Columns: []string{user.CreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(creation.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(uuo.modifiers...)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
