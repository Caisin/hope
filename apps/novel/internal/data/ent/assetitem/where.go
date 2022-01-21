// Code generated by entc, DO NOT EDIT.

package assetitem

import (
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AssetItemId applies equality check predicate on the "assetItemId" field. It's identical to AssetItemIdEQ.
func AssetItemId(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetItemId), v))
	})
}

// AssetName applies equality check predicate on the "assetName" field. It's identical to AssetNameEQ.
func AssetName(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetName), v))
	})
}

// CashTag applies equality check predicate on the "cashTag" field. It's identical to CashTagEQ.
func CashTag(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCashTag), v))
	})
}

// ValidDays applies equality check predicate on the "validDays" field. It's identical to ValidDaysEQ.
func ValidDays(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidDays), v))
	})
}

// EffectTime applies equality check predicate on the "effectTime" field. It's identical to EffectTimeEQ.
func EffectTime(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffectTime), v))
	})
}

// ExpiredTime applies equality check predicate on the "expiredTime" field. It's identical to ExpiredTimeEQ.
func ExpiredTime(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiredTime), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreateBy applies equality check predicate on the "createBy" field. It's identical to CreateByEQ.
func CreateBy(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// UpdateBy applies equality check predicate on the "updateBy" field. It's identical to UpdateByEQ.
func UpdateBy(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// TenantId applies equality check predicate on the "tenantId" field. It's identical to TenantIdEQ.
func TenantId(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// AssetItemIdEQ applies the EQ predicate on the "assetItemId" field.
func AssetItemIdEQ(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetItemId), v))
	})
}

// AssetItemIdNEQ applies the NEQ predicate on the "assetItemId" field.
func AssetItemIdNEQ(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAssetItemId), v))
	})
}

// AssetItemIdIn applies the In predicate on the "assetItemId" field.
func AssetItemIdIn(vs ...int32) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAssetItemId), v...))
	})
}

// AssetItemIdNotIn applies the NotIn predicate on the "assetItemId" field.
func AssetItemIdNotIn(vs ...int32) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAssetItemId), v...))
	})
}

// AssetItemIdGT applies the GT predicate on the "assetItemId" field.
func AssetItemIdGT(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAssetItemId), v))
	})
}

// AssetItemIdGTE applies the GTE predicate on the "assetItemId" field.
func AssetItemIdGTE(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAssetItemId), v))
	})
}

// AssetItemIdLT applies the LT predicate on the "assetItemId" field.
func AssetItemIdLT(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAssetItemId), v))
	})
}

// AssetItemIdLTE applies the LTE predicate on the "assetItemId" field.
func AssetItemIdLTE(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAssetItemId), v))
	})
}

// AssetNameEQ applies the EQ predicate on the "assetName" field.
func AssetNameEQ(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetName), v))
	})
}

// AssetNameNEQ applies the NEQ predicate on the "assetName" field.
func AssetNameNEQ(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAssetName), v))
	})
}

// AssetNameIn applies the In predicate on the "assetName" field.
func AssetNameIn(vs ...string) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAssetName), v...))
	})
}

// AssetNameNotIn applies the NotIn predicate on the "assetName" field.
func AssetNameNotIn(vs ...string) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAssetName), v...))
	})
}

// AssetNameGT applies the GT predicate on the "assetName" field.
func AssetNameGT(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAssetName), v))
	})
}

// AssetNameGTE applies the GTE predicate on the "assetName" field.
func AssetNameGTE(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAssetName), v))
	})
}

// AssetNameLT applies the LT predicate on the "assetName" field.
func AssetNameLT(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAssetName), v))
	})
}

// AssetNameLTE applies the LTE predicate on the "assetName" field.
func AssetNameLTE(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAssetName), v))
	})
}

// AssetNameContains applies the Contains predicate on the "assetName" field.
func AssetNameContains(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAssetName), v))
	})
}

// AssetNameHasPrefix applies the HasPrefix predicate on the "assetName" field.
func AssetNameHasPrefix(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAssetName), v))
	})
}

// AssetNameHasSuffix applies the HasSuffix predicate on the "assetName" field.
func AssetNameHasSuffix(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAssetName), v))
	})
}

// AssetNameIsNil applies the IsNil predicate on the "assetName" field.
func AssetNameIsNil() predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAssetName)))
	})
}

// AssetNameNotNil applies the NotNil predicate on the "assetName" field.
func AssetNameNotNil() predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAssetName)))
	})
}

// AssetNameEqualFold applies the EqualFold predicate on the "assetName" field.
func AssetNameEqualFold(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAssetName), v))
	})
}

// AssetNameContainsFold applies the ContainsFold predicate on the "assetName" field.
func AssetNameContainsFold(v string) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAssetName), v))
	})
}

// CashTagEQ applies the EQ predicate on the "cashTag" field.
func CashTagEQ(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCashTag), v))
	})
}

// CashTagNEQ applies the NEQ predicate on the "cashTag" field.
func CashTagNEQ(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCashTag), v))
	})
}

// CashTagIn applies the In predicate on the "cashTag" field.
func CashTagIn(vs ...int32) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCashTag), v...))
	})
}

// CashTagNotIn applies the NotIn predicate on the "cashTag" field.
func CashTagNotIn(vs ...int32) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCashTag), v...))
	})
}

// CashTagGT applies the GT predicate on the "cashTag" field.
func CashTagGT(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCashTag), v))
	})
}

// CashTagGTE applies the GTE predicate on the "cashTag" field.
func CashTagGTE(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCashTag), v))
	})
}

// CashTagLT applies the LT predicate on the "cashTag" field.
func CashTagLT(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCashTag), v))
	})
}

// CashTagLTE applies the LTE predicate on the "cashTag" field.
func CashTagLTE(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCashTag), v))
	})
}

// CashTagIsNil applies the IsNil predicate on the "cashTag" field.
func CashTagIsNil() predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCashTag)))
	})
}

// CashTagNotNil applies the NotNil predicate on the "cashTag" field.
func CashTagNotNil() predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCashTag)))
	})
}

// ValidDaysEQ applies the EQ predicate on the "validDays" field.
func ValidDaysEQ(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValidDays), v))
	})
}

// ValidDaysNEQ applies the NEQ predicate on the "validDays" field.
func ValidDaysNEQ(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValidDays), v))
	})
}

// ValidDaysIn applies the In predicate on the "validDays" field.
func ValidDaysIn(vs ...int32) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func ValidDaysNotIn(vs ...int32) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func ValidDaysGT(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValidDays), v))
	})
}

// ValidDaysGTE applies the GTE predicate on the "validDays" field.
func ValidDaysGTE(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValidDays), v))
	})
}

// ValidDaysLT applies the LT predicate on the "validDays" field.
func ValidDaysLT(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValidDays), v))
	})
}

// ValidDaysLTE applies the LTE predicate on the "validDays" field.
func ValidDaysLTE(v int32) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValidDays), v))
	})
}

// ValidDaysIsNil applies the IsNil predicate on the "validDays" field.
func ValidDaysIsNil() predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValidDays)))
	})
}

// ValidDaysNotNil applies the NotNil predicate on the "validDays" field.
func ValidDaysNotNil() predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValidDays)))
	})
}

// EffectTimeEQ applies the EQ predicate on the "effectTime" field.
func EffectTimeEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEffectTime), v))
	})
}

// EffectTimeNEQ applies the NEQ predicate on the "effectTime" field.
func EffectTimeNEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEffectTime), v))
	})
}

// EffectTimeIn applies the In predicate on the "effectTime" field.
func EffectTimeIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEffectTime), v...))
	})
}

// EffectTimeNotIn applies the NotIn predicate on the "effectTime" field.
func EffectTimeNotIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEffectTime), v...))
	})
}

// EffectTimeGT applies the GT predicate on the "effectTime" field.
func EffectTimeGT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEffectTime), v))
	})
}

// EffectTimeGTE applies the GTE predicate on the "effectTime" field.
func EffectTimeGTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEffectTime), v))
	})
}

// EffectTimeLT applies the LT predicate on the "effectTime" field.
func EffectTimeLT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEffectTime), v))
	})
}

// EffectTimeLTE applies the LTE predicate on the "effectTime" field.
func EffectTimeLTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEffectTime), v))
	})
}

// ExpiredTimeEQ applies the EQ predicate on the "expiredTime" field.
func ExpiredTimeEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiredTime), v))
	})
}

// ExpiredTimeNEQ applies the NEQ predicate on the "expiredTime" field.
func ExpiredTimeNEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiredTime), v))
	})
}

// ExpiredTimeIn applies the In predicate on the "expiredTime" field.
func ExpiredTimeIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpiredTime), v...))
	})
}

// ExpiredTimeNotIn applies the NotIn predicate on the "expiredTime" field.
func ExpiredTimeNotIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpiredTime), v...))
	})
}

// ExpiredTimeGT applies the GT predicate on the "expiredTime" field.
func ExpiredTimeGT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiredTime), v))
	})
}

// ExpiredTimeGTE applies the GTE predicate on the "expiredTime" field.
func ExpiredTimeGTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiredTime), v))
	})
}

// ExpiredTimeLT applies the LT predicate on the "expiredTime" field.
func ExpiredTimeLT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiredTime), v))
	})
}

// ExpiredTimeLTE applies the LTE predicate on the "expiredTime" field.
func ExpiredTimeLTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiredTime), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreateByEQ applies the EQ predicate on the "createBy" field.
func CreateByEQ(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// CreateByNEQ applies the NEQ predicate on the "createBy" field.
func CreateByNEQ(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateBy), v))
	})
}

// CreateByIn applies the In predicate on the "createBy" field.
func CreateByIn(vs ...int64) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func CreateByNotIn(vs ...int64) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func CreateByGT(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateBy), v))
	})
}

// CreateByGTE applies the GTE predicate on the "createBy" field.
func CreateByGTE(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateBy), v))
	})
}

// CreateByLT applies the LT predicate on the "createBy" field.
func CreateByLT(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateBy), v))
	})
}

// CreateByLTE applies the LTE predicate on the "createBy" field.
func CreateByLTE(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateBy), v))
	})
}

// UpdateByEQ applies the EQ predicate on the "updateBy" field.
func UpdateByEQ(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByNEQ applies the NEQ predicate on the "updateBy" field.
func UpdateByNEQ(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByIn applies the In predicate on the "updateBy" field.
func UpdateByIn(vs ...int64) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func UpdateByNotIn(vs ...int64) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func UpdateByGT(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByGTE applies the GTE predicate on the "updateBy" field.
func UpdateByGTE(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLT applies the LT predicate on the "updateBy" field.
func UpdateByLT(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLTE applies the LTE predicate on the "updateBy" field.
func UpdateByLTE(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateBy), v))
	})
}

// TenantIdEQ applies the EQ predicate on the "tenantId" field.
func TenantIdEQ(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// TenantIdNEQ applies the NEQ predicate on the "tenantId" field.
func TenantIdNEQ(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantId), v))
	})
}

// TenantIdIn applies the In predicate on the "tenantId" field.
func TenantIdIn(vs ...int64) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func TenantIdNotIn(vs ...int64) predicate.AssetItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AssetItem(func(s *sql.Selector) {
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
func TenantIdGT(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenantId), v))
	})
}

// TenantIdGTE applies the GTE predicate on the "tenantId" field.
func TenantIdGTE(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenantId), v))
	})
}

// TenantIdLT applies the LT predicate on the "tenantId" field.
func TenantIdLT(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenantId), v))
	})
}

// TenantIdLTE applies the LTE predicate on the "tenantId" field.
func TenantIdLTE(v int64) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenantId), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AssetItem) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AssetItem) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
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
func Not(p predicate.AssetItem) predicate.AssetItem {
	return predicate.AssetItem(func(s *sql.Selector) {
		p(s.Not())
	})
}