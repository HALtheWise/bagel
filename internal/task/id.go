package task

var namesToID = make(map[string]uint32)

const INVALID_ID uint32 = 0

var maxID = INVALID_ID

func getNextID() uint32 {
	maxID += 1
	return maxID
}

func getMaxID() uint32 {
	return maxID
}

func idForName(name string) uint32 {
	if id, ok := namesToID[name]; ok {
		return id
	}
	maxID += 1
	namesToID[name] = maxID
	return maxID
}
