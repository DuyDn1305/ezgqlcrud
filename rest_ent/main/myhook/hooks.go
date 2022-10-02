package myhook

import (
	"context"
	"restent/ent"
)

func OnCreate(fn func(context.Context, ent.Mutation)) func(ent.Mutator) ent.Mutator {
	bien := func(next ent.Mutator) ent.Mutator {
		mutate_fn := func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if ent.OpCreate.Is(m.Op()) {
				fn(ctx, m)
			}
			// fmt.Println("Not set: ", m.Op())
			return next.Mutate(ctx, m)
		}

		return ent.MutateFunc(mutate_fn)
	}
	return bien
}

type HookParam = func(context.Context, ent.Mutation)
type HookReturn = func(ent.Mutator) ent.Mutator

func On(fn HookParam, ops ent.Op) HookReturn {
	on_result := func(next ent.Mutator) ent.Mutator {
		mutate_param := func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if ops.Is(m.Op()) {
				fn(ctx, m)
			}
			return next.Mutate(ctx, m)
		}
		return ent.MutateFunc(mutate_param)
	}
	return on_result
}
