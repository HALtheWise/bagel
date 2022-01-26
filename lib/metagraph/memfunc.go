package metagraph

func MemoFunc1[Arg1 internKey, V any](name string, f func(Arg1) V) func(Arg1) V {
	memoTable := make(map[Arg1]V)
	return func(a1 Arg1) V {}
}
