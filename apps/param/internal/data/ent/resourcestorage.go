// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"hope/apps/param/internal/data/ent/resourcestorage"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ResourceStorage is the model entity for the ResourceStorage schema.
type ResourceStorage struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// GroupId holds the value of the "groupId" field.
	// 分组
	GroupId int32 `json:"groupId,omitempty"`
	// StorageType holds the value of the "storageType" field.
	// 存储类型 1
	StorageType int32 `json:"storageType,omitempty"`
	// RealName holds the value of the "realName" field.
	// 文件真实的名称
	RealName string `json:"realName,omitempty"`
	// Bucket holds the value of the "bucket" field.
	// Bucket 识别符(七牛)
	Bucket string `json:"bucket,omitempty"`
	// Name holds the value of the "name" field.
	// 文件名称
	Name string `json:"name,omitempty"`
	// Suffix holds the value of the "suffix" field.
	// 后缀
	Suffix string `json:"suffix,omitempty"`
	// Path holds the value of the "path" field.
	// 路径
	Path string `json:"path,omitempty"`
	// Type holds the value of the "type" field.
	// 类型
	Type string `json:"type,omitempty"`
	// Size holds the value of the "size" field.
	// 大小
	Size string `json:"size,omitempty"`
	// DeleteUrl holds the value of the "deleteUrl" field.
	// (sm.sm)删除的URL
	DeleteUrl string `json:"deleteUrl,omitempty"`
	// Filename holds the value of the "filename" field.
	// (sm.sm)图片名称
	Filename string `json:"filename,omitempty"`
	// Key holds the value of the "key" field.
	// 文件名(七牛)
	Key string `json:"key,omitempty"`
	// Height holds the value of the "height" field.
	// (sm.sm)图片高度
	Height string `json:"height,omitempty"`
	// URL holds the value of the "url" field.
	// (sm.sm)图片地址
	URL string `json:"url,omitempty"`
	// Username holds the value of the "username" field.
	// (sm.sm)用户名称
	Username string `json:"username,omitempty"`
	// Width holds the value of the "width" field.
	// (sm.sm)图片宽度
	Width string `json:"width,omitempty"`
	// Md5code holds the value of the "md5code" field.
	// (sm.sm)文件的MD5值
	Md5code string `json:"md5code,omitempty"`
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
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResourceStorage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case resourcestorage.FieldID, resourcestorage.FieldGroupId, resourcestorage.FieldStorageType, resourcestorage.FieldCreateBy, resourcestorage.FieldUpdateBy, resourcestorage.FieldTenantId:
			values[i] = new(sql.NullInt64)
		case resourcestorage.FieldRealName, resourcestorage.FieldBucket, resourcestorage.FieldName, resourcestorage.FieldSuffix, resourcestorage.FieldPath, resourcestorage.FieldType, resourcestorage.FieldSize, resourcestorage.FieldDeleteUrl, resourcestorage.FieldFilename, resourcestorage.FieldKey, resourcestorage.FieldHeight, resourcestorage.FieldURL, resourcestorage.FieldUsername, resourcestorage.FieldWidth, resourcestorage.FieldMd5code, resourcestorage.FieldRemark:
			values[i] = new(sql.NullString)
		case resourcestorage.FieldCreatedAt, resourcestorage.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ResourceStorage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourceStorage fields.
func (rs *ResourceStorage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resourcestorage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rs.ID = int64(value.Int64)
		case resourcestorage.FieldGroupId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field groupId", values[i])
			} else if value.Valid {
				rs.GroupId = int32(value.Int64)
			}
		case resourcestorage.FieldStorageType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field storageType", values[i])
			} else if value.Valid {
				rs.StorageType = int32(value.Int64)
			}
		case resourcestorage.FieldRealName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field realName", values[i])
			} else if value.Valid {
				rs.RealName = value.String
			}
		case resourcestorage.FieldBucket:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bucket", values[i])
			} else if value.Valid {
				rs.Bucket = value.String
			}
		case resourcestorage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rs.Name = value.String
			}
		case resourcestorage.FieldSuffix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field suffix", values[i])
			} else if value.Valid {
				rs.Suffix = value.String
			}
		case resourcestorage.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				rs.Path = value.String
			}
		case resourcestorage.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				rs.Type = value.String
			}
		case resourcestorage.FieldSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				rs.Size = value.String
			}
		case resourcestorage.FieldDeleteUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleteUrl", values[i])
			} else if value.Valid {
				rs.DeleteUrl = value.String
			}
		case resourcestorage.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				rs.Filename = value.String
			}
		case resourcestorage.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				rs.Key = value.String
			}
		case resourcestorage.FieldHeight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				rs.Height = value.String
			}
		case resourcestorage.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				rs.URL = value.String
			}
		case resourcestorage.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				rs.Username = value.String
			}
		case resourcestorage.FieldWidth:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				rs.Width = value.String
			}
		case resourcestorage.FieldMd5code:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field md5code", values[i])
			} else if value.Valid {
				rs.Md5code = value.String
			}
		case resourcestorage.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				rs.Remark = value.String
			}
		case resourcestorage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				rs.CreatedAt = value.Time
			}
		case resourcestorage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				rs.UpdatedAt = value.Time
			}
		case resourcestorage.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field createBy", values[i])
			} else if value.Valid {
				rs.CreateBy = value.Int64
			}
		case resourcestorage.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updateBy", values[i])
			} else if value.Valid {
				rs.UpdateBy = value.Int64
			}
		case resourcestorage.FieldTenantId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenantId", values[i])
			} else if value.Valid {
				rs.TenantId = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ResourceStorage.
// Note that you need to call ResourceStorage.Unwrap() before calling this method if this ResourceStorage
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *ResourceStorage) Update() *ResourceStorageUpdateOne {
	return (&ResourceStorageClient{config: rs.config}).UpdateOne(rs)
}

// Unwrap unwraps the ResourceStorage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rs *ResourceStorage) Unwrap() *ResourceStorage {
	tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("ent: ResourceStorage is not a transactional entity")
	}
	rs.config.driver = tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *ResourceStorage) String() string {
	var builder strings.Builder
	builder.WriteString("ResourceStorage(")
	builder.WriteString(fmt.Sprintf("id=%v", rs.ID))
	builder.WriteString(", groupId=")
	builder.WriteString(fmt.Sprintf("%v", rs.GroupId))
	builder.WriteString(", storageType=")
	builder.WriteString(fmt.Sprintf("%v", rs.StorageType))
	builder.WriteString(", realName=")
	builder.WriteString(rs.RealName)
	builder.WriteString(", bucket=")
	builder.WriteString(rs.Bucket)
	builder.WriteString(", name=")
	builder.WriteString(rs.Name)
	builder.WriteString(", suffix=")
	builder.WriteString(rs.Suffix)
	builder.WriteString(", path=")
	builder.WriteString(rs.Path)
	builder.WriteString(", type=")
	builder.WriteString(rs.Type)
	builder.WriteString(", size=")
	builder.WriteString(rs.Size)
	builder.WriteString(", deleteUrl=")
	builder.WriteString(rs.DeleteUrl)
	builder.WriteString(", filename=")
	builder.WriteString(rs.Filename)
	builder.WriteString(", key=")
	builder.WriteString(rs.Key)
	builder.WriteString(", height=")
	builder.WriteString(rs.Height)
	builder.WriteString(", url=")
	builder.WriteString(rs.URL)
	builder.WriteString(", username=")
	builder.WriteString(rs.Username)
	builder.WriteString(", width=")
	builder.WriteString(rs.Width)
	builder.WriteString(", md5code=")
	builder.WriteString(rs.Md5code)
	builder.WriteString(", remark=")
	builder.WriteString(rs.Remark)
	builder.WriteString(", createdAt=")
	builder.WriteString(rs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(rs.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", createBy=")
	builder.WriteString(fmt.Sprintf("%v", rs.CreateBy))
	builder.WriteString(", updateBy=")
	builder.WriteString(fmt.Sprintf("%v", rs.UpdateBy))
	builder.WriteString(", tenantId=")
	builder.WriteString(fmt.Sprintf("%v", rs.TenantId))
	builder.WriteByte(')')
	return builder.String()
}

// ResourceStorages is a parsable slice of ResourceStorage.
type ResourceStorages []*ResourceStorage

func (rs ResourceStorages) config(cfg config) {
	for _i := range rs {
		rs[_i].config = cfg
	}
}
