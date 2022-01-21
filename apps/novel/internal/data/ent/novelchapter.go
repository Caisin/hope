// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/novel/internal/data/ent/novel"
	"hope/apps/novel/internal/data/ent/novelchapter"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// NovelChapter is the model entity for the NovelChapter schema.
type NovelChapter struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// NovelId holds the value of the "novelId" field.
	// 小说编号
	NovelId int64 `json:"novelId,omitempty"`
	// OrderNum holds the value of the "orderNum" field.
	// 章节序号
	OrderNum int32 `json:"orderNum,omitempty"`
	// ChapterName holds the value of the "chapterName" field.
	// 章节名称
	ChapterName string `json:"chapterName,omitempty"`
	// Content holds the value of the "content" field.
	// 章节内容
	Content string `json:"content,omitempty"`
	// MediaKey holds the value of the "mediaKey" field.
	// 阿里云音频目录
	MediaKey string `json:"mediaKey,omitempty"`
	// Duration holds the value of the "duration" field.
	// 音频时长
	Duration string `json:"duration,omitempty"`
	// PublishTime holds the value of the "publishTime" field.
	// 发布时间
	PublishTime time.Time `json:"publishTime,omitempty"`
	// Status holds the value of the "status" field.
	// 状态：0 草稿 ，1 发布
	Status int32 `json:"status,omitempty"`
	// IsFree holds the value of the "isFree" field.
	// 0
	IsFree bool `json:"isFree,omitempty"`
	// Price holds the value of the "price" field.
	// 价格
	Price int64 `json:"price,omitempty"`
	// WordNum holds the value of the "wordNum" field.
	// 章节字数
	WordNum int32 `json:"wordNum,omitempty"`
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
	// The values are being populated by the NovelChapterQuery when eager-loading is set.
	Edges              NovelChapterEdges `json:"edges"`
	novel_chapter_next *int64
}

// NovelChapterEdges holds the relations/edges for other nodes in the graph.
type NovelChapterEdges struct {
	// Prev holds the value of the prev edge.
	Prev *NovelChapter `json:"prev,omitempty"`
	// Next holds the value of the next edge.
	Next *NovelChapter `json:"next,omitempty"`
	// Novel holds the value of the novel edge.
	Novel *Novel `json:"novel,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PrevOrErr returns the Prev value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NovelChapterEdges) PrevOrErr() (*NovelChapter, error) {
	if e.loadedTypes[0] {
		if e.Prev == nil {
			// The edge prev was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: novelchapter.Label}
		}
		return e.Prev, nil
	}
	return nil, &NotLoadedError{edge: "prev"}
}

// NextOrErr returns the Next value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NovelChapterEdges) NextOrErr() (*NovelChapter, error) {
	if e.loadedTypes[1] {
		if e.Next == nil {
			// The edge next was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: novelchapter.Label}
		}
		return e.Next, nil
	}
	return nil, &NotLoadedError{edge: "next"}
}

// NovelOrErr returns the Novel value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NovelChapterEdges) NovelOrErr() (*Novel, error) {
	if e.loadedTypes[2] {
		if e.Novel == nil {
			// The edge novel was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: novel.Label}
		}
		return e.Novel, nil
	}
	return nil, &NotLoadedError{edge: "novel"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NovelChapter) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case novelchapter.FieldIsFree:
			values[i] = new(sql.NullBool)
		case novelchapter.FieldID, novelchapter.FieldNovelId, novelchapter.FieldOrderNum, novelchapter.FieldStatus, novelchapter.FieldPrice, novelchapter.FieldWordNum, novelchapter.FieldCreateBy, novelchapter.FieldUpdateBy, novelchapter.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case novelchapter.FieldChapterName, novelchapter.FieldContent, novelchapter.FieldMediaKey, novelchapter.FieldDuration, novelchapter.FieldRemark:
			values[i] = new(sql.NullString)
		case novelchapter.FieldPublishTime, novelchapter.FieldCreatedAt, novelchapter.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case novelchapter.ForeignKeys[0]: // novel_chapter_next
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type NovelChapter", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NovelChapter fields.
func (nc *NovelChapter) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case novelchapter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			nc.ID = int64(value.Int64)
		case novelchapter.FieldNovelId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field novelId", values[i])
			} else if value.Valid {
				nc.NovelId = value.Int64
			}
		case novelchapter.FieldOrderNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field orderNum", values[i])
			} else if value.Valid {
				nc.OrderNum = int32(value.Int64)
			}
		case novelchapter.FieldChapterName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chapterName", values[i])
			} else if value.Valid {
				nc.ChapterName = value.String
			}
		case novelchapter.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				nc.Content = value.String
			}
		case novelchapter.FieldMediaKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mediaKey", values[i])
			} else if value.Valid {
				nc.MediaKey = value.String
			}
		case novelchapter.FieldDuration:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				nc.Duration = value.String
			}
		case novelchapter.FieldPublishTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field publishTime", values[i])
			} else if value.Valid {
				nc.PublishTime = value.Time
			}
		case novelchapter.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				nc.Status = int32(value.Int64)
			}
		case novelchapter.FieldIsFree:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field isFree", values[i])
			} else if value.Valid {
				nc.IsFree = value.Bool
			}
		case novelchapter.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				nc.Price = value.Int64
			}
		case novelchapter.FieldWordNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wordNum", values[i])
			} else if value.Valid {
				nc.WordNum = int32(value.Int64)
			}
		case novelchapter.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				nc.Remark = value.String
			}
		case novelchapter.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				nc.CreatedAt = value.Time
			}
		case novelchapter.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				nc.UpdatedAt = value.Time
			}
		case novelchapter.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				nc.CreateBy = value.Int64
			}
		case novelchapter.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				nc.UpdateBy = value.Int64
			}
		case novelchapter.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				nc.TenantId = value.Int64
			}
		case novelchapter.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field novel_chapter_next", value)
			} else if value.Valid {
				nc.novel_chapter_next = new(int64)
				*nc.novel_chapter_next = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryPrev queries the "prev" edge of the NovelChapter entity.
func (nc *NovelChapter) QueryPrev() *NovelChapterQuery {
	return (&NovelChapterClient{config: nc.config}).QueryPrev(nc)
}

// QueryNext queries the "next" edge of the NovelChapter entity.
func (nc *NovelChapter) QueryNext() *NovelChapterQuery {
	return (&NovelChapterClient{config: nc.config}).QueryNext(nc)
}

// QueryNovel queries the "novel" edge of the NovelChapter entity.
func (nc *NovelChapter) QueryNovel() *NovelQuery {
	return (&NovelChapterClient{config: nc.config}).QueryNovel(nc)
}

// Update returns a builder for updating this NovelChapter.
// Note that you need to call NovelChapter.Unwrap() before calling this method if this NovelChapter
// was returned from a transaction, and the transaction was committed or rolled back.
func (nc *NovelChapter) Update() *NovelChapterUpdateOne {
	return (&NovelChapterClient{config: nc.config}).UpdateOne(nc)
}

// Unwrap unwraps the NovelChapter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nc *NovelChapter) Unwrap() *NovelChapter {
	tx, ok := nc.config.driver.(*txDriver)
	if !ok {
		panic("ent: NovelChapter is not a transactional entity")
	}
	nc.config.driver = tx.drv
	return nc
}

// String implements the fmt.Stringer.
func (nc *NovelChapter) String() string {
	var builder strings.Builder
	builder.WriteString("NovelChapter(")
	builder.WriteString(fmt.Sprintf("id=%v", nc.ID))
	builder.WriteString(", novelId=")
	builder.WriteString(fmt.Sprintf("%v", nc.NovelId))
	builder.WriteString(", orderNum=")
	builder.WriteString(fmt.Sprintf("%v", nc.OrderNum))
	builder.WriteString(", chapterName=")
	builder.WriteString(nc.ChapterName)
	builder.WriteString(", content=")
	builder.WriteString(nc.Content)
	builder.WriteString(", mediaKey=")
	builder.WriteString(nc.MediaKey)
	builder.WriteString(", duration=")
	builder.WriteString(nc.Duration)
	builder.WriteString(", publishTime=")
	builder.WriteString(nc.PublishTime.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", nc.Status))
	builder.WriteString(", isFree=")
	builder.WriteString(fmt.Sprintf("%v", nc.IsFree))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", nc.Price))
	builder.WriteString(", wordNum=")
	builder.WriteString(fmt.Sprintf("%v", nc.WordNum))
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

// NovelChapters is a parsable slice of NovelChapter.
type NovelChapters []*NovelChapter

func (nc NovelChapters) config(cfg config) {
	for _i := range nc {
		nc[_i].config = cfg
	}
}
