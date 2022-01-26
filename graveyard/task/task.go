package task

import (
	capnp "zombiezen.com/go/capnproto2"

	"github.com/HALtheWise/bagel/lib/cache/graph"
)

type packer interface{ Pack() uint32 }
type diskResult interface {
	ToPtr() capnp.Ptr
	~struct{ capnp.Struct }
}

func getFuncData(ctx *globalContext, kind uint32, arg packer) (graph.FuncObj, uint32, bool) {
	idx, created := ctx.Cache.InternFunc(kind, arg.Pack())
	funcs, err := ctx.Cache.Funcs()
	if err != nil {
		panic(err)
	}
	data := funcs.At(int(idx))
	return data, idx, created
}

func DiskTask[Arg packer, Result diskResult](name string, f func(*Context, Arg, *capnp.Segment) Result) func(*Context, Arg) Result {
	kind := getNextID()
	return func(parent *Context, arg Arg) Result {
		funcData, _, created := getFuncData(parent.Global, kind, arg)

		if !created {
			// This data is already cached.
			// TODO(eric): Check for invalidation
			parent.Global.stats.cacheHits += 1

			ptr, err := funcData.ResultPtr()
			if err != nil {
				panic(err)
			}
			result := struct{ capnp.Struct }{ptr.Struct()}
			return Result(result)
		}

		parent.Global.stats.cacheMisses += 1
		ctx := Context{
			name:   name,
			Global: parent.Global,
		}

		result := f(&ctx, arg, parent.Global.Cache.Segment())

		// TODO(eric): Save called function refs
		funcData.SetResultPtr(result.ToPtr())
		return result
	}
}

func GoTask[Arg packer, Result any](name string, f func(*Context, Arg) Result) func(*Context, Arg) Result {
	kind := getNextID()
	return func(parent *Context, arg Arg) Result {
		_, idx, created := getFuncData(parent.Global, kind, arg)

		if !created {
			// This function is in the disk cache
			if memCache, ok := parent.Global.Cache.FuncExtraData[idx]; ok {
				parent.Global.stats.cacheHits += 1
				return memCache.(Result)
			}
		}

		parent.Global.stats.cacheMisses += 1

		ctx := Context{
			name:   name,
			Global: parent.Global,
		}

		result := f(&ctx, arg)

		parent.Global.Cache.FuncExtraData[idx] = result

		return result
	}
}
