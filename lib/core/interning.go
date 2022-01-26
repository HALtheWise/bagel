package core

type internKey interface{ ~uint32 }

const INVALID = 1<<32 - 1

// TODO(eric): Change this to store data in the Context
type InternTable[K internKey, V comparable] struct {
	data    []V
	mapping map[V]K
}

func (i *InternTable[K, V]) Insert(c *Context, value V) K {
	if len(i.data) == 0 {
		i.mapping = make(map[V]K)
	}

	if key, ok := i.mapping[value]; ok {
		return key
	}

	key := K(len(i.data))

	i.data = append(i.data, value)
	i.mapping[value] = key

	return key
}

func (i *InternTable[K, V]) Get(c *Context, key K) V {
	if key == INVALID {
		panic("INVALID key")
	}
	if key >= K(len(i.data)) {
		panic("Key too large")
	}
	return i.data[key]
}
