package execution

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HALtheWise/bagel/lib/analysis"
	"github.com/HALtheWise/bagel/lib/core"
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
