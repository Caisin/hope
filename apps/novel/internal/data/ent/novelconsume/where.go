// Code generated by entc, DO NOT EDIT.

package novelconsume

import (
	"hope/apps/novel/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func IDNotIn(ids ...int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func IDGT(id int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// NovelId applies equality check predicate on the "novelId" field. It's identical to NovelIdEQ.
func NovelId(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNovelId), v))
	})
}

// Coin applies equality check predicate on the "coin" field. It's identical to CoinEQ.
func Coin(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoin), v))
	})
}

// Coupon applies equality check predicate on the "coupon" field. It's identical to CouponEQ.
func Coupon(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoupon), v))
	})
}

// Discount applies equality check predicate on the "discount" field. It's identical to DiscountEQ.
func Discount(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// Reward applies equality check predicate on the "reward" field. It's identical to RewardEQ.
func Reward(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReward), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// CreateBy applies equality check predicate on the "createBy" field. It's identical to CreateByEQ.
func CreateBy(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// UpdateBy applies equality check predicate on the "updateBy" field. It's identical to UpdateByEQ.
func UpdateBy(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// TenantId applies equality check predicate on the "tenantId" field. It's identical to TenantIdEQ.
func TenantId(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// NovelIdEQ applies the EQ predicate on the "novelId" field.
func NovelIdEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNovelId), v))
	})
}

// NovelIdNEQ applies the NEQ predicate on the "novelId" field.
func NovelIdNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNovelId), v))
	})
}

// NovelIdIn applies the In predicate on the "novelId" field.
func NovelIdIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNovelId), v...))
	})
}

// NovelIdNotIn applies the NotIn predicate on the "novelId" field.
func NovelIdNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNovelId), v...))
	})
}

// NovelIdGT applies the GT predicate on the "novelId" field.
func NovelIdGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNovelId), v))
	})
}

// NovelIdGTE applies the GTE predicate on the "novelId" field.
func NovelIdGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNovelId), v))
	})
}

// NovelIdLT applies the LT predicate on the "novelId" field.
func NovelIdLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNovelId), v))
	})
}

// NovelIdLTE applies the LTE predicate on the "novelId" field.
func NovelIdLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNovelId), v))
	})
}

// CoinEQ applies the EQ predicate on the "coin" field.
func CoinEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoin), v))
	})
}

// CoinNEQ applies the NEQ predicate on the "coin" field.
func CoinNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoin), v))
	})
}

// CoinIn applies the In predicate on the "coin" field.
func CoinIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoin), v...))
	})
}

// CoinNotIn applies the NotIn predicate on the "coin" field.
func CoinNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoin), v...))
	})
}

// CoinGT applies the GT predicate on the "coin" field.
func CoinGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoin), v))
	})
}

// CoinGTE applies the GTE predicate on the "coin" field.
func CoinGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoin), v))
	})
}

// CoinLT applies the LT predicate on the "coin" field.
func CoinLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoin), v))
	})
}

// CoinLTE applies the LTE predicate on the "coin" field.
func CoinLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoin), v))
	})
}

// CoinIsNil applies the IsNil predicate on the "coin" field.
func CoinIsNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoin)))
	})
}

// CoinNotNil applies the NotNil predicate on the "coin" field.
func CoinNotNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoin)))
	})
}

// CouponEQ applies the EQ predicate on the "coupon" field.
func CouponEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoupon), v))
	})
}

// CouponNEQ applies the NEQ predicate on the "coupon" field.
func CouponNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoupon), v))
	})
}

// CouponIn applies the In predicate on the "coupon" field.
func CouponIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoupon), v...))
	})
}

// CouponNotIn applies the NotIn predicate on the "coupon" field.
func CouponNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoupon), v...))
	})
}

// CouponGT applies the GT predicate on the "coupon" field.
func CouponGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoupon), v))
	})
}

// CouponGTE applies the GTE predicate on the "coupon" field.
func CouponGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoupon), v))
	})
}

// CouponLT applies the LT predicate on the "coupon" field.
func CouponLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoupon), v))
	})
}

// CouponLTE applies the LTE predicate on the "coupon" field.
func CouponLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoupon), v))
	})
}

// CouponIsNil applies the IsNil predicate on the "coupon" field.
func CouponIsNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoupon)))
	})
}

// CouponNotNil applies the NotNil predicate on the "coupon" field.
func CouponNotNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoupon)))
	})
}

// DiscountEQ applies the EQ predicate on the "discount" field.
func DiscountEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// DiscountNEQ applies the NEQ predicate on the "discount" field.
func DiscountNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscount), v))
	})
}

// DiscountIn applies the In predicate on the "discount" field.
func DiscountIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscount), v...))
	})
}

// DiscountNotIn applies the NotIn predicate on the "discount" field.
func DiscountNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscount), v...))
	})
}

// DiscountGT applies the GT predicate on the "discount" field.
func DiscountGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscount), v))
	})
}

// DiscountGTE applies the GTE predicate on the "discount" field.
func DiscountGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscount), v))
	})
}

// DiscountLT applies the LT predicate on the "discount" field.
func DiscountLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscount), v))
	})
}

// DiscountLTE applies the LTE predicate on the "discount" field.
func DiscountLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscount), v))
	})
}

// DiscountIsNil applies the IsNil predicate on the "discount" field.
func DiscountIsNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDiscount)))
	})
}

// DiscountNotNil applies the NotNil predicate on the "discount" field.
func DiscountNotNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDiscount)))
	})
}

// RewardEQ applies the EQ predicate on the "reward" field.
func RewardEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReward), v))
	})
}

// RewardNEQ applies the NEQ predicate on the "reward" field.
func RewardNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReward), v))
	})
}

// RewardIn applies the In predicate on the "reward" field.
func RewardIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReward), v...))
	})
}

// RewardNotIn applies the NotIn predicate on the "reward" field.
func RewardNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReward), v...))
	})
}

// RewardGT applies the GT predicate on the "reward" field.
func RewardGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReward), v))
	})
}

// RewardGTE applies the GTE predicate on the "reward" field.
func RewardGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReward), v))
	})
}

// RewardLT applies the LT predicate on the "reward" field.
func RewardLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReward), v))
	})
}

// RewardLTE applies the LTE predicate on the "reward" field.
func RewardLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReward), v))
	})
}

// RewardIsNil applies the IsNil predicate on the "reward" field.
func RewardIsNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReward)))
	})
}

// RewardNotNil applies the NotNil predicate on the "reward" field.
func RewardNotNil() predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReward)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// CreateByEQ applies the EQ predicate on the "createBy" field.
func CreateByEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateBy), v))
	})
}

// CreateByNEQ applies the NEQ predicate on the "createBy" field.
func CreateByNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateBy), v))
	})
}

// CreateByIn applies the In predicate on the "createBy" field.
func CreateByIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func CreateByNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func CreateByGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateBy), v))
	})
}

// CreateByGTE applies the GTE predicate on the "createBy" field.
func CreateByGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateBy), v))
	})
}

// CreateByLT applies the LT predicate on the "createBy" field.
func CreateByLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateBy), v))
	})
}

// CreateByLTE applies the LTE predicate on the "createBy" field.
func CreateByLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateBy), v))
	})
}

// UpdateByEQ applies the EQ predicate on the "updateBy" field.
func UpdateByEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByNEQ applies the NEQ predicate on the "updateBy" field.
func UpdateByNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateBy), v))
	})
}

// UpdateByIn applies the In predicate on the "updateBy" field.
func UpdateByIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func UpdateByNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func UpdateByGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByGTE applies the GTE predicate on the "updateBy" field.
func UpdateByGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLT applies the LT predicate on the "updateBy" field.
func UpdateByLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateBy), v))
	})
}

// UpdateByLTE applies the LTE predicate on the "updateBy" field.
func UpdateByLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateBy), v))
	})
}

// TenantIdEQ applies the EQ predicate on the "tenantId" field.
func TenantIdEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTenantId), v))
	})
}

// TenantIdNEQ applies the NEQ predicate on the "tenantId" field.
func TenantIdNEQ(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTenantId), v))
	})
}

// TenantIdIn applies the In predicate on the "tenantId" field.
func TenantIdIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func TenantIdNotIn(vs ...int64) predicate.NovelConsume {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func TenantIdGT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTenantId), v))
	})
}

// TenantIdGTE applies the GTE predicate on the "tenantId" field.
func TenantIdGTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTenantId), v))
	})
}

// TenantIdLT applies the LT predicate on the "tenantId" field.
func TenantIdLT(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTenantId), v))
	})
}

// TenantIdLTE applies the LTE predicate on the "tenantId" field.
func TenantIdLTE(v int64) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTenantId), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NovelConsume) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NovelConsume) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
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
func Not(p predicate.NovelConsume) predicate.NovelConsume {
	return predicate.NovelConsume(func(s *sql.Selector) {
		p(s.Not())
	})
}