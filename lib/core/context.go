package core

import "fmt"

type Context struct {
	hits, misses int
}

// DefaultContext should only be used for debug printing, and must not affect the output of the build
var DefaultContext *Context

func NewContext() *Context {
	c := &Context{}
	DefaultContext = c

	return c
}

func (c *Context) String() string {
	return fmt.Sprintf("ctx{hits=%d misses=%d}", c.hits, c.misses)
}
