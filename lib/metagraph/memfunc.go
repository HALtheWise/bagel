package metagraph

func MemoFunc1[Arg1 internKey, V any](name string, f func(Context, Arg1) V) func(Context, Arg1) V {
	memoTable := make(map[Arg1]V)
	return func(c Context, a1 Arg1) V {
		if val, ok := memoTable[a1]; ok {
			// Return previously computed value, if available
			c.hits += 1
			return val
		}

		c.misses += 1
		result := f(c, a1)

		memoTable[a1] = result

		return result
	}
}
