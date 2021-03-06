// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/novel/internal/data/ent/novelcomment"
	"hope/apps/novel/internal/data/ent/socialuser"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// NovelComment is the model entity for the NovelComment schema.
type NovelComment struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// NovelId holds the value of the "novelId" field.
	// 小说编号
	NovelId int64 `json:"novelId,omitempty"`
	// UserId holds the value of the "userId" field.
	// 用户Id
	UserId int64 `json:"userId,omitempty"`
	// Avatar holds the value of the "avatar" field.
	// 评论用户头像
	Avatar string `json:"avatar,omitempty"`
	// UserName holds the value of the "userName" field.
	// 用户名
	UserName string `json:"userName,omitempty"`
	// RepUserId holds the value of the "repUserId" field.
	// 回复用户
	RepUserId int64 `json:"repUserId,omitempty"`
	// RepUserName holds the value of the "repUserName" field.
	// 回复用户ID
	RepUserName string `json:"repUserName,omitempty"`
	// Content holds the value of the "content" field.
	// 回复内容
	Content string `json:"content,omitempty"`
	// Score holds the value of the "score" field.
	// 评分,除以10
	Score int32 `json:"score,omitempty"`
	// PId holds the value of the "pId" field.
	// 回复评论ID
	PId int64 `json:"pId,omitempty"`
	// IsTop holds the value of the "isTop" field.
	// 置顶
	IsTop bool `json:"isTop,omitempty"`
	// State holds the value of the "state" field.
	// 状态,0
	State novelcomment.State `json:"state,omitempty"`
	// IsHighlight holds the value of the "isHighlight" field.
	// 高亮
	IsHighlight bool `json:"isHighlight,omitempty"`
	// IsHot holds the value of the "isHot" field.
	// 是否热门
	IsHot bool `json:"isHot,omitempty"`
	// Remark holds the value of the "remark" field.
	// 备注
	Remark string `json:"remark,omitempty"`
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
	// The values are being populated by the NovelCommentQuery when eager-loading is set.
	Edges                 NovelCommentEdges `json:"edges"`
	novel_comment_childes *int64
}

// NovelCommentEdges holds the relations/edges for other nodes in the graph.
type NovelCommentEdges struct {
	// Parent holds the value of the parent edge.
	Parent *NovelComment `json:"parent,omitempty"`
	// Childes holds the value of the childes edge.
	Childes []*NovelComment `json:"childes,omitempty"`
	// User holds the value of the user edge.
	User *SocialUser `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NovelCommentEdges) ParentOrErr() (*NovelComment, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: novelcomment.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildesOrErr returns the Childes value or an error if the edge
// was not loaded in eager-loading.
func (e NovelCommentEdges) ChildesOrErr() ([]*NovelComment, error) {
	if e.loadedTypes[1] {
		return e.Childes, nil
	}
	return nil, &NotLoadedError{edge: "childes"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NovelCommentEdges) UserOrErr() (*SocialUser, error) {
	if e.loadedTypes[2] {
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
func (*NovelComment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case novelcomment.FieldIsTop, novelcomment.FieldIsHighlight, novelcomment.FieldIsHot:
			values[i] = new(sql.NullBool)
		case novelcomment.FieldID, novelcomment.FieldNovelId, novelcomment.FieldUserId, novelcomment.FieldRepUserId, novelcomment.FieldScore, novelcomment.FieldPId, novelcomment.FieldCreateBy, novelcomment.FieldUpdateBy, novelcomment.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case novelcomment.FieldAvatar, novelcomment.FieldUserName, novelcomment.FieldRepUserName, novelcomment.FieldContent, novelcomment.FieldState, novelcomment.FieldRemark:
			values[i] = new(sql.NullString)
		case novelcomment.FieldCreatedAt, novelcomment.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case novelcomment.ForeignKeys[0]: // novel_comment_childes
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NovelComment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NovelComment fields.
func (nc *NovelComment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case novelcomment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nc.ID = int64(value.Int64)
		case novelcomment.FieldNovelId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field novelId", values[i])
			} else if value.Valid {
				nc.NovelId = value.Int64
			}
		case novelcomment.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				nc.UserId = value.Int64
			}
		case novelcomment.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				nc.Avatar = value.String
			}
		case novelcomment.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field userName", values[i])
			} else if value.Valid {
				nc.UserName = value.String
			}
		case novelcomment.FieldRepUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field repUserId", values[i])
			} else if value.Valid {
				nc.RepUserId = value.Int64
			}
		case novelcomment.FieldRepUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repUserName", values[i])
			} else if value.Valid {
				nc.RepUserName = value.String
			}
		case novelcomment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				nc.Content = value.String
			}
		case novelcomment.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				nc.Score = int32(value.Int64)
			}
		case novelcomment.FieldPId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pId", values[i])
			} else if value.Valid {
				nc.PId = value.Int64
			}
		case novelcomment.FieldIsTop:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isTop", values[i])
			} else if value.Valid {
				nc.IsTop = value.Bool
			}
		case novelcomment.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				nc.State = novelcomment.State(value.String)
			}
		case novelcomment.FieldIsHighlight:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isHighlight", values[i])
			} else if value.Valid {
				nc.IsHighlight = value.Bool
			}
		case novelcomment.FieldIsHot:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isHot", values[i])
			} else if value.Valid {
				nc.IsHot = value.Bool
			}
		case novelcomment.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				nc.Remark = value.String
			}
		case novelcomment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				nc.CreatedAt = value.Time
			}
		case novelcomment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				nc.UpdatedAt = value.Time
			}
		case novelcomment.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				nc.CreateBy = value.Int64
			}
		case novelcomment.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				nc.UpdateBy = value.Int64
			}
		case novelcomment.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				nc.TenantId = value.Int64
			}
		case novelcomment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field novel_comment_childes", value)
			} else if value.Valid {
				nc.novel_comment_childes = new(int64)
				*nc.novel_comment_childes = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the NovelComment entity.
func (nc *NovelComment) QueryParent() *NovelCommentQuery {
	return (&NovelCommentClient{config: nc.config}).QueryParent(nc)
}

// QueryChildes queries the "childes" edge of the NovelComment entity.
func (nc *NovelComment) QueryChildes() *NovelCommentQuery {
	return (&NovelCommentClient{config: nc.config}).QueryChildes(nc)
}

// QueryUser queries the "user" edge of the NovelComment entity.
func (nc *NovelComment) QueryUser() *SocialUserQuery {
	return (&NovelCommentClient{config: nc.config}).QueryUser(nc)
}

// Update returns a builder for updating this NovelComment.
// Note that you need to call NovelComment.Unwrap() before calling this method if this NovelComment
// was returned from a transaction, and the transaction was committed or rolled back.
func (nc *NovelComment) Update() *NovelCommentUpdateOne {
	return (&NovelCommentClient{config: nc.config}).UpdateOne(nc)
}

// Unwrap unwraps the NovelComment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nc *NovelComment) Unwrap() *NovelComment {
	tx, ok := nc.config.driver.(*txDriver)
	if !ok {
		panic("ent: NovelComment is not a transactional entity")
	}
	nc.config.driver = tx.drv
	return nc
}

// String implements the fmt.Stringer.
func (nc *NovelComment) String() string {
	var builder strings.Builder
	builder.WriteString("NovelComment(")
	builder.WriteString(fmt.Sprintf("id=%v", nc.ID))
	builder.WriteString(", novelId=")
	builder.WriteString(fmt.Sprintf("%v", nc.NovelId))
	builder.WriteString(", userId=")
	builder.WriteString(fmt.Sprintf("%v", nc.UserId))
	builder.WriteString(", avatar=")
	builder.WriteString(nc.Avatar)
	builder.WriteString(", userName=")
	builder.WriteString(nc.UserName)
	builder.WriteString(", repUserId=")
	builder.WriteString(fmt.Sprintf("%v", nc.RepUserId))
	builder.WriteString(", repUserName=")
	builder.WriteString(nc.RepUserName)
	builder.WriteString(", content=")
	builder.WriteString(nc.Content)
	builder.WriteString(", score=")
	builder.WriteString(fmt.Sprintf("%v", nc.Score))
	builder.WriteString(", pId=")
	builder.WriteString(fmt.Sprintf("%v", nc.PId))
	builder.WriteString(", isTop=")
	builder.WriteString(fmt.Sprintf("%v", nc.IsTop))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", nc.State))
	builder.WriteString(", isHighlight=")
	builder.WriteString(fmt.Sprintf("%v", nc.IsHighlight))
	builder.WriteString(", isHot=")
	builder.WriteString(fmt.Sprintf("%v", nc.IsHot))
	builder.WriteString(", remark=")
	builder.WriteString(nc.Remark)
	builder.WriteString(", createdAt=")
	builder.WriteString(nc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(nc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", nc.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", nc.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", nc.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// NovelComments is a parsable slice of NovelComment.
type NovelComments []*NovelComment

func (nc NovelComments) config(cfg config) {
	for _i := range nc {
		nc[_i].config = cfg
	}
}
