package refs

import (
	"fmt"

	"github.com/HALtheWise/bagel/lib/core"
)

// A Config refers to a unique compile configuration
type ConfigRef uint32

const (
	TARGET_CONFIG ConfigRef = iota
	EXEC_CONFIG

	// Rules cannot be built in this configuration, it's just a marker
	FROM_FILESYSTEM_CONFIG
)

var ConfigTable = core.NewInternTable(map[string]ConfigRef{
	"target":              TARGET_CONFIG,
	"exec":                EXEC_CONFIG,
	"__from_filesystem__": FROM_FILESYSTEM_CONFIG,
})

func (r ConfigRef) Get(c *core.Context) string { return ConfigTable.Get(c, r) }

// A CLabelRef refers to a unique analyzed target, specified by its label and the active config.
type CLabelRef uint32
type CLabel struct {
	Label  LabelRef
	Config ConfigRef
}

const FILESYSTEM_SOURCE CLabelRef = iota

var CLabelTable = core.NewInternTable(map[CLabel]CLabelRef{
	{INVALID_LABEL, FROM_FILESYSTEM_CONFIG}: FILESYSTEM_SOURCE,
})

func (r CLabelRef) Get(c *core.Context) CLabel { return CLabelTable.Get(c, r) }

// A CFileRef refers to a file that can be an input of an Action.
// If the file is not generated, the Source will be FILESYSTEM_SOURCE
type CFileRef uint32
type CFile struct {
	Location LabelRef
	Source   CLabelRef
}

var CFileTable = core.NewInternTable(map[CFile]CFileRef{})

func (r CFileRef) Get(c *core.Context) CFile { return CFileTable.Get(c, r) }

// An ActionRef refers to a specific action instantiated by the listed target
type ActionRef uint32
type CAction struct {
	Source CLabelRef
	Id     uint32
}

const NO_ACTION ActionRef = iota

var CActionTable = core.NewInternTable(map[CAction]ActionRef{
	{FILESYSTEM_SOURCE, 0}: NO_ACTION,
})

func (r ActionRef) Get(c *core.Context) CAction { return CActionTable.Get(c, r) }

// String functions

func (r ConfigRef) String() string { return fmt.Sprintf(`r"%s"`, r.Get(core.DefaultContext)) }

func (r CLabelRef) String() string { return fmt.Sprintf("r%+v", r.Get(core.DefaultContext)) }

func (l CLabel) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("%s (%s)", l.Label.Get(c), l.Config.Get(c))
}

func (r CFileRef) String() string { return fmt.Sprintf("r%+v", r.Get(core.DefaultContext)) }

func (f CFile) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("File(%s, %s)", f.Location.Get(c), f.Source.Get(c))
}

func (r ActionRef) String() string { return fmt.Sprintf("r%+v", r.Get(core.DefaultContext)) }

func (f CAction) String() string {
	c := core.DefaultContext
	return fmt.Sprintf("Action(%s, #%d)", f.Source.Get(c), f.Id)
}
