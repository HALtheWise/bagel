package task

var namesToID = make(map[string]int)
var maxID = -1

func getNextID() int {
	maxID += 1
	return maxID
}

func getMaxID() int {
	return maxID
}

func idForName(name string) int {
	if id, ok := namesToID[name]; ok {
		return id
	}
	maxID += 1
	namesToID[name] = maxID
	return maxID
}
