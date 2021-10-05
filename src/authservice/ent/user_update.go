// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/predicate"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/schema"
	"github.com/ducnguyen96/ducnguyen96.xyz-apis/authservice/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// SetRemindMe sets the "remind_me" field.
func (uu *UserUpdate) SetRemindMe(b bool) *UserUpdate {
	uu.mutation.SetRemindMe(b)
	return uu
}

// SetNillableRemindMe sets the "remind_me" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRemindMe(b *bool) *UserUpdate {
	if b != nil {
		uu.SetRemindMe(*b)
	}
	return uu
}

// SetWakeUpTime sets the "wake_up_time" field.
func (uu *UserUpdate) SetWakeUpTime(t time.Time) *UserUpdate {
	uu.mutation.SetWakeUpTime(t)
	return uu
}

// SetSleepTime sets the "sleep_time" field.
func (uu *UserUpdate) SetSleepTime(t time.Time) *UserUpdate {
	uu.mutation.SetSleepTime(t)
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(s schema.Gender) *UserUpdate {
	uu.mutation.SetGender(s)
	return uu
}

// SetWeight sets the "weight" field.
func (uu *UserUpdate) SetWeight(f float32) *UserUpdate {
	uu.mutation.ResetWeight()
	uu.mutation.SetWeight(f)
	return uu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWeight(f *float32) *UserUpdate {
	if f != nil {
		uu.SetWeight(*f)
	}
	return uu
}

// AddWeight adds f to the "weight" field.
func (uu *UserUpdate) AddWeight(f float32) *UserUpdate {
	uu.mutation.AddWeight(f)
	return uu
}

// SetDailyIntake sets the "daily_intake" field.
func (uu *UserUpdate) SetDailyIntake(i int32) *UserUpdate {
	uu.mutation.ResetDailyIntake()
	uu.mutation.SetDailyIntake(i)
	return uu
}

// SetNillableDailyIntake sets the "daily_intake" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDailyIntake(i *int32) *UserUpdate {
	if i != nil {
		uu.SetDailyIntake(*i)
	}
	return uu
}

// AddDailyIntake adds i to the "daily_intake" field.
func (uu *UserUpdate) AddDailyIntake(i int32) *UserUpdate {
	uu.mutation.AddDailyIntake(i)
	return uu
}

// SetContainerImage sets the "container_image" field.
func (uu *UserUpdate) SetContainerImage(s string) *UserUpdate {
	uu.mutation.SetContainerImage(s)
	return uu
}

// SetNillableContainerImage sets the "container_image" field if the given value is not nil.
func (uu *UserUpdate) SetNillableContainerImage(s *string) *UserUpdate {
	if s != nil {
		uu.SetContainerImage(*s)
	}
	return uu
}

// SetContainerVolume sets the "container_volume" field.
func (uu *UserUpdate) SetContainerVolume(i int32) *UserUpdate {
	uu.mutation.ResetContainerVolume()
	uu.mutation.SetContainerVolume(i)
	return uu
}

// SetNillableContainerVolume sets the "container_volume" field if the given value is not nil.
func (uu *UserUpdate) SetNillableContainerVolume(i *int32) *UserUpdate {
	if i != nil {
		uu.SetContainerVolume(*i)
	}
	return uu
}

// AddContainerVolume adds i to the "container_volume" field.
func (uu *UserUpdate) AddContainerVolume(i int32) *UserUpdate {
	uu.mutation.AddContainerVolume(i)
	return uu
}

// SetDrinkAtATime sets the "drink_at_a_time" field.
func (uu *UserUpdate) SetDrinkAtATime(i int32) *UserUpdate {
	uu.mutation.ResetDrinkAtATime()
	uu.mutation.SetDrinkAtATime(i)
	return uu
}

// SetNillableDrinkAtATime sets the "drink_at_a_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDrinkAtATime(i *int32) *UserUpdate {
	if i != nil {
		uu.SetDrinkAtATime(*i)
	}
	return uu
}

// AddDrinkAtATime adds i to the "drink_at_a_time" field.
func (uu *UserUpdate) AddDrinkAtATime(i int32) *UserUpdate {
	uu.mutation.AddDrinkAtATime(i)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := uu.defaults(); err != nil {
		return 0, err
	}
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
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
func (uu *UserUpdate) defaults() error {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatar,
		})
	}
	if value, ok := uu.mutation.RemindMe(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldRemindMe,
		})
	}
	if value, ok := uu.mutation.WakeUpTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldWakeUpTime,
		})
	}
	if value, ok := uu.mutation.SleepTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldSleepTime,
		})
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uu.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uu.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uu.mutation.DailyIntake(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDailyIntake,
		})
	}
	if value, ok := uu.mutation.AddedDailyIntake(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDailyIntake,
		})
	}
	if value, ok := uu.mutation.ContainerImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldContainerImage,
		})
	}
	if value, ok := uu.mutation.ContainerVolume(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldContainerVolume,
		})
	}
	if value, ok := uu.mutation.AddedContainerVolume(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldContainerVolume,
		})
	}
	if value, ok := uu.mutation.DrinkAtATime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDrinkAtATime,
		})
	}
	if value, ok := uu.mutation.AddedDrinkAtATime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDrinkAtATime,
		})
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// SetRemindMe sets the "remind_me" field.
func (uuo *UserUpdateOne) SetRemindMe(b bool) *UserUpdateOne {
	uuo.mutation.SetRemindMe(b)
	return uuo
}

// SetNillableRemindMe sets the "remind_me" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRemindMe(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetRemindMe(*b)
	}
	return uuo
}

// SetWakeUpTime sets the "wake_up_time" field.
func (uuo *UserUpdateOne) SetWakeUpTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetWakeUpTime(t)
	return uuo
}

// SetSleepTime sets the "sleep_time" field.
func (uuo *UserUpdateOne) SetSleepTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetSleepTime(t)
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(s schema.Gender) *UserUpdateOne {
	uuo.mutation.SetGender(s)
	return uuo
}

// SetWeight sets the "weight" field.
func (uuo *UserUpdateOne) SetWeight(f float32) *UserUpdateOne {
	uuo.mutation.ResetWeight()
	uuo.mutation.SetWeight(f)
	return uuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWeight(f *float32) *UserUpdateOne {
	if f != nil {
		uuo.SetWeight(*f)
	}
	return uuo
}

// AddWeight adds f to the "weight" field.
func (uuo *UserUpdateOne) AddWeight(f float32) *UserUpdateOne {
	uuo.mutation.AddWeight(f)
	return uuo
}

// SetDailyIntake sets the "daily_intake" field.
func (uuo *UserUpdateOne) SetDailyIntake(i int32) *UserUpdateOne {
	uuo.mutation.ResetDailyIntake()
	uuo.mutation.SetDailyIntake(i)
	return uuo
}

// SetNillableDailyIntake sets the "daily_intake" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDailyIntake(i *int32) *UserUpdateOne {
	if i != nil {
		uuo.SetDailyIntake(*i)
	}
	return uuo
}

// AddDailyIntake adds i to the "daily_intake" field.
func (uuo *UserUpdateOne) AddDailyIntake(i int32) *UserUpdateOne {
	uuo.mutation.AddDailyIntake(i)
	return uuo
}

// SetContainerImage sets the "container_image" field.
func (uuo *UserUpdateOne) SetContainerImage(s string) *UserUpdateOne {
	uuo.mutation.SetContainerImage(s)
	return uuo
}

// SetNillableContainerImage sets the "container_image" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableContainerImage(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetContainerImage(*s)
	}
	return uuo
}

// SetContainerVolume sets the "container_volume" field.
func (uuo *UserUpdateOne) SetContainerVolume(i int32) *UserUpdateOne {
	uuo.mutation.ResetContainerVolume()
	uuo.mutation.SetContainerVolume(i)
	return uuo
}

// SetNillableContainerVolume sets the "container_volume" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableContainerVolume(i *int32) *UserUpdateOne {
	if i != nil {
		uuo.SetContainerVolume(*i)
	}
	return uuo
}

// AddContainerVolume adds i to the "container_volume" field.
func (uuo *UserUpdateOne) AddContainerVolume(i int32) *UserUpdateOne {
	uuo.mutation.AddContainerVolume(i)
	return uuo
}

// SetDrinkAtATime sets the "drink_at_a_time" field.
func (uuo *UserUpdateOne) SetDrinkAtATime(i int32) *UserUpdateOne {
	uuo.mutation.ResetDrinkAtATime()
	uuo.mutation.SetDrinkAtATime(i)
	return uuo
}

// SetNillableDrinkAtATime sets the "drink_at_a_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDrinkAtATime(i *int32) *UserUpdateOne {
	if i != nil {
		uuo.SetDrinkAtATime(*i)
	}
	return uuo
}

// AddDrinkAtATime adds i to the "drink_at_a_time" field.
func (uuo *UserUpdateOne) AddDrinkAtATime(i int32) *UserUpdateOne {
	uuo.mutation.AddDrinkAtATime(i)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if err := uuo.defaults(); err != nil {
		return nil, err
	}
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
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
func (uuo *UserUpdateOne) defaults() error {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		if user.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized user.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Gender(); ok {
		if err := user.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
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
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldUsername,
		})
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldAvatar,
		})
	}
	if value, ok := uuo.mutation.RemindMe(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldRemindMe,
		})
	}
	if value, ok := uuo.mutation.WakeUpTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldWakeUpTime,
		})
	}
	if value, ok := uuo.mutation.SleepTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldSleepTime,
		})
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: user.FieldGender,
		})
	}
	if value, ok := uuo.mutation.Weight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uuo.mutation.AddedWeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: user.FieldWeight,
		})
	}
	if value, ok := uuo.mutation.DailyIntake(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDailyIntake,
		})
	}
	if value, ok := uuo.mutation.AddedDailyIntake(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDailyIntake,
		})
	}
	if value, ok := uuo.mutation.ContainerImage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldContainerImage,
		})
	}
	if value, ok := uuo.mutation.ContainerVolume(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldContainerVolume,
		})
	}
	if value, ok := uuo.mutation.AddedContainerVolume(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldContainerVolume,
		})
	}
	if value, ok := uuo.mutation.DrinkAtATime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDrinkAtATime,
		})
	}
	if value, ok := uuo.mutation.AddedDrinkAtATime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: user.FieldDrinkAtATime,
		})
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldPassword,
		})
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
