// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/creation"
	"github.com/huoayi/lfgp-ent/pkg/ent_work/user"
	"github.com/huoayi/lfgp-ent/pkg/enum"
)

// Creation is the model entity for the Creation schema.
type Creation struct {
	config `json:"-"`
	// ID of the ent.
	// 19 位雪花 ID
	ID int64 `json:"id,string"`
	// 创建者 ID
	CreatedBy int64 `json:"created_by"`
	// 更新者 ID
	UpdatedBy int64 `json:"updated_by"`
	// 创建时刻，带时区
	CreatedAt time.Time `json:"created_at"`
	// 更新时刻，带时区
	UpdatedAt time.Time `json:"updated_at"`
	// 软删除时刻，带时区
	DeletedAt time.Time `json:"deleted_at"`
	// 创作类型
	CreationType enum.CreationType `json:"creation_type"`
	// 参数
	Parameter string `json:"parameter"`
	// 作品链接
	URL string `json:"url"`
	// 用户 id
	UserID int64 `json:"user_id"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CreationQuery when eager-loading is set.
	Edges        CreationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CreationEdges holds the relations/edges for other nodes in the graph.
type CreationEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CreationEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Creation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case creation.FieldID, creation.FieldCreatedBy, creation.FieldUpdatedBy, creation.FieldUserID:
			values[i] = new(sql.NullInt64)
		case creation.FieldCreationType, creation.FieldParameter, creation.FieldURL:
			values[i] = new(sql.NullString)
		case creation.FieldCreatedAt, creation.FieldUpdatedAt, creation.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Creation fields.
func (c *Creation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case creation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case creation.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = value.Int64
			}
		case creation.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = value.Int64
			}
		case creation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case creation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case creation.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case creation.FieldCreationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creation_type", values[i])
			} else if value.Valid {
				c.CreationType = enum.CreationType(value.String)
			}
		case creation.FieldParameter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parameter", values[i])
			} else if value.Valid {
				c.Parameter = value.String
			}
		case creation.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				c.URL = value.String
			}
		case creation.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				c.UserID = value.Int64
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Creation.
// This includes values selected through modifiers, order, etc.
func (c *Creation) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Creation entity.
func (c *Creation) QueryUser() *UserQuery {
	return NewCreationClient(c.config).QueryUser(c)
}

// Update returns a builder for updating this Creation.
// Note that you need to call Creation.Unwrap() before calling this method if this Creation
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Creation) Update() *CreationUpdateOne {
	return NewCreationClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Creation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Creation) Unwrap() *Creation {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent_work: Creation is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Creation) String() string {
	var builder strings.Builder
	builder.WriteString("Creation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("creation_type=")
	builder.WriteString(fmt.Sprintf("%v", c.CreationType))
	builder.WriteString(", ")
	builder.WriteString("parameter=")
	builder.WriteString(c.Parameter)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(c.URL)
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", c.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Creations is a parsable slice of Creation.
type Creations []*Creation
