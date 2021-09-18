package task

func Task1[Arg comparable, Result any](name string, f func(*Context, Arg) Result) func(*Context, Arg) Result {
	id := getNextID()
	return func(parent *Context, arg Arg) Result {
		typedMemo := getTypedMemo[Arg, Result](parent.global, id)

		if cached, ok := typedMemo[arg]; ok {
			parent.global.stats.cacheHits += 1
			return cached
		}
		parent.global.stats.cacheMisses += 1

		ctx := Context{
			name:   name,
			global: parent.global,
		}

		result := f(&ctx, arg)

		typedMemo[arg] = result

		return result
	}
}

type ManyArgs[Arg1, Arg2, Arg3 comparable] struct {
	arg1 Arg1
	arg2 Arg2
	arg3 Arg3
}

func Task2[Arg1, Arg2 comparable, Result any](
	name string, f func(*Context, Arg1, Arg2) Result,
) func(*Context, Arg1, Arg2) Result {
	id := getNextID()
	return func(parent *Context, arg1 Arg1, arg2 Arg2) Result {
		typedMemo := getTypedMemo[ManyArgs[Arg1, Arg2, struct{}], Result](parent.global, name)
		args := ManyArgs[Arg1, Arg2, struct{}]{arg1, arg2, struct{}{}}

		if cached, ok := typedMemo[args]; ok {
			parent.global.stats.cacheHits += 1
			return cached
		}
		parent.global.stats.cacheMisses += 1

		ctx := Context{
			name:   name,
			global: parent.global,
		}

		result := f(&ctx, arg1, arg2)

		typedMemo[args] = result

		return result
	}
}

func Task3[Arg1, Arg2, Arg3 comparable, Result any](
	name string, f func(*Context, Arg1, Arg2, Arg3) Result,
) func(*Context, Arg1, Arg2, Arg3) Result {
	return func(parent *Context, arg1 Arg1, arg2 Arg2, arg3 Arg3) Result {
		typedMemo := getTypedMemo[ManyArgs[Arg1, Arg2, Arg3], Result](parent.global, name)
		args := ManyArgs[Arg1, Arg2, Arg3]{arg1, arg2, arg3}

		if cached, ok := typedMemo[args]; ok {
			parent.global.stats.cacheHits += 1
			return cached
		}
		parent.global.stats.cacheMisses += 1

		ctx := Context{
			name:   name,
			global: parent.global,
		}

		result := f(&ctx, arg1, arg2, arg3)

		typedMemo[args] = result

		return result
	}
}
