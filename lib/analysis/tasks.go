package analysis

import (
	"go.starlark.net/starlark"

	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/loading"
	"github.com/HALtheWise/bagel/lib/refs"
)

type FileInfo struct {
	// Which action generated this file, or NO_ACTION
	Generator refs.ActionRef

	// Is this executable? (may not be necessary)
	Executable bool
}

var T_FileInfo = core.Task1("T_FileInfo", func(c *core.Context, ref refs.CFileRef) *FileInfo {
	file := ref.Get(c)

	if file.Source == refs.FILESYSTEM_SOURCE {
		return &FileInfo{
			Generator:  refs.NO_ACTION,
			Executable: false, // TODO?
		}
	}

	targetInfo := T_AnalyzeTarget(c, file.Source)

	var generatingAction *Action
	var generatingActionIndex int

	for i, action := range targetInfo.Actions {
		for _, output := range action.Outputs {
			if output == ref {
				if generatingAction != nil {
					panic("Multiple actions produced " + file.String())
				}
				generatingAction = targetInfo.Actions[i]
				generatingActionIndex = i
			}
		}
	}
	if generatingAction == nil {
		panic("No action produced " + file.String())
	}

	return &FileInfo{
		Generator: refs.CActionTable.Insert(c, refs.CAction{
			Source: file.Source,
			Id:     uint32(generatingActionIndex),
		}),
		Executable: generatingAction.Executable,
	}
})

type ActionKind int8

const (
	WRITE ActionKind = iota
)

type Action struct {
	Kind            ActionKind
	Inputs, Outputs []refs.CFileRef

	Executable   bool
	WriteContent string
}

var T_ActionInfo = core.Task1("T_ActionInfo", func(c *core.Context, ref refs.ActionRef) *Action {
	action := ref.Get(c)

	targetInfo := T_AnalyzeTarget(c, action.Source)
	return targetInfo.Actions[action.Id]
})

type analyzedTarget struct {
	Name, Kind refs.StringRef
	Providers  []loading.Provider
	Actions    []*Action
}

var T_AnalyzeTarget = core.Task1("T_AnalyzeTarget", func(c *core.Context, label refs.CLabelRef) *analyzedTarget {
	clabel_v := label.Get(c)
	label_v := clabel_v.Label.Get(c)

	if clabel_v.Config != refs.TARGET_CONFIG {
		panic("We only support target config right now... Sorry...")
	}

	unconfigured := loading.T_LoadTarget(c, clabel_v.Label)
	if unconfigured == nil {
		return nil
	}

	thread := &starlark.Thread{
		Name: "Rule evaluation thread: " + label.String(),
		Load: loading.LoadFunc(c, label_v.Pkg)}

	bzlCtx := &BzlCtx{ctx: c,
		clabel: label,
		pkg:    label_v.Pkg}

	bzlResult, err := starlark.Call(thread, unconfigured.Rule.Impl, starlark.Tuple{bzlCtx}, nil)
	if err != nil {
		panic(err)
	}

	var providers []loading.Provider

	if seq, ok := bzlResult.(starlark.Indexable); ok {
		for i := 0; i < seq.Len(); i++ {
			providers = append(providers, *seq.Index(i).(*loading.Provider))
		}
	}

	actions := bzlCtx.actions

	return &analyzedTarget{
		Kind:      refs.StringTable.Insert(c, unconfigured.Rule.Kind),
		Name:      label_v.Name,
		Providers: providers,
		Actions:   actions,
	}
})
