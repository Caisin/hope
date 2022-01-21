// Code generated by entc, DO NOT EDIT.

package clienterror

import (
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppVersion applies equality check predicate on the "appVersion" field. It's identical to AppVersionEQ.
func AppVersion(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppVersion), v))
	})
}

// DeviceName applies equality check predicate on the "deviceName" field. It's identical to DeviceNameEQ.
func DeviceName(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceName), v))
	})
}

// OsName applies equality check predicate on the "osName" field. It's identical to OsNameEQ.
func OsName(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOsName), v))
	})
}

// ErrorInfo applies equality check predicate on the "errorInfo" field. It's identical to ErrorInfoEQ.
func ErrorInfo(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorInfo), v))
	})
}

// UserId applies equality check predicate on the "userId" field. It's identical to UserIdEQ.
func UserId(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreateBy applies equality check predicate on the "createBy" field. It's identical to CreateByEQ.
func CreateBy(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// UpdateBy applies equality check predicate on the "updateBy" field. It's identical to UpdateByEQ.
func UpdateBy(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// TenantId applies equality check predicate on the "tenantId" field. It's identical to TenantIdEQ.
func TenantId(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// AppVersionEQ applies the EQ predicate on the "appVersion" field.
func AppVersionEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppVersion), v))
	})
}

// AppVersionNEQ applies the NEQ predicate on the "appVersion" field.
func AppVersionNEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppVersion), v))
	})
}

// AppVersionIn applies the In predicate on the "appVersion" field.
func AppVersionIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppVersion), v...))
	})
}

// AppVersionNotIn applies the NotIn predicate on the "appVersion" field.
func AppVersionNotIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppVersion), v...))
	})
}

// AppVersionGT applies the GT predicate on the "appVersion" field.
func AppVersionGT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppVersion), v))
	})
}

// AppVersionGTE applies the GTE predicate on the "appVersion" field.
func AppVersionGTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppVersion), v))
	})
}

// AppVersionLT applies the LT predicate on the "appVersion" field.
func AppVersionLT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppVersion), v))
	})
}

// AppVersionLTE applies the LTE predicate on the "appVersion" field.
func AppVersionLTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppVersion), v))
	})
}

// AppVersionContains applies the Contains predicate on the "appVersion" field.
func AppVersionContains(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAppVersion), v))
	})
}

// AppVersionHasPrefix applies the HasPrefix predicate on the "appVersion" field.
func AppVersionHasPrefix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAppVersion), v))
	})
}

// AppVersionHasSuffix applies the HasSuffix predicate on the "appVersion" field.
func AppVersionHasSuffix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAppVersion), v))
	})
}

// AppVersionIsNil applies the IsNil predicate on the "appVersion" field.
func AppVersionIsNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppVersion)))
	})
}

// AppVersionNotNil applies the NotNil predicate on the "appVersion" field.
func AppVersionNotNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppVersion)))
	})
}

// AppVersionEqualFold applies the EqualFold predicate on the "appVersion" field.
func AppVersionEqualFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAppVersion), v))
	})
}

// AppVersionContainsFold applies the ContainsFold predicate on the "appVersion" field.
func AppVersionContainsFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAppVersion), v))
	})
}

// DeviceNameEQ applies the EQ predicate on the "deviceName" field.
func DeviceNameEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceName), v))
	})
}

// DeviceNameNEQ applies the NEQ predicate on the "deviceName" field.
func DeviceNameNEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeviceName), v))
	})
}

// DeviceNameIn applies the In predicate on the "deviceName" field.
func DeviceNameIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeviceName), v...))
	})
}

// DeviceNameNotIn applies the NotIn predicate on the "deviceName" field.
func DeviceNameNotIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeviceName), v...))
	})
}

// DeviceNameGT applies the GT predicate on the "deviceName" field.
func DeviceNameGT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeviceName), v))
	})
}

// DeviceNameGTE applies the GTE predicate on the "deviceName" field.
func DeviceNameGTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeviceName), v))
	})
}

// DeviceNameLT applies the LT predicate on the "deviceName" field.
func DeviceNameLT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeviceName), v))
	})
}

// DeviceNameLTE applies the LTE predicate on the "deviceName" field.
func DeviceNameLTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeviceName), v))
	})
}

// DeviceNameContains applies the Contains predicate on the "deviceName" field.
func DeviceNameContains(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDeviceName), v))
	})
}

// DeviceNameHasPrefix applies the HasPrefix predicate on the "deviceName" field.
func DeviceNameHasPrefix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDeviceName), v))
	})
}

// DeviceNameHasSuffix applies the HasSuffix predicate on the "deviceName" field.
func DeviceNameHasSuffix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDeviceName), v))
	})
}

// DeviceNameIsNil applies the IsNil predicate on the "deviceName" field.
func DeviceNameIsNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeviceName)))
	})
}

// DeviceNameNotNil applies the NotNil predicate on the "deviceName" field.
func DeviceNameNotNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeviceName)))
	})
}

// DeviceNameEqualFold applies the EqualFold predicate on the "deviceName" field.
func DeviceNameEqualFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDeviceName), v))
	})
}

// DeviceNameContainsFold applies the ContainsFold predicate on the "deviceName" field.
func DeviceNameContainsFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDeviceName), v))
	})
}

// OsNameEQ applies the EQ predicate on the "osName" field.
func OsNameEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOsName), v))
	})
}

// OsNameNEQ applies the NEQ predicate on the "osName" field.
func OsNameNEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOsName), v))
	})
}

// OsNameIn applies the In predicate on the "osName" field.
func OsNameIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOsName), v...))
	})
}

// OsNameNotIn applies the NotIn predicate on the "osName" field.
func OsNameNotIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOsName), v...))
	})
}

// OsNameGT applies the GT predicate on the "osName" field.
func OsNameGT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOsName), v))
	})
}

// OsNameGTE applies the GTE predicate on the "osName" field.
func OsNameGTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOsName), v))
	})
}

// OsNameLT applies the LT predicate on the "osName" field.
func OsNameLT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOsName), v))
	})
}

// OsNameLTE applies the LTE predicate on the "osName" field.
func OsNameLTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOsName), v))
	})
}

// OsNameContains applies the Contains predicate on the "osName" field.
func OsNameContains(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOsName), v))
	})
}

// OsNameHasPrefix applies the HasPrefix predicate on the "osName" field.
func OsNameHasPrefix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOsName), v))
	})
}

// OsNameHasSuffix applies the HasSuffix predicate on the "osName" field.
func OsNameHasSuffix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOsName), v))
	})
}

// OsNameIsNil applies the IsNil predicate on the "osName" field.
func OsNameIsNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOsName)))
	})
}

// OsNameNotNil applies the NotNil predicate on the "osName" field.
func OsNameNotNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOsName)))
	})
}

// OsNameEqualFold applies the EqualFold predicate on the "osName" field.
func OsNameEqualFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOsName), v))
	})
}

// OsNameContainsFold applies the ContainsFold predicate on the "osName" field.
func OsNameContainsFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOsName), v))
	})
}

// ErrorInfoEQ applies the EQ predicate on the "errorInfo" field.
func ErrorInfoEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoNEQ applies the NEQ predicate on the "errorInfo" field.
func ErrorInfoNEQ(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoIn applies the In predicate on the "errorInfo" field.
func ErrorInfoIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldErrorInfo), v...))
	})
}

// ErrorInfoNotIn applies the NotIn predicate on the "errorInfo" field.
func ErrorInfoNotIn(vs ...string) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldErrorInfo), v...))
	})
}

// ErrorInfoGT applies the GT predicate on the "errorInfo" field.
func ErrorInfoGT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoGTE applies the GTE predicate on the "errorInfo" field.
func ErrorInfoGTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoLT applies the LT predicate on the "errorInfo" field.
func ErrorInfoLT(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoLTE applies the LTE predicate on the "errorInfo" field.
func ErrorInfoLTE(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoContains applies the Contains predicate on the "errorInfo" field.
func ErrorInfoContains(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoHasPrefix applies the HasPrefix predicate on the "errorInfo" field.
func ErrorInfoHasPrefix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoHasSuffix applies the HasSuffix predicate on the "errorInfo" field.
func ErrorInfoHasSuffix(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoIsNil applies the IsNil predicate on the "errorInfo" field.
func ErrorInfoIsNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldErrorInfo)))
	})
}

// ErrorInfoNotNil applies the NotNil predicate on the "errorInfo" field.
func ErrorInfoNotNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldErrorInfo)))
	})
}

// ErrorInfoEqualFold applies the EqualFold predicate on the "errorInfo" field.
func ErrorInfoEqualFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldErrorInfo), v))
	})
}

// ErrorInfoContainsFold applies the ContainsFold predicate on the "errorInfo" field.
func ErrorInfoContainsFold(v string) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldErrorInfo), v))
	})
}

// UserIdEQ applies the EQ predicate on the "userId" field.
func UserIdEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// UserIdNEQ applies the NEQ predicate on the "userId" field.
func UserIdNEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserId), v))
	})
}

// UserIdIn applies the In predicate on the "userId" field.
func UserIdIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserId), v...))
	})
}

// UserIdNotIn applies the NotIn predicate on the "userId" field.
func UserIdNotIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserId), v...))
	})
}

// UserIdGT applies the GT predicate on the "userId" field.
func UserIdGT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserId), v))
	})
}

// UserIdGTE applies the GTE predicate on the "userId" field.
func UserIdGTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserId), v))
	})
}

// UserIdLT applies the LT predicate on the "userId" field.
func UserIdLT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserId), v))
	})
}

// UserIdLTE applies the LTE predicate on the "userId" field.
func UserIdLTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserId), v))
	})
}

// UserIdIsNil applies the IsNil predicate on the "userId" field.
func UserIdIsNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserId)))
	})
}

// UserIdNotNil applies the NotNil predicate on the "userId" field.
func UserIdNotNil() predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserId)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreateByEQ applies the EQ predicate on the "createBy" field.
func CreateByEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// CreateByNEQ applies the NEQ predicate on the "createBy" field.
func CreateByNEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateBy), v))
	})
}

// CreateByIn applies the In predicate on the "createBy" field.
func CreateByIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateBy), v...))
	})
}

// CreateByNotIn applies the NotIn predicate on the "createBy" field.
func CreateByNotIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateBy), v...))
	})
}

// CreateByGT applies the GT predicate on the "createBy" field.
func CreateByGT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateBy), v))
	})
}

// CreateByGTE applies the GTE predicate on the "createBy" field.
func CreateByGTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateBy), v))
	})
}

// CreateByLT applies the LT predicate on the "createBy" field.
func CreateByLT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateBy), v))
	})
}

// CreateByLTE applies the LTE predicate on the "createBy" field.
func CreateByLTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateBy), v))
	})
}

// UpdateByEQ applies the EQ predicate on the "updateBy" field.
func UpdateByEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByNEQ applies the NEQ predicate on the "updateBy" field.
func UpdateByNEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByIn applies the In predicate on the "updateBy" field.
func UpdateByIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateBy), v...))
	})
}

// UpdateByNotIn applies the NotIn predicate on the "updateBy" field.
func UpdateByNotIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateBy), v...))
	})
}

// UpdateByGT applies the GT predicate on the "updateBy" field.
func UpdateByGT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByGTE applies the GTE predicate on the "updateBy" field.
func UpdateByGTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLT applies the LT predicate on the "updateBy" field.
func UpdateByLT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLTE applies the LTE predicate on the "updateBy" field.
func UpdateByLTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateBy), v))
	})
}

// TenantIdEQ applies the EQ predicate on the "tenantId" field.
func TenantIdEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// TenantIdNEQ applies the NEQ predicate on the "tenantId" field.
func TenantIdNEQ(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantId), v))
	})
}

// TenantIdIn applies the In predicate on the "tenantId" field.
func TenantIdIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTenantId), v...))
	})
}

// TenantIdNotIn applies the NotIn predicate on the "tenantId" field.
func TenantIdNotIn(vs ...int64) predicate.ClientError {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ClientError(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTenantId), v...))
	})
}

// TenantIdGT applies the GT predicate on the "tenantId" field.
func TenantIdGT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenantId), v))
	})
}

// TenantIdGTE applies the GTE predicate on the "tenantId" field.
func TenantIdGTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenantId), v))
	})
}

// TenantIdLT applies the LT predicate on the "tenantId" field.
func TenantIdLT(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenantId), v))
	})
}

// TenantIdLTE applies the LTE predicate on the "tenantId" field.
func TenantIdLTE(v int64) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenantId), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ClientError) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ClientError) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ClientError) predicate.ClientError {
	return predicate.ClientError(func(s *sql.Selector) {
		p(s.Not())
	})
}