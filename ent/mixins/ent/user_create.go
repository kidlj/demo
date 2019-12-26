// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kidlj/demo/ent/mixins/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	created_at *time.Time
	updated_at *time.Time
	age        *int
	name       *string
	nickname   *string
}

// SetCreatedAt sets the created_at field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.created_at = &t
	return uc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the updated_at field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.updated_at = &t
	return uc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.age = &i
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetNickname sets the nickname field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.nickname = &s
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.created_at == nil {
		v := user.DefaultCreatedAt()
		uc.created_at = &v
	}
	if uc.updated_at == nil {
		v := user.DefaultUpdatedAt()
		uc.updated_at = &v
	}
	if uc.age == nil {
		return nil, errors.New("ent: missing required field \"age\"")
	}
	if err := user.AgeValidator(*uc.age); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"age\": %v", err)
	}
	if uc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := user.NameValidator(*uc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	if uc.nickname == nil {
		return nil, errors.New("ent: missing required field \"nickname\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u    = &User{config: uc.config}
		spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.created_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldCreatedAt,
		})
		u.CreatedAt = *value
	}
	if value := uc.updated_at; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldUpdatedAt,
		})
		u.UpdatedAt = *value
	}
	if value := uc.age; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: user.FieldAge,
		})
		u.Age = *value
	}
	if value := uc.name; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
		u.Name = *value
	}
	if value := uc.nickname; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
		u.Nickname = *value
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
