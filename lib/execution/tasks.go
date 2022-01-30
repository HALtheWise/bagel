package execution

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/HALtheWise/bagel/lib/analysis"
	"github.com/HALtheWise/bagel/lib/core"
	"github.com/HALtheWise/bagel/lib/refs"
)

var T_FileExecute func(c *core.Context, ref refs.CFileRef) error

func init() {
	T_FileExecute = core.Task1("T_FileExecute", func(c *core.Context, ref refs.CFileRef) error {
		file := ref.Get(c)
		if file.Source == refs.FILESYSTEM_SOURCE {
			// Don't need to build code already on the filesystem
			return nil
		}

		info := analysis.T_FileInfo(c, ref)

		return T_ActionExecute(c, info.Generator)
	})
}

var T_ActionExecute = core.Task1("T_ActionExecute", func(c *core.Context, ref refs.ActionRef) error {
	info := analysis.T_ActionInfo(c, ref)

	for _, file := range info.Inputs {
		err := T_FileExecute(c, file)
		if err != nil {
			return err
		}
	}

	switch info.Kind {
	case analysis.WRITE:
		return doWrite(c, info)
	case analysis.EXPAND_TEMPLATE:
		return doExpandTemplate(c, info)
	default:
		panic(fmt.Sprintf("Unknown action kind: %d", info.Kind))
	}
})

func doExpandTemplate(c *core.Context, info *analysis.Action) error {
	if len(info.Outputs) != 1 || len(info.Inputs) != 1 || info.ExpandTemplateSubstitutions == nil {
		panic("action invariant error")
	}
	output := info.Outputs[0]
	template := info.Inputs[0]

	outPath := refs.T_FilepathForCFile(c, output)
	templatePath := refs.T_FilepathForCFile(c, template)

	err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("While making directories for %v: %w", outPath, err)
	}

	var oldnew []string
	for k, v := range info.ExpandTemplateSubstitutions {
		oldnew = append(oldnew, k, v)
	}

	replacer := strings.NewReplacer(oldnew...)

	contents, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}
	replaced := replacer.Replace(string(contents))
	err = os.WriteFile(outPath, []byte(replaced), os.ModePerm)

	fmt.Println("Wrote template output to ", outPath)

	return err
}

func doWrite(c *core.Context, info *analysis.Action) error {
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
}

var T_BuildDefaultInfo = core.Task1("T_BuildDefaultInfo", func(c *core.Context, ref refs.LabelRef) error {
	clabel := refs.CLabelTable.Insert(c, refs.CLabel{Label: ref, Config: refs.TARGET_CONFIG})

	targetInfo := analysis.T_AnalyzeTarget(c, clabel)

	files, err := analysis.GetDefaultFiles(targetInfo.Providers)
	if err != nil {
		return fmt.Errorf("While building %s: %w", clabel, err)
	}

	for _, file := range files {
		err := T_FileExecute(c, file)
		if err != nil {
			return fmt.Errorf("while building %s: %w", file, err)
		}
	}

	return nil
})
