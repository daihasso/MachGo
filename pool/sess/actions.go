package sess

import (
	"MachGo/base"
)


func Saved(object base.Base) bool {
	// NOTE: This requires a pointer because objectIsSaved and further
	//       calees assume a ptr, is this appropriate?
	saved, _ := objectIsSaved(object)
	return saved
}

