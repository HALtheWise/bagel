package core

type internKey interface{ ~uint32 }

const MAX_RESERVED_ID = 3

// TODO(eric): Change this to store data in the Context
type InternTable[K internKey, V comparable] struct {
	data    []V
	mapping map[V]K
}

// Predeclared must map consecutive integers starting at 0
func NewInternTable[K internKey, V comparable](predeclared map[V]K) InternTable[K, V] {
	table := InternTable[K, V]{
		make([]V, len(predeclared)),
		predeclared,
	}

	for v, k := range predeclared {
		table.data[k] = v
	}

	return table
}

func (i *InternTable[K, V]) Insert(c *Context, value V) K {
	if key, ok := i.mapping[value]; ok {
		return key
	}

	key := K(len(i.data))

	i.data = append(i.data, value)
	i.mapping[value] = key

	return key
}

func (i *InternTable[K, V]) Get(c *Context, key K) V {
	return i.data[key]
}
