// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/novel/internal/data/ent/novelautobuy"
	"hope/apps/novel/internal/data/ent/socialuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// NovelAutoBuy is the model entity for the NovelAutoBuy schema.
type NovelAutoBuy struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	// 用户ID
	UserId int64 `json:"userId,omitempty"`
	// NovelId holds the value of the "novelId" field.
	// 小说编号
	NovelId int64 `json:"novelId,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	// 创建时间
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	// 更新时间
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// CreateBy holds the value of the "createBy" field.
	// 创建者
	CreateBy int64 `json:"createBy,omitempty"`
	// UpdateBy holds the value of the "updateBy" field.
	// 更新者
	UpdateBy int64 `json:"updateBy,omitempty"`
	// TenantId holds the value of the "tenantId" field.
	// 租户
	TenantId int64 `json:"tenantId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NovelAutoBuyQuery when eager-loading is set.
	Edges                       NovelAutoBuyEdges `json:"edges"`
	social_user_auto_buy_novels *int64
}

// NovelAutoBuyEdges holds the relations/edges for other nodes in the graph.
type NovelAutoBuyEdges struct {
	// User holds the value of the user edge.
	User *SocialUser `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NovelAutoBuyEdges) UserOrErr() (*SocialUser, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: socialuser.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NovelAutoBuy) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case novelautobuy.FieldID, novelautobuy.FieldUserId, novelautobuy.FieldNovelId, novelautobuy.FieldCreateBy, novelautobuy.FieldUpdateBy, novelautobuy.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case novelautobuy.FieldCreatedAt, novelautobuy.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case novelautobuy.ForeignKeys[0]: // social_user_auto_buy_novels
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NovelAutoBuy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NovelAutoBuy fields.
func (nab *NovelAutoBuy) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case novelautobuy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nab.ID = int64(value.Int64)
		case novelautobuy.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				nab.UserId = value.Int64
			}
		case novelautobuy.FieldNovelId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field novelId", values[i])
			} else if value.Valid {
				nab.NovelId = value.Int64
			}
		case novelautobuy.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				nab.CreatedAt = value.Time
			}
		case novelautobuy.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				nab.UpdatedAt = value.Time
			}
		case novelautobuy.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				nab.CreateBy = value.Int64
			}
		case novelautobuy.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				nab.UpdateBy = value.Int64
			}
		case novelautobuy.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				nab.TenantId = value.Int64
			}
		case novelautobuy.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field social_user_auto_buy_novels", value)
			} else if value.Valid {
				nab.social_user_auto_buy_novels = new(int64)
				*nab.social_user_auto_buy_novels = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the NovelAutoBuy entity.
func (nab *NovelAutoBuy) QueryUser() *SocialUserQuery {
	return (&NovelAutoBuyClient{config: nab.config}).QueryUser(nab)
}

// Update returns a builder for updating this NovelAutoBuy.
// Note that you need to call NovelAutoBuy.Unwrap() before calling this method if this NovelAutoBuy
// was returned from a transaction, and the transaction was committed or rolled back.
func (nab *NovelAutoBuy) Update() *NovelAutoBuyUpdateOne {
	return (&NovelAutoBuyClient{config: nab.config}).UpdateOne(nab)
}

// Unwrap unwraps the NovelAutoBuy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nab *NovelAutoBuy) Unwrap() *NovelAutoBuy {
	tx, ok := nab.config.driver.(*txDriver)
	if !ok {
		panic("ent: NovelAutoBuy is not a transactional entity")
	}
	nab.config.driver = tx.drv
	return nab
}

// String implements the fmt.Stringer.
func (nab *NovelAutoBuy) String() string {
	var builder strings.Builder
	builder.WriteString("NovelAutoBuy(")
	builder.WriteString(fmt.Sprintf("id=%v", nab.ID))
	builder.WriteString(", userId=")
	builder.WriteString(fmt.Sprintf("%v", nab.UserId))
	builder.WriteString(", novelId=")
	builder.WriteString(fmt.Sprintf("%v", nab.NovelId))
	builder.WriteString(", createdAt=")
	builder.WriteString(nab.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(nab.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", nab.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", nab.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", nab.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// NovelAutoBuys is a parsable slice of NovelAutoBuy.
type NovelAutoBuys []*NovelAutoBuy

func (nab NovelAutoBuys) config(cfg config) {
	for _i := range nab {
		nab[_i].config = cfg
	}
}