package execution

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HALtheWise/bagel/lib/analysis"
	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/loading"
	"github.com/HALtheWise/bagel/lib/refs"
)

var T_FileExecute = core.Task1("T_FileExecute", func(c *core.Context, ref refs.CFileRef) error {
	file := ref.Get(c)
	if file.Source == refs.FILESYSTEM_SOURCE {
		// Don't need to build code already on the filesystem
		return nil
	}

	info := analysis.T_FileInfo(c, ref)

	return T_ActionExecute(c, info.Generator)
})

var T_ActionExecute = core.Task1("T_ActionExecute", func(c *core.Context, ref refs.ActionRef) error {
	info := analysis.T_ActionInfo(c, ref)

	// TODO: build inputs

	switch info.Kind {
	case analysis.WRITE:
		if len(info.Outputs) != 1 {
			panic("Write must have one output")
		}

		output := info.Outputs[0]

		outPath := refs.T_FilepathForCFile(c, output)

		err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
		if err != nil {
			return fmt.Errorf("While making directories for %v: %w", outPath, err)
		}

		err = os.WriteFile(outPath, []byte(info.WriteContent), os.ModePerm)
		if err != nil {
			return fmt.Errorf("While writing %s: %w", outPath, err)
		}
		fmt.Println("Wrote file", outPath)
		return nil
	default:
		panic(fmt.Sprintf("Unknown action kind: %d", info.Kind))
	}
})

var T_BuildDefaultInfo = core.Task1("T_BuildDefaultInfo", func(c *core.Context, ref refs.LabelRef) error {
	clabel := refs.CLabelTable.Insert(c, refs.CLabel{ref, refs.TARGET_CONFIG})

	targetInfo := analysis.T_AnalyzeTarget(c, clabel)

	var defaultInfo *loading.Provider

	for i, provider := range targetInfo.Providers {
		if provider.Kind == loading.DefaultInfo {
			if defaultInfo != nil {
				return fmt.Errorf("Multiple copies of %s returned from %s", provider.Kind.Name(), ref)
			}
			defaultInfo = &targetInfo.Providers[i]
		}
	}

	if defaultInfo == nil {
		return fmt.Errorf("%s did not return DefaultInfo provider", clabel)
	}

	files := defaultInfo.Data["files"].(*loading.Depset)

	for _, value := range files.Items {
		file := value.(*analysis.BzlFile).Ref
		err := T_FileExecute(c, file)
		if err != nil {
			return fmt.Errorf("while building %s: %w", file, err)
		}
	}
	return nil
})
