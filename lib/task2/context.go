package task2

// TODO: This should be backed by disk using capnp or something
type GlobalContext struct {
	strings    []string
	stringsMap map[string]anyRef
	structs    []anyStruct
	structsMap map[anyStruct]anyRef
}

type anyStruct struct{ left, right anyRef }

func NewGlobalContext() *GlobalContext {
	return &GlobalContext{
		stringsMap: map[string]anyRef{},
		structsMap: map[anyStruct]anyRef{},
	}
}

type Context struct {
	TaskID int
	Global *GlobalContext
}
