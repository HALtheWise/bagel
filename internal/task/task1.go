package task

func Task1[Arg comparable, Result any](name string, f func(*Context, Arg) Result) func(*Context, Arg) Result {
	return func(parent *Context, arg Arg) Result {
		typedMemo := getTypedMemo[Arg, Result](parent.global, name)

		if cached, ok := typedMemo[arg]; ok {
			parent.global.Stats.cacheHits += 1
			return cached
		}
		parent.global.Stats.cacheMisses += 1

		ctx := Context{
			name:   name,
			global: parent.global,
		}

		result := f(&ctx, arg)

		typedMemo[arg] = result

		return result
	}
}
