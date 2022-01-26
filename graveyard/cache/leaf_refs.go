package cache

// A FileRef refers to a specific file visible to the Bagel build
// The left Ref refers to a workspace name (or "" for the default workspace)
// The right Ref refers to a Unix filepath within that workspace
type FileRef struct {
	genRef3[StringRef, StringRef]
}

func (FileRef) unpack(data uint32) FileRef {
	var result FileRef
	result.fillFromUnpack(data)
	return result
}

func InternFile(c *GlobalCache, workspace, path StringRef) FileRef {
	var result FileRef
	result.fillFromIntern(c, workspace, path)
	return result
}

// A DirRef refers to a specific directory visible to the Bagel build
// The left Ref refers to a workspace name (or "" for the default workspace)
// The right Ref refers to a Unix filepath within that workspace (no trailing /)
type DirRef struct {
	genRef3[StringRef, StringRef]
}

func (DirRef) unpack(data uint32) DirRef {
	var result DirRef
	result.fillFromUnpack(data)
	return result
}

func InternDir(c *GlobalCache, workspace, path StringRef) DirRef {
	var result DirRef
	result.fillFromIntern(c, workspace, path)
	return result
}
