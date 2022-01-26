package core

import "fmt"

type C struct {
	hits, misses int
}

func NewContext() *C {
	return &C{}
}

func (c *C) String() string {
	return fmt.Sprintf("hits=%d misses=%d", c.hits, c.misses)
}
