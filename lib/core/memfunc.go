package core

func MemoFunc1[Arg1 internKey, V any](name string, f func(*Context, Arg1) V) func(*Context, Arg1) V {
	memoTable := make(map[Arg1]V)
	return func(c *Context, a1 Arg1) V {
		if val, ok := memoTable[a1]; ok {
			// Return previously computed value, if available
			c.hits++
			return val
		}

		c.misses++
		result := f(c, a1)

		memoTable[a1] = result

		return result
	}
}

type key2[A1, A2 any] struct {
	a1 A1
	a2 A2
}

func MemoFunc2[Arg1, Arg2 internKey, V any](name string, f func(*Context, Arg1, Arg2) V) func(*Context, Arg1, Arg2) V {

	memoTable := make(map[key2[Arg1, Arg2]]V)

	return func(c *Context, a1 Arg1, a2 Arg2) V {
		k := key2[Arg1, Arg2]{a1, a2}

		if val, ok := memoTable[k]; ok {
			// Return previously computed value, if available
			c.hits++
			return val
		}

		c.misses++
		result := f(c, a1, a2)

		memoTable[k] = result

		return result
	}
}
