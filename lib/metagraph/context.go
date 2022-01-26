package metagraph

type Context struct {
	hits, misses int
}

func NewContext() *Context {
	return &Context{}
}
