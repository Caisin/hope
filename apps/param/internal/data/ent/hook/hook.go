// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"hope/apps/param/internal/data/ent"
)

// The NovelPayConfigFunc type is an adapter to allow the use of ordinary
// function as NovelPayConfig mutator.
type NovelPayConfigFunc func(context.Context, *ent.NovelPayConfigMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NovelPayConfigFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.NovelPayConfigMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NovelPayConfigMutation", m)
	}
	return f(ctx, mv)
}

// The NovelTagFunc type is an adapter to allow the use of ordinary
// function as NovelTag mutator.
type NovelTagFunc func(context.Context, *ent.NovelTagMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NovelTagFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.NovelTagMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NovelTagMutation", m)
	}
	return f(ctx, mv)
}

// The PageConfigFunc type is an adapter to allow the use of ordinary
// function as PageConfig mutator.
type PageConfigFunc func(context.Context, *ent.PageConfigMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PageConfigFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PageConfigMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PageConfigMutation", m)
	}
	return f(ctx, mv)
}

// The QiniuConfigFunc type is an adapter to allow the use of ordinary
// function as QiniuConfig mutator.
type QiniuConfigFunc func(context.Context, *ent.QiniuConfigMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f QiniuConfigFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.QiniuConfigMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.QiniuConfigMutation", m)
	}
	return f(ctx, mv)
}

// The ResourceGroupFunc type is an adapter to allow the use of ordinary
// function as ResourceGroup mutator.
type ResourceGroupFunc func(context.Context, *ent.ResourceGroupMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ResourceGroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ResourceGroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ResourceGroupMutation", m)
	}
	return f(ctx, mv)
}

// The ResourceStorageFunc type is an adapter to allow the use of ordinary
// function as ResourceStorage mutator.
type ResourceStorageFunc func(context.Context, *ent.ResourceStorageMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ResourceStorageFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ResourceStorageMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ResourceStorageMutation", m)
	}
	return f(ctx, mv)
}

// The ScoreProductFunc type is an adapter to allow the use of ordinary
// function as ScoreProduct mutator.
type ScoreProductFunc func(context.Context, *ent.ScoreProductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScoreProductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScoreProductMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScoreProductMutation", m)
	}
	return f(ctx, mv)
}

// The TaskFunc type is an adapter to allow the use of ordinary
// function as Task mutator.
type TaskFunc func(context.Context, *ent.TaskMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TaskFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TaskMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TaskMutation", m)
	}
	return f(ctx, mv)
}

// The UserAnalysisStatisticsFunc type is an adapter to allow the use of ordinary
// function as UserAnalysisStatistics mutator.
type UserAnalysisStatisticsFunc func(context.Context, *ent.UserAnalysisStatisticsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserAnalysisStatisticsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserAnalysisStatisticsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserAnalysisStatisticsMutation", m)
	}
	return f(ctx, mv)
}

// The UserConsumeFunc type is an adapter to allow the use of ordinary
// function as UserConsume mutator.
type UserConsumeFunc func(context.Context, *ent.UserConsumeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserConsumeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserConsumeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserConsumeMutation", m)
	}
	return f(ctx, mv)
}

// The UserResourceFunc type is an adapter to allow the use of ordinary
// function as UserResource mutator.
type UserResourceFunc func(context.Context, *ent.UserResourceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserResourceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserResourceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserResourceMutation", m)
	}
	return f(ctx, mv)
}

// The UserResourceRecordFunc type is an adapter to allow the use of ordinary
// function as UserResourceRecord mutator.
type UserResourceRecordFunc func(context.Context, *ent.UserResourceRecordMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserResourceRecordFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserResourceRecordMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserResourceRecordMutation", m)
	}
	return f(ctx, mv)
}

// The VipTypeFunc type is an adapter to allow the use of ordinary
// function as VipType mutator.
type VipTypeFunc func(context.Context, *ent.VipTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VipTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.VipTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VipTypeMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}