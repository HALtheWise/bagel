package core

import "fmt"

type C struct {
	hits, misses int
}

// DefaultContext should only be used for debug printing, and must not affect the output of the build
var DefaultContext *C

func NewContext() *C {
	c := &C{}
	DefaultContext = c

	return c
}

func (c *C) String() string {
	return fmt.Sprintf("ctx{hits=%d misses=%d}", c.hits, c.misses)
}
