// Code generated by entc, DO NOT EDIT.

package viptype

import (
	"hope/apps/param/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// VipName applies equality check predicate on the "vipName" field. It's identical to VipNameEQ.
func VipName(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVipName), v))
	})
}

// IsSuper applies equality check predicate on the "isSuper" field. It's identical to IsSuperEQ.
func IsSuper(v bool) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsSuper), v))
	})
}

// ValidDays applies equality check predicate on the "validDays" field. It's identical to ValidDaysEQ.
func ValidDays(v int32) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidDays), v))
	})
}

// DiscountRate applies equality check predicate on the "discountRate" field. It's identical to DiscountRateEQ.
func DiscountRate(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountRate), v))
	})
}

// AvatarId applies equality check predicate on the "avatarId" field. It's identical to AvatarIdEQ.
func AvatarId(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatarId), v))
	})
}

// Summary applies equality check predicate on the "summary" field. It's identical to SummaryEQ.
func Summary(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreateBy applies equality check predicate on the "createBy" field. It's identical to CreateByEQ.
func CreateBy(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// UpdateBy applies equality check predicate on the "updateBy" field. It's identical to UpdateByEQ.
func UpdateBy(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// TenantId applies equality check predicate on the "tenantId" field. It's identical to TenantIdEQ.
func TenantId(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// VipNameEQ applies the EQ predicate on the "vipName" field.
func VipNameEQ(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVipName), v))
	})
}

// VipNameNEQ applies the NEQ predicate on the "vipName" field.
func VipNameNEQ(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVipName), v))
	})
}

// VipNameIn applies the In predicate on the "vipName" field.
func VipNameIn(vs ...string) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVipName), v...))
	})
}

// VipNameNotIn applies the NotIn predicate on the "vipName" field.
func VipNameNotIn(vs ...string) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVipName), v...))
	})
}

// VipNameGT applies the GT predicate on the "vipName" field.
func VipNameGT(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVipName), v))
	})
}

// VipNameGTE applies the GTE predicate on the "vipName" field.
func VipNameGTE(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVipName), v))
	})
}

// VipNameLT applies the LT predicate on the "vipName" field.
func VipNameLT(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVipName), v))
	})
}

// VipNameLTE applies the LTE predicate on the "vipName" field.
func VipNameLTE(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVipName), v))
	})
}

// VipNameContains applies the Contains predicate on the "vipName" field.
func VipNameContains(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVipName), v))
	})
}

// VipNameHasPrefix applies the HasPrefix predicate on the "vipName" field.
func VipNameHasPrefix(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVipName), v))
	})
}

// VipNameHasSuffix applies the HasSuffix predicate on the "vipName" field.
func VipNameHasSuffix(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVipName), v))
	})
}

// VipNameIsNil applies the IsNil predicate on the "vipName" field.
func VipNameIsNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldVipName)))
	})
}

// VipNameNotNil applies the NotNil predicate on the "vipName" field.
func VipNameNotNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldVipName)))
	})
}

// VipNameEqualFold applies the EqualFold predicate on the "vipName" field.
func VipNameEqualFold(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVipName), v))
	})
}

// VipNameContainsFold applies the ContainsFold predicate on the "vipName" field.
func VipNameContainsFold(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVipName), v))
	})
}

// IsSuperEQ applies the EQ predicate on the "isSuper" field.
func IsSuperEQ(v bool) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsSuper), v))
	})
}

// IsSuperNEQ applies the NEQ predicate on the "isSuper" field.
func IsSuperNEQ(v bool) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsSuper), v))
	})
}

// IsSuperIsNil applies the IsNil predicate on the "isSuper" field.
func IsSuperIsNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIsSuper)))
	})
}

// IsSuperNotNil applies the NotNil predicate on the "isSuper" field.
func IsSuperNotNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIsSuper)))
	})
}

// ValidDaysEQ applies the EQ predicate on the "validDays" field.
func ValidDaysEQ(v int32) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidDays), v))
	})
}

// ValidDaysNEQ applies the NEQ predicate on the "validDays" field.
func ValidDaysNEQ(v int32) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValidDays), v))
	})
}

// ValidDaysIn applies the In predicate on the "validDays" field.
func ValidDaysIn(vs ...int32) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValidDays), v...))
	})
}

// ValidDaysNotIn applies the NotIn predicate on the "validDays" field.
func ValidDaysNotIn(vs ...int32) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValidDays), v...))
	})
}

// ValidDaysGT applies the GT predicate on the "validDays" field.
func ValidDaysGT(v int32) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValidDays), v))
	})
}

// ValidDaysGTE applies the GTE predicate on the "validDays" field.
func ValidDaysGTE(v int32) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValidDays), v))
	})
}

// ValidDaysLT applies the LT predicate on the "validDays" field.
func ValidDaysLT(v int32) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValidDays), v))
	})
}

// ValidDaysLTE applies the LTE predicate on the "validDays" field.
func ValidDaysLTE(v int32) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValidDays), v))
	})
}

// ValidDaysIsNil applies the IsNil predicate on the "validDays" field.
func ValidDaysIsNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValidDays)))
	})
}

// ValidDaysNotNil applies the NotNil predicate on the "validDays" field.
func ValidDaysNotNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValidDays)))
	})
}

// DiscountRateEQ applies the EQ predicate on the "discountRate" field.
func DiscountRateEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountRate), v))
	})
}

// DiscountRateNEQ applies the NEQ predicate on the "discountRate" field.
func DiscountRateNEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscountRate), v))
	})
}

// DiscountRateIn applies the In predicate on the "discountRate" field.
func DiscountRateIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscountRate), v...))
	})
}

// DiscountRateNotIn applies the NotIn predicate on the "discountRate" field.
func DiscountRateNotIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscountRate), v...))
	})
}

// DiscountRateGT applies the GT predicate on the "discountRate" field.
func DiscountRateGT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscountRate), v))
	})
}

// DiscountRateGTE applies the GTE predicate on the "discountRate" field.
func DiscountRateGTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscountRate), v))
	})
}

// DiscountRateLT applies the LT predicate on the "discountRate" field.
func DiscountRateLT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscountRate), v))
	})
}

// DiscountRateLTE applies the LTE predicate on the "discountRate" field.
func DiscountRateLTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscountRate), v))
	})
}

// DiscountRateIsNil applies the IsNil predicate on the "discountRate" field.
func DiscountRateIsNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDiscountRate)))
	})
}

// DiscountRateNotNil applies the NotNil predicate on the "discountRate" field.
func DiscountRateNotNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDiscountRate)))
	})
}

// AvatarIdEQ applies the EQ predicate on the "avatarId" field.
func AvatarIdEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatarId), v))
	})
}

// AvatarIdNEQ applies the NEQ predicate on the "avatarId" field.
func AvatarIdNEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatarId), v))
	})
}

// AvatarIdIn applies the In predicate on the "avatarId" field.
func AvatarIdIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvatarId), v...))
	})
}

// AvatarIdNotIn applies the NotIn predicate on the "avatarId" field.
func AvatarIdNotIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvatarId), v...))
	})
}

// AvatarIdGT applies the GT predicate on the "avatarId" field.
func AvatarIdGT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatarId), v))
	})
}

// AvatarIdGTE applies the GTE predicate on the "avatarId" field.
func AvatarIdGTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatarId), v))
	})
}

// AvatarIdLT applies the LT predicate on the "avatarId" field.
func AvatarIdLT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatarId), v))
	})
}

// AvatarIdLTE applies the LTE predicate on the "avatarId" field.
func AvatarIdLTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatarId), v))
	})
}

// AvatarIdIsNil applies the IsNil predicate on the "avatarId" field.
func AvatarIdIsNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAvatarId)))
	})
}

// AvatarIdNotNil applies the NotNil predicate on the "avatarId" field.
func AvatarIdNotNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAvatarId)))
	})
}

// SummaryEQ applies the EQ predicate on the "summary" field.
func SummaryEQ(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSummary), v))
	})
}

// SummaryNEQ applies the NEQ predicate on the "summary" field.
func SummaryNEQ(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSummary), v))
	})
}

// SummaryIn applies the In predicate on the "summary" field.
func SummaryIn(vs ...string) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSummary), v...))
	})
}

// SummaryNotIn applies the NotIn predicate on the "summary" field.
func SummaryNotIn(vs ...string) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSummary), v...))
	})
}

// SummaryGT applies the GT predicate on the "summary" field.
func SummaryGT(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSummary), v))
	})
}

// SummaryGTE applies the GTE predicate on the "summary" field.
func SummaryGTE(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSummary), v))
	})
}

// SummaryLT applies the LT predicate on the "summary" field.
func SummaryLT(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSummary), v))
	})
}

// SummaryLTE applies the LTE predicate on the "summary" field.
func SummaryLTE(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSummary), v))
	})
}

// SummaryContains applies the Contains predicate on the "summary" field.
func SummaryContains(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSummary), v))
	})
}

// SummaryHasPrefix applies the HasPrefix predicate on the "summary" field.
func SummaryHasPrefix(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSummary), v))
	})
}

// SummaryHasSuffix applies the HasSuffix predicate on the "summary" field.
func SummaryHasSuffix(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSummary), v))
	})
}

// SummaryIsNil applies the IsNil predicate on the "summary" field.
func SummaryIsNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSummary)))
	})
}

// SummaryNotNil applies the NotNil predicate on the "summary" field.
func SummaryNotNil() predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSummary)))
	})
}

// SummaryEqualFold applies the EqualFold predicate on the "summary" field.
func SummaryEqualFold(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSummary), v))
	})
}

// SummaryContainsFold applies the ContainsFold predicate on the "summary" field.
func SummaryContainsFold(v string) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSummary), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreateByEQ applies the EQ predicate on the "createBy" field.
func CreateByEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// CreateByNEQ applies the NEQ predicate on the "createBy" field.
func CreateByNEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateBy), v))
	})
}

// CreateByIn applies the In predicate on the "createBy" field.
func CreateByIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func CreateByNotIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func CreateByGT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateBy), v))
	})
}

// CreateByGTE applies the GTE predicate on the "createBy" field.
func CreateByGTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateBy), v))
	})
}

// CreateByLT applies the LT predicate on the "createBy" field.
func CreateByLT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateBy), v))
	})
}

// CreateByLTE applies the LTE predicate on the "createBy" field.
func CreateByLTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateBy), v))
	})
}

// UpdateByEQ applies the EQ predicate on the "updateBy" field.
func UpdateByEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByNEQ applies the NEQ predicate on the "updateBy" field.
func UpdateByNEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByIn applies the In predicate on the "updateBy" field.
func UpdateByIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func UpdateByNotIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func UpdateByGT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByGTE applies the GTE predicate on the "updateBy" field.
func UpdateByGTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLT applies the LT predicate on the "updateBy" field.
func UpdateByLT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLTE applies the LTE predicate on the "updateBy" field.
func UpdateByLTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateBy), v))
	})
}

// TenantIdEQ applies the EQ predicate on the "tenantId" field.
func TenantIdEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// TenantIdNEQ applies the NEQ predicate on the "tenantId" field.
func TenantIdNEQ(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantId), v))
	})
}

// TenantIdIn applies the In predicate on the "tenantId" field.
func TenantIdIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func TenantIdNotIn(vs ...int64) predicate.VipType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.VipType(func(s *sql.Selector) {
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
func TenantIdGT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenantId), v))
	})
}

// TenantIdGTE applies the GTE predicate on the "tenantId" field.
func TenantIdGTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenantId), v))
	})
}

// TenantIdLT applies the LT predicate on the "tenantId" field.
func TenantIdLT(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenantId), v))
	})
}

// TenantIdLTE applies the LTE predicate on the "tenantId" field.
func TenantIdLTE(v int64) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenantId), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VipType) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VipType) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
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
func Not(p predicate.VipType) predicate.VipType {
	return predicate.VipType(func(s *sql.Selector) {
		p(s.Not())
	})
}
