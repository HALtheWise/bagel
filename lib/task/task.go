package task

import (
	capnp "zombiezen.com/go/capnproto2"

	"github.com/HALtheWise/bagel/lib/cache/graph"
)

type packer interface{ pack() uint32 }
type diskResult interface {
	ToPtr() capnp.Ptr
	~struct{ capnp.Struct }
}

func getFuncData(ctx *globalContext, kind uint32, arg packer) (graph.FuncObj, uint32, bool) {
	argPacked := arg.pack()
	idx := ctx.cache.InternFunc(kind, argPacked)
	funcs, err := ctx.cache.Funcs()
	if err != nil {
		panic(err)
	}
	data := funcs.At(int(idx))
	if data.Kind() == kind {
		return data, idx, true
	}
	data.SetKind(kind)
	data.SetArg(argPacked)
	return data, idx, false
}

func DiskTask[Arg packer, Result diskResult](name string, f func(*Context, Arg, *capnp.Segment) Result) func(*Context, Arg) Result {
	kind := getNextID()
	return func(parent *Context, arg Arg) Result {
		funcData, _, ok := getFuncData(parent.global, kind, arg)

		if ok {
			// This data is already cached.
			// TODO(eric): Check for invalidation
			ptr, err := funcData.ResultPtr()
			if err != nil {
				panic(err)
			}
			result := struct{ capnp.Struct }{ptr.Struct()}
			parent.global.stats.cacheHits += 1
			return Result(result)
		}

		parent.global.stats.cacheMisses += 1
		ctx := Context{
			name:   name,
			global: parent.global,
		}

		result := f(&ctx, arg, parent.global.cache.Segment())

		// TODO(eric): Save called function refs
		funcData.SetResultPtr(result.ToPtr())
		return result
	}
}

func GoTask[Arg packer, Result any](name string, f func(*Context, Arg) Result) func(*Context, Arg) Result {
	kind := getNextID()
	return func(parent *Context, arg Arg) Result {
		_, idx, ok := getFuncData(parent.global, kind, arg)

		if ok {
			// This function is in the disk cache
			if memCache, ok := parent.global.cache.FuncExtraData[idx]; ok {
				parent.global.stats.cacheHits += 1
				return memCache.(Result)
			}
		}

		parent.global.stats.cacheMisses += 1

		ctx := Context{
			name:   name,
			global: parent.global,
		}

		result := f(&ctx, arg)

		parent.global.cache.FuncExtraData[idx] = result

		return result
	}
}
