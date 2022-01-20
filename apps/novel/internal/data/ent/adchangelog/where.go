// Code generated by entc, DO NOT EDIT.

package adchangelog

import (
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserId applies equality check predicate on the "userId" field. It's identical to UserIdEQ.
func UserId(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// AdId applies equality check predicate on the "adId" field. It's identical to AdIdEQ.
func AdId(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdId), v))
	})
}

// ChId applies equality check predicate on the "chId" field. It's identical to ChIdEQ.
func ChId(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChId), v))
	})
}

// DeviceId applies equality check predicate on the "deviceId" field. It's identical to DeviceIdEQ.
func DeviceId(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceId), v))
	})
}

// ExtInfo applies equality check predicate on the "extInfo" field. It's identical to ExtInfoEQ.
func ExtInfo(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtInfo), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreateBy applies equality check predicate on the "createBy" field. It's identical to CreateByEQ.
func CreateBy(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// UpdateBy applies equality check predicate on the "updateBy" field. It's identical to UpdateByEQ.
func UpdateBy(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// TenantId applies equality check predicate on the "tenantId" field. It's identical to TenantIdEQ.
func TenantId(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// UserIdEQ applies the EQ predicate on the "userId" field.
func UserIdEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserId), v))
	})
}

// UserIdNEQ applies the NEQ predicate on the "userId" field.
func UserIdNEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserId), v))
	})
}

// UserIdIn applies the In predicate on the "userId" field.
func UserIdIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func UserIdNotIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func UserIdGT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserId), v))
	})
}

// UserIdGTE applies the GTE predicate on the "userId" field.
func UserIdGTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserId), v))
	})
}

// UserIdLT applies the LT predicate on the "userId" field.
func UserIdLT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserId), v))
	})
}

// UserIdLTE applies the LTE predicate on the "userId" field.
func UserIdLTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserId), v))
	})
}

// UserIdIsNil applies the IsNil predicate on the "userId" field.
func UserIdIsNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserId)))
	})
}

// UserIdNotNil applies the NotNil predicate on the "userId" field.
func UserIdNotNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserId)))
	})
}

// AdIdEQ applies the EQ predicate on the "adId" field.
func AdIdEQ(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdId), v))
	})
}

// AdIdNEQ applies the NEQ predicate on the "adId" field.
func AdIdNEQ(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdId), v))
	})
}

// AdIdIn applies the In predicate on the "adId" field.
func AdIdIn(vs ...string) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdId), v...))
	})
}

// AdIdNotIn applies the NotIn predicate on the "adId" field.
func AdIdNotIn(vs ...string) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdId), v...))
	})
}

// AdIdGT applies the GT predicate on the "adId" field.
func AdIdGT(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdId), v))
	})
}

// AdIdGTE applies the GTE predicate on the "adId" field.
func AdIdGTE(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdId), v))
	})
}

// AdIdLT applies the LT predicate on the "adId" field.
func AdIdLT(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdId), v))
	})
}

// AdIdLTE applies the LTE predicate on the "adId" field.
func AdIdLTE(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdId), v))
	})
}

// AdIdContains applies the Contains predicate on the "adId" field.
func AdIdContains(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAdId), v))
	})
}

// AdIdHasPrefix applies the HasPrefix predicate on the "adId" field.
func AdIdHasPrefix(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAdId), v))
	})
}

// AdIdHasSuffix applies the HasSuffix predicate on the "adId" field.
func AdIdHasSuffix(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAdId), v))
	})
}

// AdIdIsNil applies the IsNil predicate on the "adId" field.
func AdIdIsNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAdId)))
	})
}

// AdIdNotNil applies the NotNil predicate on the "adId" field.
func AdIdNotNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAdId)))
	})
}

// AdIdEqualFold applies the EqualFold predicate on the "adId" field.
func AdIdEqualFold(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAdId), v))
	})
}

// AdIdContainsFold applies the ContainsFold predicate on the "adId" field.
func AdIdContainsFold(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAdId), v))
	})
}

// ChIdEQ applies the EQ predicate on the "chId" field.
func ChIdEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChId), v))
	})
}

// ChIdNEQ applies the NEQ predicate on the "chId" field.
func ChIdNEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChId), v))
	})
}

// ChIdIn applies the In predicate on the "chId" field.
func ChIdIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChId), v...))
	})
}

// ChIdNotIn applies the NotIn predicate on the "chId" field.
func ChIdNotIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChId), v...))
	})
}

// ChIdGT applies the GT predicate on the "chId" field.
func ChIdGT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChId), v))
	})
}

// ChIdGTE applies the GTE predicate on the "chId" field.
func ChIdGTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChId), v))
	})
}

// ChIdLT applies the LT predicate on the "chId" field.
func ChIdLT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChId), v))
	})
}

// ChIdLTE applies the LTE predicate on the "chId" field.
func ChIdLTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChId), v))
	})
}

// ChIdIsNil applies the IsNil predicate on the "chId" field.
func ChIdIsNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChId)))
	})
}

// ChIdNotNil applies the NotNil predicate on the "chId" field.
func ChIdNotNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChId)))
	})
}

// DeviceIdEQ applies the EQ predicate on the "deviceId" field.
func DeviceIdEQ(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceId), v))
	})
}

// DeviceIdNEQ applies the NEQ predicate on the "deviceId" field.
func DeviceIdNEQ(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeviceId), v))
	})
}

// DeviceIdIn applies the In predicate on the "deviceId" field.
func DeviceIdIn(vs ...string) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeviceId), v...))
	})
}

// DeviceIdNotIn applies the NotIn predicate on the "deviceId" field.
func DeviceIdNotIn(vs ...string) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeviceId), v...))
	})
}

// DeviceIdGT applies the GT predicate on the "deviceId" field.
func DeviceIdGT(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeviceId), v))
	})
}

// DeviceIdGTE applies the GTE predicate on the "deviceId" field.
func DeviceIdGTE(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeviceId), v))
	})
}

// DeviceIdLT applies the LT predicate on the "deviceId" field.
func DeviceIdLT(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeviceId), v))
	})
}

// DeviceIdLTE applies the LTE predicate on the "deviceId" field.
func DeviceIdLTE(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeviceId), v))
	})
}

// DeviceIdContains applies the Contains predicate on the "deviceId" field.
func DeviceIdContains(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDeviceId), v))
	})
}

// DeviceIdHasPrefix applies the HasPrefix predicate on the "deviceId" field.
func DeviceIdHasPrefix(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDeviceId), v))
	})
}

// DeviceIdHasSuffix applies the HasSuffix predicate on the "deviceId" field.
func DeviceIdHasSuffix(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDeviceId), v))
	})
}

// DeviceIdIsNil applies the IsNil predicate on the "deviceId" field.
func DeviceIdIsNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeviceId)))
	})
}

// DeviceIdNotNil applies the NotNil predicate on the "deviceId" field.
func DeviceIdNotNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeviceId)))
	})
}

// DeviceIdEqualFold applies the EqualFold predicate on the "deviceId" field.
func DeviceIdEqualFold(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDeviceId), v))
	})
}

// DeviceIdContainsFold applies the ContainsFold predicate on the "deviceId" field.
func DeviceIdContainsFold(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDeviceId), v))
	})
}

// ExtInfoEQ applies the EQ predicate on the "extInfo" field.
func ExtInfoEQ(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtInfo), v))
	})
}

// ExtInfoNEQ applies the NEQ predicate on the "extInfo" field.
func ExtInfoNEQ(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExtInfo), v))
	})
}

// ExtInfoIn applies the In predicate on the "extInfo" field.
func ExtInfoIn(vs ...string) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExtInfo), v...))
	})
}

// ExtInfoNotIn applies the NotIn predicate on the "extInfo" field.
func ExtInfoNotIn(vs ...string) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExtInfo), v...))
	})
}

// ExtInfoGT applies the GT predicate on the "extInfo" field.
func ExtInfoGT(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExtInfo), v))
	})
}

// ExtInfoGTE applies the GTE predicate on the "extInfo" field.
func ExtInfoGTE(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExtInfo), v))
	})
}

// ExtInfoLT applies the LT predicate on the "extInfo" field.
func ExtInfoLT(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExtInfo), v))
	})
}

// ExtInfoLTE applies the LTE predicate on the "extInfo" field.
func ExtInfoLTE(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExtInfo), v))
	})
}

// ExtInfoContains applies the Contains predicate on the "extInfo" field.
func ExtInfoContains(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExtInfo), v))
	})
}

// ExtInfoHasPrefix applies the HasPrefix predicate on the "extInfo" field.
func ExtInfoHasPrefix(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExtInfo), v))
	})
}

// ExtInfoHasSuffix applies the HasSuffix predicate on the "extInfo" field.
func ExtInfoHasSuffix(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExtInfo), v))
	})
}

// ExtInfoIsNil applies the IsNil predicate on the "extInfo" field.
func ExtInfoIsNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExtInfo)))
	})
}

// ExtInfoNotNil applies the NotNil predicate on the "extInfo" field.
func ExtInfoNotNil() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExtInfo)))
	})
}

// ExtInfoEqualFold applies the EqualFold predicate on the "extInfo" field.
func ExtInfoEqualFold(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExtInfo), v))
	})
}

// ExtInfoContainsFold applies the ContainsFold predicate on the "extInfo" field.
func ExtInfoContainsFold(v string) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExtInfo), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreateByEQ applies the EQ predicate on the "createBy" field.
func CreateByEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// CreateByNEQ applies the NEQ predicate on the "createBy" field.
func CreateByNEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateBy), v))
	})
}

// CreateByIn applies the In predicate on the "createBy" field.
func CreateByIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func CreateByNotIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func CreateByGT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateBy), v))
	})
}

// CreateByGTE applies the GTE predicate on the "createBy" field.
func CreateByGTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateBy), v))
	})
}

// CreateByLT applies the LT predicate on the "createBy" field.
func CreateByLT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateBy), v))
	})
}

// CreateByLTE applies the LTE predicate on the "createBy" field.
func CreateByLTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateBy), v))
	})
}

// UpdateByEQ applies the EQ predicate on the "updateBy" field.
func UpdateByEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByNEQ applies the NEQ predicate on the "updateBy" field.
func UpdateByNEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByIn applies the In predicate on the "updateBy" field.
func UpdateByIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func UpdateByNotIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func UpdateByGT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByGTE applies the GTE predicate on the "updateBy" field.
func UpdateByGTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLT applies the LT predicate on the "updateBy" field.
func UpdateByLT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLTE applies the LTE predicate on the "updateBy" field.
func UpdateByLTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateBy), v))
	})
}

// TenantIdEQ applies the EQ predicate on the "tenantId" field.
func TenantIdEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// TenantIdNEQ applies the NEQ predicate on the "tenantId" field.
func TenantIdNEQ(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantId), v))
	})
}

// TenantIdIn applies the In predicate on the "tenantId" field.
func TenantIdIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func TenantIdNotIn(vs ...int64) predicate.AdChangeLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func TenantIdGT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenantId), v))
	})
}

// TenantIdGTE applies the GTE predicate on the "tenantId" field.
func TenantIdGTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenantId), v))
	})
}

// TenantIdLT applies the LT predicate on the "tenantId" field.
func TenantIdLT(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenantId), v))
	})
}

// TenantIdLTE applies the LTE predicate on the "tenantId" field.
func TenantIdLTE(v int64) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenantId), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.SocialUser) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AdChangeLog) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AdChangeLog) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
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
func Not(p predicate.AdChangeLog) predicate.AdChangeLog {
	return predicate.AdChangeLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
